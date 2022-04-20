package service

import (
	"context"

	fsspb "github.com/essayZW/hpcmanager/fss/proto"
	"github.com/essayZW/hpcmanager/logger"
	publicproto "github.com/essayZW/hpcmanager/proto"
	"go-micro.dev/v4/client"
)

type FssService struct {
}

// Ping ping测试
func (fss *FssService) Ping(
	ctx context.Context,
	req *publicproto.Empty,
	resp *publicproto.PingResponse,
) error {
	logger.Info("Ping: ", req.BaseRequest)
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	return nil
}

var _ fsspb.FssServiceHandler = (*FssService)(nil)

func NewFss(client client.Client) *FssService {
	return &FssService{}
}
