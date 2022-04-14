package logic

import (
	"context"
	"errors"
	"time"

	"github.com/essayZW/hpcmanager/node/db"
)

// NodeUsageTime 计算节点使用时长记录的操作逻辑结构体
type NodeUsageTime struct {
	nodeUsageTimeDB *db.NodeUsageTimeDB
}

// AddRecord 添加新的记录
func (n *NodeUsageTime) AddRecord(ctx context.Context, info *db.HpcUsageTime) (int64, error) {
	if info == nil {
		return 0, errors.New("info can't be nil")
	}
	if info.QueueName == "" {
		return 0, errors.New("invalid queueName")
	}
	if info.UserID == 0 {
		return 0, errors.New("invalid userID")
	}
	if info.TutorID == 0 {
		return 0, errors.New("invalid tutorID")
	}
	if info.HpcUsername == "" {
		return 0, errors.New("invalid hpc user name")
	}
	if info.HpcGroupName == "" {
		return 0, errors.New("invalid hpc group name")
	}
	if info.WallTime < 0 || info.GWallTime < 0 {
		return 0, errors.New("wall time and gwall time must larger than 0")
	}
	if info.EndTime.IsZero() || info.StartTime.IsZero() {
		return 0, errors.New("invalid time")
	}
	if info.StartTime.After(info.EndTime) {
		return 0, errors.New("start time can't after than end time")
	}
	info.CreateTime = time.Now()
	return n.nodeUsageTimeDB.Insert(ctx, info)
}

// PaginationGetNodeUsageRecordResult 分页查询机器使用时间记录的结果
type PaginationGetNodeUsageRecordResult struct {
	Data  []*db.HpcUsageTime
	Count int
}

// PaginationGetNodeUsageRecord 分页查询一段时间内的机器节点机器时长记录
func (n *NodeUsageTime) PaginationGetNodeUsageRecord(
	ctx context.Context,
	pageIndex, pageSize int,
	startUnixMilli, endUnixMilli int64,
) (*PaginationGetNodeUsageRecordResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}
	startDate := time.UnixMilli(startUnixMilli)
	endDate := time.UnixMilli(endUnixMilli)
	count, err := n.nodeUsageTimeDB.QueryAllCount(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}
	limit := pageSize * (pageIndex - 1)
	data, err := n.nodeUsageTimeDB.QueryAllWithLimit(ctx, limit, pageSize, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return &PaginationGetNodeUsageRecordResult{
		Data:  data,
		Count: count,
	}, nil
}

// PaginationGetNodeUsageRecordByUserID 通过用户ID分页查询对应用户在某段时间内的所有的机器节点使用情况记录
func (n *NodeUsageTime) PaginationGetNodeUsageRecordByUserID(
	ctx context.Context,
	userID, pageIndex, pageSize int,
	startUnixMilli, endUnixMilli int64,
) (*PaginationGetNodeUsageRecordResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}
	startDate := time.UnixMilli(startUnixMilli)
	endDate := time.UnixMilli(endUnixMilli)
	count, err := n.nodeUsageTimeDB.QueryCountByUserID(ctx, userID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	limit := pageSize * (pageIndex - 1)
	data, err := n.nodeUsageTimeDB.QueryWithLimitByUserID(ctx, userID, limit, pageSize, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return &PaginationGetNodeUsageRecordResult{
		Data:  data,
		Count: count,
	}, nil
}

// PaginationGetNodeUsageRecordByTutorID 通过导师ID分页查询一段时间内的机器使用时间节点
func (n *NodeUsageTime) PaginationGetNodeUsageRecordByTutorID(
	ctx context.Context,
	tutorID, pageIndex, pageSize int,
	startUnixMilli, endUnixMilli int64,
) (*PaginationGetNodeUsageRecordResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}
	startDate := time.UnixMilli(startUnixMilli)
	endDate := time.UnixMilli(endUnixMilli)
	count, err := n.nodeUsageTimeDB.QueryCountByTutorID(ctx, tutorID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	limit := pageSize * (pageIndex - 1)
	data, err := n.nodeUsageTimeDB.QueryWithLimitByTutorID(ctx, tutorID, limit, pageSize, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return &PaginationGetNodeUsageRecordResult{
		Data:  data,
		Count: count,
	}, nil
}

// NewNodeUsageTime 创建新的计算节点使用时长记录的操作逻辑
func NewNodeUsageTime(nodeUsageTimeDB *db.NodeUsageTimeDB) *NodeUsageTime {
	return &NodeUsageTime{
		nodeUsageTimeDB: nodeUsageTimeDB,
	}
}
