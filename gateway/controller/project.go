package controller

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/essayZW/hpcmanager/gateway/middleware"
	gatewaypb "github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/gateway/request/json"
	"github.com/essayZW/hpcmanager/gateway/response"
	"github.com/essayZW/hpcmanager/gateway/utils"
	projectpb "github.com/essayZW/hpcmanager/project/proto"
	"github.com/essayZW/hpcmanager/proto"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/client"
)

// Project 项目信息控制器
type Project struct {
	projectService projectpb.ProjectService
}

// ping /api/project/ping GET project服务ping测试
func (p *Project) ping(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)
	res, err := p.projectService.Ping(context.Background(), &proto.Empty{
		BaseRequest: baseRequest,
	})
	var resp *response.Response
	if err != nil {
		resp = response.New(500, err, false, "ping fail!")
	} else {
		resp = response.New(200, res, true, "success")
	}
	resp.Send(ctx)
}

// createProject /api/project POST 创建新的项目记录
func (p *Project) createProject(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)
	var param json.CreateProjectParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		httpResp := response.New(200, nil, false, "参数验证失败")
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	resp, err := p.projectService.CreateProject(c, &projectpb.CreateProjectRequest{
		BaseRequest: baseRequest,
		ProjectInfo: &projectpb.ProjectInfo{
			Name:        param.Name,
			From:        param.From,
			Numbering:   param.Numbering,
			Expenses:    param.Expenses,
			Description: param.Description,
		},
	})
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("创建项目信息失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}
	httpResp := response.New(200, map[string]interface{}{
		"id": resp.ProjectID,
	}, true, "success")
	httpResp.Send(ctx)
}

func (p *Project) paginationGet(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	pageIndex, pageSize, err := utils.ParsePagination(ctx)
	if err != nil {
		httpResp := response.New(200, nil, false, err.Error())
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	resp, err := p.projectService.PaginationGetProjectInfos(
		c,
		&projectpb.PaginationGetProjectInfosRequest{
			PageIndex:   int32(pageIndex),
			PageSize:    int32(pageSize),
			BaseRequest: baseRequest,
		},
	)
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("查询项目信息失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}
	res := response.PaginationQueryResponse{
		Data:  resp.Infos,
		Count: int(resp.Count),
	}
	if resp.Infos == nil {
		res.Data = make([]*projectpb.ProjectInfo, 0)
	}
	httpResp := response.New(200, res, true, "success")
	httpResp.Send(ctx)
}

// getProjectInfoByID /api/project/:id GET 通过ID查询项目信息
func (p *Project) getProjectInfoByID(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	idStr, ok := ctx.Params.Get("id")
	if !ok {
		httpResp := response.New(200, nil, false, "invalid id param")
		httpResp.Send(ctx)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpResp := response.New(200, nil, false, "invalid id param")
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	resp, err := p.projectService.GetProjectInfoByID(c, &projectpb.GetProjectInfoByIDRequest{
		Id:          int32(id),
		BaseRequest: baseRequest,
	})
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("查询项目信息错误: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}
	httpResp := response.New(200, resp.Data, true, "success")
	httpResp.Send(ctx)
}

// Registry 注册控制器方法
func (p *Project) Registry(router *gin.RouterGroup) *gin.RouterGroup {
	projectGroup := router.Group("/project")
	projectGroup.GET("/ping", p.ping)
	middleware.RegistryExcludeAPIPath("GET:/api/project/ping")

	projectGroup.POST("", p.createProject)
	projectGroup.GET("", p.paginationGet)
	projectGroup.GET("/:id", p.getProjectInfoByID)
	return projectGroup
}

// NewProject 创建新的项目控制器
func NewProject(client client.Client) Controller {
	projectService := projectpb.NewProjectService("project", client)
	return &Project{
		projectService: projectService,
	}
}
