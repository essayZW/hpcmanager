package controller

import (
	"context"
	"time"

	"github.com/essayZW/hpcmanager/gateway/middleware"
	gatewaypb "github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/gateway/request/json"
	"github.com/essayZW/hpcmanager/gateway/response"
	"github.com/essayZW/hpcmanager/gateway/utils"
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

type paginationGetGroupInfoResponse struct {
	Count int
	Data  []*userpb.GroupInfo
}

// paginationGetGroupInfo 分页查询用户组信息
func (ug *UserGroup) paginationGetGroupInfo(ctx *gin.Context) {
	pageIndex, pageSize, err := utils.ParsePagination(ctx)
	if err != nil {
		res := response.New(200, nil, false, err.Error())
		res.Send(ctx)
		return
	}
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	resp, err := ug.userGroupService.PaginationGetGroupInfo(c, &userpb.PaginationGetGroupInfoRequest{
		PageSize:    int32(pageSize),
		PageIndex:   int32(pageIndex),
		BaseRequest: baseRequest,
	})
	if err != nil {
		res := response.New(200, nil, false, "信息查询失败")
		res.Send(ctx)
		return
	}
	res := paginationGetGroupInfoResponse{
		Count: int(resp.Count),
		Data:  resp.GroupInfos,
	}
	if res.Data == nil {
		res.Data = make([]*userpb.GroupInfo, 0)
	}
	sendResponse := response.New(200, res, true, "success")
	sendResponse.Send(ctx)
}

func (ug *UserGroup) createGroup(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	param := json.CreateGroupParam{}
	if err := ctx.ShouldBindJSON(&param); err != nil {
		httpResp := response.New(200, nil, false, err.Error())
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	resp, err := ug.userGroupService.CreateGroup(c, &userpb.CreateGroupRequest{
		TutorID:     int32(param.TutorID),
		Name:        param.GroupName,
		QueueName:   param.QueueName,
		BaseRequest: baseRequest,
	})
	var httpResp *response.Response
	if err != nil {
		httpResp = response.New(200, nil, false, "创建组失败")
	} else {
		httpResp = response.New(200, resp, true, "success")
	}
	httpResp.Send(ctx)
}

// Registry 为用户组控制器注册相应的接口
func (ug *UserGroup) Registry(router *gin.RouterGroup) *gin.RouterGroup {
	logger.Info("registry gateway controller UserGroup")
	userGroup := router.Group("/group")

	userGroup.GET("/ping", ug.ping)
	middleware.RegistryExcludeAPIPath("GET:/api/group/ping")

	userGroup.GET("", ug.paginationGetGroupInfo)
	userGroup.POST("", ug.createGroup)
	return userGroup
}

// NewUserGroup 创建用户组接口控制器
func NewUserGroup(client client.Client) Controller {
	return &UserGroup{
		userGroupService: userpb.NewGroupService("user", client),
	}
}
