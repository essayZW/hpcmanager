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

// NewNodeApply 创建机器节点申请相关的逻辑操作
func NewNodeApply(nodeApplyDB *db.NodeApplyDB) *NodeApply {
	return &NodeApply{
		nodeApplyDB: nodeApplyDB,
	}
}
