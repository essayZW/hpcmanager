package logic

import (
	"context"
	"errors"
	"time"

	"github.com/essayZW/hpcmanager/fee/db"
)

// NodeDistributeBill 机器独占账单操作逻辑
type NodeDistributeBill struct {
	ndb *db.NodeDistributeBillDB
}

// Create 创建新的机器节点独占账单
func (ndbl *NodeDistributeBill) Create(
	ctx context.Context,
	applyID int,
	nodeDistributeID int,
	fee float64,
	userID int,
	username, userName string,
) (int64, error) {
	if applyID <= 0 {
		return 0, errors.New("invalid apply id")
	}
	if nodeDistributeID <= 0 {
		return 0, errors.New("invalid node distribute id")
	}
	if fee < 0 {
		return 0, errors.New("invalid fee")
	}
	if userID <= 0 {
		return 0, errors.New("invalid user id")
	}
	return ndbl.ndb.Insert(ctx, &db.NodeDistributeBill{
		ApplyID:          applyID,
		NodeDistributeID: nodeDistributeID,
		Fee:              fee,
		UserID:           userID,
		Username:         username,
		UserName:         userName,
		CreateTime:       time.Now(),
	})
}

// CalFee 计算账单的费用
func (ndbl *NodeDistributeBill) CalFee() (float64, error) {
	// TODO: 具体的实现
	return 0.0, nil
}

// NewNodeDistributeBill 创建新的机器独占账单操作逻辑结构体
func NewNodeDistributeBill(ndb *db.NodeDistributeBillDB) *NodeDistributeBill {
	return &NodeDistributeBill{
		ndb: ndb,
	}
}
