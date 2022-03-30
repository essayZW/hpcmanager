package main

import (
	"context"
	"time"

	gatewaypb "github.com/essayZW/hpcmanager/gateway/proto"
	hpcpb "github.com/essayZW/hpcmanager/hpc/proto"
	nodepb "github.com/essayZW/hpcmanager/node/proto"
	"github.com/essayZW/hpcmanager/verify"
	"go-micro.dev/v4/client"
)

// RPCLoader 通过RPC协议进行数据的获取以及同步
type RPCLoader struct {
	nodeService nodepb.NodeService
	hpcService  hpcpb.HpcService
}

func (loader *RPCLoader) generateBaseRequest() *gatewaypb.BaseRequest {
	return &gatewaypb.BaseRequest{
		UserInfo: &gatewaypb.UserInfo{
			Levels: []int32{
				int32(verify.SuperAdmin),
			},
		},
		RequestInfo: &gatewaypb.RequestInfo{
			Id: "__SYSTEM_CRONTAB__",
		},
	}
}

// Sync 从hpc服务同步数据到node服务
func (loader *RPCLoader) Sync(ctx context.Context, startDate, endDate time.Time) error {
	baseRequest := loader.generateBaseRequest()
	infos, err := loader.hpcService.GetNodeUsage(ctx, &hpcpb.GetNodeUsageRequest{
		BaseRequest:   baseRequest,
		StartTimeUnix: startDate.Unix(),
		EndTimeUnix:   endDate.Unix(),
	})
	if err != nil {
		return err
	}
	for _, usage := range infos.Usages {
		resp, err := loader.nodeService.AddNodeUsageTimeRecord(ctx, &nodepb.AddNodeUsageTimeRecordRequest{
			BaseRequest: baseRequest,
			// TODO: 需要一个通过HPC用户节点的用户名查询对应用户信息的接口
		})
	}
	return nil
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
