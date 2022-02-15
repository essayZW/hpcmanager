package db

import (
	"time"

	"github.com/essayZW/hpcmanager/db"
)

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
