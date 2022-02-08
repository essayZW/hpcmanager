package logic

import (
	"context"
	"os"
	"testing"

	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
	userdb "github.com/essayZW/hpcmanager/user/db"
	"github.com/go-redis/redis/v8"
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
	userLogic = NewUser(userdb.New(sqlConn), etcdConfig, redisConn)
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
			status, err := userLogic.LoginCheck(test.Username, test.Password)
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
			token := userLogic.CreateToken(test.Username)
			if queryed := userLogic.QueryToken(test.Username); queryed != token {
				t.Errorf("Except %s Get %s", queryed, token)
			}
			userLogic.DeleteToken(test.Username)
			if queryed := userLogic.QueryToken(test.Username); queryed != "" {
				t.Errorf("Except %s Get ", queryed)
			}
		})
	}
}
