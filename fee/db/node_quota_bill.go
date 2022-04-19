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

// QueryAllCount 查询所有的账单的数量
func (this *NodeQuotaBillDB) QueryAllCount(ctx context.Context) (int, error) {
	row, err := this.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `node_quota_bill`")
	if err != nil {
		logger.Warn("QueryAllCount error: ", err)
		return 0, errors.New("QueryAllCount error")
	}
	var count int
	if err := row.Scan(&count); err != nil {
		logger.Warn("QueryAllCount scan error: ", err)
		return 0, errors.New("QueryAllCount scan error")
	}
	return count, nil
}

// QueryAllWithLimit 分页查询所有的数据
func (this *NodeQuotaBillDB) QueryAllWithLimit(ctx context.Context, limit, offset int) ([]*NodeQuotaBill, error) {
	rows, err := this.conn.Query(ctx, "SELECT * FROM `node_quota_bill` LIMIT ?,?", limit, offset)
	if err != nil {
		logger.Warn("QueryAllWithLimit error: ", err)
		return nil, errors.New("QueryAllWithLimit error")
	}
	infos := make([]*NodeQuotaBill, 0)
	for rows.Next() {
		var info NodeQuotaBill
		if err := rows.StructScan(&info); err != nil {
			logger.Warn("QueryAllWithLimit struct scan error: ", err)
			return nil, errors.New("QueryAllWithLimit struct scan error")
		}
		infos = append(infos, &info)
	}
	return infos, nil
}

// QueryAllCountByUserID 查询某个用户所有的账单的数量
func (this *NodeQuotaBillDB) QueryAllCountByUserID(ctx context.Context, userID int) (int, error) {
	row, err := this.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `node_quota_bill` WHERE `user_id`=?", userID)
	if err != nil {
		logger.Warn("QueryAllCountByUserID error: ", err)
		return 0, errors.New("QueryAllCountByUserID error")
	}
	var count int
	if err := row.Scan(&count); err != nil {
		logger.Warn("QueryAllCountByUserID scan error: ", err)
		return 0, errors.New("QueryAllCountByUserID scan error")
	}
	return count, nil
}

// QueryAllWithLimitByUserID 分页查询某个用户所有的数据
func (this *NodeQuotaBillDB) QueryAllWithLimitByUserID(
	ctx context.Context,
	limit, offset int,
	userID int,
) ([]*NodeQuotaBill, error) {
	rows, err := this.conn.Query(ctx, "SELECT * FROM `node_quota_bill` WHERE `user_id`=? LIMIT ?,?", userID, limit, offset)
	if err != nil {
		logger.Warn("QueryAllWithLimitByUserID error: ", err)
		return nil, errors.New("QueryAllWithLimitByUserID error")
	}
	infos := make([]*NodeQuotaBill, 0)
	for rows.Next() {
		var info NodeQuotaBill
		if err := rows.StructScan(&info); err != nil {
			logger.Warn("QueryAllWithLimitByUserID struct scan error: ", err)
			return nil, errors.New("QueryAllWithLimitByUserID struct scan error")
		}
		infos = append(infos, &info)
	}
	return infos, nil
}

// QueryAllCountByUserGroupID 查询某个组所有的账单的数量
func (this *NodeQuotaBillDB) QueryAllCountByUserGroupID(ctx context.Context, groupID int) (int, error) {
	row, err := this.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `node_quota_bill` WHERE `user_group_id`=?", groupID)
	if err != nil {
		logger.Warn("QueryAllCountByUserGroupID error: ", err)
		return 0, errors.New("QueryAllCountByUserGroupID error")
	}
	var count int
	if err := row.Scan(&count); err != nil {
		logger.Warn("QueryAllCountByUserGroupID scan error: ", err)
		return 0, errors.New("QueryAllCountByUserGroupID scan error")
	}
	return count, nil
}

// QueryAllWithLimitByUserGroupID 分页查询某个组所有的数据
func (this *NodeQuotaBillDB) QueryAllWithLimitByUserGroupID(
	ctx context.Context,
	limit, offset int,
	groupID int,
) ([]*NodeQuotaBill, error) {
	rows, err := this.conn.Query(
		ctx,
		"SELECT * FROM `node_quota_bill` WHERE `user_group_id`=? LIMIT ?,?",
		groupID,
		limit,
		offset,
	)
	if err != nil {
		logger.Warn("QueryAllWithLimitByUserGroupID error: ", err)
		return nil, errors.New("QueryAllWithLimitByUserGroupID error")
	}
	infos := make([]*NodeQuotaBill, 0)
	for rows.Next() {
		var info NodeQuotaBill
		if err := rows.StructScan(&info); err != nil {
			logger.Warn("QueryAllWithLimitByUserGroupID struct scan error: ", err)
			return nil, errors.New("QueryAllWithLimitByUserGroupID struct scan error")
		}
		infos = append(infos, &info)
	}
	return infos, nil
}

// UpdatePayStatus 更新支付状态
func (this *NodeQuotaBillDB) UpdatePayStatus(ctx context.Context, newBillInfo *NodeQuotaBill) (bool, error) {
	res, err := this.conn.Exec(ctx, "UPDATE `node_quota_bill` SET "+
		" `pay_flag`=1, `pay_fee`=?, `pay_time`=?, `pay_type`=?, `pay_message`=? WHERE `id`=? AND `pay_flag`=0",
		newBillInfo.PayFee, newBillInfo.PayTime, newBillInfo.PayType, newBillInfo.PayMessage, newBillInfo.ID,
	)
	if err != nil {
		logger.Warn("UpdatePayStatus error: ", err)
		return false, errors.New("UpdatePayStatus error")
	}
	count, err := res.RowsAffected()
	if err != nil {
		logger.Warn("UpdatePayStatus error: ", err)
		return false, errors.New("UpdatePayStatus error")
	}
	return count > 0, nil
}

// QueryByID 通过ID查询记录
func (this *NodeQuotaBillDB) QueryByID(ctx context.Context, billID int) (*NodeQuotaBill, error) {
	row, err := this.conn.QueryRow(ctx, "SELECT * FROM `node_quota_bill` WHERE `id`=?", billID)
	if err != nil {
		logger.Warn("QueryByID error: ", err)
		return nil, errors.New("QueryByID error")
	}
	var info NodeQuotaBill
	if err := row.StructScan(&info); err != nil {
		logger.Warn("QueryByID struct scan error: ", err)
		return nil, errors.New("QueryByID struct scan error")
	}
	return &info, nil
}

// NewNodeQuotaBill 创建新的机器存储账单数据库操作
func NewNodeQuotaBill(conn *db.DB) *NodeQuotaBillDB {
	return &NodeQuotaBillDB{
		conn: conn,
	}
}
