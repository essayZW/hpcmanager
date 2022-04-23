package controller

import (
	"context"
	"fmt"
	"time"

	awardpb "github.com/essayZW/hpcmanager/award/proto"
	"github.com/essayZW/hpcmanager/gateway/middleware"
	gatewaypb "github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/gateway/request/json"
	"github.com/essayZW/hpcmanager/gateway/response"
	"github.com/essayZW/hpcmanager/gateway/utils"
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

// createPaperAwardApply /api/award/paper 创建论文奖励申请
func (a *award) createPaperAwardApply(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	var param json.CreatePaperAwardApplyParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		httpResp := response.New(200, nil, false, "参数验证失败")
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	resp, err := a.awardService.CreatePaperAward(c, &awardpb.CreatePaperAwardRequest{
		BaseRequest:         baseRequest,
		Title:               param.Title,
		Category:            param.Category,
		Partition:           param.Partition,
		FirstPageImageName:  param.FirstPageImageName,
		ThanksPageImageName: param.ThanksPageImageName,
		RemarkMessage:       param.RemarkMessage,
	})
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("创建论文奖励申请失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}

	httpResp := response.New(200, map[string]interface{}{
		"id": resp.Id,
	}, true, "success")
	httpResp.Send(ctx)
}

// paginationGetPaperApply /api/award/paper GET 分页查询论文奖励申请信息
func (a *award) paginationGetPaperApply(ctx *gin.Context) {
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

	resp, err := a.awardService.PaginationGetPaperApply(c, &awardpb.PaginationGetPaperApplyRequest{
		BaseRequest: baseRequest,
		PageIndex:   int32(pageIndex),
		PageSize:    int32(pageSize),
	})
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("查询论文奖励申请信息失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}
	respData := &response.PaginationQueryResponse{
		Data:  resp.Applies,
		Count: int(resp.Count),
	}
	if resp.Applies == nil {
		respData.Data = make([]*awardpb.PaperApply, 0)
	}
	httpResp := response.New(200, respData, true, "success")
	httpResp.Send(ctx)

}

// checkPaperApply /api/award/paper PUT 审核论文奖励申请
func (a *award) checkPaperApply(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	var param json.CheckPaperApplyParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		httpResp := response.New(200, nil, false, "参数验证失败")
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	resp, err := a.awardService.CheckPaperApplyByID(c, &awardpb.CheckPaperApplyByIDRequest{
		BaseRequest:  baseRequest,
		ApplyID:      int32(param.ID),
		Money:        param.CheckMoney,
		CheckMessage: param.CheckMessage,
		Accept:       param.Accept,
	})

	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("审核论文奖励申请失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}
	if !resp.Success {
		httpResp := response.New(200, nil, false, "审核论文奖励申请失败")
		httpResp.Send(ctx)
		return
	}
	httpResp := response.New(200, nil, true, "success")
	httpResp.Send(ctx)
	return
}

// createTechnologyAwardApply /api/award/technology POST 创建新的科技奖励申请记录
func (a *award) createTechnologyAwardApply(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	var param json.CreateTechnologyAwardApplyParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		httpResp := response.New(200, nil, false, "参数验证失败")
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	resp, err := a.awardService.CreateTechnologyAwardApply(c, &awardpb.CreateTechnologyAwardApplyRequest{
		BaseRequest:    baseRequest,
		ProjectID:      int32(param.ProjectID),
		PrizeLevel:     param.PrizeLevel,
		PrizeImageName: param.PrizeImageName,
		RemarkMessage:  param.RemarkMessage,
	})
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("创建科技奖励申请失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}

	httpResp := response.New(200, map[string]interface{}{
		"id": resp.Id,
	}, true, "success")
	httpResp.Send(ctx)
}

func (a *award) Registry(router *gin.RouterGroup) *gin.RouterGroup {
	awardRouter := router.Group("/award")
	awardRouter.GET("/ping", a.ping)
	middleware.RegistryExcludeAPIPath("GET:/api/award/ping")

	awardRouter.POST("/paper", a.createPaperAwardApply)
	awardRouter.GET("/paper", a.paginationGetPaperApply)
	awardRouter.PUT("/paper", a.checkPaperApply)

	awardRouter.POST("/technology", a.createTechnologyAwardApply)
	return awardRouter
}
func NewAward(client client.Client) Controller {
	awardService := awardpb.NewAwardService("award", client)
	return &award{
		awardService: awardService,
	}
}
