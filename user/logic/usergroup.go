package logic

import (
	"context"
	"errors"

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

// PaginationGroupResult 分页查询用户组信息的结果
type PaginationGroupResult struct {
	Infos []*db.Group
	Count int
}

// PaginationGetGroupInfo 分页查询用户组信息
func (group *UserGroup) PaginationGetGroupInfo(ctx context.Context, pageIndex int, pageSize int) (*PaginationGroupResult, error) {
	if pageIndex < 1 {
		return nil, errors.New("pageIndex must large than 0")
	}
	if pageSize <= 0 || pageSize > 200 {
		return nil, errors.New("pageSize must large than 0 and less than 200")
	}
	// 首先查询数据数量
	count, err := group.userGroupDB.GetGroupCount(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return &PaginationGroupResult{
			Infos: make([]*db.Group, 0),
			Count: 0,
		}, nil
	}
	offset := pageSize * (pageIndex - 1)
	infos, err := group.userGroupDB.PaginationQuery(ctx, offset, pageSize)
	if err != nil {
		return nil, err
	}
	return &PaginationGroupResult{
		Infos: infos,
		Count: count,
	}, nil
}

// NewUserGroup 创建一个新的用户组的操作逻辑
func NewUserGroup(userGroupDB *db.UserGroupDB) *UserGroup {
	return &UserGroup{
		userGroupDB: userGroupDB,
	}
}
