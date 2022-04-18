package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/db"
	feedb "github.com/essayZW/hpcmanager/fee/db"
	"github.com/essayZW/hpcmanager/fee/logic"
	feepb "github.com/essayZW/hpcmanager/fee/proto"
	"github.com/essayZW/hpcmanager/fee/service"
	"github.com/essayZW/hpcmanager/logger"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func init() {
	logger.SetName("fee")
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
		micro.Name("fee"),
		micro.Registry(etcdRegistry),
	)

	serviceClient := srv.Client()

	// 创建数据库连接
	sqldb, err := db.NewDB()
	if err != nil {
		logger.Fatal("MySQL conn error: ", err)
	}

	// 创建动态配置源
	etcdConfig, err := config.NewEtcd()
	if err != nil {
		logger.Fatal("Etcd config create error: ", err)
	}

	nodeDistributeBillLogic, err := logic.NewNodeDistributeBill(feedb.NewNodeDistributeBill(sqldb), etcdConfig)
	if err != nil {
		logger.Fatal("create logic error: ", err)
	}
	nodeWeekUsageBillLogic, err := logic.NewNodeWeekUsageBill(feedb.NewNodeWeekUsageBill(sqldb), etcdConfig)
	if err != nil {
		logger.Fatal("create logic error: ", err)
	}

	nodeQuotaBillLogic, err := logic.NewNodeQuotaBill(feedb.NewNodeQuotaBill(sqldb), etcdConfig)
	if err != nil {
		logger.Fatal("create logic error: ", err)
	}

	feeService := service.NewFee(serviceClient, nodeDistributeBillLogic, nodeWeekUsageBillLogic, nodeQuotaBillLogic)
	feepb.RegisterFeeHandler(srv.Server(), feeService)

	srv.Init()
	if err := srv.Run(); err != nil {
		logger.Fatal("Service run error: ", err)
	}

}
