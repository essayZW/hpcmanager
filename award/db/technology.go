package db

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
)

// TechnologyAwardApplyDB 科技奖励申请表的数据库操作
type TechnologyAwardApplyDB struct {
	conn *db.DB
}

// Insert 插入新的记录
func (this *TechnologyAwardApplyDB) Insert(ctx context.Context, newInfo *TechnologyApply) (int64, error) {
	res, err := this.conn.Exec(ctx,
		"INSERT INTO `technology_apply` "+
			"(`creater_id`, `creater_username`, `creater_name`, `user_group_id`, `tutor_id`, `tutor_username`, `tutor_name`, `project_id`, `project_name`, `project_description`, `prize_level`, `prize_img`, `remark_message`, `create_time`, `extraAttributes`) "+
			" VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		newInfo.CreaterID,
		newInfo.CreaterUsername,
		newInfo.CreaterName,
		newInfo.UserGroupID,
		newInfo.TutorID,
		newInfo.TutorUsername,
		newInfo.TutorName,
		newInfo.ProjectID,
		newInfo.ProjectName,
		newInfo.ProjectDescription,
		newInfo.PrizeLevel,
		newInfo.PrizeImageName,
		newInfo.RemarkMessage,
		newInfo.CreateTime,
		newInfo.ExtraAttributes,
	)
	if err != nil {
		logger.Warn("Insert technology apply error: ", err)
		return 0, errors.New("Insert technology apply error")
	}
	id, err := res.LastInsertId()
	if err != nil {
		logger.Warn("Insert technology apply error: ", err)
		return 0, errors.New("Insert technology apply error")
	}
	return id, nil
}

// NewTechnologyAwardApply 创建TechnologyAwardApplyDB
func NewTechnologyAwardApply(conn *db.DB) *TechnologyAwardApplyDB {
	return &TechnologyAwardApplyDB{
		conn: conn,
	}
}
