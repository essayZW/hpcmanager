package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
	projectdb "github.com/essayZW/hpcmanager/project/db"
	"github.com/essayZW/hpcmanager/project/logic"
	projectpb "github.com/essayZW/hpcmanager/project/proto"
	projectservice "github.com/essayZW/hpcmanager/project/service"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func init() {
	logger.SetName("project")
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
		micro.Name("project"),
		micro.Registry(etcdRegistry),
	)

	serviceClient := srv.Client()

	// 创建数据库连接
	sqldb, err := db.NewDB()
	if err != nil {
		logger.Fatal("MySQL conn error: ", err)
	}

	projectLogic := logic.NewProject(projectdb.NewProject(sqldb))

	projectService := projectservice.NewProject(serviceClient, projectLogic)

	projectpb.RegisterProjectHandler(srv.Server(), projectService)

	srv.Init()
	if err := srv.Run(); err != nil {
		logger.Fatal("Service run error: ", err)
	}
}
