package db

import (
	"github.com/essayZW/hpcmanager/config"
	"github.com/jmoiron/sqlx"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// NewDB 创建一个新的数据库连接
func NewDB() (*sqlx.DB, error) {
	dbConfig, err := config.LoadDatabase()
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Open("mysql", dbConfig.Dsn())
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
