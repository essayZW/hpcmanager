package db

import (
	"context"
	"errors"
	"time"

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

// QueryCountWithTimeRange 查询一段时间内的所有记录的数量
func (this *NodeWeekUsageBillDB) QueryCountWithTimeRange(ctx context.Context, startTime, endTime time.Time) (int, error) {
	row, err := this.conn.QueryRow(
		ctx,
		"SELECT COUNT(*) FROM `week_usage_bill` WHERE `start_time`>=? AND `end_time`<=?",
		startTime,
		endTime,
	)
	if err != nil {
		logger.Warn("QueryCountWithTimeRange error: ", err)
		return 0, errors.New("QueryCountWithTimeRange error")
	}
	var count int
	if err := row.Scan(&count); err != nil {
		logger.Warn("QueryCountWithTimeRange struct scan error: ", err)
		return 0, errors.New("QueryCountWithTimeRange struct scan error")
	}
	return count, nil
}

// QueryLimitWithTimeRange 分页查询一段时间内的记录
func (this *NodeWeekUsageBillDB) QueryLimitWithTimeRange(
	ctx context.Context,
	limit, offset int,
	startTime, endTime time.Time,
) ([]*NodeWeekUsageBill, error) {
	rows, err := this.conn.Query(
		ctx,
		"SELECT * FROM `week_usage_bill` WHERE `start_time`>=? AND `end_time`<=? LIMIT ?,?",
		startTime,
		endTime,
		limit,
		offset,
	)
	if err != nil {
		logger.Warn("QueryLimitWithTimeRange error: ", err)
		return nil, errors.New("QueryLimitWithTimeRange error")
	}
	res := make([]*NodeWeekUsageBill, 0)
	for rows.Next() {
		var info NodeWeekUsageBill
		if err := rows.StructScan(&info); err != nil {
			logger.Warn("QueryLimitWithTimeRange struct scan error: ", err)
			return nil, errors.New("QueryLimitWithTimeRange struct scan error")
		}
		res = append(res, &info)
	}
	return res, nil
}

// QueryCountWithTimeRangeByUserID 查询一段时间内的某个用户的账单记录的数量
func (this *NodeWeekUsageBillDB) QueryCountWithTimeRangeByUserID(
	ctx context.Context,
	userID int,
	startTime, endTime time.Time,
) (int, error) {
	row, err := this.conn.QueryRow(
		ctx,
		"SELECT COUNT(*) FROM `week_usage_bill` WHERE `user_id`=? AND `start_time`>=? AND `end_time`<=?",
		userID,
		startTime,
		endTime,
	)
	if err != nil {
		logger.Warn("QueryCountWithTimeRangeByUserID error: ", err)
		return 0, errors.New("QueryCountWithTimeRangeByUserID error")
	}
	var count int
	if err := row.Scan(&count); err != nil {
		logger.Warn("QueryCountWithTimeRangeByUserID struct scan error: ", err)
		return 0, errors.New("QueryCountWithTimeRangeByUserID struct scan error")
	}
	return count, nil
}

// QueryLimitWithTimeRangeByUserID 分页查询一段时间范围内的某一个用户的账单
func (this *NodeWeekUsageBillDB) QueryLimitWithTimeRangeByUserID(
	ctx context.Context,
	userID int,
	limit, offset int,
	startTime, endTime time.Time,
) ([]*NodeWeekUsageBill, error) {
	rows, err := this.conn.Query(
		ctx,
		"SELECT * FROM `week_usage_bill` WHERE `user_id`=? AND `start_time`>=? AND `end_time`<=? LIMIT ?, ?",
		userID,
		startTime,
		endTime,
		limit,
		offset,
	)
	if err != nil {
		logger.Warn("QueryLimitWithTimeRangeByUserID error: ", err)
		return nil, errors.New("QueryLimitWithTimeRangeByUserID error")
	}
	res := make([]*NodeWeekUsageBill, 0)
	for rows.Next() {
		var info NodeWeekUsageBill
		if err := rows.StructScan(&info); err != nil {
			logger.Warn("QueryLimitWithTimeRangeByUserID struct scan error: ", err)
			return nil, errors.New("QueryLimitWithTimeRangeByUserID struct scan error")
		}
		res = append(res, &info)
	}
	return res, nil
}

// QueryCountWithTimeRangeByGroupID 查询一段时间内的某个用户的账单记录的数量
func (this *NodeWeekUsageBillDB) QueryCountWithTimeRangeByGroupID(
	ctx context.Context,
	groupID int,
	startTime, endTime time.Time,
) (int, error) {
	row, err := this.conn.QueryRow(
		ctx,
		"SELECT COUNT(*) FROM `week_usage_bill` WHERE `user_group_id`=? AND `start_time`>=? AND `end_time`<=?",
		groupID,
		startTime,
		endTime,
	)
	if err != nil {
		logger.Warn("QueryCountWithTimeRangeByGroupID error: ", err)
		return 0, errors.New("QueryCountWithTimeRangeByGroupID error")
	}
	var count int
	if err := row.Scan(&count); err != nil {
		logger.Warn("QueryCountWithTimeRangeByGroupID struct scan error: ", err)
		return 0, errors.New("QueryCountWithTimeRangeByGroupID struct scan error")
	}
	return count, nil
}

// QueryLimitWithTimeRangeByGroupID 分页查询一段时间范围内的某一个用户的账单
func (this *NodeWeekUsageBillDB) QueryLimitWithTimeRangeByGroupID(
	ctx context.Context,
	groupID int,
	limit, offset int,
	startTime, endTime time.Time,
) ([]*NodeWeekUsageBill, error) {
	rows, err := this.conn.Query(
		ctx,
		"SELECT * FROM `week_usage_bill` WHERE `user_group_id`=? AND `start_time`>=? AND `end_time`<=? LIMIT ?, ?",
		groupID,
		startTime,
		endTime,
		limit,
		offset,
	)
	if err != nil {
		logger.Warn("QueryLimitWithTimeRangeByGroupID error: ", err)
		return nil, errors.New("QueryLimitWithTimeRangeByGroupID error")
	}
	res := make([]*NodeWeekUsageBill, 0)
	for rows.Next() {
		var info NodeWeekUsageBill
		if err := rows.StructScan(&info); err != nil {
			logger.Warn("QueryLimitWithTimeRangeByGroupID struct scan error: ", err)
			return nil, errors.New("QueryLimitWithTimeRangeByGroupID struct scan error")
		}
		res = append(res, &info)
	}
	return res, nil
}

// QueryGroupByGroupIDWithLimit 分页查询按照userGroupID分组的组账单记录,并根据是否支付的状态进行筛选
func (this *NodeWeekUsageBillDB) QueryGroupByGroupIDWithLimit(
	ctx context.Context,
	limit, offset int,
	payFlag int8,
) ([]*NodeWeekUsageBillForUserGroup, error) {
	rows, err := this.conn.Query(
		ctx,
		"SELECT SUM(`wall_time`) AS `wall_time`, SUM(`gwall_time`) AS `gwall_time`, SUM(`fee`) AS `fee`, SUM(`pay_fee`) AS `pay_fee`, `pay_flag`, `user_group_id` "+
			"FROM `week_usage_bill` GROUP BY `user_group_id`, `pay_flag` HAVING `pay_flag`=? ORDER BY `user_group_id` ASC LIMIT ?,?",
		payFlag,
		limit,
		offset,
	)
	if err != nil {
		logger.Warn("QueryGroupByGroupIDWithLimit error: ", err)
		return nil, errors.New("QueryGroupByGroupIDWithLimit error")
	}
	infos := make([]*NodeWeekUsageBillForUserGroup, 0)
	for rows.Next() {
		var info NodeWeekUsageBillForUserGroup
		if err := rows.StructScan(&info); err != nil {
			logger.Warn("QueryGroupByGroupIDWithLimit struct scan error: ", err)
			return nil, errors.New("QueryGroupByGroupIDWithLimit struct scan error")
		}
		infos = append(infos, &info)
	}
	return infos, nil
}

// QueryCountGroupByGroupID 查询某种支付状态下的账单的组的数量
func (this *NodeWeekUsageBillDB) QueryCountGroupByGroupID(ctx context.Context, payFlag int8) (int, error) {
	row, err := this.conn.QueryRow(
		ctx,
		"SELECT COUNT(DISTINCT `user_group_id`) FROM `week_usage_bill` WHERE `pay_flag`=?",
		payFlag,
	)
	if err != nil {
		logger.Warn("QueryCountGroupByGroupID error: ", err)
		return 0, errors.New("QueryCountGroupByGroupID error")
	}
	var count int
	if err := row.Scan(&count); err != nil {
		logger.Warn("QueryCountGroupByGroupID struct scan error: ", err)
		return 0, errors.New("QueryCountGroupByGroupID struct scan error")
	}
	return count, nil
}

// QueryGroupBillByGroupID 通过用户组ID查询某个用户组的账单情况
func (this *NodeWeekUsageBillDB) QueryGroupBillByGroupID(
	ctx context.Context,
	groupID int,
	payFlag int8,
) (*NodeWeekUsageBillForUserGroup, error) {
	row, err := this.conn.QueryRow(
		ctx,
		"SELECT SUM(`wall_time`) AS `wall_time`, SUM(`gwall_time`) AS `gwall_time`, SUM(`fee`) AS `fee`, SUM(`pay_fee`) AS `pay_fee`, `pay_flag`, `user_group_id` "+
			"FROM `week_usage_bill` WHERE `user_group_id`=? AND `pay_flag`=?",
		groupID,
		payFlag,
	)
	if err != nil {
		logger.Warn("QueryGroupBillByGroupID error: ", err)
		return nil, errors.New("QueryGroupBillByGroupID error")
	}
	var info NodeWeekUsageBillForUserGroup
	if err := row.StructScan(&info); err != nil {
		logger.Warn("QueryGroupBillByGroupID struct scan error: ", err)
		return nil, errors.New("QueryGroupBillByGroupID struct scan error")
	}
	return &info, nil
}

// UpdateBillsPayStatusByGroupID 通过用户组ID更新用户组的账单的支付状态
func (this *NodeWeekUsageBillDB) UpdateBillsPayStatusByGroupID(
	ctx context.Context,
	groupID int,
	payMessage string,
	payType int8,
	payTime time.Time,
) (int64, error) {
	res, err := this.conn.Exec(
		ctx,
		"UPDATE `week_usage_bill` SET `pay_fee`=`fee`, `pay_flag`=1, `pay_message`=?, `pay_type`=?, `pay_time`=? WHERE `user_group_id`=? AND `pay_flag`=0",
		payMessage,
		payType,
		payTime,
		groupID,
	)
	if err != nil {
		logger.Warn("UpdateBillsPayStatusByGroupID error: ", err)
		return 0, errors.New("UpdateBillsPayStatusByGroupID error")
	}
	count, err := res.RowsAffected()
	if err != nil {
		logger.Warn("UpdateBillsPayStatusByGroupID update error: ", err)
		return 0, errors.New("UpdateBillsPayStatusByGroupID update error")
	}
	return count, nil
}

// NewNodeWeekUsageBill 创建新的机器节点机时周账单数据库操作映射结构体
func NewNodeWeekUsageBill(conn *db.DB) *NodeWeekUsageBillDB {
	return &NodeWeekUsageBillDB{
		conn: conn,
	}
}
