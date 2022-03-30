package db

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
)

// NodeUsageTimeDB 计算节点使用记录的数据库表操作
type NodeUsageTimeDB struct {
	conn *db.DB
}

// Insert 插入新的记录
func (n *NodeUsageTimeDB) Insert(ctx context.Context, info *HpcUsageTime) (int64, error) {
	res, err := n.conn.Exec(ctx, "INSERT INTO `hpc_usagetime` "+
		"(`user_id`, `username`, `user_name`,`hpc_username`, `tutor_id`, `tutor_username`, `tutor_user_name`, `hpc_group_name`, `queue_name`, `wall_time`, `gwall_time`, `start_time`, `end_time`, `create_time`, `extraAttributes`) "+
		" VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", info.UserID, info.Username, info.UserName, info.HpcUsername, info.TutorID, info.TutorUsername,
		info.TutorUserName, info.HpcGroupName, info.QueueName, info.WallTime, info.GWallTime, info.StartTime,
		info.EndTime, info.CreateTime, info.ExtraAttributes)
	if err != nil {
		logger.Warn("Insert error: ", err)
		return 0, errors.New("Insert error")
	}
	id, err := res.LastInsertId()
	if err != nil {
		logger.Warn("Insert error: ", err)
		return 0, errors.New("Insert error")
	}
	return id, nil
}

// NewNodeUsageTime 创建新的操作计算节点记录的结构体
func NewNodeUsageTime(conn *db.DB) *NodeUsageTimeDB {
	return &NodeUsageTimeDB{
		conn: conn,
	}
}
