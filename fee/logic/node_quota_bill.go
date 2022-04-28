package logic

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/fee/db"
	"github.com/essayZW/hpcmanager/fee/utils"
	"gopkg.in/guregu/null.v4"
)

const (
	basicKey = "fee_rate_Quota_basic"
	extraKey = "fee_rate_Quota_extra"
)

// NodeQuotaFeeRate 机器节点存储费率
type NodeQuotaFeeRate struct {
	dynamicConfig config.DynamicConfig
	mutex         sync.Mutex
	// basicPerYearPerTB 基本存储1TB/年,指的是最初初始化的1TB 1年的空间的费率
	basicPerYearPerTB float64
	// extraPerYearPerTB 额外的存储1TB/年
	extraPerYearPerTB float64
}

// GetBasic 获取基本存储费率
func (this *NodeQuotaFeeRate) GetBasic() float64 {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	return this.basicPerYearPerTB
}

// SetBasic 设置基本存储费率
func (this *NodeQuotaFeeRate) SetBasic(ctx context.Context, value float64) error {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	return this.dynamicConfig.Put(ctx, basicKey, value)
}

// GetExtra 获得额外存储的费率
func (this *NodeQuotaFeeRate) GetExtra() float64 {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	return this.extraPerYearPerTB
}

// SetExtra 设置额外存储的费率
func (this *NodeQuotaFeeRate) SetExtra(ctx context.Context, value float64) error {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	return this.dynamicConfig.Put(ctx, extraKey, value)
}

func (this *NodeQuotaFeeRate) Registry() error {
	var basicFeeRate float64
	var extraFeeRate float64

	err := this.dynamicConfig.Registry(basicKey, &basicFeeRate, func(newV interface{}) {
		this.mutex.Lock()
		defer this.mutex.Unlock()
		this.basicPerYearPerTB = basicFeeRate
	})
	if err != nil {
		return err
	}

	err = this.dynamicConfig.Registry(extraKey, &extraFeeRate, func(newV interface{}) {
		this.mutex.Lock()
		defer this.mutex.Unlock()
		this.extraPerYearPerTB = extraFeeRate
	})

	if err != nil {
		return err
	}
	return nil
}

type NodeQuotaBill struct {
	NodeQuotaFeeRate

	nodeQuotaBillDB *db.NodeQuotaBillDB
}

// CreateNewBillParam 创建账单参数
type CreateNewBillParam struct {
	UserID      int
	UserName    string
	Username    string
	UserGroupID int
	OperType    QuoatOperationType
	OldSize     int
	NewSize     int
	OldEndTime  int64
	NewEndTime  int64
}

// CreateNewBill 创建新的账单
func (this *NodeQuotaBill) CreateNewBill(ctx context.Context, param *CreateNewBillParam) (int64, error) {
	if param == nil {
		return 0, errors.New("need param")
	}
	if param.UserID <= 0 {
		return 0, errors.New("invalid user id")
	}
	if param.UserGroupID <= 0 {
		return 0, errors.New("invalid user group id")
	}
	if param.OperType != ChangeQuotaSize && param.OperType != ChangeEndTime {
		return 0, errors.New("invalid oper type")
	}
	oldEndTime := time.Unix(param.OldEndTime, 0)
	newEndTime := time.Unix(param.NewEndTime, 0)

	fee := this.CalFee(param.OldSize, param.NewSize, oldEndTime, newEndTime)
	return this.nodeQuotaBillDB.Insert(ctx, &db.NodeQuotaBill{
		UserID:      param.UserID,
		Username:    param.Username,
		UserName:    param.UserName,
		UserGroupID: param.UserGroupID,
		OperType:    int8(param.OperType),
		OldSize:     param.OldSize,
		NewSize:     param.NewSize,
		OldEndTime:  oldEndTime,
		NewEndTime:  newEndTime,
		Fee:         fee,
		CreateTime:  time.Now(),
	})
}

// CalFee 计算费用
func (this *NodeQuotaBill) CalFee(oldSize, newSize int, oldEndTime, newEndTime time.Time) float64 {
	yearDuration := utils.CalYearDuration(oldEndTime, newEndTime)

	// 现有的容量每年需要的费用,用来计算延期的费用
	var quotaFeePerYear float64
	if oldSize > 0 {
		quotaFeePerYear = this.basicPerYearPerTB + float64(oldSize-1)*this.extraPerYearPerTB
	} else {
		quotaFeePerYear = 0
	}

	expandSize := newSize - oldSize
	// 扩展的容量每年需要的费用
	var expandFeePerYear float64
	if oldSize == 0 {
		expandFeePerYear = this.basicPerYearPerTB + this.extraPerYearPerTB*float64(expandSize-1)
	} else {
		expandFeePerYear = this.extraPerYearPerTB * float64(expandSize)
	}
	// 由于扩展的容量在今后才被使用,所以计算现在到结束的时间差
	expandSizeDuration := utils.CalYearDuration(time.Now(), newEndTime)
	// 延期带来的费用加上扩容带来的费用
	return quotaFeePerYear*yearDuration + expandFeePerYear*expandSizeDuration
}

// PaginationGetNodeQuotaBillsResult 分页查询机器存储账单的结果
type PaginationGetNodeQuotaBillsResult struct {
	Count int
	Data  []*db.NodeQuotaBill
}

// PaginationGetAllNodeQuotaBill 分页查询所有的机器存储账单
func (this *NodeQuotaBill) PaginationGetAllNodeQuotaBill(
	ctx context.Context,
	pageIndex, pageSize int,
) (*PaginationGetNodeQuotaBillsResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}
	count, err := this.nodeQuotaBillDB.QueryAllCount(ctx)
	if err != nil {
		return nil, err
	}
	limit := pageSize * (pageIndex - 1)
	data, err := this.nodeQuotaBillDB.QueryAllWithLimit(ctx, limit, pageSize)
	return &PaginationGetNodeQuotaBillsResult{
		Count: count,
		Data:  data,
	}, nil
}

// PaginationGetNodeQuotaBillByUserID 分页查询某个用户的所有的机器存储账单
func (this *NodeQuotaBill) PaginationGetNodeQuotaBillByUserID(
	ctx context.Context,
	userID int,
	pageIndex, pageSize int,
) (*PaginationGetNodeQuotaBillsResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}
	count, err := this.nodeQuotaBillDB.QueryAllCountByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	limit := pageSize * (pageIndex - 1)
	data, err := this.nodeQuotaBillDB.QueryAllWithLimitByUserID(ctx, limit, pageSize, userID)
	return &PaginationGetNodeQuotaBillsResult{
		Count: count,
		Data:  data,
	}, nil
}

// PaginationGetNodeQuotaBillByGroupID 分页查询某个组的所有机器节点账单
func (this *NodeQuotaBill) PaginationGetNodeQuotaBillByGroupID(
	ctx context.Context,
	groupID int,
	pageIndex, pageSize int,
) (*PaginationGetNodeQuotaBillsResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}
	count, err := this.nodeQuotaBillDB.QueryAllCountByUserGroupID(ctx, groupID)
	if err != nil {
		return nil, err
	}
	limit := pageSize * (pageIndex - 1)
	data, err := this.nodeQuotaBillDB.QueryAllWithLimitByUserGroupID(ctx, limit, pageSize, groupID)
	return &PaginationGetNodeQuotaBillsResult{
		Count: count,
		Data:  data,
	}, nil
}

// PayBill 支付账单
func (this *NodeQuotaBill) PayBill(
	ctx context.Context,
	billID int,
	payMoney float64,
	payType PayType,
	payMessage string,
) (bool, error) {
	if billID <= 0 {
		return false, errors.New("invalid bill id")
	}
	if payType != OfflinePay && payType != BalancePay {
		return false, errors.New("invalid pay type")
	}
	return this.nodeQuotaBillDB.UpdatePayStatus(ctx, &db.NodeQuotaBill{
		ID:         billID,
		PayFee:     payMoney,
		PayType:    null.IntFrom(int64(payType)),
		PayMessage: null.StringFrom(payMessage),
		PayTime:    null.TimeFrom(time.Now()),
	})
}

// GetInfoByID 通过ID查询
func (this *NodeQuotaBill) GetInfoByID(ctx context.Context, billID int) (*db.NodeQuotaBill, error) {
	return this.nodeQuotaBillDB.QueryByID(ctx, billID)
}

// QuotaOperationType 存储操作类型
type QuoatOperationType int8

const (
	// ChangeQuotaSize 修改容量
	ChangeQuotaSize QuoatOperationType = 1
	// ChangeEndTime 修改期限的最后的时间
	ChangeEndTime QuoatOperationType = 2
)

// NewNodeQuotaBill 创建新的节点存储账单操作逻辑
func NewNodeQuotaBill(nodeQuotaBillDB *db.NodeQuotaBillDB, dynamicConfig config.DynamicConfig) (*NodeQuotaBill, error) {
	res := &NodeQuotaBill{
		nodeQuotaBillDB: nodeQuotaBillDB,
		NodeQuotaFeeRate: NodeQuotaFeeRate{
			dynamicConfig: dynamicConfig,
		},
	}
	if err := res.Registry(); err != nil {
		return nil, err
	}
	return res, nil
}
