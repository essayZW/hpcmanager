package controller

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/gateway/middleware"
	gatewaypb "github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/gateway/request/json"
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

// queryUserQuota /api/hpc/quota/:hpcID 通过计算节点用户ID查询用户存储信息
func (hpc *Hpc) queryUserQuota(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	idStr := ctx.Param("hpcID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpResp := response.New(200, nil, false, "错误的参数id")
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	resp, err := hpc.hpcService.GetQuotaByHpcUserID(c, &hpcpb.GetQuotaByHpcUserIDRequest{
		BaseRequest: baseRequest,
		HpcUserID:   int32(id),
	})
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("查询信息失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}
	httpResp := response.New(200, map[string]interface{}{
		"used":          resp.Used,
		"max":           resp.Max,
		"startTimeUnix": resp.StartTimeUnix,
		"endTimeUnix":   resp.EndTimeUnix,
	}, true, "success")
	httpResp.Send(ctx)
}

func (hpc *Hpc) setUserQuota(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	var param json.SetUserQuotaParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		httpResp := response.New(200, nil, false, err.Error())
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	newEndTime := time.UnixMilli(param.NewEndTimeMilliUnix)
	resp, err := hpc.hpcService.SetQuotaByHpcUserID(c, &hpcpb.SetQuotaByHpcUserIDRequest{
		BaseRequest:    baseRequest,
		HpcUserID:      int32(param.HpcUserID),
		NewMaxQuotaTB:  int32(param.NewSize),
		NewEndTimeUnix: newEndTime.Unix(),
		SetDate:        param.ModifyData,
	})
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("修改用户存储信息失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}

	if !resp.Success {
		httpResp := response.New(200, nil, false, "修改用户存储信息失败")
		httpResp.Send(ctx)
		return
	}

	httpResp := response.New(200, nil, true, "success")
	httpResp.Send(ctx)
}

// Registry 注册相应的处理函数
func (hpc *Hpc) Registry(router *gin.RouterGroup) *gin.RouterGroup {
	logger.Info("registry gateway controller hpc")
	hpcRouter := router.Group("/hpc")
	hpcRouter.GET("/ping", hpc.ping)
	middleware.RegistryExcludeAPIPath("GET:/api/hpc/ping")

	hpcRouter.GET("/user/:id", hpc.getUserByID)
	hpcRouter.GET("/group/:id", hpc.getGroupByID)
	hpcRouter.GET("/quota/:hpcID", hpc.queryUserQuota)
	hpcRouter.PUT("/quota", hpc.setUserQuota)
	return hpcRouter
}

// NewHpc 创建新的hpc控制器
func NewHpc(client client.Client, configConn config.DynamicConfig) Controller {
	return &Hpc{
		hpcService: hpcpb.NewHpcService("hpc", client),
	}
}
