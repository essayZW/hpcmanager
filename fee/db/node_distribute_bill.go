package db

import "github.com/essayZW/hpcmanager/db"

type NodeDistributeBillDB struct {
	conn *db.DB
}

// NewNodeDistributeBill 新建一个node_distribute_bill数据表操作结构体
func NewNodeDistributeBill(conn *db.DB) *NodeDistributeBillDB {
	return &NodeDistributeBillDB{
		conn: conn,
	}
}
