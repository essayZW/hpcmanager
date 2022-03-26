package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/logger"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func init() {
	logger.SetName("crontab")
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
}
