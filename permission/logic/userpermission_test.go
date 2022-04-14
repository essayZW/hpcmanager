package logic

import (
	"context"
	"os"
	"testing"

	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
	permissiondb "github.com/essayZW/hpcmanager/permission/db"
	"github.com/essayZW/hpcmanager/verify"
)

var userPermissionLogic *UserPermission

func init() {
	err := os.Setenv(hpcmanager.EnvName, "testing")

	// 创建数据库连接
	sqldb, err := db.NewDB()
	if err != nil {
		logger.Fatal("MySQL conn error: ", err)
	}
	permissionLogic := NewPermission(permissiondb.NewPermission(sqldb))
	userPermissionLogic = NewUserPermission(permissiondb.NewUserPermission(sqldb), permissionLogic)
}

func TestRemoveUserPermission(t *testing.T) {
	tests := []struct {
		Name   string
		UserID int
		level  verify.Level

		Except bool
	}{
		{
			Name:   "test success",
			UserID: 2,
			level:  verify.Guest,
			Except: true,
		},
		{
			Name:   "test del no exists",
			UserID: 5,
			level:  verify.Common,
			Except: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			err := userPermissionLogic.RemoveUserPermission(
				context.Background(),
				test.UserID,
				test.level,
			)
			if (err != nil) == test.Except {
				t.Errorf("Except %v, Get %v", test.Except, err)
			}
		})
	}
}
