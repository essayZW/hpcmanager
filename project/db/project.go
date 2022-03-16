package db

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/db"
	"go-micro.dev/v4/logger"
)

// ProjectDB project数据库操作
type ProjectDB struct {
	conn *db.DB
}

// Insert 插入新的project记录
func (pdb *ProjectDB) Insert(ctx context.Context, newProject *Project) (int64, error) {
	res, err := pdb.conn.Exec(ctx, "INSERT INTO `project` "+
		"(`name`, `from`, `numbering`, `expenses`, `description`, `creater_user_id`, `creater_username`, "+
		"`creater_user_name`, `create_time`, `modify_user_id`, `modify_username`, `modify_user_name`, `modify_time`, `extraAttributes`) "+
		" VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		newProject.Name, newProject.From, newProject.Numbering, newProject.Expenses, newProject.Description,
		newProject.CreaterUserID, newProject.CreaterUserName, newProject.CreaterUsername, newProject.CreateTime,
		newProject.ModifyUserID, newProject.ModifyUserName, newProject.ModifyUsername, newProject.ModifyTime, newProject.ExtraAttributes)
	if err != nil {
		logger.Warn("Insert project error: ", err)
		return 0, errors.New("Insert project error")
	}
	id, err := res.LastInsertId()
	if err != nil {
		logger.Warn("Insert project error: ", err)
		return 0, errors.New("Insert project error")
	}
	return id, nil
}

// QueryByID 通过ID查询记录
func (pdb *ProjectDB) QueryByID(ctx context.Context, id int) (*Project, error) {
	res, err := pdb.conn.QueryRow(ctx, "SELECT * FROM `project` WHERE `id`=?", id)
	if err != nil {
		logger.Warn("QueryByID error: ", err)
		return nil, errors.New("QueryByID error")
	}
	var info Project
	if err = res.StructScan(&info); err != nil {
		logger.Warn("QueryByID error: ", err)
		return nil, errors.New("QueryByID error")
	}
	return &info, nil
}

// NewProject 创建新的数据库操作结构
func NewProject(sqlConn *db.DB) *ProjectDB {
	return &ProjectDB{
		conn: sqlConn,
	}
}
