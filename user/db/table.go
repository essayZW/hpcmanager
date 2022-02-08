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
	GroupID         string    `db:"group_id"`
	CreateTime      time.Time `db:"create_time"`
	ExtraAttributes *db.JSON  `db:"extraAttributes"`
}
