package logic

import (
	"context"
	"errors"
	"time"

	"github.com/essayZW/hpcmanager/node/db"
	"gopkg.in/guregu/null.v4"
)

// NodeDistribute 机器节点分配处理工单的操作逻辑
type NodeDistribute struct {
	nodeDistributeDB *db.NodeDistributeDB
}

// CreateNodeDistributeWO 创建节点分配处理工单
func (nodeDistribute *NodeDistribute) CreateNodeDistributeWO(
	ctx context.Context,
	applyID int,
) (int64, error) {
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

// PaginationQueryNodeDistribute 分页查询的结果
type PaginationQueryNodeDistribute struct {
	Data  []*db.NodeDistribute
	Count int
}

// PaginationGet 分页查询
func (nodeDistribute *NodeDistribute) PaginationGet(
	ctx context.Context,
	pageIndex, pageSize int,
) (*PaginationQueryNodeDistribute, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid page index")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid page size")
	}

	count, err := nodeDistribute.nodeDistributeDB.QueryCount(ctx)
	if err != nil {
		return nil, errors.New("query count error")
	}

	if count == 0 {
		return &PaginationQueryNodeDistribute{
			Count: 0,
			Data:  make([]*db.NodeDistribute, 0),
		}, nil
	}
	limit := pageSize * (pageIndex - 1)
	data, err := nodeDistribute.nodeDistributeDB.QueryLimit(ctx, limit, pageSize)
	if err != nil {
		return nil, errors.New("data query error")
	}
	return &PaginationQueryNodeDistribute{
		Count: count,
		Data:  data,
	}, nil

}

// SimpleUserInfo 简单的用户信息
type SimpleUserInfo struct {
	ID       int
	Username string
	Name     string
}

// FinishByID 通过ID处理机器节点分配工单
func (nodeDistribute *NodeDistribute) FinishByID(
	ctx context.Context,
	id int,
	userInfo *SimpleUserInfo,
) (bool, error) {
	return nodeDistribute.nodeDistributeDB.UpdateHandlerFlag(ctx, &db.NodeDistribute{
		ID:              id,
		HandlerUserID:   null.IntFrom(int64(userInfo.ID)),
		HandlerUsername: null.StringFrom(userInfo.Username),
		HandlerUserName: null.StringFrom(userInfo.Name),
	})
}

// GetInfoByID 通过ID获取记录的信息
func (nodeDistribute *NodeDistribute) GetInfoByID(ctx context.Context, id int32) (*db.NodeDistribute, error) {
	return nodeDistribute.nodeDistributeDB.QueryByID(ctx, id)
}

// NewNodeDistribute 创建新的机器节点分配处理工单的操作逻辑
func NewNodeDistribute(nodeDistributeDB *db.NodeDistributeDB) *NodeDistribute {
	return &NodeDistribute{
		nodeDistributeDB: nodeDistributeDB,
	}
}
