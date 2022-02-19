package logic

import (
	"context"
	"errors"
	"time"

	"github.com/essayZW/hpcmanager/user/db"
)

// UserGroup 用户组相关的操作逻辑
type UserGroup struct {
	userGroupDB      *db.UserGroupDB
	userGroupApplyDB *db.UserGroupApplyDB
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

// CreateUserJoinGroupApply 创建用户申请加入组的申请记录
func (group *UserGroup) CreateUserJoinGroupApply(ctx context.Context, userInfo *db.User, applyGroupID int) (int64, error) {
	if userInfo.GroupID != 0 {
		return 0, errors.New("user has a group, can't apply new group")
	}
	// 判断申请的组是否存在
	groupInfo, err := group.GetGroupInfoByID(ctx, applyGroupID)
	if err != nil {
		return 0, errors.New("error applyGroupID")
	}
	// 判断是否已经存在申请记录
	exists := group.userGroupApplyDB.ExistsApply(ctx, userInfo.ID, applyGroupID)
	if exists {
		return 0, errors.New("apply has created, can't create a new application")
	}
	return group.userGroupApplyDB.Insert(ctx, &db.UserGroupApply{
		UserID:        userInfo.ID,
		UserUsername:  userInfo.Username,
		UserName:      userInfo.Name,
		ApplyGroupID:  applyGroupID,
		TutorID:       groupInfo.TutorID,
		TutorUsername: groupInfo.TutorUsername,
		TutorName:     groupInfo.TutorName,
		CreateTime:    time.Now(),
	})
}

// GetByTutorUsername 通过导师用户名查询导师的其他信息
func (group *UserGroup) GetByTutorUsername(ctx context.Context, username string) (*db.Group, error) {
	return group.userGroupDB.QueryByTutorUsername(ctx, username)
}

// NewUserGroup 创建一个新的用户组的操作逻辑
func NewUserGroup(userGroupDB *db.UserGroupDB, userGroupApplyDB *db.UserGroupApplyDB) *UserGroup {
	return &UserGroup{
		userGroupDB:      userGroupDB,
		userGroupApplyDB: userGroupApplyDB,
	}
}
