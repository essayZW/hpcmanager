package db

import (
	"github.com/jmoiron/sqlx"
)

// UserDB 用户数据库操作
type UserDB struct {
	conn *sqlx.DB
}

// LoginQuery 用于登录的查询,需要用户名和密码
func (db *UserDB) LoginQuery(username, md5password string) (*User, error) {
	row := db.conn.QueryRowx("SELECT * FROM `user` WHERE `username`=? AND `password`=?", username, md5password)
	var res User
	err := row.StructScan(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// New 创建一个新的操作用户数据库结构体
func New(conn *sqlx.DB) *UserDB {
	return &UserDB{
		conn: conn,
	}
}
