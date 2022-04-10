package service

import (
	"context"

	"github.com/essayZW/hpcmanager/fee/logic"
	feepb "github.com/essayZW/hpcmanager/fee/proto"
	"github.com/essayZW/hpcmanager/logger"
	publicproto "github.com/essayZW/hpcmanager/proto"
	"go-micro.dev/v4/client"
)

type FeeService struct {
	nodeDistributeBillLogic *logic.NodeDistributeBill
}

// Ping ping测试
func (fs *FeeService) Ping(ctx context.Context, req *publicproto.Empty, resp *publicproto.PingResponse) error {
	logger.Info("NodeService PING ", req)
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	return nil
}

var _ feepb.FeeHandler = (*FeeService)(nil)

// NewFee 创建新的fee服务
func NewFee(client client.Client, nodeDistributeBillLogic *logic.NodeDistributeBill) *FeeService {
	return &FeeService{
		nodeDistributeBillLogic: nodeDistributeBillLogic,
	}
}
