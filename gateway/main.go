package main

import (
	"flag"
	"strconv"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/gateway/controller"
	"github.com/essayZW/hpcmanager/gateway/middleware"
	"github.com/essayZW/hpcmanager/logger"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
)

func init() {
	logger.SetName("gateway")
}

func newServiceClient(etcdAddr string) client.Client {
	etcdRegistry := etcd.NewRegistry(
		registry.Addrs(etcdAddr),
	)
	srv := micro.NewService(
		micro.Registry(etcdRegistry),
	)
	return srv.Client()
}

func main() {
	var port int
	var debug bool
	flag.IntVar(&port, "port", 80, "port to listen")
	flag.BoolVar(&debug, "debug", true, "debug mode")
	hpcmanager.LoadCommonArgs()
	flag.Parse()
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	serviceClient := newServiceClient(hpcmanager.GetEtcdAddress())
	server := gin.New()

	api := server.Group("/api")
	middleware.Registry(api, serviceClient)

	etcdConfig, err := config.NewEtcd()
	if err != nil {
		logger.Error("Etcd config create error: ", err)
	}

	userController := controller.NewUser(serviceClient, etcdConfig)
	userController.Registry(api)

	server.Run(":" + strconv.Itoa(port))
}
