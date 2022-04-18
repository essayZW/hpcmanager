package db

import (
	"github.com/essayZW/hpcmanager/db"
	"gopkg.in/guregu/null.v4"
)

// HpcUser hpc_user表
type HpcUser struct {
	ID           int    `db:"id"`
	NodeUsername string `db:"node_username"`
	NodeUID      int    `db:"node_uid"`
	// NodeMaxQuota 废弃,该字段原本代表的信息都通过作业调度系统实时获取,不在数据库里面存储
	NodeMaxQuota    int8      `db:"node_max_quota"`
	QuotaStartTime  null.Time `db:"quota_start_time"`
	QuotaEndTime    null.Time `db:"quota_end_time"`
	ExtraAttributes *db.JSON  `db:"extraAttributes"`
}

// HpcGroup hpc_group表
type HpcGroup struct {
	ID              int      `db:"id"`
	Name            string   `db:"name"`
	QueueName       string   `db:"queue_name"`
	GID             int      `db:"gid"`
	ExtraAttributes *db.JSON `db:"extraAttributes"`
}
