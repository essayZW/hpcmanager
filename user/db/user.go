package db

import (
	"github.com/jmoiron/sqlx"
)

// UserDB 用户数据库操作
type UserDB struct {
	conn *sqlx.DB
}

// LoginQuery 用于登录的查询,需要用户名和密码，返回用户的ID
func (db *UserDB) LoginQuery(username, md5password string) (bool, error) {
	row := db.conn.QueryRowx("SELECT COUNT(*) FROM `user` WHERE `username`=? AND `password`=?", username, md5password)
	var res int
	err := row.Scan(&res)
	if err != nil {
		return false, err
	}
	return res > 0, nil
}

// QueryByUsername 通过用户名查找用户
func (db *UserDB) QueryByUsername(username string) (*User, error) {
	row := db.conn.QueryRowx("SELECT * FROM `user` WHERE `username`=?", username)
	var info User
	err := row.StructScan(&info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

// NewUser 创建一个新的操作用户数据库结构体
func NewUser(conn *sqlx.DB) *UserDB {
	return &UserDB{
		conn: conn,
	}
}
