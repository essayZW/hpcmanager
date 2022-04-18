package main

import (
	"context"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	etcdsync "github.com/asim/go-micro/plugins/sync/etcd/v3"
	hpcbroker "github.com/essayZW/hpcmanager/broker"
	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
	userbroker "github.com/essayZW/hpcmanager/user/broker"
	userdb "github.com/essayZW/hpcmanager/user/db"
	"github.com/essayZW/hpcmanager/user/logic"
	user "github.com/essayZW/hpcmanager/user/proto"
	"github.com/essayZW/hpcmanager/user/service"
	"github.com/go-redis/redis/v8"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/sync"
)

func init() {
	logger.SetName("user")
}

func main() {
	registryConf, err := config.LoadRegistry()
	if err != nil {
		logger.Fatal("load etcd config error: ", nil)
	}
	etcdRegistry := etcd.NewRegistry(
		registry.Addrs(registryConf.Etcd.Address),
	)

	srv := micro.NewService(
		micro.Name("user"),
		micro.Registry(etcdRegistry),
	)

	serviceClient := srv.Client()

	// 创建数据库连接
	sqldb, err := db.NewDB()
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

	// 初始化分布式锁
	etcdSync := etcdsync.NewSync(
		sync.Nodes(registryConf.Etcd.Address+":2379"),
		sync.Prefix("user"),
	)
	etcdSync.Init()

	rabbitmqBroker, err := hpcbroker.NewRabbitmq()
	if err != nil {
		logger.Fatal(err)
	}

	// 注册消费者
	go userbroker.RegistryCustomer(rabbitmqBroker, serviceClient)

	userLogic := logic.NewUser(userdb.NewUser(sqldb), etcdConfig, redisConn)
	userGroupLogic := logic.NewUserGroup(
		userdb.NewUserGroup(sqldb),
		userdb.NewUserGroupApply(sqldb),
		etcdSync,
	)

	serviceServer := srv.Server()

	userService := service.NewUser(serviceClient, userLogic, userGroupLogic, rabbitmqBroker)
	user.RegisterUserHandler(serviceServer, userService)

	userGroupService := service.NewGroup(serviceClient, userGroupLogic, userLogic, rabbitmqBroker)
	user.RegisterGroupServiceHandler(serviceServer, userGroupService)

	srv.Init()
	if err := srv.Run(); err != nil {
		logger.Fatal("Service run error: ", err)
	}

}
