package main

import (
	"context"
	"log"
	"time"

	hpcpb "github.com/essayZW/hpcmanager/hpc/proto"
	nodepb "github.com/essayZW/hpcmanager/node/proto"
	"go-micro.dev/v4/client"
)

// RPCLoader 通过RPC协议进行数据的获取以及同步
type RPCLoader struct {
	nodeService nodepb.NodeService
	hpcService  hpcpb.HpcService
}

// Sync 从hpc服务同步数据到node服务
func (loader *RPCLoader) Sync(ctx context.Context, startDate, endDate time.Time) {
	log.Panic("need implement")
}

// NewLoader 创建新的机器时长数据的加载器
func NewLoader(client client.Client) *RPCLoader {
	nodeService := nodepb.NewNodeService("node", client)
	hpcService := hpcpb.NewHpcService("hpc", client)
	return &RPCLoader{
		nodeService: nodeService,
		hpcService:  hpcService,
	}
}
