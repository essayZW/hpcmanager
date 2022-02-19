package service

import (
	"context"
	"os"
	"testing"

	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/db"
	gateway "github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/logger"
	userdb "github.com/essayZW/hpcmanager/user/db"
	"github.com/essayZW/hpcmanager/user/logic"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/essayZW/hpcmanager/verify"
	"github.com/go-redis/redis/v8"
)

var userGroupService *UserGroupService

func init() {
	os.Setenv(hpcmanager.EnvName, "testing")

	// 创建数据库连接
	sqlConn, err := db.NewDB()
	if err != nil {
		logger.Fatal("MySQL conn error: ", err)
	}
	// 创建动态配置源
	etcdConfig, err := config.NewEtcd()
	if err != nil {
		logger.Fatal("Etcd config create error: ", err)
	}
	// 创建redis连接
	redisConfig, err := config.LoadRedis()
	if err != nil {
		logger.Fatal("Redis conn error: ", err)
	}
	redisConn := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	ping := redisConn.Ping(context.Background())
	ok, err := ping.Result()
	if err != nil {
		logger.Fatal("Redis ping error: ", err)
	}
	if ok != "PONG" {
		logger.Fatal("Redis ping get: ", ok)
	}
	userLogic := logic.NewUser(userdb.NewUser(sqlConn), etcdConfig, redisConn)
	userGroupLogic := logic.NewUserGroup(userdb.NewUserGroup(sqlConn), userdb.NewUserGroupApply(sqlConn))
	userGroupService = &UserGroupService{
		userGroupLogic: userGroupLogic,
		userLogic:      userLogic,
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
			if test.Error && err == nil {
				t.Errorf("Except: %v Get: %v", test.Error, err)
				return
			}
			if resp.GroupInfo.Name != test.ExceptName {
				t.Errorf("Except: %v Get: %v", test.Error, resp.GroupInfo)
			}
		})
	}
}

func TestPaginationGetGroupInfo(t *testing.T) {
	tests := []struct {
		Name       string
		PageSize   int
		PageIndex  int
		UserLevels []int32

		ExceptLen   int
		ExceptCount int
		Error       bool
	}{
		{
			Name:      "test error pageSize",
			PageSize:  -1,
			PageIndex: 10,
			UserLevels: []int32{
				int32(verify.CommonAdmin),
			},
			ExceptCount: 0,
			Error:       true,
		},
		{
			Name:        "test error pageIndex",
			PageSize:    1,
			PageIndex:   0,
			ExceptCount: 0,
			UserLevels: []int32{
				int32(verify.CommonAdmin),
			},
			Error: true,
		},
		{
			Name:        "test success",
			PageSize:    1,
			PageIndex:   1,
			ExceptLen:   1,
			ExceptCount: 2,
			UserLevels: []int32{
				int32(verify.CommonAdmin),
			},
			Error: false,
		},
		{
			Name:        "test permission forbidden",
			PageSize:    1,
			PageIndex:   1,
			ExceptCount: 0,
			UserLevels: []int32{
				int32(verify.Common),
			},
			Error: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var request userpb.PaginationGetGroupInfoRequest
			request.BaseRequest = baseRequest
			request.BaseRequest.UserInfo.Levels = test.UserLevels
			request.PageIndex = int32(test.PageIndex)
			request.PageSize = int32(test.PageSize)
			var resp userpb.PaginationGetGroupInfoResponse

			err := userGroupService.PaginationGetGroupInfo(context.Background(), &request, &resp)
			if err != nil {
				if !test.Error {
					t.Errorf("Get: %v, Except: %v", err, test.Error)
				}
			}
			if test.Error && err == nil {
				t.Errorf("Except: %v Get: %v", test.Error, err)
				return
			}
			if len(resp.GroupInfos) != test.ExceptLen {
				t.Errorf("Get:%v ExceptCount: %v", resp.GroupInfos, test.ExceptLen)
			}
			if resp.Count != int32(test.ExceptCount) {
				t.Errorf("Get:%v ExceptCount: %v", resp.Count, test.ExceptCount)
			}
		})
	}
}

func TestCreateJoinGroupApply(t *testing.T) {
	tests := []struct {
		Name string

		UserID     int
		UserLevels []int32

		ApplyGroupID int

		Error bool
	}{
		{
			Name:   "permission forbidden",
			UserID: 21,
			UserLevels: []int32{
				int32(verify.Common),
			},
			ApplyGroupID: 1,
			Error:        true,
		},
		{
			Name:   "already has group",
			UserID: 1,
			UserLevels: []int32{
				int32(verify.Guest),
			},
			ApplyGroupID: 1,
			Error:        true,
		},
		{
			Name:   "test success",
			UserID: 21,
			UserLevels: []int32{
				int32(verify.Guest),
			},
			ApplyGroupID: 2,
			Error:        false,
		},
		{
			Name:   "test repeaded apply",
			UserID: 21,
			UserLevels: []int32{
				int32(verify.Guest),
			},
			ApplyGroupID: 2,
			Error:        true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var req userpb.CreateJoinGroupApplyRequest
			req.BaseRequest = baseRequest
			req.BaseRequest.UserInfo.UserId = int32(test.UserID)
			req.BaseRequest.UserInfo.Levels = test.UserLevels
			req.ApplyGroupID = int32(test.ApplyGroupID)
			var resp userpb.CreateJoinGroupApplyResponse
			err := userGroupService.CreateJoinGroupApply(context.Background(), &req, &resp)
			if err != nil {
				if !test.Error {
					t.Errorf("Except: %v Get: %v", test.Error, err)
				}
				return
			}
			if test.Error {
				t.Errorf("Except: %v Get: %v", test.Error, err)
				return
			}
		})
	}
}
