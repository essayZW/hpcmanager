package logic

import "github.com/essayZW/hpcmanager/fee/db"

// NodeDistributeBill 机器独占账单操作逻辑
type NodeDistributeBill struct {
	ndb *db.NodeDistributeBillDB
}

// NewNodeDistributeBill 创建新的机器独占账单操作逻辑结构体
func NewNodeDistributeBill(ndb *db.NodeDistributeBillDB) *NodeDistributeBill {
	return &NodeDistributeBill{
		ndb: ndb,
	}
}
