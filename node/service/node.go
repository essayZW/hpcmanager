package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	hpcdb "github.com/essayZW/hpcmanager/db"
	feepb "github.com/essayZW/hpcmanager/fee/proto"
	"github.com/essayZW/hpcmanager/logger"
	nodebroker "github.com/essayZW/hpcmanager/node/broker"
	"github.com/essayZW/hpcmanager/node/db"
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
	nodeUsageTime  *logic.NodeUsageTime

	userGroupService userpb.GroupService
	feeService       feepb.FeeService

	rabbitmqBroker broker.Broker
}

// Ping ping测试
func (ns *NodeService) Ping(
	ctx context.Context,
	req *publicproto.Empty,
	resp *publicproto.PingResponse,
) error {
	logger.Info("NodeService PING ", req)
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	return nil
}

// CreateNodeApply 创建一个新的申请计算节点记录
func (ns *NodeService) CreateNodeApply(ctx context.Context, req *nodepb.CreateNodeApplyRequest,
	resp *nodepb.CreateNodeApplyResponse,
) error {
	logger.Info("CreateNodeApply: ", req.BaseRequest)
	if !verify.Identify(verify.CreateNodeApply, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"CreateNodeApply permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("CreateNodeApply permission forbidden")
	}
	// 查询用户所属组的相关信息
	// 临时赋予其admin权限
	req.BaseRequest.UserInfo.Levels = append(
		req.BaseRequest.UserInfo.Levels,
		int32(verify.SuperAdmin),
	)
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
		StartTime: time.UnixMilli(req.StartTime),
		EndTime:   time.UnixMilli(req.EndTime),
	}, int(req.ProjectID))
	if err != nil {
		return fmt.Errorf("create node apply info error: %s", err.Error())
	}
	resp.Id = int32(id)
	return nil
}

// PaginationGetNodeApply 分页查询用户申请机器节点包机申请表
func (ns *NodeService) PaginationGetNodeApply(
	ctx context.Context,
	req *nodepb.PaginationGetNodeApplyRequest,
	resp *nodepb.PaginationGetNodeApplyResponse,
) error {
	logger.Info("PaginationGetNodeApply: ", req.BaseRequest)
	if !verify.Identify(verify.GetNodeApplyInfo, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"PaginationGetNodeApply permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("PaginationGetNodeApply permission forbidden")
	}
	isAdmin := verify.IsAdmin(req.BaseRequest.UserInfo.Levels)
	isTutor := verify.IsTutor(req.BaseRequest.UserInfo.Levels)
	var paginationRes *logic.PaginationGetResult
	var err error
	if !isAdmin && !isTutor {
		// 普通用户,只可以查看自己的包机申请记录
		paginationRes, err = ns.nodeApplyLogic.PaginationGetByCreaterID(
			ctx,
			int(req.BaseRequest.UserInfo.UserId),
			int(req.PageIndex),
			int(req.PageSize),
		)
	} else if !isAdmin && isTutor {
		// 导师用户,可以查看自己组的所有的包机申请记录
		paginationRes, err = ns.nodeApplyLogic.PaginationGetByTutorID(ctx, int(req.BaseRequest.UserInfo.UserId), int(req.PageIndex), int(req.PageSize))
	} else {
		// 管理员用户,可以查看所有人的并且已经被导师审核通过的包机申请记录
		paginationRes, err = ns.nodeApplyLogic.PaginationGet(ctx, int(req.PageIndex), int(req.PageSize))
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
func (ns *NodeService) CheckNodeApply(
	ctx context.Context,
	req *nodepb.CheckNodeApplyRequest,
	resp *nodepb.CheckNodeApplyResponse,
) error {
	logger.Info("CheckNodeApply: ", req.BaseRequest)
	if !verify.Identify(verify.CheckNodeApply, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"CheckNodeApply permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("CheckNodeApply permission forbidden")
	}

	isAdmin := verify.IsAdmin(req.BaseRequest.UserInfo.Levels)
	isTutor := verify.IsTutor(req.BaseRequest.UserInfo.Levels)

	var status bool
	var err error
	if req.TutorCheck {
		if isTutor {
			status, err = ns.nodeApplyLogic.CheckNodeApplyByTutor(
				ctx,
				int(req.ApplyID),
				req.CheckStatus,
				req.CheckMessage,
			)
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
func (ns *NodeService) CreateNodeDistributeWO(
	ctx context.Context,
	req *nodepb.CreateNodeDistributeWORequest,
	resp *nodepb.CreateNodeDistributeWOResponse,
) error {
	logger.Info("CreateNodeDistributeWO: ", req.BaseRequest)
	if !verify.Identify(verify.CreateNodeDistributeWO, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"CreateNodeDistributeWO permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
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
func (ns *NodeService) PaginationGetNodeDistributeWO(
	ctx context.Context,
	req *nodepb.PaginationGetNodeDistributeWORequest,
	resp *nodepb.PaginationGetNodeDistributeWOResponse,
) error {
	logger.Info("PaginationGetNodeDistributeWO: ", req.BaseRequest)
	if !verify.Identify(verify.QueryNodeDistributeWO, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"PaginationGetNodeDistributeWO permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
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
			HandlerUsername:  infos.Data[i].HandlerUsername.String,
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
func (ns *NodeService) GetNodeApplyByID(
	ctx context.Context,
	req *nodepb.GetNodeApplyByIDRequest,
	resp *nodepb.GetNodeApplyByIDResponse,
) error {
	logger.Info("GetNodeApplyByID: ", req.BaseRequest)
	if !verify.Identify(verify.GetNodeApplyInfo, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"GetNodeApplyByID PaginationGetNodeApply permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
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

// FinishNodeDistributeWO 处理机器节点分配工单
func (ns *NodeService) FinishNodeDistributeWO(
	ctx context.Context,
	req *nodepb.FinishNodeDistributeWORequest,
	resp *nodepb.FinishNodeDistributeWOResponse,
) error {
	logger.Info("FinishNodeDistributeWO: ", req.BaseRequest)
	if !verify.Identify(verify.FinishNodeDistributeWO, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"FinishNodeDistributeWO PaginationGetNodeApply permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("FinishNodeDistributeWO permission forbidden")
	}

	status, err := hpcdb.Transaction(context.Background(), func(c context.Context, i ...interface{}) (interface{}, error) {
		status, err := ns.nodeDistribute.FinishByID(c, int(req.DistributeID), &logic.SimpleUserInfo{
			ID:       int(req.BaseRequest.UserInfo.UserId),
			Username: req.BaseRequest.UserInfo.Username,
			Name:     req.BaseRequest.UserInfo.Name,
		})
		if err != nil {
			return status, err
		}
		// 创建对应的机器独占账单
		res, err := ns.feeService.CreateNodeDistributeBill(c, &feepb.CreateNodeDistributeBillRequest{
			BaseRequest:      req.BaseRequest,
			NodeDistributeID: req.DistributeID,
		})
		logger.Info("create node distribute bill with id: ", res.Id)
		return status, err
	})
	resp.Success = status.(bool)
	return err
}

// RevokeNodeApply 撤销机器节点申请
func (ns *NodeService) RevokeNodeApply(
	ctx context.Context,
	req *nodepb.RevokeNodeApplyRequest,
	resp *nodepb.RevokeNodeApplyResponse,
) error {
	logger.Info("RevokeNodeApply: ", req.BaseRequest)
	if !verify.Identify(verify.RevokeNodeApply, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"RevokeNodeApply PaginationGetNodeApply permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("RevokeNodeApply permission forbidden")
	}
	info, err := ns.nodeApplyLogic.GetNodeApplyByID(ctx, int(req.ApplyID))
	if err != nil {
		return errors.New("apply info query error, invalid apply id")
	}

	if info.CreaterID != int(req.BaseRequest.UserInfo.UserId) {
		return errors.New("permission forbidden")
	}

	if info.ManagerCheckStatus != -1 || info.Status != 1 {
		return errors.New("can't revoke this apply")
	}
	status, err := ns.nodeApplyLogic.RevokeNodeApply(ctx, int(req.ApplyID))
	if err != nil {
		return err
	}
	resp.Success = status
	return nil
}

// AddNodeUsageTimeRecord 添加机器节点使用时间记录
func (ns *NodeService) AddNodeUsageTimeRecord(
	ctx context.Context,
	req *nodepb.AddNodeUsageTimeRecordRequest,
	resp *nodepb.AddNodeUsageTimeRecordResponse,
) error {
	logger.Info("AddNodeUsageTimeRecord: ", req.BaseRequest)
	if !verify.Identify(verify.AddNodeUsage, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"AddNodeUsageTimeRecord PaginationGetNodeApply permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("AddNodeUsageTimeRecord permission forbidden")
	}
	id, err := ns.nodeUsageTime.AddRecord(ctx, &db.HpcUsageTime{
		UserID:        int(req.UserID),
		Username:      req.Username,
		UserName:      req.Name,
		HpcUsername:   req.HpcUserName,
		TutorID:       int(req.TutorID),
		TutorUsername: req.TutorUsername,
		TutorUserName: req.TutorName,
		HpcGroupName:  req.HpcGroupName,
		QueueName:     req.QueueName,
		WallTime:      req.WallTime,
		GWallTime:     req.GwallTime,
		StartTime:     time.Unix(req.StartTimeUnix, 0),
		EndTime:       time.Unix(req.EndTimeUnix, 0),
	})
	if err != nil {
		return err
	}
	// 创建相应的机器时长周账单
	_, err = ns.feeService.CreateNodeWeekUsageBill(ctx, &feepb.CreateNodeWeekUsageBillRequest{
		BaseRequest:           req.BaseRequest,
		NodeWeekUsageRecordID: int32(id),
	})
	if err != nil {
		logger.Warn("create node week usage bill error: ", err)
		return err
	}
	resp.Id = int32(id)
	return nil
}

// PaginationGetNodeUsage 分页查询机器节点使用详情信息
func (ns *NodeService) PaginationGetNodeUsage(
	ctx context.Context,
	req *nodepb.PaginationGetNodeUsageRequest,
	resp *nodepb.PaginationGetNodeUsageResponse,
) error {
	logger.Info("PaginationGetNodeUsage: ", req.BaseRequest)
	if !verify.Identify(verify.QueryNodeUsage, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"QueryNodeUsage permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("QueryNodeUsage permission forbidden")
	}
	isAdmin := verify.IsAdmin(req.BaseRequest.UserInfo.Levels)
	isTutor := verify.IsTutor(req.BaseRequest.UserInfo.Levels)

	var paginationResult *logic.PaginationGetNodeUsageRecordResult
	var err error
	if !isAdmin && !isTutor {
		// 普通用户,只能查询自己的机器时间记录
		paginationResult, err = ns.nodeUsageTime.PaginationGetNodeUsageRecordByUserID(
			context.Background(),
			int(req.BaseRequest.UserInfo.UserId),
			int(req.PageIndex),
			int(req.PageSize),
			req.StartDateMilliUnix,
			req.EndDateMilliUnix,
		)
	} else if !isAdmin && isTutor {
		// 导师用户,只能查看自己组的机器时间记录
		paginationResult, err = ns.nodeUsageTime.PaginationGetNodeUsageRecordByTutorID(
			context.Background(),
			int(req.BaseRequest.UserInfo.UserId),
			int(req.PageIndex),
			int(req.PageSize),
			req.StartDateMilliUnix,
			req.EndDateMilliUnix,
		)
	} else {
		// 管理员用户,可以查看所有的用户的机器时间记录
		paginationResult, err = ns.nodeUsageTime.PaginationGetNodeUsageRecord(
			context.Background(),
			int(req.PageIndex),
			int(req.PageSize),
			req.StartDateMilliUnix,
			req.EndDateMilliUnix,
		)
	}
	if err != nil {
		return err
	}
	resp.Count = int32(paginationResult.Count)
	resp.Usages = make([]*nodepb.NodeUsageTime, len(paginationResult.Data))
	for i := range resp.Usages {
		resp.Usages[i] = &nodepb.NodeUsageTime{
			Id:            int32(paginationResult.Data[i].ID),
			UserID:        int32(paginationResult.Data[i].UserID),
			Username:      paginationResult.Data[i].Username,
			Name:          paginationResult.Data[i].UserName,
			HpcUserName:   paginationResult.Data[i].HpcUsername,
			TutorID:       int32(paginationResult.Data[i].TutorID),
			TutorUsername: paginationResult.Data[i].TutorUsername,
			TutorName:     paginationResult.Data[i].TutorUserName,
			HpcGroupName:  paginationResult.Data[i].HpcGroupName,
			QueueName:     paginationResult.Data[i].QueueName,
			WallTime:      paginationResult.Data[i].WallTime,
			GwallTime:     paginationResult.Data[i].GWallTime,
			StartTime:     paginationResult.Data[i].StartTime.Unix(),
			EndTime:       paginationResult.Data[i].EndTime.Unix(),
			CreateTime:    paginationResult.Data[i].CreateTime.Unix(),
		}
		if paginationResult.Data[i].ExtraAttributes != nil {
			resp.Usages[i].ExtraAttributes = paginationResult.Data[i].ExtraAttributes.String()
		}
	}
	return nil
}

// UpdateNodeApply 更新节点申请信息
func (ns *NodeService) UpdateNodeApply(
	ctx context.Context,
	req *nodepb.UpdateNodeApplyRequest,
	resp *nodepb.UpdateNodeApplyResponse,
) error {
	logger.Info("UpdateNodeApply: ", req.BaseRequest)
	if !verify.Identify(verify.UpdateNodeApply, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"UpdateNodeApply permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("UpdateNodeApply permission forbidden")
	}
	// 如果是管理员可以修改所有的节点申请信息,否则只能修改自己创建的节点申请信息
	isAdmin := verify.IsAdmin(req.BaseRequest.UserInfo.Levels)
	if !isAdmin {
		// 查询需要修改的机器节点申请信息
		info, err := ns.nodeApplyLogic.GetNodeApplyByID(ctx, int(req.NewInfos.Id))
		if err != nil {
			return errors.New("node apply info query fail")
		}
		if info.CreaterID != int(req.BaseRequest.UserInfo.UserId) {
			return errors.New("permission forbidden: not creater")
		}
	}

	status, err := ns.nodeApplyLogic.UpdateNodeApplyInfo(
		ctx,
		int(req.NewInfos.Id),
		int(req.BaseRequest.UserInfo.UserId),
		req.NewInfos.NodeType,
		int(req.NewInfos.NodeNum),
		req.NewInfos.StartTime,
		req.NewInfos.EndTime,
		int(req.BaseRequest.UserInfo.UserId),
		req.BaseRequest.UserInfo.Username,
		req.BaseRequest.UserInfo.Name,
	)
	if err != nil {
		return err
	}
	resp.Success = status
	return nil
}

// GetNodeDistributeInfoByID 通过ID查询机器节点分配工单信息
func (ns *NodeService) GetNodeDistributeInfoByID(
	ctx context.Context,
	req *nodepb.GetNodeDistributeInfoByIDRequest,
	resp *nodepb.GetNodeDistributeInfoByIDResponse,
) error {
	logger.Info("GetNodeDistributeInfoByID: ", req.BaseRequest)
	if !verify.Identify(verify.QueryNodeDistributeWO, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"QueryNodeDistributeWO permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("QueryNodeDistributeWO permission forbidden")
	}

	info, err := ns.nodeDistribute.GetInfoByID(ctx, req.Id)
	if err != nil {
		return err
	}
	resp.Wo = &nodepb.NodeDistribute{
		Id:               int32(info.ID),
		ApplyID:          int32(info.ApplyID),
		HandlerFlag:      int32(info.HandlerFlag),
		HandlerUserID:    int32(info.HandlerUserID.Int64),
		HandlerUsername:  info.HandlerUsername.String,
		HandlerName:      info.HandlerUserName.String,
		DistributeBillID: int32(info.DistributeBillID),
		CreateTime:       info.CreateTime.Unix(),
	}
	if info.ExtraAttributes != nil {
		resp.Wo.ExtraAttributes = info.ExtraAttributes.String()
	}
	return nil
}

// GetNodeUsageTimeRecordByID 通过ID查询机器时长记录
func (ns *NodeService) GetNodeUsageTimeRecordByID(
	ctx context.Context,
	req *nodepb.GetNodeUsageTimeRecordByIDRequest,
	resp *nodepb.GetNodeUsageTimeRecordByIDResponse,
) error {
	logger.Info("GetNodeDistributeInfoByID: ", req.BaseRequest)
	if !verify.Identify(verify.QueryNodeUsage, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"QueryNodeUsage permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("QueryNodeUsage permission forbidden")
	}

	info, err := ns.nodeUsageTime.GetRecordByID(ctx, int(req.Id))
	if err != nil {
		return err
	}

	resp.Record = &nodepb.NodeUsageTime{
		Id:            int32(info.ID),
		UserID:        int32(info.UserID),
		Username:      info.Username,
		Name:          info.UserName,
		HpcUserName:   info.HpcUsername,
		TutorID:       int32(info.TutorID),
		TutorUsername: info.TutorUsername,
		TutorName:     info.TutorUserName,
		HpcGroupName:  info.HpcGroupName,
		QueueName:     info.QueueName,
		WallTime:      info.WallTime,
		GwallTime:     info.GWallTime,
		StartTime:     info.StartTime.Unix(),
		EndTime:       info.EndTime.Unix(),
		CreateTime:    info.CreateTime.Unix(),
	}
	if info.ExtraAttributes != nil {
		resp.Record.ExtraAttributes = info.ExtraAttributes.String()
	}
	return nil
}

var _ nodepb.NodeHandler = (*NodeService)(nil)

// NewNode 创建新的机器节点管理服务
func NewNode(
	client client.Client,
	nodeApplyLogic *logic.NodeApply,
	nodeDistribute *logic.NodeDistribute,
	nodeUsageTime *logic.NodeUsageTime,
	rabbitmqBroker broker.Broker,
) *NodeService {
	userGroupService := userpb.NewGroupService("user", client)
	feeService := feepb.NewFeeService("fee", client)
	return &NodeService{
		nodeApplyLogic:   nodeApplyLogic,
		nodeDistribute:   nodeDistribute,
		userGroupService: userGroupService,
		feeService:       feeService,
		rabbitmqBroker:   rabbitmqBroker,
		nodeUsageTime:    nodeUsageTime,
	}
}
