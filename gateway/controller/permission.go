package controller

import (
	"context"

	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/gateway/middleware"
	gatewaypb "github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/gateway/response"
	"github.com/essayZW/hpcmanager/logger"
	permissionpb "github.com/essayZW/hpcmanager/permission/proto"
	"github.com/essayZW/hpcmanager/proto"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/client"
)

// Permission permission相关的控制器
type Permission struct {
	permissionService permissionpb.PermissionService
}

// ping /api/permission/ping GET ping!
func (permission *Permission) ping(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	res, err := permission.permissionService.Ping(context.Background(), &proto.Empty{
		BaseRequest: baseReq.(*gatewaypb.BaseRequest),
	})
	var resp *response.Response
	if err != nil {
		resp = response.New(500, err, false, "ping fail!")
	} else {
		resp = response.New(200, res, true, "success")
	}
	resp.Send(ctx)
}

// Registry 注册相应的处理函数
func (permission *Permission) Registry(router *gin.RouterGroup) *gin.RouterGroup {
	logger.Info("registry gateway controller permission")
	permissionRouter := router.Group("/permission")
	permissionRouter.GET("/ping", permission.ping)
	middleware.RegistryExcludeAPIPath("/api/permission/ping")
	return permissionRouter
}

// NewPermission 创建新的permission控制器
func NewPermission(client client.Client, configConn config.DynamicConfig) Controller {
	return &Permission{
		permissionService: permissionpb.NewPermissionService("permission", client),
	}
}
