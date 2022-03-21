package logic

import (
	"context"
	"errors"
	"time"

	"github.com/essayZW/hpcmanager/user/db"
	"gopkg.in/guregu/null.v4"
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
		// FIXME 使用TimeFrom函数生成,不止这一个地方需要修改
		CreateTime: null.NewTime(time.Now(), true),
	})
}

// GetByTutorUsername 通过导师用户名查询导师的其他信息
func (group *UserGroup) GetByTutorUsername(ctx context.Context, username string) (*db.Group, error) {
	return group.userGroupDB.QueryByTutorUsername(ctx, username)
}

// PaginationApplyResult 分页查询申请信息的结果
type PaginationApplyResult struct {
	Applies []*db.UserGroupApply
	Count   int
}

// AdminPageGetApplyInfo 管理员分页查询所有的申请信息
func (group *UserGroup) AdminPageGetApplyInfo(ctx context.Context, pageIndex, pageSize int) (*PaginationApplyResult, error) {
	if pageIndex < 1 {
		return nil, errors.New("pageIndex must large than 0")
	}
	if pageSize <= 0 || pageSize > 200 {
		return nil, errors.New("pageSize must large than 0 and less than 200")
	}
	// 先查询总数
	count, err := group.userGroupApplyDB.AdminLimitQueryApplyCount(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return &PaginationApplyResult{
			Applies: make([]*db.UserGroupApply, 0),
			Count:   0,
		}, nil
	}
	offset := pageSize * (pageIndex - 1)
	applis, err := group.userGroupApplyDB.AdminLimitQueryApplyInfo(ctx, offset, pageSize)
	if err != nil {
		return nil, err
	}
	return &PaginationApplyResult{
		Applies: applis,
		Count:   count,
	}, nil
}

// TutorPageGetApplyInfo 导师分页查看申请本组的所有申请信息
func (group *UserGroup) TutorPageGetApplyInfo(ctx context.Context, pageIndex, pageSize, groupID int) (*PaginationApplyResult, error) {
	if pageIndex < 1 {
		return nil, errors.New("pageIndex must large than 0")
	}
	if pageSize <= 0 || pageSize > 200 {
		return nil, errors.New("pageSize must large than 0 and less than 200")
	}
	count, err := group.userGroupApplyDB.TutorLimitQueryApplyCount(ctx, groupID)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return &PaginationApplyResult{
			Applies: make([]*db.UserGroupApply, 0),
			Count:   0,
		}, nil
	}
	offset := pageSize * (pageIndex - 1)
	applis, err := group.userGroupApplyDB.TutorLimitQueryApplyInfo(ctx, offset, pageSize, groupID)
	if err != nil {
		return nil, err
	}
	return &PaginationApplyResult{
		Applies: applis,
		Count:   count,
	}, nil
}

// CommonPageGetApplyInfo 普通用户分页查询自己创建的所有申请信息
func (group *UserGroup) CommonPageGetApplyInfo(ctx context.Context, pageIndex, pageSize, userID int) (*PaginationApplyResult, error) {
	if pageIndex < 1 {
		return nil, errors.New("pageIndex must large than 0")
	}
	if pageSize <= 0 || pageSize > 200 {
		return nil, errors.New("pageSize must large than 0 and less than 200")
	}
	count, err := group.userGroupApplyDB.CommonLimitQueryApplyCount(ctx, userID)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return &PaginationApplyResult{
			Applies: make([]*db.UserGroupApply, 0),
			Count:   0,
		}, nil
	}
	offset := pageSize * (pageIndex - 1)
	applis, err := group.userGroupApplyDB.CommonLimitQueryApplyInfo(ctx, offset, pageSize, userID)
	if err != nil {
		return nil, err
	}
	return &PaginationApplyResult{
		Applies: applis,
		Count:   count,
	}, nil
}

// GetApplyInfoByID 通过申请ID查询申请记录信息
func (group *UserGroup) GetApplyInfoByID(ctx context.Context, applyID int) (*db.UserGroupApply, error) {
	return group.userGroupApplyDB.QueryByID(ctx, applyID)
}

// TutorCheckApply 导师审核申请记录
func (group *UserGroup) TutorCheckApply(ctx context.Context, tutorID int, applyID int, checkAccept bool, checkMessage string) (bool, error) {
	applyInfo, err := group.GetApplyInfoByID(ctx, applyID)
	if err != nil {
		return false, errors.New("Invalid applyid: Get apply info error")
	}
	if applyInfo.Status != 1 {
		// 当前的申请记录状态不是正常状态，可能已经被撤销
		return false, errors.New("invalid apply status")
	}
	if applyInfo.TutorCheckStatus != -1 {
		// 已经被导师审核
		return false, errors.New("tutor has reviewed this application")
	}
	if applyInfo.TutorID != tutorID {
		// 不是这条记录的导师
		return false, errors.New("must be group tutor")
	}
	var checkStatus int8
	if checkAccept {
		checkStatus = 1
	} else {
		checkStatus = 0
	}
	return group.userGroupApplyDB.UpdateTutorCheckStatus(ctx, &db.UserGroupApply{
		TutorCheckStatus: checkStatus,
		MessageTutor:     null.NewString(checkMessage, true),
		TutorCheckTime:   null.NewTime(time.Now(), true),
		ID:               applyID,
	})
}

// AdminCheckApply 管理员审核申请记录
func (group *UserGroup) AdminCheckApply(ctx context.Context, applyID int, checkerID int, checkerUsername, checkerName string, checkAccept bool, checkMessage string) (bool, error) {
	applyInfo, err := group.GetApplyInfoByID(ctx, applyID)
	if err != nil {
		return false, errors.New("Invalid applyid: Get apply info error")
	}
	if applyInfo.Status != 1 {
		// 当前的申请记录状态不是正常状态，可能已经被撤销
		return false, errors.New("invalid apply status")
	}
	if applyInfo.TutorCheckStatus != 1 {
		// 导师未审核或者审核失败
		return false, errors.New("tutor has not reviewed or the review has not passed")
	}
	if applyInfo.ManagerCheckStatus != -1 {
		// 已经被管理员进行了审核
		return false, errors.New("manager has reviewed this application")
	}
	var checkStatus int8
	if checkAccept {
		checkStatus = 1
	} else {
		checkStatus = 0
	}
	return group.userGroupApplyDB.UpdateAdminCheckStatus(ctx, &db.UserGroupApply{
		ID:                     applyID,
		ManagerCheckStatus:     checkStatus,
		MessageManager:         null.NewString(checkMessage, true),
		ManagerCheckTime:       null.NewTime(time.Now(), true),
		ManagerCheckerID:       null.NewInt(int64(checkerID), true),
		ManagerCheckerUsername: null.NewString(checkerUsername, true),
		ManagerCheckerName:     null.NewString(checkerName, true),
	})
}

// CreateGroup 创建一个新的用户组
func (group *UserGroup) CreateGroup(ctx context.Context, createrInfo, tutorInfo *db.User, name string, hpcGroupID int) (int64, error) {
	return group.userGroupDB.Insert(ctx, &db.Group{
		HpcGroupID:      hpcGroupID,
		Name:            name,
		CreateTime:      time.Now(),
		CreaterID:       createrInfo.ID,
		CreaterUsername: createrInfo.Username,
		CreaterName:     createrInfo.Name,
		TutorID:         tutorInfo.ID,
		TutorUsername:   tutorInfo.Username,
		TutorName:       tutorInfo.Name,
	})

}

// NewUserGroup 创建一个新的用户组的操作逻辑
func NewUserGroup(userGroupDB *db.UserGroupDB, userGroupApplyDB *db.UserGroupApplyDB) *UserGroup {
	return &UserGroup{
		userGroupDB:      userGroupDB,
		userGroupApplyDB: userGroupApplyDB,
	}
}
