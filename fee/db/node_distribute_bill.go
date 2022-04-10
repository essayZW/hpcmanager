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
			"(`apply_id`, `node_distribute_id`, `fee`, `user_id`, `user_username`, `user_name`, `user_group_id`, `create_time`, `extraAttributes`) "+
			"VALUES (?,?,?,?,?,?,?,?,?)",
		newInfo.ApplyID,
		newInfo.NodeDistributeID,
		newInfo.Fee,
		newInfo.UserID,
		newInfo.Username,
		newInfo.UserName,
		newInfo.UserGroupID,
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

// QueryAllCount 查询一定时间范围内的所有的记录的数量
func (ndb *NodeDistributeBillDB) QueryAllCount(ctx context.Context) (int, error) {
	res, err := ndb.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `node_distribute_bill`")
	if err != nil {
		logger.Warn("QueryAllCount error: ", err)
		return 0, errors.New("QueryAllCount error")
	}
	var count int
	if err := res.Scan(&count); err != nil {
		logger.Warn("QueryAllCount error: ", err)
		return 0, errors.New("QueryAllCount error")
	}
	return count, nil
}

// QueryAllWithLimit 分页查询账单信息
func (ndb *NodeDistributeBillDB) QueryAllWithLimit(ctx context.Context, limit, offset int) ([]*NodeDistributeBill, error) {
	rows, err := ndb.conn.Query(ctx, "SELECT * FROM `node_distribute_bill` LIMIT ?,?", limit, offset)
	if err != nil {
		logger.Warn("QueryAllWithLimit error: ", err)
		return nil, errors.New("QueryAllWithLimit error")
	}
	infos := make([]*NodeDistributeBill, 0)
	for rows.Next() {
		var info NodeDistributeBill
		if err := rows.StructScan(&info); err != nil {
			logger.Warn("QueryAllWithLimit struct scan error: ", err)
			return nil, errors.New("QueryAllWithLimit struct scan error")
		}
		infos = append(infos, &info)
	}
	return infos, nil
}

// QueryCountByGroupID 查询某个组ID下的记录的数量
func (ndb *NodeDistributeBillDB) QueryCountByGroupID(ctx context.Context, groupID int) (int, error) {
	res, err := ndb.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `node_distribute_bill` WHERE `user_group_id`=?", groupID)
	if err != nil {
		logger.Warn("QueryCountByGroupID error: ", err)
		return 0, errors.New("QueryCountByGroupID error")
	}
	var count int
	if err := res.Scan(&count); err != nil {
		logger.Warn("QueryCountByGroupID error: ", err)
		return 0, errors.New("QueryCountByGroupID error")
	}
	return count, nil
}

// QueryWithLimitByGroupID 分页查询属于某个组的记录
func (ndb *NodeDistributeBillDB) QueryWithLimitByGroupID(
	ctx context.Context,
	limit, offset, groupID int,
) ([]*NodeDistributeBill, error) {
	rows, err := ndb.conn.Query(
		ctx,
		"SELECT * FROM `node_distribute_bill` WHERE `user_group_id`=? LIMIT ?,?",
		groupID,
		limit,
		offset,
	)
	if err != nil {
		logger.Warn("QueryWithLimitByGroupID error: ", err)
		return nil, errors.New("QueryWithLimitByGroupID error")
	}
	infos := make([]*NodeDistributeBill, 0)
	for rows.Next() {
		var info NodeDistributeBill
		if err := rows.StructScan(&info); err != nil {
			logger.Warn("QueryWithLimitByGroupID struct scan error: ", err)
			return nil, errors.New("QueryWithLimitByGroupID struct scan error")
		}
		infos = append(infos, &info)
	}
	return infos, nil
}

// QueryCountByUserID 查询某个组ID下的记录的数量
func (ndb *NodeDistributeBillDB) QueryCountByUserID(ctx context.Context, groupID int) (int, error) {
	res, err := ndb.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `node_distribute_bill` WHERE `user_id`=?", groupID)
	if err != nil {
		logger.Warn("QueryCountByUserID error: ", err)
		return 0, errors.New("QueryCountByUserID error")
	}
	var count int
	if err := res.Scan(&count); err != nil {
		logger.Warn("QueryCountByUserID error: ", err)
		return 0, errors.New("QueryCountByUserID error")
	}
	return count, nil
}

// QueryWithLimitByUserID 分页查询属于某个组的记录
func (ndb *NodeDistributeBillDB) QueryWithLimitByUserID(
	ctx context.Context,
	limit, offset, groupID int,
) ([]*NodeDistributeBill, error) {
	rows, err := ndb.conn.Query(
		ctx,
		"SELECT * FROM `node_distribute_bill` WHERE `user_id`=? LIMIT ?,?",
		groupID,
		limit,
		offset,
	)
	if err != nil {
		logger.Warn("QueryWithLimitByUserID error: ", err)
		return nil, errors.New("QueryWithLimitByUserID error")
	}
	infos := make([]*NodeDistributeBill, 0)
	for rows.Next() {
		var info NodeDistributeBill
		if err := rows.StructScan(&info); err != nil {
			logger.Warn("QueryWithLimitByUserID struct scan error: ", err)
			return nil, errors.New("QueryWithLimitByUserID struct scan error")
		}
		infos = append(infos, &info)
	}
	return infos, nil
}

// NewNodeDistributeBill 新建一个node_distribute_bill数据表操作结构体
func NewNodeDistributeBill(conn *db.DB) *NodeDistributeBillDB {
	return &NodeDistributeBillDB{
		conn: conn,
	}
}
