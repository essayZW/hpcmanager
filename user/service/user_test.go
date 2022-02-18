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
	user "github.com/essayZW/hpcmanager/user/proto"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/essayZW/hpcmanager/verify"
	"github.com/go-redis/redis/v8"
)

var userService *UserService
var baseRequest *gateway.BaseRequest

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
	userService = &UserService{
		userLogic: userLogic,
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

func TestLogin(t *testing.T) {
	tests := []struct {
		Name     string
		Username string
		Password string
		Except   *user.UserInfo
		Error    bool
	}{
		{
			Name:     "test success",
			Username: "123123123",
			Password: "essay",
			Except: &user.UserInfo{
				Username: "123123123",
				Id:       2,
				Name:     "测试",
			},
			Error: false,
		},
		{
			Name:     "test success2",
			Username: "121121121",
			Password: "essay",
			Except: &user.UserInfo{
				Username: "121121121",
				Id:       1,
				Name:     "测试",
			},
			Error: false,
		},
		{
			Name:     "test invalid password",
			Username: "123123123",
			Password: "error",
			Except:   &user.UserInfo{},
			Error:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var res user.LoginResponse
			err := userService.Login(context.Background(), &user.LoginRequest{
				BaseRequest: baseRequest,
				Username:    test.Username,
				Password:    test.Password,
			}, &res)
			if err != nil && !test.Error {
				t.Error(err)
				return
			}
			if res.UserInfo.GetId() != test.Except.Id {
				t.Errorf("Except %v, Get %v", test.Except, res.UserInfo)
			}
		})
	}
}

func TestCheckLogin(t *testing.T) {
	tests := []struct {
		Name string

		Token  string
		Except []int32
		Error  bool
	}{
		// NOTE: token会过期,单元测试之间需要重新生成最新的token
		{
			Name:   "test1",
			Token:  "49b8d8b2691144cf86ef58670229422d",
			Except: []int32{0, 1},
		},
		{
			Name:   "test2",
			Token:  "45c6053e75954184a684d6e26b44bdcf",
			Except: []int32{0},
		},
		{
			Name:   "test3",
			Token:  "",
			Except: []int32{},
			Error:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res := &user.CheckLoginResponse{}
			err := userService.CheckLogin(context.Background(), &user.CheckLoginRequest{
				Token:       test.Token,
				BaseRequest: baseRequest,
			}, res)
			if err != nil {
				if !test.Error {
					t.Error(err)
				}
				return
			}
			for index, p := range res.PermissionLevel {
				if p != test.Except[index] {
					t.Errorf("Except %v Get %v", test.Except, p)
					break
				}
			}
		})
	}
}

func TestExistUsername(t *testing.T) {
	tests := []struct {
		Name     string
		Username string
		Except   bool
	}{
		{
			Name:     "test not exists",
			Username: "no",
			Except:   false,
		},
		{
			Name:     "test exists",
			Username: "121121121",
			Except:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var res user.ExistUsernameResponse
			err := userService.ExistUsername(context.Background(), &user.ExistUsernameRequest{
				Username:    test.Username,
				BaseRequest: baseRequest,
			}, &res)
			if err != nil || res.Exist != test.Except {
				t.Errorf("Error %v Except %v Get %v", err, test.Except, res.Exist)
			}
		})
	}
}

func TestAddUser(t *testing.T) {
	tests := []struct {
		Name  string
		Data  *user.UserInfo
		Error bool
	}{
		{
			Name: "test add success",
			Data: &user.UserInfo{
				Username: "777777777",
				Password: "testing",
				Name:     "大佬",
			},
			Error: false,
		},
		{
			Name: "test add success2",
			Data: &user.UserInfo{
				Username:        "666666666",
				Password:        "testing",
				Name:            "大佬",
				ExtraAttributes: "{}",
			},
			Error: false,
		},
		{
			Name: "test add fail",
			Data: &user.UserInfo{
				Username: "123123123",
				Password: "testing",
				Name:     "大佬",
			},
			Error: true,
		},
		{
			Name: "test add fail2",
			Data: &user.UserInfo{
				Username: "888888888",
				Password: "testing",
			},
			Error: true,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var resp user.AddUserResponse
			err := userService.AddUser(context.Background(), &user.AddUserRequest{
				BaseRequest: baseRequest,
				UserInfo:    test.Data,
			}, &resp)
			if (err != nil) != test.Error {
				t.Errorf("Except: %v Get %v", test.Error, err)
			}
		})
	}
}

func TestCreateToken(t *testing.T) {
	tests := []struct {
		Name string

		Username string
		Error    bool
	}{
		{
			Name:     "test createToken success",
			Username: "121121121",
			Error:    false,
		},
		{
			Name:     "test createToken fail",
			Username: "nouser",
			Error:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var resp user.CreateTokenResponse
			err := userService.CreateToken(context.Background(), &user.CreateTokenRequest{
				BaseRequest: baseRequest,
				Username:    test.Username,
			}, &resp)
			if (err != nil) != test.Error {
				t.Errorf("Except %v Get %v UserInfo %v Token %v", test.Error, err, resp.UserInfo, resp.Token)
			}
		})
	}
}

func TestGetUserInfo(t *testing.T) {
	tests := []struct {
		Name        string
		UserID      int
		GroupID     int
		QueryUserID int
		Levels      []int32

		Error bool
	}{
		{
			Name:        "Query self success",
			UserID:      2,
			QueryUserID: 2,
			Levels: []int32{
				int32(verify.Common),
			},
			Error: false,
		},
		{
			Name:        "Query other user forbidden",
			UserID:      43,
			QueryUserID: 2,
			Levels: []int32{
				int32(verify.Common),
			},
			Error: true,
		},
		{
			Name:        "Tutor Query other user",
			UserID:      1,
			QueryUserID: 2,
			GroupID:     1,
			Levels: []int32{
				int32(verify.Tutor),
			},
			Error: false,
		},
		{
			Name:        "Tutor query other group user",
			UserID:      1,
			QueryUserID: 43,
			GroupID:     3,
			Levels: []int32{
				int32(verify.Tutor),
			},
			Error: true,
		},
		{
			Name:        "Admin query other user",
			UserID:      1,
			QueryUserID: 2,
			Levels: []int32{
				int32(verify.CommonAdmin),
			},
			Error: false,
		},
		{
			Name:        "Admin query other user2",
			UserID:      43,
			QueryUserID: 2,
			Levels: []int32{
				int32(verify.CommonAdmin),
			},
			Error: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			req := &user.GetUserInfoRequest{
				BaseRequest: baseRequest,
				Userid:      int32(test.QueryUserID),
			}
			req.BaseRequest.UserInfo.Levels = test.Levels
			req.BaseRequest.UserInfo.UserId = int32(test.UserID)
			req.BaseRequest.UserInfo.GroupId = int32(test.GroupID)
			var resp user.GetUserInfoResponse
			err := userService.GetUserInfo(context.Background(), req, &resp)
			if (err != nil) != test.Error {
				t.Errorf("Except: %v Get: %v Resp: %v", test.Error, err, resp.UserInfo)
			}
		})
	}
}

func TestPaginationGetUserInfo(t *testing.T) {
	tests := []struct {
		Name string

		PageIndex int
		PageSize  int

		UserGroupID int
		UserLevels  []int32

		ExceptCount int
		ExceptLen   int
		Error       bool
	}{
		{
			Name: "permission forbidden",
			UserLevels: []int32{
				int32(verify.Common),
			},
			Error: true,
		},
		{
			Name:        "tutor success",
			PageIndex:   1,
			PageSize:    1,
			UserGroupID: 1,
			UserLevels: []int32{
				int32(verify.Tutor),
			},
			ExceptCount: 1,
			ExceptLen:   1,
		},
		{
			Name:        "tutor success2",
			PageIndex:   2,
			PageSize:    1,
			UserGroupID: 1,
			UserLevels: []int32{
				int32(verify.Tutor),
			},
			ExceptCount: 1,
			ExceptLen:   0,
			Error:       false,
		},
		{
			Name:        "admin success",
			PageIndex:   1,
			PageSize:    1,
			UserGroupID: 1,
			UserLevels: []int32{
				int32(verify.CommonAdmin),
			},
			ExceptCount: 2,
			ExceptLen:   1,
			Error:       false,
		},
		{
			Name:        "admin success2",
			PageIndex:   1,
			PageSize:    2,
			UserGroupID: 1,
			UserLevels: []int32{
				int32(verify.CommonAdmin),
			},
			ExceptCount: 2,
			ExceptLen:   2,
			Error:       false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var req userpb.PaginationGetUserInfoRequest
			req.BaseRequest.UserInfo.Levels = test.UserLevels
			req.BaseRequest.UserInfo.GroupId = int32(test.UserGroupID)
			req.PageIndex = int32(test.PageIndex)
			req.PageSize = int32(test.PageSize)
			var resp userpb.PaginationGetUserInfoResponse
			err := userService.PaginationGetUserInfo(context.Background(), &req, &resp)
			if err != nil {
				if !test.Error {
					t.Errorf("Get: %v, Except: %v", err, test.Error)
				}
			}
			if test.Error && err == nil {
				t.Errorf("Except: %v Get: %v", test.Error, err)
				return
			}
			if len(resp.UserInfos) != test.ExceptLen {
				t.Errorf("Get: %v Except: %v", resp.UserInfos, test.ExceptLen)
			}
			if resp.Count != int32(test.ExceptCount) {
				t.Errorf("Get:%v ExceptCount: %v", resp.Count, test.ExceptCount)
			}
		})
	}
}
