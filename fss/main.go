package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/essayZW/hpcmanager/config"
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

	srv := micro.NewService(
		micro.Name("fss"),
		micro.Registry(etcdRegistry),
	)

	serviceClient := srv.Client()
	serviceServer := srv.Server()

	fssService := service.NewFss(serviceClient)
	fsspb.RegisterFssServiceHandler(serviceServer, fssService)

	srv.Init()
	if err := srv.Run(); err != nil {
		logger.Fatal("Service run error: ", err)
	}

}
