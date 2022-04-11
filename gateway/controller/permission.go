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
	"github.com/essayZW/hpcmanager/logger"
	permissionpb "github.com/essayZW/hpcmanager/permission/proto"
	"github.com/essayZW/hpcmanager/proto"
	"github.com/essayZW/hpcmanager/verify"
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

// addAdmin /api/permission/admin POST 添加新的管理员用户
func (permission *Permission) addAdmin(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	var param json.ChangeUserPermissionParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		httpResp := response.New(200, nil, false, err.Error())
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	resp, err := permission.permissionService.AddUserPermission(
		c,
		&permissionpb.AddUserPermissionRequest{
			Userid:      int32(param.UserID),
			Level:       int32(verify.CommonAdmin),
			BaseRequest: baseRequest,
		},
	)
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("添加用户权限失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}
	if !resp.Success {
		httpResp := response.New(200, nil, false, "添加用户权限失败")
		httpResp.Send(ctx)
		return
	}

	httpResp := response.New(200, nil, true, "success")
	httpResp.Send(ctx)
}

// getUserPermission /api/permission/user/:id GET 查询某个用户的所有权限
func (permission *Permission) getUserPermission(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpResp := response.New(200, nil, false, "invalid id param")
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	resp, err := permission.permissionService.GetUserPermission(
		c,
		&permissionpb.GetUserPermissionRequest{
			Id:          int32(id),
			BaseRequest: baseRequest,
		},
	)

	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("查询用户权限信息失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}

	httpResp := response.New(200, resp.Info, true, "success")
	httpResp.Send(ctx)
}

func (permission *Permission) removeAdmin(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	var param json.ChangeUserPermissionParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		httpResp := response.New(200, nil, false, err.Error())
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	resp, err := permission.permissionService.RemoveUserPermission(
		c,
		&permissionpb.RemoveUserPermissionRequest{
			BaseRequest: baseRequest,
			Userid:      int32(param.UserID),
			Level:       int32(verify.CommonAdmin),
		},
	)

	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("删除用户权限失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}

	if !resp.Success {
		httpResp := response.New(200, nil, false, "删除用户权限失败")
		httpResp.Send(ctx)
		return
	}

	httpResp := response.New(200, nil, true, "success")
	httpResp.Send(ctx)
}

// Registry 注册相应的处理函数
func (permission *Permission) Registry(router *gin.RouterGroup) *gin.RouterGroup {
	logger.Info("registry gateway controller permission")
	permissionRouter := router.Group("/permission")
	permissionRouter.GET("/ping", permission.ping)
	middleware.RegistryExcludeAPIPath("GET:/api/permission/ping")

	permissionRouter.POST("/admin", permission.addAdmin)
	permissionRouter.DELETE("/admin", permission.removeAdmin)
	permissionRouter.GET("/user/:id", permission.getUserPermission)
	return permissionRouter
}

// NewPermission 创建新的permission控制器
func NewPermission(client client.Client, configConn config.DynamicConfig) Controller {
	return &Permission{
		permissionService: permissionpb.NewPermissionService("permission", client),
	}
}
