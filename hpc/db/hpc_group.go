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

// QueryByID 通过ID查询记录
func (hpc *HpcGroupDB) QueryByID(ctx context.Context, id int) (*HpcGroup, error) {
	res, err := hpc.conn.QueryRow(ctx, "SELECT * FROM `hpc_group` WHERE `id`=?", id)
	if err != nil {
		logger.Warn("QueryByID error: ", err)
		return nil, errors.New("QueryByID error")
	}
	var info HpcGroup
	err = res.StructScan(&info)
	if err != nil {
		logger.Warn("QueryByID structScan error: ", err)
		return nil, errors.New("QueryByID structScan error")
	}
	return &info, nil
}

// QueryByName 通过用户组的名查询计算节点用户组的信息
func (hpc *HpcGroupDB) QueryByName(ctx context.Context, name string) (*HpcGroup, error) {
	res, err := hpc.conn.QueryRow(ctx, "SELECT * FROM `hpc_group` WHERE `name`=?", name)
	if err != nil {
		logger.Warn("QueryByName error: ", err)
		return nil, errors.New("QueryByName error")
	}
	var info HpcGroup
	err = res.StructScan(&info)
	if err != nil {
		logger.Warn("QueryByName structScan error: ", err)
		return nil, errors.New("QueryByName structScan error")
	}
	return &info, nil
}

// NewHpcGroup 创建新的NewHpcGroup结构体并返回指针
func NewHpcGroup(db *db.DB) *HpcGroupDB {
	return &HpcGroupDB{
		conn: db,
	}
}
