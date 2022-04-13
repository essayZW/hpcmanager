package service

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/fee/logic"
	feepb "github.com/essayZW/hpcmanager/fee/proto"
	"github.com/essayZW/hpcmanager/logger"
	nodepb "github.com/essayZW/hpcmanager/node/proto"
	publicproto "github.com/essayZW/hpcmanager/proto"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/essayZW/hpcmanager/verify"
	"go-micro.dev/v4/client"
)

type FeeService struct {
	nodeDistributeBillLogic *logic.NodeDistributeBill
	nodeService             nodepb.NodeService
	userService             userpb.UserService
	userGroupService        userpb.GroupService
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

	// 查询用户的组信息
	userResp, err := fs.userService.GetUserInfo(ctx, &userpb.GetUserInfoRequest{
		BaseRequest: req.BaseRequest,
		Userid:      nodeApplyInfo.Apply.CreaterID,
	})
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
		int(userResp.UserInfo.GroupId),
	)
	if err != nil {
		return err
	}
	resp.Id = int32(id)
	return nil
}

// PaginationGetNodeDistributeBill 分页查询计算节点独占账单
func (fs *FeeService) PaginationGetNodeDistributeBill(
	ctx context.Context,
	req *feepb.PaginationGetNodeDistributeBillRequest,
	resp *feepb.PaginationGetNodeDistributeBillResponse,
) error {
	logger.Info("PaginationGetNodeDistributeBill: ", req.BaseRequest)
	if !verify.Identify(verify.QueryNodeDistributeBill, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"QueryNodeDistributeBill permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("QueryNodeDistributeBill permission forbidden")
	}
	var infos *logic.PaginationGetNodeDistributeBillResult
	var err error

	isAdmin := verify.IsAdmin(req.BaseRequest.UserInfo.Levels)
	isTutor := verify.IsTutor(req.BaseRequest.UserInfo.Levels)
	if !isAdmin && !isTutor {
		// 普通用户,只能查询自己的账单信息
		infos, err = fs.nodeDistributeBillLogic.PaginationGetWithUserID(
			ctx,
			int(req.PageIndex),
			int(req.PageSize),
			int(req.BaseRequest.UserInfo.UserId),
		)
	} else if !isAdmin && isTutor {
		// 导师用户,只能查询自己组所有用户的账单信息
		infos, err = fs.nodeDistributeBillLogic.PaginationGetWithGroupID(
			ctx,
			int(req.PageIndex),
			int(req.PageSize),
			int(req.BaseRequest.UserInfo.GroupId),
		)
	} else {
		// 管理员用户,可以查询所有的信息
		infos, err = fs.nodeDistributeBillLogic.PaginationGetAll(
			ctx,
			int(req.PageIndex),
			int(req.PageSize),
		)
	}

	if err != nil {
		return err
	}

	resp.Count = int32(infos.Count)
	resp.Bills = make([]*feepb.NodeDistributeBill, len(infos.Data))
	for i := range infos.Data {
		resp.Bills[i] = &feepb.NodeDistributeBill{
			Id:                  int32(infos.Data[i].ID),
			ApplyID:             int32(infos.Data[i].ApplyID),
			NodeDistributeID:    int32(infos.Data[i].NodeDistributeID),
			Fee:                 infos.Data[i].Fee,
			PayFee:              infos.Data[i].PayFee,
			PayFlag:             int32(infos.Data[i].PayFlag),
			PayType:             int32(infos.Data[i].PayType.Int64),
			PayMessage:          infos.Data[i].PayMessage.String,
			UserID:              int32(infos.Data[i].UserID),
			UserUsername:        infos.Data[i].Username,
			UserName:            infos.Data[i].UserName,
			UserGroupID:         int32(infos.Data[i].UserGroupID),
			CreateTimeMilliUnix: infos.Data[i].CreateTime.UnixMilli(),
		}
		if infos.Data[i].ExtraAttributes != nil {
			resp.Bills[i].ExtraAttributes = infos.Data[i].ExtraAttributes.String()
		}
	}
	return nil
}

// PayNodeDistributeBill 支持机器节点独占账单
func (fs *FeeService) PayNodeDistributeBill(
	ctx context.Context,
	req *feepb.PayNodeDistributeBillRequest,
	resp *feepb.PayNodeDistributeBillResponse,
) error {
	logger.Info("PayNodeDistributeBill:", req.BaseRequest)
	if !verify.Identify(verify.PayNodeDistributeBill, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"PayNodeDistributeBill permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("PayNodeDistributeBill permission forbidden")
	}

	status, err := db.Transaction(context.Background(), func(c context.Context, i ...interface{}) (interface{}, error) {
		status, err := fs.nodeDistributeBillLogic.PayBill(
			c,
			int(req.Id),
			req.PayMoney,
			req.PayMessage,
			logic.PayType(req.PayType),
		)
		if err != nil {
			return status, err
		}

		if logic.OfflinePay == req.PayType {
			// 是线下缴费
			return status, nil
		}
		// 余额缴费
		bill, err := fs.nodeDistributeBillLogic.GetInfoByID(c, int(req.Id))
		if err != nil {
			return false, err
		}

		_, err = fs.userGroupService.AddBalance(c, &userpb.AddBalanceRequest{
			BaseRequest: req.BaseRequest,
			GroupID:     int32(bill.UserGroupID),
			Money:       -req.PayMoney,
		})
		if err != nil {
			return false, err
		}
		return status, nil
	})
	resp.Success = status.(bool)
	return err
}

var _ feepb.FeeHandler = (*FeeService)(nil)

// NewFee 创建新的fee服务
func NewFee(client client.Client, nodeDistributeBillLogic *logic.NodeDistributeBill) *FeeService {
	nodeService := nodepb.NewNodeService("node", client)
	userService := userpb.NewUserService("user", client)
	userGroupService := userpb.NewGroupService("user", client)
	return &FeeService{
		nodeDistributeBillLogic: nodeDistributeBillLogic,
		nodeService:             nodeService,
		userService:             userService,
		userGroupService:        userGroupService,
	}
}
