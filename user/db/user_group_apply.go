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

// AdminLimitQueryApplyCount 管理员查询所有的申请总数
func (ugadb *UserGroupApplyDB) AdminLimitQueryApplyCount(ctx context.Context) (int, error) {
	row, err := ugadb.db.QueryRow(ctx, "SELECT COUNT(*) FROM `user_group_apply` WHERE `tutor_check_status`=1")
	if err != nil {
		logger.Warn("AdminLimitQueryApplyCount error: ", err)
		return 0, errors.New("AdminLimitQueryApplyCount query fail")
	}
	var count int
	err = row.Scan(&count)
	if err != nil {
		return 0, errors.New("AdminLimitQueryApplyCount query fail")
	}
	return count, nil
}

// AdminLimitQueryApplyInfo 管理员分页查询所有的申请信息
func (ugadb *UserGroupApplyDB) AdminLimitQueryApplyInfo(ctx context.Context, offset, size int) ([]*UserGroupApply, error) {
	// 查询所有tutor_check_status等于1的记录，代表已经被导师审核通过
	// 只有导师已经审核通过的申请管理员才可以看到
	rows, err := ugadb.db.Query(ctx, "SELECT * FROM `user_group_apply` WHERE `tutor_check_status`=1 LIMIT ?,?", offset, size)
	if err != nil {
		logger.Warn("AdminLimitQueryApplyInfo error: ", err)
		return nil, errors.New("AdminLimitQueryApplyInfo query fail")
	}
	applies := make([]*UserGroupApply, 0)
	for rows.Next() {
		var apply UserGroupApply
		err := rows.StructScan(&apply)
		if err != nil {
			logger.Warn("AdminLimitQueryApplyInfo StructScan error: ", err)
			return nil, errors.New("AdminLimitQueryApplyInfo query fail")
		}
		applies = append(applies, &apply)
	}
	return applies, nil
}

// TutorLimitQueryApplyInfo 导师分页查询所有的申请信息
func (ugadb *UserGroupApplyDB) TutorLimitQueryApplyInfo(ctx context.Context, offset, size, groupID int) ([]*UserGroupApply, error) {
	rows, err := ugadb.db.Query(ctx, "SELECT * FROM `user_group_apply` WHERE `apply_group_id`=? LIMIT ?,?", groupID, offset, size)
	if err != nil {
		logger.Warn("TutorLimitQueryApplyInfo error: ", err)
		return nil, errors.New("TutorLimitQueryApplyInfo query fail")
	}
	applies := make([]*UserGroupApply, 0)
	for rows.Next() {
		var apply UserGroupApply
		err := rows.StructScan(&apply)
		if err != nil {
			logger.Warn("TutorLimitQueryApplyInfo StructScan error: ", err)
			return nil, errors.New("TutorLimitQueryApplyInfo query fail")
		}
		applies = append(applies, &apply)
	}
	return applies, nil
}

// TutorLimitQueryApplyCount 导师分页查询某一个组所有的申请的数量
func (ugadb *UserGroupApplyDB) TutorLimitQueryApplyCount(ctx context.Context, groupID int) (int, error) {
	row, err := ugadb.db.QueryRow(ctx, "SELECT COUNT(*) FROM `user_group_apply` WHERE `apply_group_id`=?", groupID)
	if err != nil {
		logger.Warn("TutorLimitQueryApplyCount error: ", err)
		return 0, errors.New("TutorLimitQueryApplyCount query fail")
	}
	var count int
	err = row.Scan(&count)
	if err != nil {
		return 0, errors.New("TutorLimitQueryApplyCount query fail")
	}
	return count, nil
}

// CommonLimitQueryApplyInfo 普通用户分页查询自己的所有申请的信息
func (ugadb *UserGroupApplyDB) CommonLimitQueryApplyInfo(ctx context.Context, offset, size, userID int) ([]*UserGroupApply, error) {
	rows, err := ugadb.db.Query(ctx, "SELECT * FROM `user_group_apply` WHERE `user_id`=? LIMIT ?,?", userID, offset, size)
	if err != nil {
		logger.Warn("CommonLimitQueryApplyInfo error: ", err)
		return nil, errors.New("CommonLimitQueryApplyInfo query fail")
	}
	applies := make([]*UserGroupApply, 0)
	for rows.Next() {
		var apply UserGroupApply
		err := rows.StructScan(&apply)
		if err != nil {
			logger.Warn("CommonLimitQueryApplyInfo StructScan error: ", err)
			return nil, errors.New("CommonLimitQueryApplyInfo query fail")
		}
		applies = append(applies, &apply)
	}
	return applies, nil
}

// CommonLimitQueryApplyCount 普通用户查询自己创建的申请的数量
func (ugadb *UserGroupApplyDB) CommonLimitQueryApplyCount(ctx context.Context, userID int) (int, error) {
	row, err := ugadb.db.QueryRow(ctx, "SELECT COUNT(*) FROM `user_group_apply` WHERE `user_id`=?", userID)
	if err != nil {
		logger.Warn("CommonLimitQueryApplyCount error: ", err)
		return 0, errors.New("CommonLimitQueryApplyCount query fail")
	}
	var count int
	err = row.Scan(&count)
	if err != nil {
		return 0, errors.New("CommonLimitQueryApplyCount query fail")
	}
	return count, nil
}

// NewUserGroupApply 创建新用户申请表数据库操作结构体
func NewUserGroupApply(db *db.DB) *UserGroupApplyDB {
	return &UserGroupApplyDB{
		db: db,
	}
}
