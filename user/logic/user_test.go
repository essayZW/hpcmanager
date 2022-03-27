package logic

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
	userdb "github.com/essayZW/hpcmanager/user/db"
	"github.com/go-redis/redis/v8"
	"github.com/mozillazg/go-pinyin"
)

var userLogic *User

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
	userLogic = NewUser(userdb.NewUser(sqlConn), etcdConfig, redisConn)
}

func getUserDB() *userdb.UserDB {
	// 创建数据库连接
	sqlConn, err := db.NewDB()
	if err != nil {
		logger.Fatal("MySQL conn error: ", err)
	}
	return userdb.NewUser(sqlConn)
}

func TestLoginCheck(t *testing.T) {
	tests := []struct {
		Name     string
		Username string
		Password string
		Except   bool
	}{
		{
			Name:     "test login success",
			Username: "123123123",
			Password: "essay",
			Except:   true,
		},
		{
			Name:     "test login fail",
			Username: "121121121",
			Password: "error",
			Except:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			status, err := userLogic.LoginCheck(context.Background(), test.Username, test.Password)
			if err != nil && test.Except {
				t.Error(err)
			}
			if status != test.Except {
				t.Errorf("Except %v Get %v", test.Except, status)
			}
		})
	}
}

func TestCreateAndQueryToken(t *testing.T) {
	tests := []struct {
		Name     string
		Username string
	}{
		{
			Name:     "test1",
			Username: "1234567890",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			token := userLogic.CreateToken(context.Background(), test.Username)
			if queryed := userLogic.GetToken(context.Background(), test.Username); queryed != token {
				t.Errorf("Except %s Get %s", queryed, token)
			}
			userLogic.DeleteToken(context.Background(), test.Username)
			if queryed := userLogic.GetToken(context.Background(), test.Username); queryed != "" {
				t.Errorf("Except %s Get ", queryed)
			}
		})
	}
}

func TestQueryByUsername(t *testing.T) {
	tests := []struct {
		Name     string
		Username string

		Except int
		Error  bool
	}{
		{
			Name:     "test not exists",
			Username: "1234567890",
			Error:    true,
		},
		{
			Name:     "test id 2",
			Username: "123123123",
			Error:    false,
			Except:   2,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			info, err := userLogic.GetByUsername(context.Background(), test.Username)
			if err != nil {
				if !test.Error {
					t.Error(err)
				}
				return
			}
			if info.ID != test.Except {
				t.Errorf("Except: %d Get %v", test.Except, info)
			}
		})
	}
}

func TestAddUser(t *testing.T) {
	tests := []struct {
		Name  string
		Data  *userdb.User
		Error bool
	}{
		{
			Name: "test add success",
			Data: &userdb.User{
				Username:   "999999999",
				Password:   "testing",
				Name:       "大佬",
				CreateTime: time.Now(),
				ExtraAttributes: &db.JSON{
					"testing": true,
				},
			},
			Error: false,
		},
		{
			Name: "test add fail",
			Data: &userdb.User{
				Username:   "123123123",
				Password:   "testing",
				Name:       "大佬",
				CreateTime: time.Now(),
			},
			Error: true,
		},
		{
			Name: "test add fail2",
			Data: &userdb.User{
				Username:   "888888888",
				Password:   "testing",
				CreateTime: time.Now(),
			},
			Error: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			_, err := userLogic.AddUser(context.Background(), test.Data)
			if (err != nil) != test.Error {
				t.Errorf("Except: %v Get %v", test.Error, err)
			}
		})
	}

}

func TestPaginationGetUserInfo(t *testing.T) {
	tests := []struct {
		Name string

		PageSize  int
		PageIndex int

		ExceptLen   int
		ExceptCount int
		Error       bool
	}{
		{
			Name:        "error pageSize",
			PageSize:    0,
			PageIndex:   1,
			ExceptCount: 0,
			Error:       true,
		},
		{
			Name:        "error pageIndex",
			PageSize:    1,
			PageIndex:   0,
			ExceptCount: 0,
			Error:       true,
		},
		{
			Name:        "success",
			PageSize:    1,
			PageIndex:   2,
			ExceptLen:   1,
			ExceptCount: 2,
			Error:       false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			infos, err := userLogic.PaginationGetUserInfo(context.Background(), test.PageIndex, test.PageSize, 0)
			if err != nil {
				if !test.Error {
					t.Errorf("Get: %v Except: %v", err, test.Error)
				}
				return
			}
			if test.Error && err == nil {
				t.Errorf("Get: %v Except: %v", err, test.Error)
			}
			if test.ExceptLen != len(infos.Infos) {
				t.Errorf("Get: %v Except: %v", infos, test.ExceptCount)
			}
			if test.ExceptCount != infos.Count {
				t.Errorf("Get: %v Except: %v", infos, test.ExceptCount)
			}
		})
	}
}

func TestChangeUserGroup(t *testing.T) {
	tests := []struct {
		Name string

		UserID  int
		GroupID int

		Error bool
	}{
		{
			Name:    "success1",
			UserID:  22,
			GroupID: 2,
			Error:   false,
		},
		{
			Name:    "success2",
			UserID:  22,
			GroupID: 0,
			Error:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			err := userLogic.ChangeUserGroup(context.Background(), test.UserID, test.GroupID)
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

func TestPinyinLib(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string

		args args
		want string
	}{
		{
			name: "success1",
			args: args{
				name: "测试",
			},
			want: "ceshi",
		},
		{
			name: "success1",
			args: args{
				name: "大佬",
			},
			want: "dalao",
		},
		{
			name: "success1",
			args: args{
				name: "大佬1",
			},
			want: "dalao1",
		},
		{
			name: "success1",
			args: args{
				name: "1大佬",
			},
			want: "1dalao",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pinyin.Fallback = func(r rune, a pinyin.Args) []string {
				return []string{string(r)}
			}
			got := pinyin.LazyPinyin(tt.args.name, pinyin.NewArgs())
			if strings.Join(got, "") != tt.want {
				t.Errorf("Want=%s, Got=%s", tt.want, got)
			}
		})
	}
}

func TestUser_addSuffixForPYName(t *testing.T) {
	type fields struct {
		userDB *userdb.UserDB
	}
	type args struct {
		ctx    context.Context
		pyName string
	}
	userDB := getUserDB()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test dalao",
			fields: fields{
				userDB: userDB,
			},
			args: args{
				ctx:    context.Background(),
				pyName: "dalao",
			},
			want:    "dalao3",
			wantErr: false,
		},
		{
			name: "test none",
			fields: fields{
				userDB: userDB,
			},
			args: args{
				ctx:    context.Background(),
				pyName: "none",
			},
			want:    "none",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				userDB: tt.fields.userDB,
			}
			got, err := u.addSuffixForPYName(tt.args.ctx, tt.args.pyName)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.addSuffixForPYName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("User.addSuffixForPYName() = %v, want %v", got, tt.want)
			}
		})
	}
}
