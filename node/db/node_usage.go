package db

import (
	"context"
	"errors"
	"time"

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

// QueryAllCount 查询一段时间内所有记录的数量
func (n *NodeUsageTimeDB) QueryAllCount(ctx context.Context, startDate, endDate time.Time) (int, error) {
	row, err := n.conn.QueryRow(
		ctx,
		"SELECT COUNT(*) FROM `hpc_usagetime` WHERE `start_time`>=? AND `end_time`<=?",
		startDate,
		endDate,
	)
	if err != nil {
		logger.Warn("QueryAllCount error: ", err)
		return 0, errors.New("QueryAllCount error")
	}
	var count int
	if err := row.Scan(&count); err != nil {
		logger.Warn("QueryAllCount error: ", err)
		return 0, errors.New("QueryAllCount error")
	}
	return count, nil
}

// QueryAllWithLimit 分页查询一段时间内所有的记录
func (n *NodeUsageTimeDB) QueryAllWithLimit(
	ctx context.Context,
	limit, offset int,
	startDate, endDate time.Time,
) ([]*HpcUsageTime, error) {
	rows, err := n.conn.Query(
		ctx,
		"SELECT * FROM `hpc_usagetime` WHERE `start_time`>=? AND `end_time`<=? LIMIT ?,?",
		startDate,
		endDate,
		limit,
		offset,
	)
	if err != nil {
		logger.Warn("QueryAllWithLimit error: ", err)
		return nil, errors.New("QueryAllWithLimit error")
	}
	infos := make([]*HpcUsageTime, 0)
	for rows.Next() {
		var info HpcUsageTime
		if err := rows.StructScan(&info); err != nil {
			logger.Warn("QueryAllWithLimit strut scan error: ", err)
			return nil, errors.New("QueryAllWithLimit strut scan error")
		}
		infos = append(infos, &info)
	}
	return infos, nil
}

// QueryCountByUserID 通过用户ID查询记录的数量
func (n *NodeUsageTimeDB) QueryCountByUserID(ctx context.Context, id int, startDate, endDate time.Time) (int, error) {
	row, err := n.conn.QueryRow(
		ctx,
		"SELECT COUNT(*) FROM `hpc_usagetime` WHERE `start_time`>=? AND `end_time`<=? AND `user_id`=?",
		startDate,
		endDate,
		id,
	)
	if err != nil {
		logger.Warn("QueryCountByUserID error: ", err)
		return 0, errors.New("QueryCountByUserID error")
	}
	var count int
	if err := row.Scan(&count); err != nil {
		logger.Warn("QueryCountByUserID error: ", err)
		return 0, errors.New("QueryCountByUserID error")
	}
	return count, nil
}

// QueryWithLimitByUserID 通过用户ID分页查询记录信息
func (n *NodeUsageTimeDB) QueryWithLimitByUserID(
	ctx context.Context,
	id, limit, offset int,
	startDate, endDate time.Time,
) ([]*HpcUsageTime, error) {
	rows, err := n.conn.Query(
		ctx,
		"SELECT * FROM `hpc_usagetime` WHERE `start_time`>=? AND `end_time`<=? AND `user_id`=? LIMIT ?,?",
		startDate,
		endDate,
		id,
		limit,
		offset,
	)
	if err != nil {
		logger.Warn("QueryWithLimitByUserID error: ", err)
		return nil, errors.New("QueryWithLimitByUserID error")
	}
	infos := make([]*HpcUsageTime, 0)
	for rows.Next() {
		var info HpcUsageTime
		if err := rows.StructScan(&info); err != nil {
			logger.Warn("QueryWithLimitByUserID strut scan error: ", err)
			return nil, errors.New("QueryWithLimitByUserID strut scan error")
		}
		infos = append(infos, &info)
	}
	return infos, nil
}

// QueryCountByTutorID 通过导师ID查询所有记录的信息
func (n *NodeUsageTimeDB) QueryCountByTutorID(ctx context.Context, id int, startDate, endDate time.Time) (int, error) {
	row, err := n.conn.QueryRow(
		ctx,
		"SELECT COUNT(*) FROM `hpc_usagetime` WHERE `start_time`>=? AND `end_time`<=? AND `tutor_id`=?",
		startDate,
		endDate,
		id,
	)
	if err != nil {
		logger.Warn("QueryCountByTutorID error: ", err)
		return 0, errors.New("QueryCountByTutorID error")
	}
	var count int
	if err := row.Scan(&count); err != nil {
		logger.Warn("QueryCountByTutorID error: ", err)
		return 0, errors.New("QueryCountByTutorID error")
	}
	return count, nil
}

func (n *NodeUsageTimeDB) QueryWithLimitByTutorID(
	ctx context.Context,
	tutorID, limit, offset int,
	startDate, endDate time.Time,
) ([]*HpcUsageTime, error) {
	rows, err := n.conn.Query(
		ctx,
		"SELECT * FROM `hpc_usagetime` WHERE `start_time`>=? AND `end_time`<=? AND `tutor_id`=? LIMIT ?,?",
		startDate,
		endDate,
		tutorID,
		limit,
		offset,
	)
	if err != nil {
		logger.Warn("QueryWithLimitByTutorID error: ", err)
		return nil, errors.New("QueryWithLimitByTutorID error")
	}
	infos := make([]*HpcUsageTime, 0)
	for rows.Next() {
		var info HpcUsageTime
		if err := rows.StructScan(&info); err != nil {
			logger.Warn("QueryWithLimitByTutorID strut scan error: ", err)
			return nil, errors.New("QueryWithLimitByTutorID strut scan error")
		}
		infos = append(infos, &info)
	}
	return infos, nil
}

// QueryByID 通过ID查询记录信息
func (n *NodeUsageTimeDB) QueryByID(ctx context.Context, id int) (*HpcUsageTime, error) {
	row, err := n.conn.QueryRow(ctx, "SELECT * FROM `hpc_usagetime` WHERE id=?", id)
	if err != nil {
		logger.Warn("QueryByID error: ", err)
		return nil, errors.New("QueryByID error")
	}
	var info HpcUsageTime
	if err := row.StructScan(&info); err != nil {
		logger.Warn("QueryByID struct scan error: ", err)
		return nil, errors.New("QueryByID struct scan error")
	}
	return &info, nil
}

// NewNodeUsageTime 创建新的操作计算节点记录的结构体
func NewNodeUsageTime(conn *db.DB) *NodeUsageTimeDB {
	return &NodeUsageTimeDB{
		conn: conn,
	}
}
