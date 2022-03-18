package service

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/logger"
	"github.com/essayZW/hpcmanager/node/logic"
	nodepb "github.com/essayZW/hpcmanager/node/proto"
	publicproto "github.com/essayZW/hpcmanager/proto"
	"github.com/essayZW/hpcmanager/verify"
	"go-micro.dev/v4/client"
)

// NodeService 机器管理服务
type NodeService struct {
	nodeApplyLogic *logic.NodeApply
}

// Ping ping测试
func (ns *NodeService) Ping(ctx context.Context, req *publicproto.Empty, resp *publicproto.PingResponse) error {
	logger.Info("NodeService PING ", req)
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	return nil
}

// CreateNodeApply 创建一个新的申请计算节点记录
func (ns *NodeService) CreateNodeApply(ctx context.Context, req *nodepb.CreateNodeApplyRequest, resp *nodepb.CreateNodeApplyResponse) error {
	logger.Info("CreateNodeApply: ", req.BaseRequest)
	if !verify.Identify(verify.CreateNodeApply, req.BaseRequest.UserInfo.Levels) {
		logger.Info("CreateNodeApply permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("CreateNodeApply permission forbidden")
	}
	return nil
}

var _ nodepb.NodeHandler = (*NodeService)(nil)

// NewNode 创建新的机器节点管理服务
func NewNode(client client.Client, nodeApplyLogic *logic.NodeApply) *NodeService {
	return &NodeService{
		nodeApplyLogic: nodeApplyLogic,
	}
}
