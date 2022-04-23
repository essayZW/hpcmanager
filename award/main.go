package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	awarddb "github.com/essayZW/hpcmanager/award/db"
	"github.com/essayZW/hpcmanager/award/logic"
	awardpb "github.com/essayZW/hpcmanager/award/proto"
	"github.com/essayZW/hpcmanager/award/service"
	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/db"
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
	// 创建数据库连接
	sqldb, err := db.NewDB()
	if err != nil {
		logger.Fatal("MySQL conn error: ", err)
	}

	serviceClient := srv.Client()
	serviceServer := srv.Server()

	paperAwardLogic := logic.NewPaper(awarddb.NewPaperAward(sqldb))

	awardService := service.NewAward(serviceClient, paperAwardLogic)
	awardpb.RegisterAwardServiceHandler(serviceServer, awardService)

	srv.Init()
	if err := srv.Run(); err != nil {
		logger.Fatal("Service run error: ", err)
	}
}
