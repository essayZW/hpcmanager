package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/essayZW/hpcmanager/logger"
	"github.com/essayZW/hpcmanager/node/logic"
	nodepb "github.com/essayZW/hpcmanager/node/proto"
	publicproto "github.com/essayZW/hpcmanager/proto"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/essayZW/hpcmanager/verify"
	"go-micro.dev/v4/client"
)

// NodeService 机器管理服务
type NodeService struct {
	nodeApplyLogic   *logic.NodeApply
	userGroupService userpb.GroupService
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
	// 查询用户所属组的相关信息
	// 临时赋予其admin权限
	req.BaseRequest.UserInfo.Levels = append(req.BaseRequest.UserInfo.Levels, int32(verify.SuperAdmin))
	groupResp, err := ns.userGroupService.GetGroupInfoByID(ctx, &userpb.GetGroupInfoByIDRequest{
		BaseRequest: req.BaseRequest,
		GroupID:     req.BaseRequest.UserInfo.GroupId,
	})
	// 取消其临时赋予的管理员权限
	req.BaseRequest.UserInfo.Levels = req.BaseRequest.UserInfo.Levels[:len(req.BaseRequest.UserInfo.Levels)-1]
	if err != nil {
		return err
	}
	id, err := ns.nodeApplyLogic.CreateNodeApply(ctx, &logic.ApplyItemUserInfo{
		ID:       int(req.BaseRequest.UserInfo.UserId),
		Username: req.BaseRequest.UserInfo.Username,
		Name:     req.BaseRequest.UserInfo.Name,
	}, &logic.ApplyItemUserInfo{
		ID:       int(groupResp.GroupInfo.TutorID),
		Username: groupResp.GroupInfo.TutorUsername,
		Name:     groupResp.GroupInfo.TutorName,
	}, &logic.ApplyNodeInfo{
		NodeType: req.NodeType,
		NodeNum:  int(req.NodeNum),
	}, int(req.ProjectID))
	if err != nil {
		return fmt.Errorf("create node apply info error: %s", err.Error())
	}
	resp.Id = int32(id)
	return nil
}

var _ nodepb.NodeHandler = (*NodeService)(nil)

// NewNode 创建新的机器节点管理服务
func NewNode(client client.Client, nodeApplyLogic *logic.NodeApply) *NodeService {
	userGroupService := userpb.NewGroupService("user", client)
	return &NodeService{
		nodeApplyLogic:   nodeApplyLogic,
		userGroupService: userGroupService,
	}
}
