package db

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
)

type PaperAwardDB struct {
	conn *db.DB
}

// Insert 插入新的记录
func (this *PaperAwardDB) Insert(ctx context.Context, newInfo *PaperApply) (int64, error) {
	res, err := this.conn.Exec(
		ctx,
		"INSERT INTO `paper_apply` "+
			"(`creater_id`, `creater_username`, `creater_name`, `create_time`, `user_group_id`, `tutor_id`, `tutor_username`, `tutor_name`, `paper_title`, `paper_category`, `paper_partition`, `paper_firstpage_img`, `paper_thankspage_img`, `remark_message`) "+
			" VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		newInfo.CreaterID,
		newInfo.CreaterUsername,
		newInfo.CreaterName,
		newInfo.CreateTime,
		newInfo.UserGroupID,
		newInfo.TutorID,
		newInfo.TutorUsername,
		newInfo.TutorName,
		newInfo.PaperTitle,
		newInfo.PaperCategory,
		newInfo.PaperPartition,
		newInfo.PaperFirstPageImageName,
		newInfo.PaperThanksPageImageName,
		newInfo.RemarkMessage,
	)

	if err != nil {
		logger.Warn("Insert paper award apply error: ", err)
		return 0, errors.New("Insert paper award apply error")
	}
	id, err := res.LastInsertId()
	if err != nil {
		logger.Warn("Insert paper award apply error: ", err)
		return 0, errors.New("Insert paper award apply error")
	}
	return id, nil
}

// QueryAllCount 查询所有记录的数量
func (this *PaperAwardDB) QueryAllCount(ctx context.Context) (int, error) {
	row, err := this.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `paper_apply`")
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
func (this *PaperAwardDB) QueryCountByCreaterID(ctx context.Context, createrID int) (int, error) {
	row, err := this.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `paper_apply` WHERE `creater_id`=?", createrID)
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
func (this *PaperAwardDB) QueryCountByGroupID(ctx context.Context, groupID int) (int, error) {
	row, err := this.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `paper_apply` WHERE `user_group_id`=?", groupID)
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
func (this *PaperAwardDB) QueryAllWithLimit(ctx context.Context, limit, offset int) ([]*PaperApply, error) {
	rows, err := this.conn.Query(ctx, "SELECT * FROM `paper_apply` LIMIT ?, ?", limit, offset)
	if err != nil {
		logger.Warn("QueryAllWithLimit error: ", err)
		return nil, errors.New("QueryAllWithLimit error")
	}
	res := make([]*PaperApply, 0)
	for rows.Next() {
		var info PaperApply
		if err := rows.StructScan(&info); err != nil {
			logger.Warn("QueryAllWithLimit struct scan error: ", err)
			return nil, errors.New("QueryAllWithLimit struct scan error")
		}
		res = append(res, &info)
	}
	return res, nil
}

// QueryWithLimitByCreaterID 分页查询某个用户创建的记录
func (this *PaperAwardDB) QueryWithLimitByCreaterID(
	ctx context.Context,
	createrID int,
	limit, offset int,
) ([]*PaperApply, error) {
	rows, err := this.conn.Query(ctx, "SELECT * FROM `paper_apply` WHERE `creater_id`=? LIMIT ?, ?", createrID, limit, offset)
	if err != nil {
		logger.Warn("QueryWithLimitByCreaterID error: ", err)
		return nil, errors.New("QueryWithLimitByCreaterID error")
	}
	res := make([]*PaperApply, 0)
	for rows.Next() {
		var info PaperApply
		if err := rows.StructScan(&info); err != nil {
			logger.Warn("QueryWithLimitByCreaterID struct scan error: ", err)
			return nil, errors.New("QueryWithLimitByCreaterID struct scan error")
		}
		res = append(res, &info)
	}
	return res, nil
}

// QueryWithLimitByGroupID 分页查询某个用户组的记录
func (this *PaperAwardDB) QueryWithLimitByGroupID(ctx context.Context, groupID int, limit, offset int) ([]*PaperApply, error) {
	rows, err := this.conn.Query(ctx, "SELECT * FROM `paper_apply` WHERE `user_group_id`=? LIMIT ?, ?", groupID, limit, offset)
	if err != nil {
		logger.Warn("QueryWithLimitByGroupID error: ", err)
		return nil, errors.New("QueryWithLimitByGroupID error")
	}
	res := make([]*PaperApply, 0)
	for rows.Next() {
		var info PaperApply
		if err := rows.StructScan(&info); err != nil {
			logger.Warn("QueryWithLimitByGroupID struct scan error: ", err)
			return nil, errors.New("QueryWithLimitByGroupID struct scan error")
		}
		res = append(res, &info)
	}
	return res, nil
}

// UpdateCheckStatus 更新记录的审核信息
func (this *PaperAwardDB) UpdateCheckStatus(ctx context.Context, checkInfos *PaperApply) (bool, error) {
	res, err := this.conn.Exec(
		ctx,
		"UPDATE `paper_apply` SET `check_status`=?, `checker_id`=?, `checker_name`=?, `checker_username`=?, `check_money`=?, `check_message`=?, `check_time`=? WHERE `id`=? AND `check_status`=-1",
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
func (this *PaperAwardDB) QueryByID(ctx context.Context, id int) (*PaperApply, error) {
	row, err := this.conn.QueryRow(ctx, "SELECT * FROM `paper_apply` WHERE `id`=?", id)
	if err != nil {
		logger.Warn("QueryByID error: ", err)
		return nil, errors.New("QueryByID error")
	}
	var info PaperApply
	if err := row.StructScan(&info); err != nil {
		logger.Warn("QueryByID struct scan error: ", err)
		return nil, errors.New("QueryByID struct scan error")
	}
	return &info, nil
}

func NewPaperAward(conn *db.DB) *PaperAwardDB {
	return &PaperAwardDB{
		conn: conn,
	}
}
