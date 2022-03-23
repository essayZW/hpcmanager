package db

import (
	"time"

	"github.com/essayZW/hpcmanager/db"
)

// Project 数据库的project表的映射
type Project struct {
	ID              int       `db:"id"`
	Name            string    `db:"name"`
	From            string    `db:"from"`
	Numbering       string    `db:"numbering"`
	Expenses        string    `db:"expenses"`
	Description     string    `db:"description"`
	CreaterUserID   int       `db:"creater_user_id"`
	CreaterUsername string    `db:"creater_username"`
	CreaterUserName string    `db:"creater_user_name"`
	CreateTime      time.Time `db:"create_time"`
	ModifyUserID    int       `db:"modify_user_id"`
	ModifyUsername  string    `db:"modify_username"`
	ModifyUserName  string    `db:"modify_user_name"`
	ModifyTime      time.Time `db:"modify_time"`
	ExtraAttributes *db.JSON  `db:"extraAttributes"`
}
