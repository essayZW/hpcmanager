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

// NewNodeUsageTime 创建新的计算节点使用时长记录的操作逻辑
func NewNodeUsageTime(nodeUsageTimeDB *db.NodeUsageTimeDB) *NodeUsageTime {
	return &NodeUsageTime{
		nodeUsageTimeDB: nodeUsageTimeDB,
	}
}
