package logic

import (
	"context"
	"errors"
	"time"

	"github.com/essayZW/hpcmanager/node/db"
)

// NodeDistribute 机器节点分配处理工单的操作逻辑
type NodeDistribute struct {
	nodeDistributeDB *db.NodeDistributeDB
}

// CreateNodeDistributeWO 创建节点分配处理工单
func (nodeDistribute *NodeDistribute) CreateNodeDistributeWO(ctx context.Context, applyID int) (int64, error) {
	// 检查是否已经存在同一条applyID对应的工单
	count, err := nodeDistribute.nodeDistributeDB.QueryCountByApply(ctx, applyID)
	if err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, errors.New("the work order for this applyID has exists")
	}
	return nodeDistribute.nodeDistributeDB.Insert(ctx, &db.NodeDistribute{
		ApplyID:    applyID,
		CreateTime: time.Now(),
	})
}

// NewNodeDistribute 创建新的机器节点分配处理工单的操作逻辑
func NewNodeDistribute(nodeDistributeDB *db.NodeDistributeDB) *NodeDistribute {
	return &NodeDistribute{
		nodeDistributeDB: nodeDistributeDB,
	}
}
