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
	CreateTime       time.Time   `db:"create_time"`
	ExtraAttributes  *db.JSON    `db:"extraAttributes"`
}
