package logic

import (
	"context"

	"github.com/essayZW/hpcmanager/permission/db"
	"github.com/essayZW/hpcmanager/verify"
)

// Permission logic实现
type Permission struct {
	db *db.PermissionDB
}

// GetIDByLevel 通过权限等级查询权限ID
func (p *Permission) GetIDByLevel(ctx context.Context, level verify.Level) (int, error) {
	return p.db.QueryIDByLevel(ctx, int32(level))
}

// NewPermission 创建一个新的permission  logic
func NewPermission(db *db.PermissionDB) *Permission {
	return &Permission{
		db: db,
	}
}
