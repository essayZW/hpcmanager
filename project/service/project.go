package service

import (
	"context"

	"github.com/essayZW/hpcmanager/logger"
	"github.com/essayZW/hpcmanager/project/logic"
	publicproto "github.com/essayZW/hpcmanager/proto"
	"go-micro.dev/v4/client"
)

// ProjectService 用户服务
type ProjectService struct {
	projectLogic *logic.Project
}

// Ping ping测试
func (ps *ProjectService) Ping(ctx context.Context, req *publicproto.Empty, resp *publicproto.PingResponse) error {
	logger.Info("Project PING ", req)
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	return nil
}

// NewProject 创建用户服务
func NewProject(client client.Client, projectLogic *logic.Project) *ProjectService {
	return &ProjectService{
		projectLogic: projectLogic,
	}
}
