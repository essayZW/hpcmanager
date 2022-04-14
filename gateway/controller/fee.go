package controller

import (
	"context"
	"fmt"
	"time"

	feepb "github.com/essayZW/hpcmanager/fee/proto"
	"github.com/essayZW/hpcmanager/gateway/middleware"
	gatewaypb "github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/gateway/request/json"
	"github.com/essayZW/hpcmanager/gateway/response"
	"github.com/essayZW/hpcmanager/gateway/utils"
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

// paginationGetNodeDistributeBill /api/fee/distribute GET 分页查询用户节点独占账单
func (f *fee) paginationGetNodeDistributeBill(ctx *gin.Context) {
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

	resp, err := f.feeService.PaginationGetNodeDistributeBill(c, &feepb.PaginationGetNodeDistributeBillRequest{
		BaseRequest: baseRequest,
		PageIndex:   int32(pageIndex),
		PageSize:    int32(pageSize),
	})
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("查询账单信息失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}

	respData := &response.PaginationQueryResponse{
		Data:  resp.Bills,
		Count: int(resp.Count),
	}
	if resp.Bills == nil {
		respData.Data = make([]*feepb.NodeDistributeBill, 0)
	}
	httpResp := response.New(200, respData, true, "success")
	httpResp.Send(ctx)
}

// payNodeDistributeBill /api/fee/distribute PUT 支付机器独占账单
func (f *fee) payNodeDistributeBill(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	var param json.PayNodeDistributeBillParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		httpResp := response.New(200, nil, false, err.Error())
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	resp, err := f.feeService.PayNodeDistributeBill(c, &feepb.PayNodeDistributeBillRequest{
		BaseRequest: baseRequest,
		Id:          int32(param.ID),
		PayMoney:    param.PayMoney,
		PayMessage:  param.PayMessage,
		PayType:     int32(param.PayType),
	})
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("支付账单失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}
	if !resp.Success {
		httpResp := response.New(200, nil, false, "账单没有发生任何变化")
		httpResp.Send(ctx)
		return
	}
	httpResp := response.New(200, nil, true, "success")
	httpResp.Send(ctx)
}

// getNodeDistributeFeeRate /api/fee/rate/distribute GET 查询机器节点独占费率
func (f *fee) getNodeDistributeFeeRate(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	resp, err := f.feeService.GetNodeDistributeFeeRate(c, &feepb.GetNodeDistributeFeeRateRequest{
		BaseRequest: baseRequest,
	})
	if err != nil {
		httpResp := response.New(200, nil, false, "信息查询失败")
		httpResp.Send(ctx)
	}

	httpResp := response.New(200, map[string]float64{
		"rate36CPU": resp.Rate36CPU,
		"rate4GPU":  resp.Rate4GPU,
		"rate8GPU":  resp.Rate8GPU,
	}, true, "success")
	httpResp.Send(ctx)
}

func (f *fee) Registry(router *gin.RouterGroup) *gin.RouterGroup {
	feeRouter := router.Group("/fee")

	feeRouter.GET("/ping", f.ping)
	middleware.RegistryExcludeAPIPath("GET:/api/fee/ping")

	feeRouter.GET("/distribute", f.paginationGetNodeDistributeBill)
	feeRouter.PUT("/distribute", f.payNodeDistributeBill)

	feeRouter.GET("/rate/distribute", f.getNodeDistributeFeeRate)
	return feeRouter
}

func NewFee(client client.Client) Controller {
	feeService := feepb.NewFeeService("fee", client)
	return &fee{
		feeService: feeService,
	}
}
