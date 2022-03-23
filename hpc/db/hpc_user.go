package db

import (
	"context"
	"errors"

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

// NewHpcUser 创建新的NewHpcUser结构体并返回指针
func NewHpcUser(db *db.DB) *HpcUserDB {
	return &HpcUserDB{
		conn: db,
	}
}
