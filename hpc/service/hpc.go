package service

import (
	"context"

	hpcproto "github.com/essayZW/hpcmanager/hpc/proto"
	"github.com/essayZW/hpcmanager/logger"
	publicproto "github.com/essayZW/hpcmanager/proto"
	"go-micro.dev/v4/client"
)

// HpcService hpc服务
type HpcService struct {
}

// Ping ping测试
func (h *HpcService) Ping(ctx context.Context, req *publicproto.Empty, resp *publicproto.PingResponse) error {
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	logger.Info("PING ", resp)
	return nil
}

var _ hpcproto.HpcHandler = (*HpcService)(nil)

// NewHpc 新建一个Hpc服务
func NewHpc(client client.Client) *HpcService {
	return &HpcService{}
}
