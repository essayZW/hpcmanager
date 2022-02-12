package db

import "github.com/jmoiron/sqlx"

// PermissionDB 权限表的相关操作
type PermissionDB struct {
	conn *sqlx.DB
}

// QueryIDByLevel 查询权限等级对应的权限ID
func (p *PermissionDB) QueryIDByLevel(level int32) (int, error) {
	row := p.conn.QueryRowx("SELECT `id` FROM permission WHERE `level`=?", level)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// NewPermission 新建一个权限表操作结构体
func NewPermission(conn *sqlx.DB) *PermissionDB {
	return &PermissionDB{
		conn: conn,
	}
}
