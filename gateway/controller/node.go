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

// checkNodeApply /api/node/apply PATCH 审核机器节点申请记录
func (n *node) checkNodeApply(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	var param json.CheckNodeApplyParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		httpResp := response.New(200, nil, false, err.Error())
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	rpcResp, err := n.nodeService.CheckNodeApply(c, &nodepb.CheckNodeApplyRequest{
		BaseRequest:  baseRequest,
		ApplyID:      int32(param.ApplyID),
		CheckStatus:  param.CheckStatus,
		CheckMessage: param.CheckMessage,
		TutorCheck:   param.TutorCheck,
	})
	if err != nil || !rpcResp.Success {
		httpResp := response.New(200, nil, false, fmt.Sprintf("审核失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}
	httpResp := response.New(200, nil, true, "success")
	httpResp.Send(ctx)
}

// paginationGetNodeDistributeWOS /api/node/distribute GET 分页查询机器节点分配处理工单信息
func (n *node) paginationGetNodeDistributeWOS(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	pageIndex, pageSize, err := utils.ParsePagination(ctx)
	if err != nil {
		httpResp := response.New(200, nil, false, err.Error())
		httpResp.Send(ctx)
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	dataResp, err := n.nodeService.PaginationGetNodeDistributeWO(
		c,
		&nodepb.PaginationGetNodeDistributeWORequest{
			BaseRequest: baseRequest,
			PageIndex:   int32(pageIndex),
			PageSize:    int32(pageSize),
		},
	)
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("查询用户信息失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}

	responseData := &response.PaginationQueryResponse{
		Count: int(dataResp.Count),
		Data:  dataResp.Wos,
	}
	if dataResp.Wos == nil {
		responseData.Data = make([]*nodepb.NodeDistribute, 0)
	}
	httpResp := response.New(200, responseData, true, "success")
	httpResp.Send(ctx)
}

// getNodeApplyByID /api/node/apply/:id GET 通过ID查询机器节点申请信息
func (n *node) getNodeApplyByID(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpResp := response.New(200, nil, false, "invalid id")
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	info, err := n.nodeService.GetNodeApplyByID(c, &nodepb.GetNodeApplyByIDRequest{
		ApplyID:     int32(id),
		BaseRequest: baseRequest,
	})

	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("查询机器节点申请信息失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}

	httpResp := response.New(200, info.Apply, true, "success")
	httpResp.Send(ctx)
}

// finishNodeDistributeByID /api/node/distribute PATCH 处理机器处理分配工单
func (n *node) finishNodeDistributeByID(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	var param json.FinishNodeDistributeParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		httpResp := response.New(200, nil, false, err.Error())
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	resp, err := n.nodeService.FinishNodeDistributeWO(c, &nodepb.FinishNodeDistributeWORequest{
		BaseRequest:  baseRequest,
		DistributeID: int32(param.ID),
	})

	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("处理工单失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}
	if !resp.Success {
		httpResp := response.New(200, nil, false, "处理工单失败,可能已经被处理")
		httpResp.Send(ctx)
		return
	}

	httpResp := response.New(200, nil, true, "success")
	httpResp.Send(ctx)
}

// revokeNodeApply /api/node/apply/:id DELETE 撤销某一个机器节点申请记录
func (n *node) revokeNodeApply(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpResp := response.New(200, nil, false, "invalid id")
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	resp, err := n.nodeService.RevokeNodeApply(c, &nodepb.RevokeNodeApplyRequest{
		BaseRequest: baseRequest,
		ApplyID:     int32(id),
	})
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("撤销申请失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}

	if !resp.Success {
		httpResp := response.New(200, nil, false, "撤销申请失败,可能已经被处理")
		httpResp.Send(ctx)
		return
	}

	httpResp := response.New(200, nil, true, "success")
	httpResp.Send(ctx)
}

func (n *node) Registry(router *gin.RouterGroup) *gin.RouterGroup {
	nodeRouter := router.Group("/node")

	nodeRouter.GET("/ping", n.ping)
	middleware.RegistryExcludeAPIPath("GET:/api/node/ping")

	nodeRouter.POST("/apply", n.createNodeApply)
	nodeRouter.GET("/apply", n.paginationGet)
	nodeRouter.PATCH("/apply", n.checkNodeApply)
	nodeRouter.GET("/apply/:id", n.getNodeApplyByID)
	nodeRouter.DELETE("/apply/:id", n.revokeNodeApply)

	nodeRouter.GET("/distribute", n.paginationGetNodeDistributeWOS)
	nodeRouter.PATCH("/distribute", n.finishNodeDistributeByID)
	return nodeRouter
}

// NewNode 创建一个新的机器节点管理控制器
func NewNode(client client.Client) Controller {
	nodeService := nodepb.NewNodeService("node", client)
	return &node{
		nodeService: nodeService,
	}
}
