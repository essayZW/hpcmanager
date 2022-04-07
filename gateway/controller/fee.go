package controller

import (
	"context"

	feepb "github.com/essayZW/hpcmanager/fee/proto"
	"github.com/essayZW/hpcmanager/gateway/middleware"
	gatewaypb "github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/gateway/response"
	"github.com/essayZW/hpcmanager/proto"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/client"
)

type fee struct {
	feeService feepb.FeeService
}

// ping /api/fee/ping GET ping测试
func (f *fee) ping(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)
	res, err := f.feeService.Ping(context.Background(), &proto.Empty{
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

func (f *fee) Registry(router *gin.RouterGroup) *gin.RouterGroup {
	feeRouter := router.Group("/fee")

	feeRouter.GET("/ping", f.ping)
	middleware.RegistryExcludeAPIPath("GET:/api/fee/ping")

	return feeRouter
}

func NewFee(client client.Client) Controller {
	feeService := feepb.NewFeeService("fee", client)
	return &fee{
		feeService: feeService,
	}
}
