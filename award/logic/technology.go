package logic

import (
	"context"
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

// NewTechnology 创建Technology
func NewTechnology(technologyAwardApplyDB *db.TechnologyAwardApplyDB) *Technology {
	return &Technology{
		technologyAwardApplyDB: technologyAwardApplyDB,
	}
}
