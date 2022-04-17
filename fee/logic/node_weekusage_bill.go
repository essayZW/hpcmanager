package logic

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/fee/db"
	"github.com/shopspring/decimal"
)

type NodeUsageFeeRate struct {
	mutex sync.Mutex
	// cpu cpu节点费率/小时×核心
	cpu float64
	// gpu gpu节点费率/小时×核心
	gpu float64
}

// GetCPURate 查询CPU节点速率
func (this *NodeUsageFeeRate) GetCPURate() float64 {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	return this.cpu
}

// GetGPURate 查询GPU节点速率
func (this *NodeUsageFeeRate) GetGPURate() float64 {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	return this.gpu
}

// NodeWeekUsageBill 机器机时周账单
type NodeWeekUsageBill struct {
	nodeWeekUsageBillDB *db.NodeWeekUsageBillDB

	NodeUsageFeeRate
}

// CreateBill 创建机器时长周账单
func (this *NodeWeekUsageBill) CreateBill(
	ctx context.Context,
	userID, userGroupID int,
	username string,
	userName string,
	wallTime, gwallTime int,
	startTimeUnix, endTimeUnix int64,
) (int64, error) {
	if userID <= 0 {
		return 0, errors.New("invalid user id")
	}
	if wallTime < 0 {
		return 0, errors.New("walltime can't less than 0")
	}
	if gwallTime < 0 {
		return 0, errors.New("gwalltime can't less than 0")
	}
	startTime := time.Unix(startTimeUnix, 0)
	endTime := time.Unix(endTimeUnix, 0)
	if endTime.Before(startTime) {
		return 0, errors.New("endTime can't before than startTime")
	}
	fee := this.CalFee(ctx, wallTime, gwallTime)
	return this.nodeWeekUsageBillDB.Insert(ctx, &db.NodeWeekUsageBill{
		UserID:      userID,
		Username:    username,
		UserName:    userName,
		WallTime:    wallTime,
		GWallTime:   gwallTime,
		Fee:         fee,
		StartTime:   startTime,
		EndTime:     endTime,
		UserGroupID: userGroupID,
		CreateTime:  time.Now(),
	})
}

// CalFee 计算机时费用
func (this *NodeWeekUsageBill) CalFee(ctx context.Context, wallTime, gwallTime int) float64 {
	gwallFee := float64(gwallTime) * this.gpu
	wallFee := float64(wallTime) * this.cpu
	var fee decimal.Decimal
	if gwallFee > 0 {
		fee = decimal.NewFromFloat(gwallFee).Round(2)
	} else {
		fee = decimal.NewFromFloat(wallFee).Round(2)
	}
	resFee, _ := fee.DivRound(decimal.NewFromInt(3600), 2).Float64()
	return resFee
}

// PaginationGetWeekUsageBillResult 分页查询机器节点周账单结果
type PaginationGetWeekUsageBillResult struct {
	Count int
	Data  []*db.NodeWeekUsageBill
}

// PaginationGetWithTimeRange 分页查询一段时间内的所有的账单记录
func (this *NodeWeekUsageBill) PaginationGetWithTimeRange(
	ctx context.Context,
	pageIndex, pageSize int,
	startTimeUnix, endTimeUnix int64,
) (*PaginationGetWeekUsageBillResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}
	startTime := time.Unix(startTimeUnix, 0)
	endTime := time.Unix(endTimeUnix, 0)
	if endTime.Before(startTime) {
		return nil, errors.New("end time can't before than start time")
	}
	count, err := this.nodeWeekUsageBillDB.QueryCountWithTimeRange(ctx, startTime, endTime)
	if err != nil {
		return nil, errors.New("query count error")
	}
	limit := pageSize * (pageIndex - 1)
	data, err := this.nodeWeekUsageBillDB.QueryLimitWithTimeRange(ctx, limit, pageSize, startTime, endTime)
	if err != nil {
		return nil, errors.New("query data error")
	}
	return &PaginationGetWeekUsageBillResult{
		Count: count,
		Data:  data,
	}, nil
}

// PaginationGetWithTimeRangeWithUserID 分页查询某一段时间内的某个用户的账单
func (this *NodeWeekUsageBill) PaginationGetWithTimeRangeWithUserID(
	ctx context.Context,
	userID int,
	pageIndex, pageSize int,
	startTimeUnix, endTimeUnix int64,
) (*PaginationGetWeekUsageBillResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}
	startTime := time.Unix(startTimeUnix, 0)
	endTime := time.Unix(endTimeUnix, 0)
	if endTime.Before(startTime) {
		return nil, errors.New("end time can't before than start time")
	}
	count, err := this.nodeWeekUsageBillDB.QueryCountWithTimeRangeByUserID(ctx, userID, startTime, endTime)
	if err != nil {
		return nil, errors.New("query count error")
	}
	limit := pageSize * (pageIndex - 1)
	data, err := this.nodeWeekUsageBillDB.QueryLimitWithTimeRangeByUserID(ctx, userID, limit, pageSize, startTime, endTime)
	if err != nil {
		return nil, errors.New("query data error")
	}
	return &PaginationGetWeekUsageBillResult{
		Count: count,
		Data:  data,
	}, nil
}

// PaginationGetWithTimeRangeWithGroupID 分页查询某一段时间内的某个用户的账单
func (this *NodeWeekUsageBill) PaginationGetWithTimeRangeWithGroupID(
	ctx context.Context,
	userGroupID int,
	pageIndex, pageSize int,
	startTimeUnix, endTimeUnix int64,
) (*PaginationGetWeekUsageBillResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}
	startTime := time.Unix(startTimeUnix, 0)
	endTime := time.Unix(endTimeUnix, 0)
	if endTime.Before(startTime) {
		return nil, errors.New("end time can't before than start time")
	}
	count, err := this.nodeWeekUsageBillDB.QueryCountWithTimeRangeByGroupID(ctx, userGroupID, startTime, endTime)
	if err != nil {
		return nil, errors.New("query count error")
	}
	limit := pageSize * (pageIndex - 1)
	data, err := this.nodeWeekUsageBillDB.QueryLimitWithTimeRangeByGroupID(
		ctx,
		userGroupID,
		limit,
		pageSize,
		startTime,
		endTime,
	)
	if err != nil {
		return nil, errors.New("query data error")
	}
	return &PaginationGetWeekUsageBillResult{
		Count: count,
		Data:  data,
	}, nil
}

// PaginationGetGroupByGroupIDResult 分页查询所有组的账单结果
type PaginationGetGroupByGroupIDResult struct {
	Data  []*db.NodeWeekUsageBillForUserGroup
	Count int
}

// PaginationGetGroupByGroupIDWithPayFlag 按照组ID进行分组，分页查询某个支付状态下的所有结果
func (this *NodeWeekUsageBill) PaginationGetGroupByGroupIDWithPayFlag(
	ctx context.Context,
	pageIndex, pageSize int,
	payFlag bool,
) (*PaginationGetGroupByGroupIDResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}
	var payFlagInt int8
	if payFlag {
		payFlagInt = 1
	} else {
		payFlagInt = 0
	}
	count, err := this.nodeWeekUsageBillDB.QueryCountGroupByGroupID(ctx, payFlagInt)
	if err != nil {
		return nil, err
	}
	limit := pageSize * (pageIndex - 1)
	data, err := this.nodeWeekUsageBillDB.QueryGroupByGroupIDWithLimit(ctx, limit, pageSize, payFlagInt)
	if err != nil {
		return nil, err
	}
	return &PaginationGetGroupByGroupIDResult{
		Count: count,
		Data:  data,
	}, nil
}

// GetGroupBillByGroupID 通过用户组ID查询用户组的账单情况
func (this *NodeWeekUsageBill) GetGroupBillByGroupID(
	ctx context.Context,
	groupID int,
	payFlag bool,
) (*db.NodeWeekUsageBillForUserGroup, error) {
	var payFlagInt int8
	if payFlag {
		payFlagInt = 1
	} else {
		payFlagInt = 0
	}
	return this.nodeWeekUsageBillDB.QueryGroupBillByGroupID(ctx, groupID, payFlagInt)
}

// PayGroupBillByGroupID 通过用户组ID支付用户组的机时未缴费账单
func (this *NodeWeekUsageBill) PayGroupBillByGroupID(
	ctx context.Context,
	groupID int,
	payType PayType,
	payMessage string,
) (int64, error) {
	if groupID <= 0 {
		return 0, errors.New("invalid groupID")
	}
	if payType != OfflinePay && payType != BalancePay {
		return 0, errors.New("invalid pay type")
	}
	return this.nodeWeekUsageBillDB.UpdateBillsPayStatusByGroupID(ctx, groupID, payMessage, int8(payType), time.Now())
}

// NewNodeWeekUsageBill 创建新的机器机时周账单数据操作逻辑结构体
func NewNodeWeekUsageBill(
	nodeWeekUsageBillDB *db.NodeWeekUsageBillDB,
	dynamicConfig config.DynamicConfig,
) (*NodeWeekUsageBill, error) {
	res := &NodeWeekUsageBill{
		nodeWeekUsageBillDB: nodeWeekUsageBillDB,
	}

	var cpu float64
	var gpu float64

	err := dynamicConfig.Registry("fee_rate_CPU", &cpu, func(newV interface{}) {
		res.mutex.Lock()
		defer res.mutex.Unlock()
		res.cpu = cpu
	})
	if err != nil {
		return nil, err
	}

	err = dynamicConfig.Registry("fee_rate_GPU", &gpu, func(newV interface{}) {
		res.mutex.Lock()
		defer res.mutex.Unlock()
		res.gpu = gpu
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
