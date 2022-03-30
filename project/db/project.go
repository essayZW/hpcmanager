package db

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/essayZW/hpcmanager/db"
	"go-micro.dev/v4/logger"
)

// ProjectDB project数据库操作
type ProjectDB struct {
	conn *db.DB
}

// Insert 插入新的project记录
func (pdb *ProjectDB) Insert(ctx context.Context, newProject *Project) (int64, error) {
	res, err := pdb.conn.Exec(
		ctx,
		"INSERT INTO `project` "+
			"(`name`, `from`, `numbering`, `expenses`, `description`, `creater_user_id`, `creater_username`, "+
			"`creater_user_name`, `create_time`, `modify_user_id`, `modify_username`, `modify_user_name`, `modify_time`, `extraAttributes`) "+
			" VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		newProject.Name,
		newProject.From,
		newProject.Numbering,
		newProject.Expenses,
		newProject.Description,
		newProject.CreaterUserID,
		newProject.CreaterUserName,
		newProject.CreaterUsername,
		newProject.CreateTime,
		newProject.ModifyUserID,
		newProject.ModifyUserName,
		newProject.ModifyUsername,
		newProject.ModifyTime,
		newProject.ExtraAttributes,
	)
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

// LimitQuery 通过limit查询一系列的项目信息
func (pdb *ProjectDB) LimitQuery(ctx context.Context, limit int, offset int) ([]*Project, error) {
	res, err := pdb.conn.Query(ctx, "SELECT * FROM `project` LIMIT ?, ?", limit, offset)
	if err != nil {
		logger.Warn("LimitQuery error: ", err)
		return nil, errors.New("LimitQuery error")
	}
	infos := make([]*Project, 0)
	for res.Next() {
		var info Project
		if err := res.StructScan(&info); err != nil {
			logger.Warn("LimitQuery struct scan error: ", err)
			return nil, errors.New("LimitQuery struct scan error")
		}
		infos = append(infos, &info)
	}
	return infos, nil
}

// QueryCount 查询所有的项目数量
func (pdb *ProjectDB) QueryCount(ctx context.Context) (int, error) {
	row, err := pdb.conn.QueryRow(ctx, "SELECT COUNT(*) FROM `project`")
	if err != nil {
		logger.Warn("QueryCount error: ", err)
		return 0, errors.New("QueryCount error")
	}
	var count int
	if err := row.Scan(&count); err != nil {
		logger.Warn("QueryCount error: ", err)
		return 0, errors.New("QueryCount error")
	}
	return count, nil
}

func (pdb *ProjectDB) expandOrSQL(ctx context.Context, field string, len int) string {
	// NOTE: 多个or语句拼接情况下也会导致全表扫描
	str := strings.Builder{}
	item := fmt.Sprintf("`%s`=?", field)
	for i := 0; i < len; i++ {
		str.WriteString(item)
		if i != len-1 {
			str.WriteString(" OR ")
		}
	}
	return str.String()
}

// QueryCountByCreaterUserID 查询某一个用户创建的所有项目的总数
func (pdb *ProjectDB) QueryCountByCreaterUserID(ctx context.Context, userID ...int) (int, error) {
	params := make([]interface{}, len(userID))
	for i := range params {
		params[i] = userID[i]
	}
	row, err := pdb.conn.QueryRow(
		ctx,
		"SELECT COUNT(*) FROM `project` WHERE "+pdb.expandOrSQL(
			ctx,
			"creater_user_id",
			len(params),
		),
		params...)
	if err != nil {
		logger.Warn("QueryCountByCreaterUserID error: ", err)
		return 0, errors.New("QueryCountByCreaterUserID error")
	}
	var count int
	if err := row.Scan(&count); err != nil {
		logger.Warn("QueryCountByCreaterUserID error: ", err)
		return 0, errors.New("QueryCountByCreaterUserID error")
	}
	return count, nil
}

// LimitQueryByCreaterUserID 通过limit查询某个用户创建的所有项目信息
func (pdb *ProjectDB) LimitQueryByCreaterUserID(
	ctx context.Context,
	limit, offset int,
	userID ...int,
) ([]*Project, error) {
	params := make([]interface{}, len(userID))
	for i := range userID {
		params[i] = userID[i]
	}
	params = append(params, limit)
	params = append(params, offset)
	sql := "SELECT * FROM `project` WHERE " + pdb.expandOrSQL(
		ctx,
		"creater_user_id",
		len(userID),
	) + " LIMIT ?, ?"
	res, err := pdb.conn.Query(ctx, sql, params...)
	if err != nil {
		logger.Warn("LimitQueryByCreaterUserID error: ", err)
		return nil, errors.New("LimitQueryByCreaterUserID error")
	}
	infos := make([]*Project, 0)
	for res.Next() {
		var info Project
		if err := res.StructScan(&info); err != nil {
			logger.Warn("LimitQueryByCreaterUserID struct scan error: ", err)
			return nil, errors.New("LimitQueryByCreaterUserID struct scan error")
		}
		infos = append(infos, &info)
	}
	return infos, nil
}

// NewProject 创建新的数据库操作结构
func NewProject(sqlConn *db.DB) *ProjectDB {
	return &ProjectDB{
		conn: sqlConn,
	}
}
