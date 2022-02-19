package db

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/db"
	"go-micro.dev/v4/logger"
)

// UserGroupApplyDB 新用户申请表数据库操作
type UserGroupApplyDB struct {
	db *db.DB
}

// Insert 插入一条新的申请记录
func (ugadb *UserGroupApplyDB) Insert(ctx context.Context, apply *UserGroupApply) (int64, error) {
	res, err := ugadb.db.Exec(ctx, "INSERT INTO `user_group_apply`"+
		"(`user_id`, `user_username`, `user_name`, `apply_group_id`, `tutor_id`, `tutor_username`, `tutor_name`, `create_time`, `extraAttributes`)"+
		"VALUES (?,?,?,?,?,?,?,?,?)", apply.UserID, apply.UserUsername, apply.UserName, apply.ApplyGroupID, apply.TutorID, apply.TutorUsername, apply.TutorName, apply.CreateTime, apply.ExtraAttributes)
	if err != nil {
		logger.Warn("Insert into database error: ", err)
		return 0, errors.New("Insert into database error")
	}
	id, err := res.LastInsertId()
	if err != nil {
		logger.Warn("Get insert id: ", err)
		return 0, errors.New("Get insert id error")
	}
	return id, nil
}

// ExistsApply 判断是否已经存在某个用户申请某个组的未处理的申请记录
func (ugadb *UserGroupApplyDB) ExistsApply(ctx context.Context, userID int, applyGroupID int) bool {
	// 查询通过userID发起的applyGroupID的申请，且status状态正常即未撤销以及管理员还未审核的记录
	// 管理员是审核的最后一环，因此管理员未审核代表该记录还在处理中
	row, err := ugadb.db.QueryRow(ctx, "SELECT COUNT(*) FROM `user_group_apply`"+
		"WHERE `user_id`=? AND `apply_group_id`=? AND `status`=1 AND `manager_check_status`=-1", userID, applyGroupID)
	if err != nil {
		logger.Warn("ExistsApply error: ", err)
		return false
	}
	var count int
	err = row.Scan(&count)
	if err != nil {
		logger.Warn("ExistsApply error: ", err)
		return false
	}
	return count > 0
}

// NewUserGroupApply 创建新用户申请表数据库操作结构体
func NewUserGroupApply(db *db.DB) *UserGroupApplyDB {
	return &UserGroupApplyDB{
		db: db,
	}
}
