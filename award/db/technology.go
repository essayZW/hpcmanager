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

// QueryAllCount 查询所有记录的数量
func (this *TechnologyAwardApplyDB) QueryAllCount(ctx context.Context) (int, error) {
	row, err := this.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `technology_apply`")
	if err != nil {
		logger.Warn("QueryAllCount error: ", err)
		return 0, errors.New("QueryAllCount error")
	}
	var count int
	if err := row.Scan(&count); err != nil {
		logger.Warn("QueryAllCount scan error")
		return 0, errors.New("QueryAllCount scan error")
	}
	return count, nil
}

// QueryCountByCreaterID 查询某个用户创建的所有记录的数量
func (this *TechnologyAwardApplyDB) QueryCountByCreaterID(ctx context.Context, createrID int) (int, error) {
	row, err := this.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `technology_apply` WHERE `creater_id`=?", createrID)
	if err != nil {
		logger.Warn("QueryCountByCreaterID error: ", err)
		return 0, errors.New("QueryCountByCreaterID error")
	}
	var count int
	if err := row.Scan(&count); err != nil {
		logger.Warn("QueryCountByCreaterID scan error")
		return 0, errors.New("QueryCountByCreaterID scan error")
	}
	return count, nil
}

// QueryCountByGroupID 查询某个用户组创建的所有记录的数量
func (this *TechnologyAwardApplyDB) QueryCountByGroupID(ctx context.Context, groupID int) (int, error) {
	row, err := this.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `technology_apply` WHERE `user_group_id`=?", groupID)
	if err != nil {
		logger.Warn("QueryCountByGroupID error: ", err)
		return 0, errors.New("QueryCountByGroupID error")
	}
	var count int
	if err := row.Scan(&count); err != nil {
		logger.Warn("QueryCountByGroupID scan error")
		return 0, errors.New("QueryCountByGroupID scan error")
	}
	return count, nil
}

// QueryAllWithLimit 分页查询所有范围内的记录
func (this *TechnologyAwardApplyDB) QueryAllWithLimit(ctx context.Context, limit, offset int) ([]*TechnologyApply, error) {
	rows, err := this.conn.Query(ctx, "SELECT * FROM `technology_apply` LIMIT ?, ?", limit, offset)
	if err != nil {
		logger.Warn("QueryAllWithLimit error: ", err)
		return nil, errors.New("QueryAllWithLimit error")
	}
	res := make([]*TechnologyApply, 0)
	for rows.Next() {
		var info TechnologyApply
		if err := rows.StructScan(&info); err != nil {
			logger.Warn("QueryAllWithLimit struct scan error: ", err)
			return nil, errors.New("QueryAllWithLimit struct scan error")
		}
		res = append(res, &info)
	}
	return res, nil
}

// QueryWithLimitByCreaterID 分页查询某个用户创建的记录
func (this *TechnologyAwardApplyDB) QueryWithLimitByCreaterID(
	ctx context.Context,
	createrID int,
	limit, offset int,
) ([]*TechnologyApply, error) {
	rows, err := this.conn.Query(
		ctx,
		"SELECT * FROM `technology_apply` WHERE `creater_id`=? LIMIT ?, ?",
		createrID,
		limit,
		offset,
	)
	if err != nil {
		logger.Warn("QueryWithLimitByCreaterID error: ", err)
		return nil, errors.New("QueryWithLimitByCreaterID error")
	}
	res := make([]*TechnologyApply, 0)
	for rows.Next() {
		var info TechnologyApply
		if err := rows.StructScan(&info); err != nil {
			logger.Warn("QueryWithLimitByCreaterID struct scan error: ", err)
			return nil, errors.New("QueryWithLimitByCreaterID struct scan error")
		}
		res = append(res, &info)
	}
	return res, nil
}

// QueryWithLimitByGroupID 分页查询某个用户组的记录
func (this *TechnologyAwardApplyDB) QueryWithLimitByGroupID(
	ctx context.Context,
	groupID int,
	limit, offset int,
) ([]*TechnologyApply, error) {
	rows, err := this.conn.Query(
		ctx,
		"SELECT * FROM `technology_apply` WHERE `user_group_id`=? LIMIT ?, ?",
		groupID,
		limit,
		offset,
	)
	if err != nil {
		logger.Warn("QueryWithLimitByGroupID error: ", err)
		return nil, errors.New("QueryWithLimitByGroupID error")
	}
	res := make([]*TechnologyApply, 0)
	for rows.Next() {
		var info TechnologyApply
		if err := rows.StructScan(&info); err != nil {
			logger.Warn("QueryWithLimitByGroupID struct scan error: ", err)
			return nil, errors.New("QueryWithLimitByGroupID struct scan error")
		}
		res = append(res, &info)
	}
	return res, nil
}

// UpdateCheckStatus 更新记录的审核信息
func (this *TechnologyAwardApplyDB) UpdateCheckStatus(ctx context.Context, checkInfos *TechnologyApply) (bool, error) {
	res, err := this.conn.Exec(
		ctx,
		"UPDATE `technology_apply` SET `check_status`=?, `checker_id`=?, `checker_name`=?, `checker_username`=?, `check_money`=?, `check_message`=?, `check_time`=? WHERE `id`=? AND `check_status`=-1",
		checkInfos.CheckStatus,
		checkInfos.CheckerID,
		checkInfos.CheckerName,
		checkInfos.CheckerUsername,
		checkInfos.CheckMoney,
		checkInfos.CheckMessage,
		checkInfos.CheckTime,
		checkInfos.ID,
	)
	if err != nil {
		logger.Warn("UpdateCheckStatus error: ", err)
		return false, errors.New("UpdateCheckStatus error")
	}
	count, err := res.RowsAffected()
	if err != nil {
		logger.Warn("UpdateCheckStatus error: ", err)
		return false, errors.New("UpdateCheckStatus error")
	}
	return count > 0, nil
}

// QueryByID 通过ID查询记录信息
func (this *TechnologyAwardApplyDB) QueryByID(ctx context.Context, id int) (*TechnologyApply, error) {
	row, err := this.conn.QueryRow(ctx, "SELECT * FROM `technology_apply` WHERE `id`=?", id)
	if err != nil {
		logger.Warn("QueryByID error: ", err)
		return nil, errors.New("QueryByID error")
	}
	var info TechnologyApply
	if err := row.StructScan(&info); err != nil {
		logger.Warn("QueryByID struct scan error: ", err)
		return nil, errors.New("QueryByID struct scan error")
	}
	return &info, nil
}

// NewTechnologyAwardApply 创建TechnologyAwardApplyDB
func NewTechnologyAwardApply(conn *db.DB) *TechnologyAwardApplyDB {
	return &TechnologyAwardApplyDB{
		conn: conn,
	}
}
