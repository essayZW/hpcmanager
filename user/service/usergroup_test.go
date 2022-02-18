package service

import (
	"context"
	"os"
	"testing"

	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/db"
	gateway "github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/logger"
	userdb "github.com/essayZW/hpcmanager/user/db"
	"github.com/essayZW/hpcmanager/user/logic"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/essayZW/hpcmanager/verify"
)

var userGroupService *UserGroupService

func init() {
	os.Setenv(hpcmanager.EnvName, "testing")

	// 创建数据库连接
	sqlConn, err := db.NewDB()
	if err != nil {
		logger.Fatal("MySQL conn error: ", err)
	}
	userGroupLogic := logic.NewUserGroup(userdb.NewUserGroup(sqlConn))
	userGroupService = &UserGroupService{
		userGroupLogic: userGroupLogic,
	}
	baseRequest = &gateway.BaseRequest{
		RequestInfo: &gateway.RequestInfo{
			Id: "testing",
		},
		UserInfo: &gateway.UserInfo{
			UserId: 0,
			Levels: []int32{
				int32(verify.SuperAdmin),
				int32(verify.CommonAdmin),
				int32(verify.Tutor),
				int32(verify.Common),
				int32(verify.Guest),
			},
		},
	}
}

func TestGetGroupInfoByID(t *testing.T) {
	tests := []struct {
		Name    string
		GroupID int

		ExceptName string
		Error      bool

		UserGroupID int
		UserLevels  []int32
	}{
		{
			Name:        "test admin success",
			GroupID:     1,
			ExceptName:  "GROUP1",
			Error:       false,
			UserGroupID: 1,
			UserLevels: []int32{
				int32(verify.CommonAdmin),
			},
		},
		{
			Name:        "test tutor success",
			GroupID:     1,
			ExceptName:  "GROUP1",
			Error:       false,
			UserGroupID: 1,
			UserLevels: []int32{
				int32(verify.Tutor),
			},
		},
		{
			Name:        "test tutor error",
			GroupID:     1,
			Error:       true,
			UserGroupID: 2,
			UserLevels: []int32{
				int32(verify.Tutor),
			},
		},
		{
			Name:        "test admin success2",
			GroupID:     1,
			ExceptName:  "GROUP1",
			UserGroupID: 2,
			Error:       false,
			UserLevels: []int32{
				int32(verify.SuperAdmin),
			},
		},
		{
			Name:    "test permission forbidden",
			GroupID: 1,
			Error:   true,
			UserLevels: []int32{
				int32(verify.Common),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var request userpb.GetGroupInfoByIDRequest
			request.BaseRequest = baseRequest
			request.BaseRequest.UserInfo.Levels = test.UserLevels
			request.BaseRequest.UserInfo.GroupId = int32(test.UserGroupID)
			request.GroupID = int32(test.GroupID)
			var resp userpb.GetGroupInfoByIDResponse
			err := userGroupService.GetGroupInfoByID(context.Background(), &request, &resp)
			if err != nil {
				if !test.Error {
					t.Errorf("Except: %v Get: %v", test.Error, err)
				}
				return
			}
			if resp.GroupInfo.Name != test.ExceptName {
				t.Errorf("Except: %v Get: %v", test.Error, resp.GroupInfo)
			}
		})
	}
}
