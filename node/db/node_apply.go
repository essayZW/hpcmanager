package db

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
)

// NodeApplyDB 机器节点相关数据库操作
type NodeApplyDB struct {
	conn *db.DB
}

// Insert 插入新的申请记录
func (node *NodeApplyDB) Insert(ctx context.Context, nodeApplyInfo *NodeApply) (int64, error) {
	res, err := node.conn.Exec(ctx, "INSERT INTO `node_apply`"+
		"(`create_time`, `creater_id`, `creater_name`, `creater_username`,`project_id`, `tutor_id`, `tutor_name`, `tutor_username`,"+
		"`modify_time`, `modify_name`, `modify_userid`, `modify_username`, `node_type`, `node_num`,"+
		" `start_time`, `end_time`, `extraAttributes`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		nodeApplyInfo.CreateTime, nodeApplyInfo.CreaterID, nodeApplyInfo.CreaterName, nodeApplyInfo.CreaterUsername, nodeApplyInfo.ProjectID,
		nodeApplyInfo.TutorID, nodeApplyInfo.TutorName, nodeApplyInfo.TutorUsername, nodeApplyInfo.ModifyTime, nodeApplyInfo.ModifyName, nodeApplyInfo.ModifyUserID,
		nodeApplyInfo.ModifyUsername, nodeApplyInfo.NodeType, nodeApplyInfo.NodeNum, nodeApplyInfo.StartTime, nodeApplyInfo.EndTime, nodeApplyInfo.ExtraAttributes)
	if err != nil {
		logger.Warn("Insert apply info error: ", err)
		return 0, errors.New("Insert apply info error")
	}
	id, err := res.LastInsertId()
	if err != nil {
		logger.Warn("Insert apply info error: ", err)
		return 0, errors.New("Insert apply info error")
	}
	return id, nil
}

// NewNodeApply 创建新的node_apply数据库操作
func NewNodeApply(conn *db.DB) *NodeApplyDB {
	return &NodeApplyDB{
		conn: conn,
	}
}
