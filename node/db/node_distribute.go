package db

import (
	"context"
	"errors"

	hpcdb "github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
)

// NodeDistributeDB 机器节点分配处理工单表的操作
type NodeDistributeDB struct {
	conn *hpcdb.DB
}

// Insert 插入新的记录
func (ndb *NodeDistributeDB) Insert(ctx context.Context, info *NodeDistribute) (int64, error) {
	res, err := ndb.conn.Exec(ctx, "INSERT INTO `node_distribute` (`apply_id`, `create_time`, `extraAttributes`) VALUES (?,?,?)",
		info.ApplyID, info.CreateTime, info.ExtraAttributes)
	if err != nil {
		logger.Warn("Insert error: ", err)
		return 0, errors.New("insert node_distribute info error")
	}

	id, err := res.LastInsertId()
	if err != nil {
		logger.Warn("Insert error: ", err)
		return 0, errors.New("insert node_distribute info error")
	}
	return id, nil
}

// QueryByApplyID 通过申请ID查询工单信息
func (ndb *NodeDistributeDB) QueryByApplyID(ctx context.Context, applyID int) (*NodeDistribute, error) {
	row, err := ndb.conn.QueryRow(ctx, "SELECT * FROM `node_distribute` WHERE `apply_id`=?", applyID)
	if err != nil {
		logger.Warn("QueryByApplyID error: ", err)
		return nil, errors.New("QueryByApplyID error")
	}
	var info NodeDistribute
	if err := row.StructScan(&info); err != nil {
		logger.Warn("QueryByApplyID struct scan error: ", err)
		return nil, errors.New("QueryByApplyID struct scan error")
	}
	return &info, nil
}

// QueryCountByApply 通过申请ID查询工单数量信息
func (ndb *NodeDistributeDB) QueryCountByApply(ctx context.Context, applyID int) (int, error) {
	row, err := ndb.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `node_distribute` WHERE `apply_id`=?", applyID)
	if err != nil {
		logger.Warn("QueryCountByApply error: ", err)
		return 0, errors.New("QueryCountByApply error")
	}
	var info int
	if err := row.Scan(&info); err != nil {
		logger.Warn("QueryCountByApply struct scan error: ", err)
		return 0, errors.New("QueryCountByApply struct scan error")
	}
	return info, nil
}

// NewNodeDistribute 创建新的机器节点分配处理工单表
func NewNodeDistribute(conn *hpcdb.DB) *NodeDistributeDB {
	return &NodeDistributeDB{
		conn: conn,
	}
}
