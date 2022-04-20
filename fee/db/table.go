package db

import (
	"time"

	"github.com/essayZW/hpcmanager/db"
	"gopkg.in/guregu/null.v4"
)

// NodeDistributeBill 机器独占账单表映射
type NodeDistributeBill struct {
	ID               int         `db:"id"`
	ApplyID          int         `db:"apply_id"`
	NodeDistributeID int         `db:"node_distribute_id"`
	Fee              float64     `db:"fee"`
	PayFee           float64     `db:"pay_fee"`
	PayFlag          int8        `db:"pay_flag"`
	PayTime          null.Time   `db:"pay_time"`
	PayType          null.Int    `db:"pay_type"`
	PayMessage       null.String `db:"pay_message"`
	UserID           int         `db:"user_id"`
	Username         string      `db:"user_username"`
	UserName         string      `db:"user_name"`
	UserGroupID      int         `db:"user_group_id"`
	CreateTime       time.Time   `db:"create_time"`
	ExtraAttributes  *db.JSON    `db:"extraAttributes"`
}

// NodeWeekUsageBill 机器时长周账单表映射
type NodeWeekUsageBill struct {
	ID              int         `db:"id"`
	UserID          int         `db:"user_id"`
	Username        string      `db:"user_username"`
	UserName        string      `db:"user_name"`
	WallTime        int         `db:"wall_time"`
	GWallTime       int         `db:"gwall_time"`
	Fee             float64     `db:"fee"`
	PayFee          float64     `db:"pay_fee"`
	StartTime       time.Time   `db:"start_time"`
	EndTime         time.Time   `db:"end_time"`
	PayFlag         int8        `db:"pay_flag"`
	PayTime         null.Time   `db:"pay_time"`
	PayType         null.Int    `db:"pay_type"`
	PayMessage      null.String `db:"pay_message"`
	UserGroupID     int         `db:"user_group_id"`
	CreateTime      time.Time   `db:"create_time"`
	ExtraAttributes *db.JSON    `db:"extraAttributes"`
}

// NodeWeekUsageBillForUserGroup 按照user_group_id分组的查询结果
type NodeWeekUsageBillForUserGroup struct {
	WallTime    int     `db:"wall_time"`
	GWallTime   int     `db:"gwall_time"`
	Fee         float64 `db:"fee"`
	PayFee      float64 `db:"pay_fee"`
	PayFlag     int8    `db:"pay_flag"`
	UserGroupID int     `db:"user_group_id"`
}

// NodeQuotaBill 数据库中node_quota_bill表的映射
type NodeQuotaBill struct {
	ID              int         `db:"id"`
	UserID          int         `db:"user_id"`
	UserName        string      `db:"user_name"`
	Username        string      `db:"user_username"`
	UserGroupID     int         `db:"user_group_id"`
	OperType        int8        `db:"oper_type"`
	OldSize         int         `db:"old_size"`
	NewSize         int         `db:"new_size"`
	OldEndTime      time.Time   `db:"old_end_time"`
	NewEndTime      time.Time   `db:"new_end_time"`
	Fee             float64     `db:"fee"`
	PayFlag         int8        `db:"pay_flag"`
	PayFee          float64     `db:"pay_fee"`
	PayTime         null.Time   `db:"pay_time"`
	PayType         null.Int    `db:"pay_type"`
	PayMessage      null.String `db:"pay_message"`
	CreateTime      time.Time   `db:"create_time"`
	ExtraAttributes *db.JSON    `db:"extraAttributes"`
}
