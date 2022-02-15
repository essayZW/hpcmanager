package logic

import (
	"context"
	"time"

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

// Add 添加一条新的权限信息描述
func (p *Permission) Add(ctx context.Context, info *db.Permission) (int, error) {
	if info.CreateTime.IsZero() {
		info.CreateTime = time.Now()
	}
	return p.db.Insert(ctx, info)
}

// NewPermission 创建一个新的permission  logic
func NewPermission(db *db.PermissionDB) *Permission {
	return &Permission{
		db: db,
	}
}
