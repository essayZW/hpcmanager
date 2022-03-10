package controller

import (
	"context"
	"strconv"
	"time"

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

// getUserByID /api/hpc/user/:id GET 通过ID查询hpc_user信息
func (hpc *Hpc) getUserByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpResp := response.New(200, nil, false, "错误的参数id")
		httpResp.Send(ctx)
		return
	}
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	resp, err := hpc.hpcService.GetUserInfoByID(c, &hpcpb.GetUserInfoByIDRequest{
		HpcUserID:   int32(id),
		BaseRequest: baseRequest,
	})
	if err != nil {
		httpResp := response.New(200, nil, false, "查询hpc用户信息失败")
		httpResp.Send(ctx)
		return
	}
	httpResp := response.New(200, resp.User, true, "success")
	httpResp.Send(ctx)
	return
}

func (hpc *Hpc) getGroupByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpResp := response.New(200, nil, false, "错误的参数id")
		httpResp.Send(ctx)
		return
	}
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	resp, err := hpc.hpcService.GetGroupInfoByID(c, &hpcpb.GetGroupInfoByIDRequest{
		HpcGroupID:  int32(id),
		BaseRequest: baseRequest,
	})
	if err != nil {
		httpResp := response.New(200, nil, false, "查询hpc用户组信息失败")
		httpResp.Send(ctx)
		return
	}
	httpResp := response.New(200, resp.Group, true, "success")
	httpResp.Send(ctx)
	return
}

// Registry 注册相应的处理函数
func (hpc *Hpc) Registry(router *gin.RouterGroup) *gin.RouterGroup {
	logger.Info("registry gateway controller hpc")
	hpcRouter := router.Group("/hpc")
	hpcRouter.GET("/ping", hpc.ping)
	middleware.RegistryExcludeAPIPath("/api/hpc/ping")

	hpcRouter.GET("/user/:id", hpc.getUserByID)
	hpcRouter.GET("/group/:id", hpc.getGroupByID)
	return hpcRouter
}

// NewHpc 创建新的hpc控制器
func NewHpc(client client.Client, configConn config.DynamicConfig) Controller {
	return &Hpc{
		hpcService: hpcpb.NewHpcService("hpc", client),
	}
}
