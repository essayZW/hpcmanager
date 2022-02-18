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
func (group *UserGroupDB) PaginationQuery(ctx context.Context, offset int, size int) ([]*Group, error) {
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
	row, err := group.db.QueryRow(ctx, "SELECT COUNT(*) FROM `user`")
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

// NewUserGroup 创建一个新的用户数据库操作实例
func NewUserGroup(conn *db.DB) *UserGroupDB {
	return &UserGroupDB{
		db: conn,
	}
}
