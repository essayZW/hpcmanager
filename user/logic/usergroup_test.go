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
			if test.Error && err == nil {
				t.Errorf("Except: %v Get: %v", test.Error, err)
			}
			if res.Name != test.ExceptName {
				t.Errorf("Get: %v, Except: %v", res, test.ExceptName)
			}
		})
	}
}

func TestPaginationGetGroupInfo(t *testing.T) {
	tests := []struct {
		Name      string
		PageSize  int
		PageIndex int

		ExceptLen   int
		ExceptCount int
		Error       bool
	}{
		{
			Name:        "test error pageSize",
			PageSize:    -1,
			PageIndex:   10,
			ExceptCount: 0,
			Error:       true,
		},
		{
			Name:        "test error pageIndex",
			PageSize:    1,
			PageIndex:   0,
			ExceptCount: 0,
			Error:       true,
		},
		{
			Name:        "test success",
			PageSize:    1,
			PageIndex:   1,
			ExceptLen:   1,
			ExceptCount: 2,
			Error:       false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			infos, err := userGroupLogic.PaginationGetGroupInfo(context.Background(), test.PageIndex, test.PageSize)
			if err != nil {
				if !test.Error {
					t.Errorf("Get:%v Exceot: %v", err, test.Error)
				}
				return
			}
			if test.Error && err == nil {
				t.Errorf("Except: %v Get: %v", test.Error, err)
			}
			if len(infos.Infos) != test.ExceptLen {
				t.Errorf("Get: %v Except: %v", infos, test.ExceptLen)
			}
			if infos.Count != test.ExceptCount {
				t.Errorf("Get: %v Except: %v", infos, test.ExceptCount)
			}
		})
	}
}
