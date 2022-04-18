package db

import (
	"context"
	"errors"
	"time"

	"github.com/essayZW/hpcmanager/db"
	"go-micro.dev/v4/logger"
)

// HpcUserDB 操作hpc_user表
type HpcUserDB struct {
	conn *db.DB
}

// Insert 插入新的记录到hpc_user表中
func (hpcdb *HpcUserDB) Insert(ctx context.Context, data *HpcUser) (int64, error) {
	res, err := hpcdb.conn.Exec(ctx, "INSERT INTO `hpc_user`"+
		"(`node_username`, `node_uid`, `extraAttributes`)"+
		"VALUES (?,?,?)", data.NodeUsername, data.NodeUID, data.ExtraAttributes)
	if err != nil {
		logger.Warn("HpcUserDB Insert error: ", err, " with data: ", data)
		return 0, errors.New("Insert hpc user error")
	}
	id, err := res.LastInsertId()
	if err != nil {
		logger.Warn("HpcUserDB Insert error: ", err, " with data: ", data)
		return 0, errors.New("Insert hpc user error")
	}
	return id, nil
}

// QueryByID 通过ID查询hpc用户表的记录
func (hpcdb *HpcUserDB) QueryByID(ctx context.Context, id int) (*HpcUser, error) {
	res, err := hpcdb.conn.QueryRow(ctx, "SELECT * FROM `hpc_user` WHERE `id`=?", id)
	if err != nil {
		logger.Warn("HpcUserDB query by id error: ", err, " with id: ", id)
		return nil, errors.New("Query error")
	}
	var info HpcUser
	err = res.StructScan(&info)
	if err != nil {
		logger.Warn("HpcUserDB query by id structscan error: ", err)
		return nil, errors.New("Query error")
	}
	return &info, nil
}

// QueryByUsername 通过用户名查询hpc用户表的记录
func (hpcdb *HpcUserDB) QueryByUsername(ctx context.Context, username string) (*HpcUser, error) {
	res, err := hpcdb.conn.QueryRow(ctx, "SELECT * FROM `hpc_user` WHERE `node_username`=?", username)
	if err != nil {
		logger.Warn("HpcUserDB query by username error: ", err, " with username: ", username)
		return nil, errors.New("Query error")
	}
	var info HpcUser
	err = res.StructScan(&info)
	if err != nil {
		logger.Warn("HpcUserDB query by username structscan error: ", err)
		return nil, errors.New("Query error")
	}
	return &info, nil
}

// UpdateQuotaEndTime 更新用户存储的结束使用时间
func (hpcdb *HpcUserDB) UpdateQuotaEndTime(ctx context.Context, hpcUserID int, endTime time.Time) (bool, error) {
	res, err := hpcdb.conn.Exec(
		ctx,
		"UPDATE `hpc_user` SET `quota_end_time`=? WHERE `id`=? AND `quota_start_time`<?",
		endTime,
		hpcUserID,
		endTime,
	)
	if err != nil {
		logger.Warn("UpdateQuotaEndTime error: ", err)
		return false, errors.New("UpdateQuotaEndTime error")
	}
	count, err := res.RowsAffected()
	if err != nil {
		logger.Warn("UpdateQuotaEndTime error: ", err)
		return false, errors.New("UpdateQuotaEndTime error")
	}
	return count > 0, nil
}

// UpdateQuotaStartTime 更新用户存储的开始使用时间
func (hpcdb *HpcUserDB) UpdateQuotaStartTime(ctx context.Context, hpcUserID int, endTime time.Time) (bool, error) {
	res, err := hpcdb.conn.Exec(
		ctx,
		"UPDATE `hpc_user` SET `quota_start_time`=? WHERE `id`=?",
		endTime,
		hpcUserID,
	)
	if err != nil {
		logger.Warn("UpdateQuotaStartTime error: ", err)
		return false, errors.New("UpdateQuotaStartTime error")
	}
	count, err := res.RowsAffected()
	if err != nil {
		logger.Warn("UpdateQuotaStartTime error: ", err)
		return false, errors.New("UpdateQuotaStartTime error")
	}
	return count > 0, nil
}

// NewHpcUser 创建新的NewHpcUser结构体并返回指针
func NewHpcUser(db *db.DB) *HpcUserDB {
	return &HpcUserDB{
		conn: db,
	}
}
