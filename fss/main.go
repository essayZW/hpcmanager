package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/fss/logic"
	fsspb "github.com/essayZW/hpcmanager/fss/proto"
	"github.com/essayZW/hpcmanager/fss/service"
	"github.com/essayZW/hpcmanager/logger"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
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
	// 读取文件存储路径配置
	configSource, err := config.LoadConfigSource()
	if err != nil {
		logger.Fatal("load config error: ", err)
	}
	var fileStorePath string
	if err := configSource.Get("fileStorePath").Scan(&fileStorePath); err != nil {
		logger.Fatal("load fileStorePath config error: ", err)
	}
	logger.Info("use fileStorePath: ", fileStorePath)

	srv := micro.NewService(
		micro.Name("fss"),
		micro.Registry(etcdRegistry),
	)

	serviceClient := srv.Client()
	serviceServer := srv.Server()

	fssLogic, err := logic.NewFss(fileStorePath)
	if err != nil {
		logger.Fatal(err)
	}

	fssService := service.NewFss(serviceClient, fssLogic)
	fsspb.RegisterFssServiceHandler(serviceServer, fssService)

	srv.Init()
	if err := srv.Run(); err != nil {
		logger.Fatal("Service run error: ", err)
	}

}
