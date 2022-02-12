package db

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

// UserPermissionDB 用户权限表操作
type UserPermissionDB struct {
	conn *sqlx.DB
}

// QueryUserPermissionLevel 通过用户ID查询用户的权限所拥有的权限级别信息
func (up *UserPermissionDB) QueryUserPermissionLevel(userid int) ([]*FullUserPermission, error) {
	rows, err := up.conn.Queryx("SELECT `user_permission`.*, `permission`.level "+
		"FROM `user_permission`, `permission` "+
		"WHERE `user_id`=? AND `user_permission`.permission_id=`permission`.id", userid)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	res := make([]*FullUserPermission, 0)
	for rows.Next() {
		var row FullUserPermission
		rows.StructScan(&row)
		res = append(res, &row)
	}
	return res, nil
}

// Insert 插入一条新的用户权限记录
func (up *UserPermissionDB) Insert(info *UserPermission) error {
	res, err := up.conn.Exec("INSERT INTO `user_permission`"+
		"(`user_id`, `user_group_id`, `permissionid`, `create_time`, `extraAttributes`)"+
		"VALUES (?,?,?,?,?)", info.UserID, info.UserGroupID, info.PermissionID, info.CreateTime, info.ExtraAttributes)
	if err != nil {
		return err
	}
	ar, err := res.RowsAffected()
	if err != nil {
		return nil
	}
	if ar == 0 {
		return errors.New("insert has no affected rows")
	}
	return nil
}

// NewUserPermission 创建新的用户权限操作结构体
func NewUserPermission(conn *sqlx.DB) *UserPermissionDB {
	return &UserPermissionDB{
		conn: conn,
	}
}
