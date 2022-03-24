package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/essayZW/hpcmanager/logger"
	nodebroker "github.com/essayZW/hpcmanager/node/broker"
	"github.com/essayZW/hpcmanager/node/logic"
	nodepb "github.com/essayZW/hpcmanager/node/proto"
	publicproto "github.com/essayZW/hpcmanager/proto"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/essayZW/hpcmanager/verify"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/client"
)

// NodeService 机器管理服务
type NodeService struct {
	nodeApplyLogic *logic.NodeApply
	nodeDistribute *logic.NodeDistribute

	userGroupService userpb.GroupService

	rabbitmqBroker broker.Broker
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
		NodeType:  req.NodeType,
		NodeNum:   int(req.NodeNum),
		StartTime: time.UnixMicro(req.StartTime),
		EndTime:   time.UnixMicro(req.EndTime),
	}, int(req.ProjectID))
	if err != nil {
		return fmt.Errorf("create node apply info error: %s", err.Error())
	}
	resp.Id = int32(id)
	return nil
}

// PaginationGetNodeApply 分页查询用户申请机器节点包机申请表
func (ns *NodeService) PaginationGetNodeApply(ctx context.Context, req *nodepb.PaginationGetNodeApplyRequest, resp *nodepb.PaginationGetNodeApplyResponse) error {
	logger.Info("PaginationGetNodeApply: ", req.BaseRequest)
	if !verify.Identify(verify.GetNodeApplyInfo, req.BaseRequest.UserInfo.Levels) {
		logger.Info("PaginationGetNodeApply permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("PaginationGetNodeApply permission forbidden")
	}
	isAdmin := verify.IsAdmin(req.BaseRequest.UserInfo.Levels)
	isTutor := verify.IsTutor(req.BaseRequest.UserInfo.Levels)
	var paginationRes *logic.PaginationGetResult
	var err error
	if !isAdmin && !isTutor {
		// 普通用户,只可以查看自己的包机申请记录
		paginationRes, err = ns.nodeApplyLogic.PaginationGetByCreaterID(ctx, int(req.BaseRequest.UserInfo.UserId), int(req.PageIndex), int(req.PageSize))
	} else if !isAdmin && isTutor {
		// 导师用户,可以查看自己组的所有的包机申请记录
		paginationRes, err = ns.nodeApplyLogic.PaginationGetByTutorID(ctx, int(req.BaseRequest.UserInfo.UserId), int(req.PageIndex), int(req.PageSize))
	} else {
		// 管理员用户,可以查看所有人的并且已经被导师审核通过的包机申请记录
		paginationRes, err = ns.nodeApplyLogic.PaginationWithTutorChecked(ctx, int(req.PageIndex), int(req.PageSize))
	}
	if err != nil {
		return err
	}
	resp.Applies = make([]*nodepb.NodeApply, 0)
	resp.Count = int32(paginationRes.Count)
	for _, apply := range paginationRes.Data {
		respApply := &nodepb.NodeApply{
			Id:                     int32(apply.ID),
			CreateTime:             apply.CreateTime.Unix(),
			CreaterID:              int32(apply.CreaterID),
			CreaterUsername:        apply.CreaterUsername,
			CreaterName:            apply.CreaterName,
			ProjectID:              int32(apply.ProjectID),
			TutorCheckStatus:       int32(apply.TutorCheckStatus),
			ManagerCheckStatus:     int32(apply.ManagerCheckStatus),
			Status:                 int32(apply.Status),
			MessageTutor:           apply.MessageTutor.String,
			MessageManager:         apply.MessageManager.String,
			TutorCheckTime:         apply.TutorCheckTime.Time.Unix(),
			TutorID:                int32(apply.TutorID),
			TutorUsername:          apply.TutorUsername,
			TutorName:              apply.TutorName,
			ManagerCheckTime:       apply.ManagerCheckTime.Time.Unix(),
			ManagerCheckerID:       int32(apply.ManagerCheckerID.Int64),
			ManagerCheckerUsername: apply.ManagerCheckerUsername.String,
			ManagerCheckerName:     apply.ManagerCheckerName.String,
			ModifyTime:             apply.ModifyTime.Time.Unix(),
			ModifyUserID:           int32(apply.ModifyUserID),
			ModifyName:             apply.ModifyName,
			ModifyUsername:         apply.ModifyUsername,
			NodeType:               apply.NodeType,
			NodeNum:                int32(apply.NodeNum),
			StartTime:              apply.StartTime.Unix(),
			EndTime:                apply.EndTime.Unix(),
		}
		if apply.ExtraAttributes != nil {
			respApply.ExtraAttributes = apply.ExtraAttributes.String()
		}
		resp.Applies = append(resp.Applies, respApply)
	}
	return nil
}

// CheckNodeApply 审核机器节点申请
func (ns *NodeService) CheckNodeApply(ctx context.Context, req *nodepb.CheckNodeApplyRequest, resp *nodepb.CheckNodeApplyResponse) error {
	logger.Info("CheckNodeApply: ", req.BaseRequest)
	if !verify.Identify(verify.CheckNodeApply, req.BaseRequest.UserInfo.Levels) {
		logger.Info("CheckNodeApply permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("CheckNodeApply permission forbidden")
	}

	isAdmin := verify.IsAdmin(req.BaseRequest.UserInfo.Levels)
	isTutor := verify.IsTutor(req.BaseRequest.UserInfo.Levels)

	var status bool
	var err error
	if req.TutorCheck {
		if isTutor {
			status, err = ns.nodeApplyLogic.CheckNodeApplyByTutor(ctx, int(req.ApplyID), req.CheckStatus, req.CheckMessage)
		} else {
			err = errors.New("must be tutor")
		}
	} else {
		if isAdmin {
			status, err = ns.nodeApplyLogic.CheckNodeApplyByAdmin(ctx, int(req.ApplyID), req.CheckStatus, req.CheckMessage, &logic.ApplyItemUserInfo{
				ID:       int(req.BaseRequest.UserInfo.UserId),
				Username: req.BaseRequest.UserInfo.Username,
				Name:     req.BaseRequest.UserInfo.Name,
			})
		} else {
			err = errors.New("must be admin")
		}
	}
	if err != nil {
		return err
	}
	if !status {
		return errors.New("check error")
	}
	resp.Success = true

	// 发送MQ消息
	message := &nodebroker.CheckApplyMessage{
		CheckStatus:  req.CheckStatus,
		ApplyID:      int(req.ApplyID),
		CheckMessage: req.CheckMessage,
		TutorCheck:   req.TutorCheck,
	}
	if err := message.Public(ns.rabbitmqBroker, req.BaseRequest); err != nil {
		logger.Warn("Message public error: ", err)
	}
	return nil
}

// CreateNodeDistributeWO 创建机器节点分配处理工单
func (ns *NodeService) CreateNodeDistributeWO(ctx context.Context, req *nodepb.CreateNodeDistributeWORequest, resp *nodepb.CreateNodeDistributeWOResponse) error {
	logger.Info("CreateNodeDistributeWO: ", req.BaseRequest)
	if !verify.Identify(verify.CreateNodeDistributeWO, req.BaseRequest.UserInfo.Levels) {
		logger.Info("CreateNodeDistributeWO permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("CreateNodeDistributeWO permission forbidden")
	}

	id, err := ns.nodeDistribute.CreateNodeDistributeWO(context.Background(), int(req.ApplyID))
	if err != nil {
		return err
	}
	resp.Id = int32(id)
	return nil
}

// PaginationGetNodeDistributeWO 分页查询包机处理工单信息
func (ns *NodeService) PaginationGetNodeDistributeWO(ctx context.Context, req *nodepb.PaginationGetNodeDistributeWORequest, resp *nodepb.PaginationGetNodeDistributeWOResponse) error {
	logger.Info("PaginationGetNodeDistributeWO: ", req.BaseRequest)
	if !verify.Identify(verify.QueryNodeDistributeWO, req.BaseRequest.UserInfo.Levels) {
		logger.Info("PaginationGetNodeDistributeWO permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("PaginationGetNodeDistributeWO permission forbidden")
	}

	infos, err := ns.nodeDistribute.PaginationGet(ctx, int(req.PageIndex), int(req.PageSize))
	if err != nil {
		return err
	}

	resp.Count = int32(infos.Count)
	resp.Wos = make([]*nodepb.NodeDistribute, len(infos.Data))

	for i := range infos.Data {
		resp.Wos[i] = &nodepb.NodeDistribute{
			Id:               int32(infos.Data[i].ID),
			ApplyID:          int32(infos.Data[i].ApplyID),
			HandlerFlag:      int32(infos.Data[i].HandlerFlag),
			HandlerUserID:    int32(infos.Data[i].HandlerUserID.Int64),
			HandlerUsername:  infos.Data[i].HandlerUserName.String,
			HandlerName:      infos.Data[i].HandlerUserName.String,
			DistributeBillID: int32(infos.Data[i].DistributeBillID),
			CreateTime:       infos.Data[i].CreateTime.Unix(),
		}
		if infos.Data[i].ExtraAttributes != nil {
			resp.Wos[i].ExtraAttributes = infos.Data[i].ExtraAttributes.String()
		}
	}
	return nil
}

// GetNodeApplyInfoByID 通过申请ID查询申请的具体信息
func (ns *NodeService) GetNodeApplyByID(ctx context.Context, req *nodepb.GetNodeApplyByIDRequest, resp *nodepb.GetNodeApplyByIDResponse) error {
	logger.Info("GetNodeApplyByID: ", req.BaseRequest)
	if !verify.Identify(verify.GetNodeApplyInfo, req.BaseRequest.UserInfo.Levels) {
		logger.Info("GetNodeApplyByID PaginationGetNodeApply permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("GetNodeApplyByID permission forbidden")
	}

	info, err := ns.nodeApplyLogic.GetNodeApplyByID(ctx, int(req.ApplyID))
	if err != nil {
		return err
	}
	// 进行不同权限的角色查询权限的校验
	isAdmin := verify.IsAdmin(req.BaseRequest.UserInfo.Levels)
	isTutor := verify.IsTutor(req.BaseRequest.UserInfo.Levels)
	if !isAdmin && !isTutor {
		// 普通用户
		if info.CreaterID != int(req.BaseRequest.UserInfo.UserId) {
			return errors.New("common user permission forbidden")
		}
	}

	if !isAdmin && isTutor {
		// 导师用户
		if info.TutorID != int(req.BaseRequest.UserInfo.UserId) {
			return errors.New("tutor permission forbidden")
		}
	}

	resp.Apply = &nodepb.NodeApply{
		Id:                     int32(info.ID),
		CreateTime:             info.CreateTime.Unix(),
		CreaterID:              int32(info.CreaterID),
		CreaterUsername:        info.CreaterUsername,
		CreaterName:            info.CreaterName,
		ProjectID:              int32(info.ProjectID),
		TutorCheckStatus:       int32(info.TutorCheckStatus),
		ManagerCheckStatus:     int32(info.ManagerCheckStatus),
		Status:                 int32(info.Status),
		MessageTutor:           info.MessageTutor.String,
		MessageManager:         info.MessageManager.String,
		TutorCheckTime:         info.TutorCheckTime.Time.Unix(),
		TutorID:                int32(info.TutorID),
		TutorUsername:          info.TutorUsername,
		TutorName:              info.TutorName,
		ManagerCheckTime:       info.ManagerCheckTime.Time.Unix(),
		ManagerCheckerID:       int32(info.ManagerCheckerID.Int64),
		ManagerCheckerUsername: info.ManagerCheckerUsername.String,
		ManagerCheckerName:     info.ManagerCheckerName.String,
		ModifyTime:             info.ModifyTime.Time.Unix(),
		ModifyUserID:           int32(info.ModifyUserID),
		ModifyName:             info.ModifyName,
		ModifyUsername:         info.ModifyUsername,
		NodeType:               info.NodeType,
		NodeNum:                int32(info.NodeNum),
		StartTime:              info.StartTime.Unix(),
		EndTime:                info.EndTime.Unix(),
	}
	if info.ExtraAttributes != nil {
		resp.Apply.ExtraAttributes = info.ExtraAttributes.String()
	}
	return nil
}

var _ nodepb.NodeHandler = (*NodeService)(nil)

// NewNode 创建新的机器节点管理服务
func NewNode(client client.Client, nodeApplyLogic *logic.NodeApply, nodeDistribute *logic.NodeDistribute, rabbitmqBroker broker.Broker) *NodeService {
	userGroupService := userpb.NewGroupService("user", client)
	return &NodeService{
		nodeApplyLogic:   nodeApplyLogic,
		nodeDistribute:   nodeDistribute,
		userGroupService: userGroupService,
		rabbitmqBroker:   rabbitmqBroker,
	}
}
