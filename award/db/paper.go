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

func NewPaperAward(conn *db.DB) *PaperAwardDB {
	return &PaperAwardDB{
		conn: conn,
	}
}
