package db

import (
	"context"

	"github.com/essayZW/hpcmanager/db"
	"go-micro.dev/v4/logger"
)

// PermissionDB 权限表的相关操作
type PermissionDB struct {
	conn *db.DB
}

// QueryIDByLevel 查询权限等级对应的权限ID
func (p *PermissionDB) QueryIDByLevel(ctx context.Context, level int32) (int, error) {
	row, err := p.conn.QueryRow(ctx, "SELECT `id` FROM permission WHERE `level`=?", level)
	if err != nil {
		return 0, err
	}
	var id int
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Insert 插入一条新的权限记录
func (p *PermissionDB) Insert(ctx context.Context, info *Permission) (int, error) {
	res, err := p.conn.Exec(
		ctx,
		"INSERT INTO `permission` (`name`, `level`, `description`, `create_time`, `extraAttributes`) VALUES (?,?,?,?,?)",
		info.Name,
		info.Level,
		info.Description,
		info.CreateTime,
		info.ExtraAttributes,
	)
	if err != nil {
		logger.Warn("insert into permission error: %v", err, " with data: ", info)
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		logger.Warn("insert into permission error: %v", err, " with data: ", info)
		return 0, err
	}
	return int(id), nil
}

// NewPermission 新建一个权限表操作结构体
func NewPermission(conn *db.DB) *PermissionDB {
	return &PermissionDB{
		conn: conn,
	}
}
