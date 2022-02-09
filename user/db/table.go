package db

import (
	"time"

	"github.com/essayZW/hpcmanager/db"
)

// User 数据库中的user表的结构体
type User struct {
	ID              int       `db:"id"`
	Username        string    `db:"username"`
	Password        string    `db:"password"`
	Tel             string    `db:"tel"`
	Email           string    `db:"email"`
	Name            string    `db:"name"`
	PinyinName      string    `db:"pinyin_name"`
	CollegeName     string    `db:"college_name"`
	GroupID         int       `db:"group_id"`
	CreateTime      time.Time `db:"create_time"`
	ExtraAttributes *db.JSON  `db:"extraAttributes"`
}

// UserPermission 数据库中user_permission表的结构体
type UserPermission struct {
	ID              int       `db:"id"`
	UserID          int       `db:"user_id"`
	UserGroupID     int       `db:"user_group_id"`
	PermissionID    int       `db:"permission_id"`
	CreateTime      time.Time `db:"create_time"`
	ExtraAttributes *db.JSON  `db:"extraAttributes"`
}

// Permission 数据库中的permission表的结构体
type Permission struct {
	ID              int       `db:"id"`
	Name            string    `db:"name"`
	Level           int8      `db:"level"`
	Description     string    `db:"description"`
	CreateTime      time.Time `db:"create_time"`
	ExtraAttributes *db.JSON  `db:"extraAttributes"`
}

// FullUserPermission 连接user_permission和permission表的完整用户权限信息
type FullUserPermission struct {
	UserPermission
	Level int8 `db:"level"`
}
