package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/essayZW/hpcmanager/db"
	"go-micro.dev/v4/logger"
)

// UserGroupDB 用户组数据库操作
type UserGroupDB struct {
	db *db.DB
}

// QueryGroupByID 通过组ID查询组信息
func (group *UserGroupDB) QueryGroupByID(ctx context.Context, groupID int) (*Group, error) {
	row, err := group.db.QueryRow(ctx, "SELECT * FROM `group` WHERE `id`=?", groupID)
	if err != nil {
		return nil, errors.New("group info query error")
	}
	var groupInfo Group
	err = row.StructScan(&groupInfo)
	if err != nil {
		logger.Warn("Group info StructScan error: ", err)
		return nil, errors.New("group info query error")
	}
	return &groupInfo, nil
}

// PaginationQuery 分页查询用户组信息记录
func (group *UserGroupDB) PaginationQuery(
	ctx context.Context,
	offset int,
	size int,
) ([]*Group, error) {
	rows, err := group.db.Query(ctx, "SELECT * FROM `group` LIMIT ?, ?", offset, size)
	if err != nil {
		return nil, errors.New("group infos query error")
	}
	infos := make([]*Group, 0)
	for rows.Next() {
		var groupInfo Group
		err := rows.StructScan(&groupInfo)
		if err != nil {
			logger.Warn("Group info StructScan error: ", err)
			return nil, fmt.Errorf("query info StructScan error: %v", err)
		}
		infos = append(infos, &groupInfo)
	}
	return infos, nil
}

// GetGroupCount 查询所有组的数量
func (group *UserGroupDB) GetGroupCount(ctx context.Context) (int, error) {
	row, err := group.db.QueryRow(ctx, "SELECT COUNT(*) FROM `group`")
	if err != nil {
		logger.Warn("Query Group count error: ", err)
		return 0, errors.New("Query group count error")
	}
	var count int
	err = row.Scan(&count)
	if err != nil {
		logger.Warn("Query Group count error: ", err)
		return 0, errors.New("Query group count error")
	}
	return count, nil
}

// QueryByTutorUsername 通过导师用户名精准查询信息
func (group *UserGroupDB) QueryByTutorUsername(
	ctx context.Context,
	username string,
) (*Group, error) {
	row, err := group.db.QueryRow(ctx, "SELECT * FROM `group` WHERE `tutor_username`=?", username)
	if err != nil {
		logger.Warn("QueryByTutorUsername error: ", err)
		return nil, errors.New("QueryByTutorUsername fail")
	}
	var info Group
	err = row.StructScan(&info)
	if err != nil {
		return nil, errors.New("QueryByTutorUsername fail")
	}
	return &info, nil
}

// Insert 插入新的用户组记录
func (group *UserGroupDB) Insert(ctx context.Context, groupInfo *Group) (int64, error) {
	res, err := group.db.Exec(ctx, "INSERT INTO `group`"+
		"(`hpc_group_id`, `name`, `create_time`, `creater_id`, `creater_username`, `creater_name`, `tutor_id`, `tutor_username`, `tutor_name`, `extraAttributes`)"+
		"VALUES (?,?,?,?,?,?,?,?,?,?)", groupInfo.HpcGroupID, groupInfo.Name, groupInfo.CreateTime, groupInfo.CreaterID,
		groupInfo.CreaterUsername, groupInfo.CreaterName,
		groupInfo.TutorID, groupInfo.TutorUsername, groupInfo.TutorName, groupInfo.ExtraAttributes)
	if err != nil {
		logger.Warn("Insert group error: ", err)
		return 0, errors.New("insert new group row error")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.New("insert new group row error")
	}
	return id, nil
}

// NewUserGroup 创建一个新的用户数据库操作实例
func NewUserGroup(conn *db.DB) *UserGroupDB {
	return &UserGroupDB{
		db: conn,
	}
}
