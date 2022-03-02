package db

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/db"
	"github.com/jmoiron/sqlx"
	"go-micro.dev/v4/logger"
)

// UserDB 用户数据库操作
type UserDB struct {
	conn *db.DB
}

// LoginQuery 用于登录的查询,需要用户名和密码，返回用户的ID
func (db *UserDB) LoginQuery(ctx context.Context, username, md5password string) (bool, error) {
	row, err := db.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `user` WHERE `username`=? AND `password`=?", username, md5password)
	if err != nil {
		return false, err
	}
	var res int
	err = row.Scan(&res)
	if err != nil {
		return false, err
	}
	return res > 0, nil
}

// QueryByUsername 通过用户名查找用户
func (db *UserDB) QueryByUsername(ctx context.Context, username string) (*User, error) {
	row, err := db.conn.QueryRow(ctx, "SELECT * FROM `user` WHERE `username`=?", username)
	if err != nil {
		return nil, err
	}
	var info User
	err = row.StructScan(&info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

// InsertUser 插入新的用户
func (db *UserDB) InsertUser(ctx context.Context, userinfo *User) (int, error) {
	result, err := db.conn.Exec(ctx, "INSERT INTO `user`"+
		"(`username`, `password`, `tel`, `email`, `name`, `pinyin_name`, `college_name`, `group_id`, `hpc_user_id`, `create_time`, `extraAttributes`)"+
		"VALUES (?,?,?,?,?,?,?,?,?,?,?)",
		userinfo.Username, userinfo.Password, userinfo.Tel, userinfo.Email, userinfo.Name, userinfo.PinyinName,
		userinfo.CollegeName, userinfo.GroupID, userinfo.HpcUserID, userinfo.CreateTime, userinfo.ExtraAttributes)
	if err != nil {
		return 0, err
	}
	if res, err := result.RowsAffected(); err == nil {
		return int(res), nil
	}
	return 0, err
}

// QueryByID 通过ID查询用户表的记录
func (db *UserDB) QueryByID(ctx context.Context, userid int) (*User, error) {
	row, err := db.conn.QueryRow(ctx, "SELECT * FROM `user` WHERE `id`=?", userid)
	if err != nil {
		return nil, err
	}
	var userInfo User
	err = row.StructScan(&userInfo)
	if err != nil {
		logger.Warn("struct scan User error: ", err)
		return nil, err
	}
	return &userInfo, nil
}

// PaginationQuery 分页查询用户信息记录,若groupID为0则查询所有用户
func (db *UserDB) PaginationQuery(ctx context.Context, offset, size, groupID int) ([]*User, error) {
	var rows *sqlx.Rows
	var err error
	if groupID == 0 {
		rows, err = db.conn.Query(ctx, "SELECT * FROM `user` LIMIT ?, ?", offset, size)
	} else {
		rows, err = db.conn.Query(ctx, "SELECT * FROM `user` WHERE `group_id`=? LIMIT ?, ?", groupID, offset, size)
	}
	if err != nil {
		return nil, errors.New("user infos query error")
	}
	infos := make([]*User, 0)
	for rows.Next() {
		var info User
		err := rows.StructScan(&info)
		if err != nil {
			logger.Warn("StructScan User error: ", err)
			return nil, errors.New("struct scan user info error")
		}
		infos = append(infos, &info)
	}
	return infos, nil
}

// QueryAllUserCount 查询所有用户的数量,若groupID为0则查询所有用户
func (db *UserDB) QueryAllUserCount(ctx context.Context, groupID int) (int, error) {
	var row *sqlx.Row
	var err error
	if groupID == 0 {
		row, err = db.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `user`")
	} else {
		row, err = db.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `user` WHERE `group_id`=?", groupID)
	}
	if err != nil {
		logger.Warn("QueryAllUserCount error: ", err)
		return 0, errors.New("Get user count error")
	}
	var count int
	err = row.Scan(&count)
	if err != nil {
		logger.Warn("QueryAllUserCount error: ", err)
		return 0, errors.New("Get user count error")
	}
	return count, nil
}

// UpdateUserGroup 更新用户所属的组信息
func (db *UserDB) UpdateUserGroup(ctx context.Context, userID, groupID int) error {
	res, err := db.conn.Exec(ctx, "UPDATE `user` SET `group_id`=? WHERE `id`=?", groupID, userID)
	if err != nil {
		logger.Warn("UpdateUserGroup error: ", err)
		return errors.New("UpdateUserGroup: update error")
	}
	_, err = res.RowsAffected()
	if err != nil {
		return errors.New("UpdateUserGroup: update error")
	}
	return nil
}

// NewUser 创建一个新的操作用户数据库结构体
func NewUser(db *db.DB) *UserDB {
	return &UserDB{
		conn: db,
	}
}
