package main

import (
	"flag"
	"strconv"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/gateway/controller"
	"github.com/essayZW/hpcmanager/gateway/middleware"
	"github.com/essayZW/hpcmanager/gateway/response"
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
	flag.Parse()
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	registryConf, err := config.LoadRegistry()
	if err != nil {
		logger.Fatal("etcd config load error: ", err)
	}

	serviceClient := newServiceClient(registryConf.Etcd.Address)
	server := gin.New()

	api := server.Group("/api")
	middleware.Registry(api, serviceClient)

	etcdConfig, err := config.NewEtcd()
	if err != nil {
		logger.Error("Etcd config create error: ", err)
	}

	userController := controller.NewUser(serviceClient, etcdConfig)
	userController.Registry(api)

	// 注册404处理
	server.NoRoute(func(c *gin.Context) {
		response.New(404, "404 not found", false, "404 not found").Send(c)
	})

	server.Run(":" + strconv.Itoa(port))
}
