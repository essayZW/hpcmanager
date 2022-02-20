package db

import (
	"time"

	"github.com/essayZW/hpcmanager/db"
	"gopkg.in/guregu/null.v4"
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

// UserGroupApply 数据库中的新用户申请表结构体
type UserGroupApply struct {
	ID                     int         `db:"id"`
	UserID                 int         `db:"user_id"`
	UserUsername           string      `db:"user_username"`
	UserName               string      `db:"user_name"`
	ApplyGroupID           int         `db:"apply_group_id"`
	TutorID                int         `db:"tutor_id"`
	TutorUsername          string      `db:"tutor_username"`
	TutorName              string      `db:"tutor_name"`
	TutorCheckStatus       int8        `db:"tutor_check_status"`
	ManagerCheckStatus     int8        `db:"manager_check_status"`
	Status                 int8        `db:"status"`
	MessageTutor           null.String `db:"message_tutor"`
	MessageManager         null.String `db:"message_manager"`
	TutorCheckTime         null.Time   `db:"tutor_check_time"`
	ManagerCheckTime       null.Time   `db:"manager_check_time"`
	ManagerCheckerID       null.Int    `db:"manager_checker_id"`
	ManagerCheckerUsername null.String `db:"manager_checker_username"`
	ManagerCheckerName     null.String `db:"manager_checker_name"`
	CreateTime             null.Time   `db:"create_time"`
	ExtraAttributes        *db.JSON    `db:"extraAttributes"`
}
