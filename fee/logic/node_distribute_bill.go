package logic

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/fee/db"
	"gopkg.in/guregu/null.v4"
)

// NodeDistributeRate 机器节点独占费率
type NodeDistributeRate struct {
	mutex sync.Mutex
	// rate36CPU 36 核心节点费率/年
	rate36CPU float64
	// rate4GPU 4 gpu核心节点费率/年
	rate4GPU float64
	// rate8GPU 8 gpu核心节点费率/年
	rate8GPU float64
}

// Get36CPURate 查询36 CPU节点的费率
func (ndr *NodeDistributeRate) Get36CPURate() float64 {
	ndr.mutex.Lock()
	defer ndr.mutex.Unlock()
	return ndr.rate36CPU
}

// Get4GPU 查询4 GPU节点的费率
func (ndr *NodeDistributeRate) Get4GPU() float64 {
	ndr.mutex.Lock()
	defer ndr.mutex.Unlock()
	return ndr.rate4GPU
}

// Get8GPU 查询8 GPU节点的费率
func (ndr *NodeDistributeRate) Get8GPU() float64 {
	ndr.mutex.Lock()
	defer ndr.mutex.Unlock()
	return ndr.rate8GPU
}

// NodeDistributeBill 机器独占账单操作逻辑
type NodeDistributeBill struct {
	ndb *db.NodeDistributeBillDB

	NodeDistributeRate
}

// Create 创建新的机器节点独占账单
func (ndbl *NodeDistributeBill) Create(
	ctx context.Context,
	applyID int,
	nodeDistributeID int,
	fee float64,
	userID int,
	username, userName string,
	groupID int,
) (int64, error) {
	if applyID <= 0 {
		return 0, errors.New("invalid apply id")
	}
	if nodeDistributeID <= 0 {
		return 0, errors.New("invalid node distribute id")
	}
	if fee < 0 {
		return 0, errors.New("invalid fee")
	}
	if userID <= 0 {
		return 0, errors.New("invalid user id")
	}
	if groupID <= 0 {
		return 0, errors.New("invalid groupID")
	}
	return ndbl.ndb.Insert(ctx, &db.NodeDistributeBill{
		ApplyID:          applyID,
		NodeDistributeID: nodeDistributeID,
		Fee:              fee,
		UserID:           userID,
		Username:         username,
		UserName:         userName,
		UserGroupID:      groupID,
		CreateTime:       time.Now(),
	})
}

// CalFee 计算账单的费用
func (ndbl *NodeDistributeBill) CalFee(startTimeUnix, endTimeUnix int64, nodeType string, nodeNum int) (float64, error) {
	year := ndbl.calTimeDurationYear(startTimeUnix, endTimeUnix)
	var res float64
	var err error = nil
	ndbl.mutex.Lock()
	defer ndbl.mutex.Unlock()
	switch nodeType {
	case "cpuc36":
		res = year * ndbl.rate36CPU
	case "gpuc4":
		res = year * ndbl.rate4GPU
	case "gpuc8":
		res = year * ndbl.rate8GPU
	default:
		res = 0
		err = errors.New("invalid nodeType")
	}
	res = res * float64(nodeNum)
	return res, err
}

func (ndbl *NodeDistributeBill) calTimeDurationYear(startTimeUnix, endTimeUnix int64) float64 {
	startTime := time.Unix(startTimeUnix, 0)
	endTime := time.Unix(endTimeUnix, 0)
	// TODO: 对于费率的计算存疑惑,目前将时间差换算为年,然后乘对应的费率
	yearDuration := endTime.Year() - startTime.Year()
	var year float64
	year = float64(yearDuration)
	year += float64((endTime.Month()+12-startTime.Month())%12) / 12
	return year
}

// PaginationGetNodeDistributeBillResult 分页查询的结果
type PaginationGetNodeDistributeBillResult struct {
	Data  []*db.NodeDistributeBill
	Count int
}

// PaginationGetAll 分页查询所有的记录
func (ndbl *NodeDistributeBill) PaginationGetAll(
	ctx context.Context,
	pageIndex, pageSize int,
) (*PaginationGetNodeDistributeBillResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}

	count, err := ndbl.ndb.QueryAllCount(ctx)
	if err != nil {
		return nil, err
	}
	limit := pageSize * (pageIndex - 1)
	data, err := ndbl.ndb.QueryAllWithLimit(ctx, limit, pageSize)
	if err != nil {
		return nil, err
	}
	return &PaginationGetNodeDistributeBillResult{
		Data:  data,
		Count: count,
	}, nil
}

// PaginationGetWithGroupID 分页查询属于某个组的记录
func (ndbl *NodeDistributeBill) PaginationGetWithGroupID(
	ctx context.Context,
	pageIndex, pageSize, groupID int,
) (*PaginationGetNodeDistributeBillResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}

	count, err := ndbl.ndb.QueryCountByGroupID(ctx, groupID)
	if err != nil {
		return nil, err
	}
	limit := pageSize * (pageIndex - 1)
	data, err := ndbl.ndb.QueryWithLimitByGroupID(ctx, limit, pageSize, groupID)
	if err != nil {
		return nil, err
	}
	return &PaginationGetNodeDistributeBillResult{
		Data:  data,
		Count: count,
	}, nil
}

// PaginationGetWithUserID 分页查询某个用户的账单记录
func (ndbl *NodeDistributeBill) PaginationGetWithUserID(
	ctx context.Context,
	pageIndex, pageSize, userID int,
) (*PaginationGetNodeDistributeBillResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}

	count, err := ndbl.ndb.QueryCountByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	limit := pageSize * (pageIndex - 1)
	data, err := ndbl.ndb.QueryWithLimitByUserID(ctx, limit, pageSize, userID)
	if err != nil {
		return nil, err
	}
	return &PaginationGetNodeDistributeBillResult{
		Data:  data,
		Count: count,
	}, nil
}

// PayBill 支付账单
func (ndbl *NodeDistributeBill) PayBill(
	ctx context.Context,
	billID int,
	money float64,
	message string,
	t PayType,
) (bool, error) {
	if billID <= 0 {
		return false, errors.New("invalid bill id")
	}
	if money < 0 {
		return false, errors.New("pay money must larger than 0")
	}
	if t != BalancePay && t != OfflinePay {
		return false, errors.New("invalid pay type")
	}
	return ndbl.ndb.UpdatePayFee(ctx, &db.NodeDistributeBill{
		ID:         billID,
		PayType:    null.IntFrom(int64(t)),
		PayTime:    null.TimeFrom(time.Now()),
		PayFee:     money,
		PayMessage: null.StringFrom(message),
	})
}

// GetInfoByID 通过ID查询信息
func (ndbl *NodeDistributeBill) GetInfoByID(ctx context.Context, id int) (*db.NodeDistributeBill, error) {
	return ndbl.ndb.QueryByID(ctx, id)
}

// GetRate 查询机器节点速率
func (ndbl *NodeDistributeBill) GetRate(ctx context.Context) *NodeDistributeRate {
	return &ndbl.NodeDistributeRate
}

// NewNodeDistributeBill 创建新的机器独占账单操作逻辑结构体
func NewNodeDistributeBill(ndb *db.NodeDistributeBillDB, dynamicConfig config.DynamicConfig) (*NodeDistributeBill, error) {
	res := &NodeDistributeBill{
		ndb: ndb,
	}

	// rate36CPU 36 核心节点费率/年
	var rate36CPU float64
	// rate4GPU 4 gpu核心节点费率/年
	var rate4GPU float64
	// rate8GPU 8 gpu核心节点费率/年
	var rate8GPU float64
	if err := dynamicConfig.Registry("fee_rate36CPU", &rate36CPU, func(newV interface{}) {
		res.mutex.Lock()
		defer res.mutex.Unlock()
		res.rate36CPU = rate36CPU
	}); err != nil {
		return nil, err
	}
	if err := dynamicConfig.Registry("fee_rate8GPU", &rate8GPU, func(newV interface{}) {
		res.mutex.Lock()
		defer res.mutex.Unlock()
		res.rate8GPU = rate8GPU
	}); err != nil {
		return nil, err
	}
	if err := dynamicConfig.Registry("fee_rate4GPU", &rate4GPU, func(newV interface{}) {
		res.mutex.Lock()
		defer res.mutex.Unlock()
		res.rate4GPU = rate4GPU
	}); err != nil {
		return nil, err
	}
	return res, nil
}

type PayType int8

const (
	// OfflinePay 线下缴费
	OfflinePay PayType = 1
	// BalancePay 余额缴费
	BalancePay PayType = 2
)
