package db

import (
	"context"

	"github.com/essayZW/hpcmanager/db"
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
		"(`username`, `password`, `tel`, `email`, `name`, `pinyin_name`, `college_name`, `group_id`, `create_time`, `extraAttributes`)"+
		"VALUES (?,?,?,?,?,?,?,?,?,?)", userinfo.Username, userinfo.Password, userinfo.Tel, userinfo.Email, userinfo.Name, userinfo.PinyinName, userinfo.CollegeName, userinfo.GroupID, userinfo.CreateTime, userinfo.ExtraAttributes)
	if err != nil {
		return 0, err
	}
	if res, err := result.RowsAffected(); err == nil {
		return int(res), nil
	}
	return 0, err
}

// NewUser 创建一个新的操作用户数据库结构体
func NewUser(db *db.DB) *UserDB {
	return &UserDB{
		conn: db,
	}
}
