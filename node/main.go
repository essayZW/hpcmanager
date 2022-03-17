package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/logger"
	nodepb "github.com/essayZW/hpcmanager/node/proto"
	"github.com/essayZW/hpcmanager/node/service"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func init() {
	logger.SetName("node")
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
		micro.Name("node"),
		micro.Registry(etcdRegistry),
	)

	serviceClient := srv.Client()

	// 创建数据库连接
	// sqldb, err := db.NewDB()
	// if err != nil {
	// logger.Fatal("MySQL conn error: ", err)
	// }

	nodeService := service.NewNode(serviceClient)
	nodepb.RegisterNodeHandler(srv.Server(), nodeService)

	srv.Init()
	if err := srv.Run(); err != nil {
		logger.Fatal("Service run error: ", err)
	}
}
