package db

import (
	"encoding/json"
	"fmt"

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

// JSON 数据库JSON类型
type JSON map[string]interface{}

// Scan 数据库json类型
func (j *JSON) Scan(src interface{}) error {
	var source map[string]interface{}
	switch src.(type) {
	case nil:
		source = make(map[string]interface{})
	case []byte:
		if err := json.Unmarshal(src.([]byte), &source); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Invalid type %t for JSON", src)
	}
	*j = JSON(source)
	return nil
}

func (j JSON) String() string {
	res, _ := json.Marshal(j)
	return string(res)
}
