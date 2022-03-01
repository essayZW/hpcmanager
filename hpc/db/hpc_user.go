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

// NewHpcUser 创建新的NewHpcUser结构体并返回指针
func NewHpcUser(db *db.DB) *HpcUserDB {
	return &HpcUserDB{
		conn: db,
	}
}
