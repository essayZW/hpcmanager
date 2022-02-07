package config

import "fmt"

// Database Mysql数据库的配置
type Database struct {
	// Host 数据库服务地址
	Host string `json:"host"`
	// Database 数据库民
	Database string `json:"database"`
	// Port 数据库服务的端口
	Port int `json:"port"`
	// Username 数据库服务的连接用户名
	Username string `json:"username"`
	// Password 数据库服务的连接密码
	Password string `json:"password"`
}

// Dsn 返回对应的Dsn
func (db *Database) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4,utf8&parseTime=true&loc=Asia%%2fShanghai", db.Username, db.Password, db.Host, db.Port, db.Database)
}
