package logic

import (
	"context"
	"errors"
	"time"

	projectdb "github.com/essayZW/hpcmanager/project/db"
)

// Project 项目logic
type Project struct {
	projectDB *projectdb.ProjectDB
}

// Create 创建新的project记录
func (p *Project) Create(ctx context.Context, createrUserID int, createrUserName string, createrUsername string, projectInfo *projectdb.Project) (int64, error) {
	if createrUserID == 0 {
		return 0, errors.New("user id can't be zero")
	}
	if projectInfo.Name == "" {
		return 0, errors.New("project name can't be empty")
	}
	return p.projectDB.Insert(ctx, &projectdb.Project{
		Name:            projectInfo.Name,
		From:            projectInfo.From,
		Numbering:       projectInfo.Numbering,
		Expenses:        projectInfo.Expenses,
		Description:     projectInfo.Description,
		CreaterUserID:   createrUserID,
		CreaterUserName: createrUserName,
		CreaterUsername: createrUsername,
		CreateTime:      time.Now(),
		ModifyUserID:    createrUserID,
		ModifyUserName:  createrUserName,
		ModifyUsername:  createrUsername,
		ModifyTime:      time.Now(),
	})
}

// GetByID 通过ID获取项目信息
func (p *Project) GetByID(ctx context.Context, id int) (*projectdb.Project, error) {
	return p.projectDB.QueryByID(ctx, id)
}

// PaginationProjectResult 分页查询项目信息的结果
type PaginationProjectResult struct {
	Data  []*projectdb.Project
	Count int
}

// PaginationGet 分页查询项目信息
func (p *Project) PaginationGet(ctx context.Context, pageIndex, pageSize int) (*PaginationProjectResult, error) {
	if pageIndex == 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize == 0 {
		return nil, errors.New("invalid pageSize")
	}
	count, err := p.projectDB.QueryCount(ctx)
	if err != nil {
		return nil, errors.New("query count error")
	}
	limit := pageSize * (pageIndex - 1)
	infos, err := p.projectDB.LimitQuery(ctx, limit, pageSize)
	if err != nil {
		return nil, err
	}
	return &PaginationProjectResult{
		Data:  infos,
		Count: count,
	}, nil
}

// PaginationGetByCreaterUserID 分页查询某个用户创建的所有项目信息
func (p *Project) PaginationGetByCreaterUserID(ctx context.Context, pageIndex, pageSize int, userID ...int) (*PaginationProjectResult, error) {
	if pageIndex == 0 {
		return nil, errors.New("invalid pageIndex")
	}
	if pageSize == 0 {
		return nil, errors.New("invalid pageSize")
	}
	count, err := p.projectDB.QueryCountByCreaterUserID(ctx, userID...)
	if err != nil {
		return nil, errors.New("query count error")
	}
	limit := pageSize * (pageIndex - 1)
	infos, err := p.projectDB.LimitQueryByCreaterUserID(ctx, limit, pageSize, userID...)
	if err != nil {
		return nil, err
	}
	return &PaginationProjectResult{
		Data:  infos,
		Count: count,
	}, nil
}

// NewProject 创建新的project的logic
func NewProject(projectDB *projectdb.ProjectDB) *Project {
	return &Project{
		projectDB: projectDB,
	}
}
