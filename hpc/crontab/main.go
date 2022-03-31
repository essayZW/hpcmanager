package main

import (
	"context"
	"flag"
	"time"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/logger"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func init() {
	logger.SetName("crontab")
}

func parseFlag() (time.Time, time.Time, error) {
	var startDateStr string
	var endDateStr string
	flag.StringVar(&startDateStr, "dt_start", "", "start date")
	flag.StringVar(&endDateStr, "dt_end", "", "end date")
	flag.Parse()

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	startDate, err := time.ParseInLocation("2006-01-02", startDateStr, location)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	endDate, err := time.ParseInLocation("2006-01-02", endDateStr, location)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	return startDate, endDate, nil
}

func main() {
	startDate, endDate, err := parseFlag()
	if err != nil {
		logger.Fatal("parse flag error: ", err)
	}

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

	loader := NewLoader(serviceClient, 3)
	err = loader.Sync(context.Background(), startDate, endDate)
	if err != nil {
		logger.Fatalf("sync error: %v with date: %v %v", err, startDate, endDate)
	}
}
