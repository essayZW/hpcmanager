package db

import (
	"time"

	"github.com/essayZW/hpcmanager/db"
	"gopkg.in/guregu/null.v4"
)

// NodeApply 计算节点申请表的结构映射
type NodeApply struct {
	ID                     int         `db:"id"`
	CreateTime             time.Time   `db:"create_time"`
	CreaterID              int         `db:"creater_id"`
	CreaterUsername        string      `db:"creater_username"`
	CreaterName            string      `db:"creater_name"`
	ProjectID              int         `db:"project_id"`
	TutorCheckStatus       int8        `db:"tutor_check_status"`
	ManagerCheckStatus     int8        `db:"manager_check_status"`
	Status                 int8        `db:"status"`
	MessageTutor           null.String `db:"message_tutor"`
	MessageManager         null.String `db:"message_manager"`
	TutorCheckTime         null.Time   `db:"tutor_check_time"`
	TutorID                int         `db:"tutor_id"`
	TutorName              string      `db:"tutor_name"`
	TutorUsername          string      `db:"tutor_username"`
	ManagerCheckTime       null.Time   `db:"manager_check_time"`
	ManagerCheckerID       null.Int    `db:"manager_checker_id"`
	ManagerCheckerName     null.String `db:"manager_checker_name"`
	ManagerCheckerUsername null.String `db:"manager_checker_username"`
	ModifyTime             null.Time   `db:"modify_time"`
	ModifyUserID           int         `db:"modify_userid"`
	ModifyName             string      `db:"modify_name"`
	ModifyUsername         string      `db:"modify_username"`
	NodeType               string      `db:"node_type"`
	NodeNum                int         `db:"node_num"`
	StartTime              time.Time   `db:"start_time"`
	EndTime                time.Time   `db:"end_time"`
	ExtraAttributes        *db.JSON    `db:"extraAttributes"`
}

// NodeDistribute 机器节点分配工单
type NodeDistribute struct {
	ID               int         `db:"id"`
	ApplyID          int         `db:"apply_id"`
	HandlerFlag      int8        `db:"handler_flag"`
	HandlerUserID    null.Int    `db:"handler_userid"`
	HandlerUsername  null.String `db:"handler_username"`
	HandlerUserName  null.String `db:"handler_user_name"`
	DistributeBillID int         `db:"distribute_bill_id"`
	CreateTime       time.Time   `db:"create_time"`
	ExtraAttributes  *db.JSON    `db:"extraAttributes"`
}

// HpcUsageTime 机器节点使用时间记录表
type HpcUsageTime struct {
	ID              int       `db:"id"`
	UserID          int       `db:"user_id"`
	QueueName       string    `db:"queue_name"`
	WallTime        float64   `db:"wall_time"`
	GWallTime       float64   `db:"gwall_time"`
	StartTime       time.Time `db:"start_time"`
	EndTime         time.Time `db:"end_time"`
	CreateTime      time.Time `db:"create_time"`
	ExtraAttributes *db.JSON  `db:"extraAttributes"`
}
