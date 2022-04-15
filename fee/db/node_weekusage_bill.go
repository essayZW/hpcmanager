package db

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
)

// NodeWeekUsageBillDB 机器节点时长周账单数据库操作映射结构体
type NodeWeekUsageBillDB struct {
	conn *db.DB
}

// Insert 插入新的记录
func (this *NodeWeekUsageBillDB) Insert(ctx context.Context, newInfo *NodeWeekUsageBill) (int64, error) {
	res, err := this.conn.Exec(
		ctx,
		"INSERT INTO `week_usage_bill` "+
			"(`user_id`, `user_username`, `user_name`, `wall_time`, `gwall_time`, `fee`, `start_time`, `end_time`, `user_group_id`, `create_time`, `extraAttributes`) "+
			" VALUES (?,?,?,?,?,?,?,?,?,?,?)",
		newInfo.UserID,
		newInfo.Username,
		newInfo.UserName,
		newInfo.WallTime,
		newInfo.GWallTime,
		newInfo.Fee,
		newInfo.StartTime,
		newInfo.EndTime,
		newInfo.UserGroupID,
		newInfo.CreateTime,
		newInfo.ExtraAttributes,
	)
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

// NewNodeWeekUsageBill 创建新的机器节点机时周账单数据库操作映射结构体
func NewNodeWeekUsageBill(conn *db.DB) *NodeWeekUsageBillDB {
	return &NodeWeekUsageBillDB{
		conn: conn,
	}
}
