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

// QueryLimit 分页查询记录
func (ndb *NodeDistributeDB) QueryLimit(ctx context.Context, limit, offset int) ([]*NodeDistribute, error) {
	row, err := ndb.conn.Query(ctx, "SELECT * FROM `node_distribute` ORDER BY `id` DESC LIMIT ?,?", limit, offset)
	if err != nil {
		logger.Warn("QueryLimit error: ", err)
		return nil, errors.New("QueryLimit error")
	}
	infos := make([]*NodeDistribute, 0)
	for row.Next() {
		var info NodeDistribute
		if err := row.StructScan(&info); err != nil {
			logger.Warn("QueryLimit struct scan error: ", err)
			return nil, errors.New("QueryLimit struct scan error")
		}
		infos = append(infos, &info)
	}
	return infos, nil
}

// QueryCount 查询总的记录数量
func (ndb *NodeDistributeDB) QueryCount(ctx context.Context) (int, error) {
	row, err := ndb.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `node_distribute`")
	if err != nil {
		logger.Warn("QueryCount error: ", err)
		return 0, errors.New("QueryCount error")
	}
	var count int
	if err := row.Scan(&count); err != nil {
		logger.Warn("QueryCount scan error: ", err)
		return 0, errors.New("QueryCount scan error")
	}
	return count, nil
}

// UpdateHandlerFlag 更新处理标记,包括处理人的相关信息
func (ndb *NodeDistributeDB) UpdateHandlerFlag(ctx context.Context, newInfo *NodeDistribute) (bool, error) {
	res, err := ndb.conn.Exec(ctx, "UPDATE `node_distribute` SET "+
		"`handler_flag`=1, `handler_userid`=?, `handler_username`=?, `handler_user_name`=? WHERE `id`=?",
		newInfo.HandlerUserID, newInfo.HandlerUsername, newInfo.HandlerUserName, newInfo.ID)
	if err != nil {
		logger.Warn("UpdateHandlerFlag error: ", err)
		return false, errors.New("UpdateHandlerFlag error")
	}

	count, err := res.RowsAffected()
	if err != nil {
		logger.Warn("UpdateHandlerFlag error: ", err)
		return false, errors.New("UpdateHandlerFlag error")
	}
	return count > 0, nil
}

// NewNodeDistribute 创建新的机器节点分配处理工单表
func NewNodeDistribute(conn *hpcdb.DB) *NodeDistributeDB {
	return &NodeDistributeDB{
		conn: conn,
	}
}
