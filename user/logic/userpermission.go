package logic

import (
	"context"
	"time"

	"github.com/essayZW/hpcmanager/user/db"
	"github.com/essayZW/hpcmanager/verify"
)

// UserPermission 用户权限相关的主要逻辑操作
type UserPermission struct {
	db  *db.UserPermissionDB
	pdb *db.PermissionDB
}

// GetUserPermissionByID 通过用户ID查询用户拥有的权限信息
func (u *UserPermission) GetUserPermissionByID(ctx context.Context, id int) ([]*db.FullUserPermission, error) {
	return u.db.QueryUserPermissionLevel(ctx, id)
}

// AddUserPermission 添加用户权限
func (u *UserPermission) AddUserPermission(ctx context.Context, info *db.UserPermission, level verify.Level) error {
	// 查询权限level对应的权限ID
	id, err := u.pdb.QueryIDByLevel(ctx, int32(level))
	if err != nil {
		return err
	}
	if info.CreateTime.IsZero() {
		info.CreateTime = time.Now()
	}
	info.PermissionID = id
	err = u.db.Insert(ctx, info)
	if err != nil {
		return err
	}
	return nil
}

// NewUserPermission 创建UserPermission结构体指针
func NewUserPermission(db *db.UserPermissionDB, pdb *db.PermissionDB) *UserPermission {
	return &UserPermission{
		db:  db,
		pdb: pdb,
	}
}
