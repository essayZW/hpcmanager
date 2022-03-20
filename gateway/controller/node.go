package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/essayZW/hpcmanager/gateway/middleware"
	gatewaypb "github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/gateway/request/json"
	"github.com/essayZW/hpcmanager/gateway/response"
	"github.com/essayZW/hpcmanager/gateway/utils"
	nodepb "github.com/essayZW/hpcmanager/node/proto"
	"github.com/essayZW/hpcmanager/proto"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/client"
)

// node 机器节点管理控制器
type node struct {
	nodeService nodepb.NodeService
}

// ping /api/node/ping GET node服务的ping测试
func (n *node) ping(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)
	res, err := n.nodeService.Ping(context.Background(), &proto.Empty{
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

// createNodeApply /api/node/apply POST 创建新的机器节点申请信息
func (n *node) createNodeApply(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	var param json.CreateNodeApplyParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		httpResp := response.New(200, nil, false, err.Error())
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	resp, err := n.nodeService.CreateNodeApply(c, &nodepb.CreateNodeApplyRequest{
		ProjectID:   int32(param.ProjectID),
		NodeType:    param.NodeType,
		NodeNum:     int32(param.NodeNum),
		StartTime:   param.StartTime,
		EndTime:     param.StartTime,
		BaseRequest: baseRequest,
	})
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("创建机器节点申请信息失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}
	httpResp := response.New(200, map[string]interface{}{
		"id": resp.Id,
	}, true, "success")
	httpResp.Send(ctx)
}

// paginationGet /api/node/apply GET 分页查询机器节点申请信息
func (n *node) paginationGet(ctx *gin.Context) {
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
	resp, err := n.nodeService.PaginationGetNodeApply(c, &nodepb.PaginationGetNodeApplyRequest{
		BaseRequest: baseRequest,
		PageIndex:   int32(pageIndex),
		PageSize:    int32(pageSize),
	})
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("查询机器节点申请记录失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}
	respData := &response.PaginationQueryResponse{
		Data:  resp.Applies,
		Count: int(resp.Count),
	}
	if resp.Applies == nil {
		respData.Data = make([]*nodepb.NodeApply, 0)
	}
	httpResp := response.New(200, respData, true, "success")
	httpResp.Send(ctx)

}

func (n *node) Registry(router *gin.RouterGroup) *gin.RouterGroup {
	nodeRouter := router.Group("/node")

	nodeRouter.GET("/ping", n.ping)
	middleware.RegistryExcludeAPIPath("GET:/api/node/ping")

	nodeRouter.POST("/apply", n.createNodeApply)
	nodeRouter.GET("/apply", n.paginationGet)
	return nodeRouter
}

// NewNode 创建一个新的机器节点管理控制器
func NewNode(client client.Client) Controller {
	nodeService := nodepb.NewNodeService("node", client)
	return &node{
		nodeService: nodeService,
	}
}
