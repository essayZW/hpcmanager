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
	"github.com/essayZW/hpcmanager/verify"
	"github.com/go-redis/redis/v8"
)

var userService *UserService
var baseRequest *gateway.BaseRequest

func init() {
	os.Setenv(hpcmanager.EnvName, "testing")
	hpcmanager.LoadCommonArgs()

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
	userpLogic := logic.NewUserPermission(userdb.NewUserPermission(sqlConn))
	userService = &UserService{
		userLogic:  userLogic,
		userpLogic: userpLogic,
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
				Username: test.Username,
				Password: test.Password,
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
