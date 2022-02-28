package service

import (
	"context"

	"github.com/essayZW/hpcmanager/hpc/logic"
	hpcproto "github.com/essayZW/hpcmanager/hpc/proto"
	"github.com/essayZW/hpcmanager/logger"
	publicproto "github.com/essayZW/hpcmanager/proto"
	"go-micro.dev/v4/client"
)

// HpcService hpc服务
type HpcService struct {
	hpcLogic *logic.HpcLogic
}

// Ping ping测试
func (h *HpcService) Ping(ctx context.Context, req *publicproto.Empty, resp *publicproto.PingResponse) error {
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	logger.Info("PING ", resp)
	return nil
}

// CreateGroup 创建
func (h *HpcService) CreateGroup(ctx context.Context, req *hpcproto.CreateGroupRequest, resp *hpcproto.CreateGroupResponse) error {
	return nil
}

var _ hpcproto.HpcHandler = (*HpcService)(nil)

// NewHpc 新建一个Hpc服务
func NewHpc(client client.Client, hpcLogic *logic.HpcLogic) *HpcService {
	return &HpcService{
		hpcLogic: hpcLogic,
	}
}
