package logic

import (
	"context"
	"os"
	"testing"

	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
	userdb "github.com/essayZW/hpcmanager/user/db"
)

var userGroupLogic *UserGroup

func init() {
	os.Setenv(hpcmanager.EnvName, "testing")
	hpcmanager.LoadCommonArgs()

	// 创建数据库连接
	sqlConn, err := db.NewDB()
	if err != nil {
		logger.Fatal("MySQL conn error: ", err)
	}
	userGroupLogic = NewUserGroup(userdb.NewUserGroup(sqlConn))
}

func TestGetGroupByID(t *testing.T) {
	tests := []struct {
		Name    string
		GroupID int

		ExceptName string
		Error      bool
	}{
		{
			Name:       "test success1",
			GroupID:    1,
			ExceptName: "GROUP1",
			Error:      false,
		},
		{
			Name:       "test success2",
			GroupID:    2,
			ExceptName: "GROUP2",
			Error:      false,
		},
		{
			Name:    "test fail",
			GroupID: 45,
			Error:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res, err := userGroupLogic.GetGroupInfoByID(context.Background(), test.GroupID)
			if err != nil {
				if !test.Error {
					t.Errorf("Except: %v Get: %v", test.Error, err)
				}
				return
			}
			if res.Name != test.ExceptName {
				t.Errorf("Get: %v, Except: %v", res, test.ExceptName)
			}
		})
	}
}
