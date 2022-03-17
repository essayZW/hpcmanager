package controller

import (
	"context"

	"github.com/essayZW/hpcmanager/gateway/middleware"
	gatewaypb "github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/gateway/response"
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

func (n *node) Registry(router *gin.RouterGroup) *gin.RouterGroup {
	nodeRouter := router.Group("/node")

	nodeRouter.GET("/ping", n.ping)
	middleware.RegistryExcludeAPIPath("GET:/api/node/ping")
	return nodeRouter
}

// NewNode 创建一个新的机器节点管理控制器
func NewNode(client client.Client) Controller {
	nodeService := nodepb.NewNodeService("node", client)
	return &node{
		nodeService: nodeService,
	}
}
