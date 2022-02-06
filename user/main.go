package main

import (
	"flag"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/logger"
	user "github.com/essayZW/hpcmanager/user/proto"
	"github.com/essayZW/hpcmanager/user/service"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func init() {
	logger.SetName("user")
}

func main() {
	hpcmanager.LoadCommonArgs()
	flag.Parse()

	etcdRegistry := etcd.NewRegistry(
		registry.Addrs(hpcmanager.GetEtcdAddress()),
	)

	srv := micro.NewService(
		micro.Name("user"),
		micro.Registry(etcdRegistry),
	)

	userService := service.NewUser()
	user.RegisterUserHandler(srv.Server(), userService)

	srv.Init()
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}

}
