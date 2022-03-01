package db

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/db"
	"go-micro.dev/v4/logger"
)

// HpcGroupDB 操作hpc_group表
type HpcGroupDB struct {
	conn *db.DB
}

// Insert 插入新的hpc_group表的记录
func (hpc *HpcGroupDB) Insert(ctx context.Context, data *HpcGroup) (int64, error) {
	res, err := hpc.conn.Exec(ctx, "INSERT INTO `hpc_group`"+
		"(`name`, `queue_name`, `gid`, `extraAttributes`) VALUES (?,?,?,?)",
		data.Name, data.QueueName, data.GID, data.ExtraAttributes)
	if err != nil {
		logger.Warn("Insert error: ", err)
		return 0, errors.New("Insert group error")
	}
	id, err := res.LastInsertId()
	if err != nil {
		logger.Warn("Insert error: ", err, " with data: ", data)
		return 0, errors.New("Insert group error")
	}
	return id, nil
}

// NewHpcGroup 创建新的NewHpcGroup结构体并返回指针
func NewHpcGroup(db *db.DB) *HpcGroupDB {
	return &HpcGroupDB{
		conn: db,
	}
}
