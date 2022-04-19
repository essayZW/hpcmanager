package controller

import (
	"context"
	"fmt"
	"strconv"
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
		return
	}

	httpResp := response.New(200, map[string]float64{
		"rate36CPU": resp.Rate36CPU,
		"rate4GPU":  resp.Rate4GPU,
		"rate8GPU":  resp.Rate8GPU,
	}, true, "success")
	httpResp.Send(ctx)
}

// paginationGetNodeWeekUsageBills /api/fee/usage/week GET 分页查询机器节点机时周账单
func (f *fee) paginationGetNodeWeekUsageBills(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	pageIndex, pageSize, err := utils.ParsePagination(ctx)
	if err != nil {
		httpResp := response.New(200, nil, false, err.Error())
		httpResp.Send(ctx)
		return
	}

	startTime, endTime, err := utils.ParseDateRange(ctx)
	if err != nil {
		httpResp := response.New(200, nil, false, err.Error())
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	resp, err := f.feeService.PaginationGetNodeWeekUsageBillRecords(c, &feepb.PaginationGetNodeWeekUsageBillRecordsResquest{
		BaseRequest:   baseRequest,
		PageIndex:     int32(pageIndex),
		PageSize:      int32(pageSize),
		StartTimeUnix: startTime.Unix(),
		EndTimeUnix:   endTime.Unix(),
	})
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("查询账单信息失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}
	responseData := response.PaginationQueryResponse{
		Data:  resp.Bills,
		Count: int(resp.Count),
	}
	if resp.Bills == nil {
		responseData.Data = make([]*feepb.NodeWeekUsageBill, 0)
	}
	httpResp := response.New(200, responseData, true, "success")
	httpResp.Send(ctx)
}

// paginationGetNodeWeekUsageBillsGroupByGroupID 分页查询机时周账单并按照组ID进行分组
func (f *fee) paginationGetNodeWeekUsageBillsGroupByGroupID(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	pageIndex, pageSize, err := utils.ParsePagination(ctx)
	if err != nil {
		httpResp := response.New(200, nil, false, err.Error())
		httpResp.Send(ctx)
		return
	}

	payFlagStr, ok := ctx.GetQuery("payFlag")
	if !ok {
		httpResp := response.New(200, nil, false, "缺少payFlag参数")
		httpResp.Send(ctx)
		return
	}
	payFlag, err := strconv.ParseBool(payFlagStr)
	if err != nil {
		httpResp := response.New(200, nil, false, "payFlag 必须是一个Bool值")
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	resp, err := f.feeService.PaginationGetUserGroupUsageBillRecords(c, &feepb.PaginationGetUserGroupUsageBillRecordsRequest{
		BaseRequest: baseRequest,
		PageIndex:   int32(pageIndex),
		PageSize:    int32(pageSize),
		PayFlag:     payFlag,
	})

	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("查询账单失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}

	responseData := response.PaginationQueryResponse{
		Count: int(resp.Count),
		Data:  resp.Bills,
	}
	if resp.Bills == nil {
		responseData.Data = make([]*feepb.NodeWeekUsageBillForUserGroup, 0)
	}
	httpResp := response.New(200, responseData, true, "success")
	httpResp.Send(ctx)
}

// payGroupNodeUsageBill /api/fee/usage/group/bill PUT 支付用户组机器节点机时账单
func (f *fee) payGroupNodeUsageBill(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	var param json.PayGroupNodeUsageBillParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		httpResp := response.New(200, nil, false, "参数验证失败")
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	resp, err := f.feeService.PayGroupNodeUsageBill(c, &feepb.PayGroupNodeUsageBillRequest{
		BaseRequest: baseRequest,
		UserGroupID: int32(param.UserGroupID),
		PayType:     int32(param.PayType),
		PayMessage:  param.PayMessage,
		NeedFee:     param.NeedFee,
	})
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("缴费失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}

	httpResp := response.New(200, map[string]int{
		"count": int(resp.PayCount),
	}, true, "success")
	httpResp.Send(ctx)
}

// getNodeUsageFeeRate /api/fee/rate/usage GET 查询机器时长费率信息
func (f *fee) getNodeUsageFeeRate(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	resp, err := f.feeService.GetNodeUsageFeeRate(c, &feepb.GetNodeUsageFeeRateRequest{
		BaseRequest: baseRequest,
	})
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("查询费率信息失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}

	httpResp := response.New(200, map[string]float64{
		"cpu": resp.Cpu,
		"gpu": resp.Gpu,
	}, true, "success")
	httpResp.Send(ctx)
}

// paginationGetNodeQuotaBills /api/fee/quota GET 分页查询机器存储账单
func (f *fee) paginationGetNodeQuotaBills(ctx *gin.Context) {
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

	resp, err := f.feeService.PaginationGetNodeQuotaBill(c, &feepb.PaginationGetNodeQuotaBillRequest{
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
		respData.Data = make([]*feepb.NodeQuotaBill, 0)
	}
	httpResp := response.New(200, respData, true, "success")
	httpResp.Send(ctx)
}

// getNodeQuotaFeeRate /api/fee/rate/quota GET  查询机器存储费率
func (f *fee) getNodeQuotaFeeRate(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	resp, err := f.feeService.GetNodeQuotaFeeRate(c, &feepb.GetNodeQuotaFeeRateRequest{
		BaseRequest: baseRequest,
	})
	if err != nil {
		httpResp := response.New(200, nil, false, "信息查询失败")
		httpResp.Send(ctx)
		return
	}

	httpResp := response.New(200, map[string]interface{}{
		"basic": resp.Basic,
		"extra": resp.Extra,
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
	feeRouter.GET("/rate/usage", f.getNodeUsageFeeRate)
	feeRouter.GET("/rate/quota", f.getNodeQuotaFeeRate)

	feeRouter.GET("/usage/week", f.paginationGetNodeWeekUsageBills)
	feeRouter.GET("/usage/group/week", f.paginationGetNodeWeekUsageBillsGroupByGroupID)
	feeRouter.PUT("/usage/group/bill", f.payGroupNodeUsageBill)

	feeRouter.GET("/quota", f.paginationGetNodeQuotaBills)
	return feeRouter
}

func NewFee(client client.Client) Controller {
	feeService := feepb.NewFeeService("fee", client)
	return &fee{
		feeService: feeService,
	}
}
