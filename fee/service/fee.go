package service

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/fee/logic"
	feepb "github.com/essayZW/hpcmanager/fee/proto"
	"github.com/essayZW/hpcmanager/logger"
	nodepb "github.com/essayZW/hpcmanager/node/proto"
	publicproto "github.com/essayZW/hpcmanager/proto"
	"github.com/essayZW/hpcmanager/verify"
	"go-micro.dev/v4/client"
)

type FeeService struct {
	nodeDistributeBillLogic *logic.NodeDistributeBill
	nodeService             nodepb.NodeService
}

// Ping ping测试
func (fs *FeeService) Ping(ctx context.Context, req *publicproto.Empty, resp *publicproto.PingResponse) error {
	logger.Info("NodeService PING ", req)
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	return nil
}

// CreateNodeDistributeBill 创建节点分配处理之后对应的账单
func (fs *FeeService) CreateNodeDistributeBill(
	ctx context.Context,
	req *feepb.CreateNodeDistributeBillRequest,
	resp *feepb.CreateNodeDistributeBillResponse,
) error {
	logger.Info("CreateNodeDistributeBill: ", req.BaseRequest)
	if !verify.Identify(verify.CreateNodeDistributeBill, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"CreateNodeDistributeBill permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("CreateNodeDistributeBill permission forbidden")
	}

	nodeDistributeResp, err := fs.nodeService.GetNodeDistributeInfoByID(ctx, &nodepb.GetNodeDistributeInfoByIDRequest{
		BaseRequest: req.BaseRequest,
		Id:          req.NodeDistributeID,
	})
	if err != nil {
		return errors.New("can't find node distribute by id")
	}

	nodeApplyInfo, err := fs.nodeService.GetNodeApplyByID(ctx, &nodepb.GetNodeApplyByIDRequest{
		BaseRequest: req.BaseRequest,
		ApplyID:     nodeDistributeResp.Wo.ApplyID,
	})
	if err != nil {
		return errors.New("can't find node apply by id")
	}

	fee, err := fs.nodeDistributeBillLogic.CalFee(
		nodeApplyInfo.Apply.StartTime,
		nodeApplyInfo.Apply.EndTime,
		nodeApplyInfo.Apply.NodeType,
	)
	if err != nil {
		return err
	}

	id, err := fs.nodeDistributeBillLogic.Create(
		ctx,
		int(nodeApplyInfo.Apply.Id),
		int(nodeDistributeResp.Wo.Id),
		fee,
		int(nodeApplyInfo.Apply.CreaterID),
		nodeApplyInfo.Apply.CreaterUsername,
		nodeApplyInfo.Apply.CreaterName,
	)
	if err != nil {
		return err
	}
	resp.Id = int32(id)
	return nil
}

var _ feepb.FeeHandler = (*FeeService)(nil)

// NewFee 创建新的fee服务
func NewFee(client client.Client, nodeDistributeBillLogic *logic.NodeDistributeBill) *FeeService {
	nodeService := nodepb.NewNodeService("node", client)
	return &FeeService{
		nodeDistributeBillLogic: nodeDistributeBillLogic,
		nodeService:             nodeService,
	}
}
