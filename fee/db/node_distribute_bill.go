package db

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
)

type NodeDistributeBillDB struct {
	conn *db.DB
}

// Insert 插入新的账单记录
func (ndb *NodeDistributeBillDB) Insert(ctx context.Context, newInfo *NodeDistributeBill) (int64, error) {
	res, err := ndb.conn.Exec(
		ctx,
		"INSERT INTO `node_distribute_bill` "+
			"(`apply_id`, `node_distribute_id`, `fee`, `user_id`, `user_username`, `user_name`, `create_time`, `extraAttributes`) "+
			"VALUES (?,?,?,?,?,?,?,?)",
		newInfo.ApplyID,
		newInfo.NodeDistributeID,
		newInfo.Fee,
		newInfo.UserID,
		newInfo.Username,
		newInfo.UserName,
		newInfo.CreateTime,
		newInfo.ExtraAttributes,
	)
	if err != nil {
		logger.Warn("Insert node distribute bill error: ", err)
		return 0, errors.New("Insert node distribute bill error")
	}
	id, err := res.LastInsertId()
	if err != nil {
		logger.Warn("Insert node distribute bill error: ", err)
		return 0, errors.New("Insert node distribute bill error")
	}
	return id, nil
}

// NewNodeDistributeBill 新建一个node_distribute_bill数据表操作结构体
func NewNodeDistributeBill(conn *db.DB) *NodeDistributeBillDB {
	return &NodeDistributeBillDB{
		conn: conn,
	}
}
