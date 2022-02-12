package controller

import (
	"context"

	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/gateway/middleware"
	gatewaypb "github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/gateway/response"
	hpcpb "github.com/essayZW/hpcmanager/hpc/proto"
	"github.com/essayZW/hpcmanager/logger"
	"github.com/essayZW/hpcmanager/proto"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/client"
)

// Hpc hpc相关的控制器
type Hpc struct {
	hpcService hpcpb.HpcService
}

// ping /api/hpc/ping GET ping!
func (hpc *Hpc) ping(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	res, err := hpc.hpcService.Ping(context.Background(), &proto.Empty{
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
func (hpc *Hpc) Registry(router *gin.RouterGroup) *gin.RouterGroup {
	logger.Info("registry gateway controller hpc")
	hpcRouter := router.Group("/hpc")
	hpcRouter.GET("/ping", hpc.ping)
	middleware.RegistryExcludeAPIPath("/api/hpc/ping")
	return hpcRouter
}

// NewHpc 创建新的hpc控制器
func NewHpc(client client.Client, configConn config.DynamicConfig) *Hpc {
	return &Hpc{
		hpcService: hpcpb.NewHpcService("hpc", client),
	}
}
