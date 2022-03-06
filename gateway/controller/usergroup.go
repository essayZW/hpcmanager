package controller

import (
	"context"

	"github.com/essayZW/hpcmanager/gateway/middleware"
	gatewaypb "github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/gateway/response"
	"github.com/essayZW/hpcmanager/proto"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/logger"
)

// UserGroup 用户组相关接口控制器
type UserGroup struct {
	userGroupService userpb.GroupService
}

// ping /api/group/ping GET ping!
func (ug *UserGroup) ping(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	res, err := ug.userGroupService.Ping(context.Background(), &proto.Empty{
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

// Registry 为用户组控制器注册相应的接口
func (ug *UserGroup) Registry(router *gin.RouterGroup) *gin.RouterGroup {
	logger.Info("registry gateway controller UserGroup")
	userGroup := router.Group("/group")

	userGroup.GET("/ping", ug.ping)
	middleware.RegistryExcludeAPIPath("GET:/api/group/ping")
	return userGroup
}

// NewUserGroup 创建用户组接口控制器
func NewUserGroup(client client.Client) Controller {
	return &UserGroup{
		userGroupService: userpb.NewGroupService("user", client),
	}
}
