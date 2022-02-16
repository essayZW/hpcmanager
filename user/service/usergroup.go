package service

import (
	"context"

	"github.com/essayZW/hpcmanager/logger"
	publicpb "github.com/essayZW/hpcmanager/proto"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"go-micro.dev/v4/client"
)

// UserGroupService 提供关于用户组方面的接口
type UserGroupService struct {
}

// Ping 用户组服务ping测试
func (group *UserGroupService) Ping(ctx context.Context, req *publicpb.Empty, resp *publicpb.PingResponse) error {
	logger.Info("UserGroup PING ", req)
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	return nil
}

var _ userpb.GroupServiceHandler = (*UserGroupService)(nil)

// NewGroup 创建一个新的group服务
func NewGroup(client client.Client) *UserGroupService {
	return &UserGroupService{}
}
