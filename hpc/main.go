package main

import (
	"flag"
	"os"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/db"
	hpcdb "github.com/essayZW/hpcmanager/hpc/db"
	"github.com/essayZW/hpcmanager/hpc/logic"
	hpcpb "github.com/essayZW/hpcmanager/hpc/proto"
	"github.com/essayZW/hpcmanager/hpc/service"
	"github.com/essayZW/hpcmanager/hpc/source"
	"github.com/essayZW/hpcmanager/logger"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func init() {
	logger.SetName("hpc")
}

func main() {
	var hpcCmdBaseDir string
	flag.StringVar(&hpcCmdBaseDir, "cmdBaseDir", "", "used to found cmd file")
	flag.Parse()

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
	sqlConn, err := db.NewDB()
	if err != nil {
		logger.Fatal("MySQL conn error: ", err)
	}
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

	env := os.Getenv(hpcmanager.EnvName)
	// TODO: 添加上数据库配置的加载
	hpcSource, err := source.New(
		source.WithCmdBaseDir(hpcCmdBaseDir),
		source.WithDevSource(env == "dev"),
	)
	if err != nil {
		logger.Fatal("hpcSource init error: ", err)
	}

	hpcLogic := logic.NewHpc(hpcSource, hpcdb.NewHpcUser(sqlConn), hpcdb.NewHpcGroup(sqlConn))

	hpcService := service.NewHpc(serviceClient, hpcLogic)
	hpcpb.RegisterHpcHandler(srv.Server(), hpcService)

	srv.Init()
	if err := srv.Run(); err != nil {
		logger.Fatal("Service run error: ", err)
	}
}
