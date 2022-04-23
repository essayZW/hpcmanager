package logic

import (
	"context"
	"errors"
	"time"

	"github.com/essayZW/hpcmanager/award/db"
	"gopkg.in/guregu/null.v4"
)

// Technology 科技奖励申请记录操作逻辑
type Technology struct {
	technologyAwardApplyDB *db.TechnologyAwardApplyDB
}

// Create 创建新的科技奖励申请记录
func (this *Technology) Create(
	ctx context.Context,
	createrInfo *UserInfoParam,
	tutorInfo *UserInfoParam,
	userGroupID int,
	technology *TechnologyParam,
	projectInfo *ProjectInfoParam,
) (int64, error) {
	return this.technologyAwardApplyDB.Insert(ctx, &db.TechnologyApply{
		CreaterID:          createrInfo.ID,
		CreaterUsername:    createrInfo.Username,
		CreaterName:        createrInfo.Name,
		UserGroupID:        userGroupID,
		TutorID:            tutorInfo.ID,
		TutorUsername:      tutorInfo.Username,
		TutorName:          tutorInfo.Name,
		ProjectID:          projectInfo.ID,
		ProjectName:        projectInfo.Name,
		ProjectDescription: null.StringFrom(projectInfo.Description),
		PrizeLevel:         technology.Level,
		PrizeImageName:     technology.ImageName,
		RemarkMessage:      technology.RemarkMessage,
		CreateTime:         time.Now(),
	})
}

// PaginationGetTechnologyApplyResult 分页查询科技奖励申请的结果
type PaginationGetTechnologyApplyResult struct {
	Count int
	Data  []*db.TechnologyApply
}

// PaginationGetAll 分页查询所有范围内的记录
func (this *Technology) PaginationGetAll(
	ctx context.Context,
	pageIndex, pageSize int,
) (*PaginationGetTechnologyApplyResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}
	count, err := this.technologyAwardApplyDB.QueryAllCount(ctx)
	if err != nil {
		return nil, err
	}
	limit := pageSize * (pageIndex - 1)
	data, err := this.technologyAwardApplyDB.QueryAllWithLimit(ctx, limit, pageSize)
	return &PaginationGetTechnologyApplyResult{
		Count: count,
		Data:  data,
	}, nil
}

// PaginationGetByCreaterID 分页查询某个用户创建的科技申请记录
func (this *Technology) PaginationGetByCreaterID(
	ctx context.Context,
	createrID int,
	pageIndex, pageSize int,
) (*PaginationGetTechnologyApplyResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}
	count, err := this.technologyAwardApplyDB.QueryCountByCreaterID(ctx, createrID)
	if err != nil {
		return nil, err
	}
	limit := pageSize * (pageIndex - 1)
	data, err := this.technologyAwardApplyDB.QueryWithLimitByCreaterID(ctx, createrID, limit, pageSize)
	return &PaginationGetTechnologyApplyResult{
		Count: count,
		Data:  data,
	}, nil
}

// PaginationGetByGroupID 分页查询某个用户组创建的科技奖励申请
func (this *Technology) PaginationGetByGroupID(
	ctx context.Context,
	groupID int,
	pageIndex, pageSize int,
) (*PaginationGetTechnologyApplyResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}
	count, err := this.technologyAwardApplyDB.QueryCountByGroupID(ctx, groupID)
	if err != nil {
		return nil, err
	}
	limit := pageSize * (pageIndex - 1)
	data, err := this.technologyAwardApplyDB.QueryWithLimitByGroupID(ctx, groupID, limit, pageSize)
	return &PaginationGetTechnologyApplyResult{
		Count: count,
		Data:  data,
	}, nil
}

// NewTechnology 创建Technology
func NewTechnology(technologyAwardApplyDB *db.TechnologyAwardApplyDB) *Technology {
	return &Technology{
		technologyAwardApplyDB: technologyAwardApplyDB,
	}
}
