package service

import (
	"context"

	"github.com/essayZW/hpcmanager/logger"
	nodepb "github.com/essayZW/hpcmanager/node/proto"
	publicproto "github.com/essayZW/hpcmanager/proto"
	"go-micro.dev/v4/client"
)

// NodeService 机器管理服务
type NodeService struct {
}

// Ping ping测试
func (ns *NodeService) Ping(ctx context.Context, req *publicproto.Empty, resp *publicproto.PingResponse) error {
	logger.Info("NodeService PING ", req)
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	return nil
}

var _ nodepb.NodeHandler = (*NodeService)(nil)

// NewNode 创建新的机器节点管理服务
func NewNode(client client.Client) *NodeService {
	return &NodeService{}
}
