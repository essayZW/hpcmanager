package db

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
)

type NodeQuotaBillDB struct {
	conn *db.DB
}

// Insert 插入新的记录
func (this *NodeQuotaBillDB) Insert(ctx context.Context, newInfo *NodeQuotaBill) (int64, error) {
	res, err := this.conn.Exec(
		ctx,
		"INSERT INTO `node_quota_bill` "+
			"(`user_id`, `user_name`, `user_username`, `user_group_id`, `oper_type`, `old_size`, `new_size`, `old_end_time`, `new_end_time`, `fee`, `create_time`, `extraAttributes`) "+
			" VALUES (?,?,?,?,?,?,?,?,?,?,?,?)",
		newInfo.UserID,
		newInfo.UserName,
		newInfo.Username,
		newInfo.UserGroupID,
		newInfo.OperType,
		newInfo.OldSize,
		newInfo.NewSize,
		newInfo.OldEndTime,
		newInfo.NewEndTime,
		newInfo.Fee,
		newInfo.CreateTime,
		newInfo.ExtraAttributes,
	)
	if err != nil {
		logger.Warn("Insert node quota bill error: ", err)
		return 0, errors.New("Insert node quota bill error")
	}

	id, err := res.LastInsertId()
	if err != nil {
		logger.Warn("Insert node quota bill error: ", err)
		return 0, errors.New("Insert node quota bill error")
	}
	return id, nil
}

// NewNodeQuotaBill 创建新的机器存储账单数据库操作
func NewNodeQuotaBill(conn *db.DB) *NodeQuotaBillDB {
	return &NodeQuotaBillDB{
		conn: conn,
	}
}
