package service

import (
	"context"

	"github.com/essayZW/hpcmanager/award/logic"
	awardpb "github.com/essayZW/hpcmanager/award/proto"
	"github.com/essayZW/hpcmanager/logger"
	publicproto "github.com/essayZW/hpcmanager/proto"
	"go-micro.dev/v4/client"
)

type AwardService struct {
	paperAwardLogic *logic.Paper
}

// Ping ping测试
func (as *AwardService) Ping(ctx context.Context, req *publicproto.Empty, resp *publicproto.PingResponse) error {
	logger.Info("AwardService PING ", req)
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	return nil
}

var _ awardpb.AwardServiceHandler = (*AwardService)(nil)

func NewAward(client client.Client, paperAwardLogic *logic.Paper) *AwardService {
	return &AwardService{
		paperAwardLogic: paperAwardLogic,
	}
}
