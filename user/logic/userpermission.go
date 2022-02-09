package logic

import "github.com/essayZW/hpcmanager/user/db"

// UserPermission 用户权限相关的主要逻辑操作
type UserPermission struct {
	db *db.UserPermissionDB
}

// GetUserPermissionByID 通过用户ID查询用户拥有的权限信息
func (u *UserPermission) GetUserPermissionByID(id int) ([]*db.FullUserPermission, error) {
	return u.db.QueryUserPermissionLevel(id)
}

// NewUserPermission 创建UserPermission结构体指针
func NewUserPermission(db *db.UserPermissionDB) *UserPermission {
	return &UserPermission{
		db: db,
	}
}
