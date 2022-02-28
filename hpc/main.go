package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/hpc/logic"
	hpcpb "github.com/essayZW/hpcmanager/hpc/proto"
	"github.com/essayZW/hpcmanager/hpc/service"
	"github.com/essayZW/hpcmanager/logger"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func init() {
	logger.SetName("hpc")
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
		micro.Name("hpc"),
		micro.Registry(etcdRegistry),
	)

	serviceClient := srv.Client()

	// 创建数据库连接
	//sqlConn, err := db.NewDB()
	//if err != nil {
	//logger.Fatal("MySQL conn error: ", err)
	//}
	// 创建动态配置源
	//etcdConfig, err := config.NewEtcd()
	//if err != nil {
	//logger.Fatal("Etcd config create error: ", err)
	//}
	// 创建redis连接
	//redisConfig, err := config.LoadRedis()
	//if err != nil {
	//logger.Fatal("Redis conn error: ", err)
	//}
	//redisConn := redis.NewClient(&redis.Options{
	//Addr:     redisConfig.Address,
	//Password: redisConfig.Password,
	//DB:       redisConfig.DB,
	//})
	//ping := redisConn.Ping(context.Background())
	//ok, err := ping.Result()
	//if err != nil {
	//logger.Fatal("Redis ping error: ", err)
	//}
	//if ok != "PONG" {
	//logger.Fatal("Redis ping get: ", ok)
	//}

	hpcLogic := logic.NewHpc()

	hpcService := service.NewHpc(serviceClient, hpcLogic)
	hpcpb.RegisterHpcHandler(srv.Server(), hpcService)

	srv.Init()
	if err := srv.Run(); err != nil {
		logger.Fatal("Service run error: ", err)
	}
}
