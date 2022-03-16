package db

import (
	"github.com/essayZW/hpcmanager/db"
)

// ProjectDB project数据库操作
type ProjectDB struct {
	conn *db.DB
}

// NewProject 创建新的数据库操作结构
func NewProject(sqlConn *db.DB) *ProjectDB {
	return &ProjectDB{
		conn: sqlConn,
	}
}
