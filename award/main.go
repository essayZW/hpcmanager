package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	awardpb "github.com/essayZW/hpcmanager/award/proto"
	"github.com/essayZW/hpcmanager/award/service"
	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/logger"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func init() {
	logger.SetName("award")
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
		micro.Name("award"),
		micro.Registry(etcdRegistry),
	)

	serviceClient := srv.Client()
	serviceServer := srv.Server()

	awardService := service.NewAward(serviceClient)
	awardpb.RegisterAwardServiceHandler(serviceServer, awardService)

	srv.Init()
	if err := srv.Run(); err != nil {
		logger.Fatal("Service run error: ", err)
	}
}
