package logic

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/fee/db"
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
	// TODO: 需要确真实的计算算法
	return this.cpu*float64(wallTime) + this.gpu*float64(gwallTime)
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
