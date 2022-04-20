package controller

import (
	"context"

	awardpb "github.com/essayZW/hpcmanager/award/proto"
	"github.com/essayZW/hpcmanager/gateway/middleware"
	gatewaypb "github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/gateway/response"
	"github.com/essayZW/hpcmanager/proto"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/client"
)

type award struct {
	awardService awardpb.AwardService
}

// ping /api/award/ping GET ping测试
func (a *award) ping(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)
	res, err := a.awardService.Ping(context.Background(), &proto.Empty{
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

func (a *award) Registry(router *gin.RouterGroup) *gin.RouterGroup {
	awardRouter := router.Group("/award")
	awardRouter.GET("/ping", a.ping)
	middleware.RegistryExcludeAPIPath("GET:/api/award/ping")

	return awardRouter
}
func NewAward(client client.Client) Controller {
	awardService := awardpb.NewAwardService("award", client)
	return &award{
		awardService: awardService,
	}
}
