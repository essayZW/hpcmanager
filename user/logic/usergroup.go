package logic

import (
	"context"

	"github.com/essayZW/hpcmanager/user/db"
)

// UserGroup 用户组相关的操作逻辑
type UserGroup struct {
	userGroupDB *db.UserGroupDB
}

// GetGroupInfoByID 通过ID查询组信息
func (group *UserGroup) GetGroupInfoByID(ctx context.Context, groupID int) (*db.Group, error) {
	return group.userGroupDB.QueryGroupByID(ctx, groupID)
}

// NewUserGroup 创建一个新的用户组的操作逻辑
func NewUserGroup(userGroupDB *db.UserGroupDB) *UserGroup {
	return &UserGroup{
		userGroupDB: userGroupDB,
	}
}
