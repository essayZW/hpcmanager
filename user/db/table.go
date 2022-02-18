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

// Group 数据库中的Group表的结构体
type Group struct {
	ID                int       `db:"id"`
	Name              string    `db:"name"`
	QueueName         string    `db:"queue_name"`
	NodeUserGroupName string    `db:"node_usergroup"`
	CreateTime        time.Time `db:"create_time"`
	CreaterID         int       `db:"creater_id"`
	CreaterUsername   string    `db:"creater_username"`
	CreaterName       string    `db:"creater_name"`
	TutorID           int       `db:"tutor_id"`
	TutorUsername     string    `db:"tutor_username"`
	TutorName         string    `db:"tutor_name"`
	Balance           float64   `db:"balance"`
	ExtraAttributes   *db.JSON  `db:"extraAttributes"`
}
