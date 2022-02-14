package service

import (
	"context"

	"github.com/essayZW/hpcmanager/logger"
	permissionpb "github.com/essayZW/hpcmanager/permission/proto"
	publicpb "github.com/essayZW/hpcmanager/proto"
	"go-micro.dev/v4/client"
)

// PermissionService 权限服务
type PermissionService struct {
}

// Ping ping测试
func (permission *PermissionService) Ping(ctx context.Context, req *publicpb.Empty, resp *publicpb.PingResponse) error {
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	logger.Info("PING ", resp)
	return nil
}

var _ permissionpb.PermissionHandler = (*PermissionService)(nil)

// NewPermission 创建一个新的Permission服务
func NewPermission(client client.Client) *PermissionService {
	return &PermissionService{}
}
