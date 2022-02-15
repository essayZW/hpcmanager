package db

import (
	"context"

	"github.com/essayZW/hpcmanager/db"
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

// NewPermission 新建一个权限表操作结构体
func NewPermission(conn *db.DB) *PermissionDB {
	return &PermissionDB{
		conn: conn,
	}
}
