package db

import (
	"context"
	"errors"

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

// NewUserGroup 创建一个新的用户数据库操作实例
func NewUserGroup(conn *db.DB) *UserGroupDB {
	return &UserGroupDB{
		db: conn,
	}
}
