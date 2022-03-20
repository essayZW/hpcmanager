package logic

import (
	"context"
	"errors"
	"time"

	"github.com/essayZW/hpcmanager/node/db"
	"gopkg.in/guregu/null.v4"
)

// NodeApply 机器节点申请相关的操作逻辑
type NodeApply struct {
	nodeApplyDB *db.NodeApplyDB
}

// ApplyItemUserInfo 用于创建申请记录的相关角色用户信息集合
type ApplyItemUserInfo struct {
	ID       int
	Username string
	Name     string
}

// ApplyNodeInfo 申请的节点信息
type ApplyNodeInfo struct {
	NodeType  string
	NodeNum   int
	StartTime time.Time
	EndTime   time.Time
}

// CreateNodeApply 创建新的机器节点申请记录
func (node *NodeApply) CreateNodeApply(ctx context.Context, user, tutor *ApplyItemUserInfo, nodeInfo *ApplyNodeInfo, projectID int) (int64, error) {
	if user.ID <= 0 {
		return 0, errors.New("invalid user info")
	}
	if tutor.ID <= 0 {
		return 0, errors.New("invalid tutor info")
	}
	if nodeInfo.StartTime.IsZero() || nodeInfo.EndTime.IsZero() {
		return 0, errors.New("invalid start or end time")
	}
	if nodeInfo.EndTime.Before(nodeInfo.StartTime) {
		return 0, errors.New("end time can't earlier than start time")
	}
	if nodeInfo.NodeType == "" || nodeInfo.NodeNum <= 0 {
		return 0, errors.New("invalid node info")
	}
	if projectID <= 0 {
		return 0, errors.New("must have valid project id info")
	}
	return node.nodeApplyDB.Insert(ctx, &db.NodeApply{
		CreaterID:       user.ID,
		CreaterUsername: user.Username,
		CreaterName:     user.Name,
		CreateTime:      time.Now(),
		ProjectID:       projectID,
		TutorID:         tutor.ID,
		TutorUsername:   tutor.Username,
		TutorName:       tutor.Name,
		ModifyTime:      null.TimeFrom(time.Now()),
		ModifyUserID:    user.ID,
		ModifyName:      user.Name,
		ModifyUsername:  user.Username,
		NodeType:        nodeInfo.NodeType,
		NodeNum:         nodeInfo.NodeNum,
		StartTime:       nodeInfo.StartTime,
		EndTime:         nodeInfo.EndTime,
	})
}

// PaginationGetResult 分页查询的结果
type PaginationGetResult struct {
	Count int
	Data  []*db.NodeApply
}

// PaginationGetByCreaterID 分页查询某一个用户创建的申请记录
func (node *NodeApply) PaginationGetByCreaterID(ctx context.Context, createrID, pageIndex, pageSize int) (*PaginationGetResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}
	count, err := node.nodeApplyDB.QueryCountByCreaterID(ctx, createrID)
	if err != nil {
		return nil, err
	}
	limit := pageSize * (pageIndex - 1)
	data, err := node.nodeApplyDB.LimitQueryByCreaterUserID(ctx, createrID, limit, pageSize)
	if err != nil {
		return nil, err
	}
	return &PaginationGetResult{
		Data:  data,
		Count: count,
	}, nil
}

// PaginationGetByTutorID 分页查询某个导师用户组下的所有申请记录信息
func (node *NodeApply) PaginationGetByTutorID(ctx context.Context, tutorID, pageIndex, pageSize int) (*PaginationGetResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}
	count, err := node.nodeApplyDB.QueryCountByTutorID(ctx, tutorID)
	if err != nil {
		return nil, err
	}
	limit := pageSize * (pageIndex - 1)
	data, err := node.nodeApplyDB.LimitQueryByTutorID(ctx, tutorID, limit, pageSize)
	if err != nil {
		return nil, err
	}
	return &PaginationGetResult{
		Data:  data,
		Count: count,
	}, nil
}

// PaginationWithTutorChecked 分页 查询所有已经经过导师审核通过的申请信息
func (node *NodeApply) PaginationWithTutorChecked(ctx context.Context, pageIndex, pageSize int) (*PaginationGetResult, error) {
	if pageIndex <= 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize <= 0 {
		return nil, errors.New("invalid pageSize")
	}
	count, err := node.nodeApplyDB.QueryCountWithTutorChecked(ctx)
	if err != nil {
		return nil, err
	}
	limit := pageSize * (pageIndex - 1)
	data, err := node.nodeApplyDB.LimitQueryWithTutorChecked(ctx, limit, pageSize)
	if err != nil {
		return nil, err
	}
	return &PaginationGetResult{
		Data:  data,
		Count: count,
	}, nil
}

// NewNodeApply 创建机器节点申请相关的逻辑操作
func NewNodeApply(nodeApplyDB *db.NodeApplyDB) *NodeApply {
	return &NodeApply{
		nodeApplyDB: nodeApplyDB,
	}
}
