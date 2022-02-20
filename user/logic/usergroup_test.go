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
	userGroupLogic = NewUserGroup(userdb.NewUserGroup(sqlConn), userdb.NewUserGroupApply(sqlConn))
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

func TestCreateUserJoinGroupApply(t *testing.T) {
	tests := []struct {
		Name string

		ApplyGroupID int
		UserInfo     *userdb.User
		Error        bool
	}{
		{
			Name:         "error user groupID",
			ApplyGroupID: 1,
			UserInfo: &userdb.User{
				ID:       1,
				Username: "121121121",
				Name:     "大佬",
				GroupID:  2,
			},
			Error: true,
		},
		{
			Name:         "success",
			ApplyGroupID: 1,
			UserInfo: &userdb.User{
				ID:       20,
				Username: "456456456",
				Name:     "申请组",
			},
			Error: false,
		},
		{
			Name:         "not exists groupid",
			ApplyGroupID: 5,
			UserInfo: &userdb.User{
				ID:       20,
				Username: "456456456",
				Name:     "申请组",
			},
			Error: true,
		},
		{
			Name:         "repeated apply",
			ApplyGroupID: 1,
			UserInfo: &userdb.User{
				ID:       20,
				Username: "456456456",
				Name:     "申请组",
			},
			Error: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			_, err := userGroupLogic.CreateUserJoinGroupApply(context.Background(), test.UserInfo, test.ApplyGroupID)
			if err != nil {
				if !test.Error {
					t.Errorf("Get: %v Except: %v", err, test.Error)
				}
				return
			}
			if test.Error {
				t.Errorf("Get: %v Except: %v", err, test.Error)
				return

			}
		})
	}
}

func TestGetByTutorUsername(t *testing.T) {
	tests := []struct {
		Name string

		Username string

		Error         bool
		ExceptGroupID int
	}{
		{
			Name:          "success",
			Username:      "123123123",
			ExceptGroupID: 1,
			Error:         false,
		},
		{
			Name:     "does not exists",
			Username: "no",
			Error:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			group, err := userGroupLogic.GetByTutorUsername(context.Background(), test.Username)
			if err != nil {
				if !test.Error {
					t.Errorf("Get: %v Except: %v", err, test.Error)
				}
				return
			}
			if test.Error {
				t.Errorf("Get: %v Except: %v", err, test.Error)
				return
			}
			if test.ExceptGroupID != group.ID {
				t.Errorf("Get: %v, Except: %v", group, test.ExceptGroupID)
			}
		})
	}
}

func TestTutorCheckApply(t *testing.T) {
	tests := []struct {
		Name string

		ApplyID      int
		CheckStatus  bool
		TutorID      int
		CheckMessage string

		Error bool
	}{
		{
			Name:         "test success",
			ApplyID:      9,
			CheckStatus:  true,
			TutorID:      2,
			CheckMessage: "测试审核通过",
			Error:        false,
		},
		{
			Name:         "test success2",
			ApplyID:      10,
			CheckStatus:  false,
			TutorID:      2,
			CheckMessage: "测试审核不通过",
			Error:        false,
		},
		{
			Name:         "test fail",
			ApplyID:      11,
			CheckStatus:  true,
			TutorID:      2,
			CheckMessage: "审核通过",
			Error:        true,
		},
		{
			Name:         "test repeated check",
			ApplyID:      1,
			CheckStatus:  true,
			TutorID:      2,
			CheckMessage: "审核通过",
			Error:        true,
		},
		{
			Name:         "test other group apply",
			ApplyID:      12,
			CheckStatus:  false,
			CheckMessage: "测试消息",
			Error:        true,
		},
		{
			Name:    "test invalid apply id",
			ApplyID: 99,
			Error:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			_, err := userGroupLogic.TutorCheckApply(context.Background(), test.TutorID, test.ApplyID, test.CheckStatus, test.CheckMessage)
			if err != nil {
				if !test.Error {
					t.Errorf("Get: %v, Except: %v", err, test.Error)
				}
				return
			}
			if test.Error {
				t.Errorf("Get: %v, Except: %v", err, test.Error)
			}
		})
	}
}

func TestAdminCheckApply(t *testing.T) {
	tests := []struct {
		Name string

		ApplyID         int
		CheckStatus     bool
		CheckMessage    string
		CheckerID       int
		CheckerUsername string
		CheckerName     string

		Error bool
	}{
		{
			Name:            "test check success",
			ApplyID:         9,
			CheckStatus:     true,
			CheckMessage:    "管理员审核通过",
			CheckerID:       1,
			CheckerUsername: "121121121",
			CheckerName:     "管理员",
			Error:           false,
		},
		{
			Name:            "check tutor check not accept",
			ApplyID:         10,
			CheckStatus:     true,
			CheckMessage:    "管理员审核通过",
			CheckerID:       1,
			CheckerUsername: "121121121",
			CheckerName:     "管理员",
			Error:           true,
		},
		{
			Name:            "check tutor not check",
			ApplyID:         1,
			CheckStatus:     true,
			CheckMessage:    "管理员审核通过",
			CheckerID:       1,
			CheckerUsername: "121121121",
			Error:           true,
			CheckerName:     "管理员",
		},
		{
			Name:            "repeated check",
			ApplyID:         9,
			CheckStatus:     true,
			CheckMessage:    "管理员审核通过",
			CheckerID:       1,
			CheckerUsername: "121121121",
			CheckerName:     "管理员",
			Error:           true,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			_, err := userGroupLogic.AdminCheckApply(context.Background(), test.ApplyID, test.CheckerID, test.CheckerUsername, test.CheckerName, test.CheckStatus, test.CheckMessage)
			if err != nil {
				if !test.Error {
					t.Errorf("Get: %v, Except: %v", err, test.Error)
				}
				return
			}
			if test.Error {
				t.Errorf("Get: %v, Except: %v", err, test.Error)
			}
		})
	}
}
