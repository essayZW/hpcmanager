package logic

import (
	projectdb "github.com/essayZW/hpcmanager/project/db"
)

// Project 项目logic
type Project struct {
	projectDB *projectdb.ProjectDB
}

// NewProject 创建新的project的logic
func NewProject(projectDB *projectdb.ProjectDB) *Project {
	return &Project{
		projectDB: projectDB,
	}
}
