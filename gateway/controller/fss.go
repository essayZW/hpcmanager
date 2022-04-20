package controller

import (
	"context"

	fsspb "github.com/essayZW/hpcmanager/fss/proto"
	"github.com/essayZW/hpcmanager/gateway/middleware"
	gatewaypb "github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/gateway/response"
	"github.com/essayZW/hpcmanager/proto"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/client"
)

type fss struct {
	fssService fsspb.FssService
}

// ping /api/fss/ping GET ping测试
func (f *fss) ping(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)
	res, err := f.fssService.Ping(context.Background(), &proto.Empty{
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

func (f *fss) Registry(router *gin.RouterGroup) *gin.RouterGroup {
	fssRouter := router.Group("/fss")
	fssRouter.GET("/ping", f.ping)
	middleware.RegistryExcludeAPIPath("GET:/api/fss/ping")

	return fssRouter
}

func NewFss(client client.Client) Controller {
	fssService := fsspb.NewFssService("fss", client)
	return &fss{
		fssService: fssService,
	}
}
