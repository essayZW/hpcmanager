package logic

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/fee/db"
	"github.com/essayZW/hpcmanager/fee/utils"
)

type NodeQuotaFeeRate struct {
	mutex sync.Mutex
	// basicPerYearPerTB 基本存储1TB/年,指的是最初初始化的1TB 1年的空间的费率
	basicPerYearPerTB float64
	// extraPerYearPerTB 额外的存储1TB/年
	extraPerYearPerTB float64
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
	if yearDuration == 0 {
		yearDuration = 1
	}

	var quotaFee float64
	if newSize > 1 {
		quotaFee = this.basicPerYearPerTB + float64(newSize-1)*this.extraPerYearPerTB
	} else {
		quotaFee = this.basicPerYearPerTB
	}
	return quotaFee * yearDuration
}

// NewNodeQuotaBill 创建新的节点存储账单操作逻辑
func NewNodeQuotaBill(nodeQuotaBillDB *db.NodeQuotaBillDB, dynamicConfig config.DynamicConfig) (*NodeQuotaBill, error) {
	res := &NodeQuotaBill{
		nodeQuotaBillDB: nodeQuotaBillDB,
	}

	var basicFeeRate float64
	var extraFeeRate float64

	err := dynamicConfig.Registry("fee_rate_Quota_basic", &basicFeeRate, func(newV interface{}) {
		res.mutex.Lock()
		defer res.mutex.Unlock()
		res.basicPerYearPerTB = basicFeeRate
	})
	if err != nil {
		return nil, err
	}

	err = dynamicConfig.Registry("fee_rate_Quota_extra", &extraFeeRate, func(newV interface{}) {
		res.mutex.Lock()
		defer res.mutex.Unlock()
		res.extraPerYearPerTB = extraFeeRate
	})
	return res, nil
}

// QuotaOperationType 存储操作类型
type QuoatOperationType int8

const (
	// ChangeQuotaSize 修改容量
	ChangeQuotaSize QuoatOperationType = 1
	// ChangeEndTime 修改期限的最后的时间
	ChangeEndTime QuoatOperationType = 2
)
