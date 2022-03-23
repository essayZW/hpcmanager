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

// LimitQueryByCreaterUserID 通过创建者的ID分页查询机器节点申请信息
func (node *NodeApplyDB) LimitQueryByCreaterUserID(ctx context.Context, userID, limit, offset int) ([]*NodeApply, error) {
	row, err := node.conn.Query(ctx, "SELECT * FROM `node_apply` WHERE `creater_id`=? LIMIT ?, ?", userID, limit, offset)
	if err != nil {
		logger.Warn("LimitQueryByCreaterUserID error: ", err)
		return nil, errors.New("LimitQueryByCreaterUserID error")
	}
	infos := make([]*NodeApply, 0)
	for row.Next() {
		var info NodeApply
		if err := row.StructScan(&info); err != nil {
			logger.Warn("LimitQueryByCreaterUserID struct scan error: ", err)
			return nil, errors.New("LimitQueryByCreaterUserID struct scan error")
		}
		infos = append(infos, &info)
	}
	return infos, nil
}

// QueryCountByCreaterID 查询某个用户创建的记录的数量
func (node *NodeApplyDB) QueryCountByCreaterID(ctx context.Context, userID int) (int, error) {
	res, err := node.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `node_apply` WHERE `creater_id`=?", userID)
	if err != nil {
		logger.Warn("QueryCountByCreaterID error: ", err)
		return 0, errors.New("QueryCountByCreaterID error")
	}
	var count int
	if err := res.Scan(&count); err != nil {
		logger.Warn("QueryCountByCreaterID int scan error: ", err)
		return 0, errors.New("QueryCountByCreaterID int scan error")
	}
	return count, nil
}

// LimitQueryByTutorID 通过导师ID分页查询申请信息
func (node *NodeApplyDB) LimitQueryByTutorID(ctx context.Context, tutorID, limit, offset int) ([]*NodeApply, error) {
	row, err := node.conn.Query(ctx, "SELECT * FROM `node_apply` WHERE `tutor_id`=? LIMIT ?, ?", tutorID, limit, offset)
	if err != nil {
		logger.Warn("LimitQueryByTutorID error: ", err)
		return nil, errors.New("LimitQueryByTutorID error")
	}
	infos := make([]*NodeApply, 0)
	for row.Next() {
		var info NodeApply
		if err := row.StructScan(&info); err != nil {
			logger.Warn("LimitQueryByTutorID struct scan error: ", err)
			return nil, errors.New("LimitQueryByTutorID struct scan error")
		}
		infos = append(infos, &info)
	}
	return infos, nil
}

// QueryCountByTutorID 查询某个导师的用户组下的所有用户的申请记录的数量
func (node *NodeApplyDB) QueryCountByTutorID(ctx context.Context, tutorID int) (int, error) {
	res, err := node.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `node_apply` WHERE `tutor_id`=?", tutorID)
	if err != nil {
		logger.Warn("QueryCountByTutorID error: ", err)
		return 0, errors.New("QueryCountByTutorID error")
	}
	var count int
	if err := res.Scan(&count); err != nil {
		logger.Warn("QueryCountByTutorID int scan error: ", err)
		return 0, errors.New("QueryCountByTutorID int scan error")
	}
	return count, nil
}

// LimitQueryWithTutorChecked 分页查询所有的申请记录信息
func (node *NodeApplyDB) LimitQueryWithTutorChecked(ctx context.Context, limit, offset int) ([]*NodeApply, error) {
	// 查询被导师审核通过的所有申请
	row, err := node.conn.Query(ctx, "SELECT * FROM `node_apply` WHERE `tutor_check_status`=1 LIMIT ?, ?", limit, offset)
	if err != nil {
		logger.Warn("LimitQuery error: ", err)
		return nil, errors.New("LimitQuery error")
	}
	infos := make([]*NodeApply, 0)
	for row.Next() {
		var info NodeApply
		if err := row.StructScan(&info); err != nil {
			logger.Warn("LimitQuery struct scan error: ", err)
			return nil, errors.New("LimitQuery struct scan error")
		}
		infos = append(infos, &info)
	}
	return infos, nil
}

// QueryCountWithTutorChecked 查询申请总条数
func (node *NodeApplyDB) QueryCountWithTutorChecked(ctx context.Context) (int, error) {
	res, err := node.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `node_apply` WHERE `tutor_check_status`=1")
	if err != nil {
		logger.Warn("QueryCountByTutorID error: ", err)
		return 0, errors.New("QueryCountByTutorID error")
	}
	var count int
	if err := res.Scan(&count); err != nil {
		logger.Warn("QueryCountByTutorID int scan error: ", err)
		return 0, errors.New("QueryCountByTutorID int scan error")
	}
	return count, nil
}

// UpdateTutorCheckStatus 更新导师审核的状态
func (node *NodeApplyDB) UpdateTutorCheckStatus(ctx context.Context, newStatus *NodeApply) (bool, error) {
	res, err := node.conn.Exec(ctx, "UPDATE `node_apply` SET `tutor_check_status`=?, `tutor_check_time`=?, `message_tutor`=? WHERE `tutor_check_status`=-1 "+
		"AND `id`=? AND `status`=1", newStatus.TutorCheckStatus, newStatus.TutorCheckTime, newStatus.MessageTutor, newStatus.ID)
	if err != nil {
		logger.Warn("UpdateTutorCheckStatus error: ", err)
		return false, errors.New("UpdateTutorCheckStatus error")
	}
	count, err := res.RowsAffected()
	if err != nil {
		logger.Warn("UpdateTutorCheckStatus error: ", err)
		return false, errors.New("UpdateTutorCheckStatus error")
	}
	return count != 0, nil
}

// UpdateAdminCheckStatus 更新管理员审核的状态
func (node *NodeApplyDB) UpdateAdminCheckStatus(ctx context.Context, newStatus *NodeApply) (bool, error) {
	res, err := node.conn.Exec(ctx, "UPDATE `node_apply` SET `manager_check_status`=?, `manager_check_time`=?, `message_manager`=?, "+
		"`manager_checker_id`=?, `manager_checker_name`=?, `manager_checker_username`=? WHERE "+
		"`id`=? AND `tutor_check_status`=1 AND `status`=1", newStatus.ManagerCheckStatus, newStatus.ManagerCheckTime, newStatus.MessageManager,
		newStatus.ManagerCheckerID, newStatus.ManagerCheckerName, newStatus.ManagerCheckerUsername, newStatus.ID)
	if err != nil {
		logger.Warn("UpdateAdminCheckStatus error: ", err)
		return false, errors.New("UpdateAdminCheckStatus error")
	}
	count, err := res.RowsAffected()
	if err != nil {
		logger.Warn("UpdateAdminCheckStatus error: ", err)
		return false, errors.New("UpdateAdminCheckStatus error")
	}
	return count != 0, nil
}

// NewNodeApply 创建新的node_apply数据库操作
func NewNodeApply(conn *db.DB) *NodeApplyDB {
	return &NodeApplyDB{
		conn: conn,
	}
}
