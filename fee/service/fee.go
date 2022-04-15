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
	nodeWeekUsageBillLogic  *logic.NodeWeekUsageBill

	nodeService      nodepb.NodeService
	userService      userpb.UserService
	userGroupService userpb.GroupService
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
		int(nodeApplyInfo.Apply.NodeNum),
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

// GetNodeDistributeFeeRate 查询机器独占费用的价格
func (fs *FeeService) GetNodeDistributeFeeRate(
	ctx context.Context,
	req *feepb.GetNodeDistributeFeeRateRequest,
	resp *feepb.GetNodeDistributeFeeRateResponse,
) error {
	logger.Info("GetNodeDistributeFeeRate: ", req.BaseRequest)
	rates := fs.nodeDistributeBillLogic.GetRate(ctx)
	resp.Rate36CPU = rates.Get36CPURate()
	resp.Rate4GPU = rates.Get4GPU()
	resp.Rate8GPU = rates.Get8GPU()
	return nil
}

// CreateNodeWeekUsageBill 创建机器机时周账单
func (fs *FeeService) CreateNodeWeekUsageBill(
	ctx context.Context,
	req *feepb.CreateNodeWeekUsageBillRequest,
	resp *feepb.CreateNodeWeekUsageBillResponse,
) error {
	logger.Info("CreateNodeWeekUsageBill: ", req.BaseRequest)
	if !verify.Identify(verify.CreateNodeWeekUsageBill, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"CreateNodeWeekUsageBill permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("CreateNodeWeekUsageBill permission forbidden")
	}

	// 查询记录拥有者的信息
	userInfoResp, err := fs.userService.GetUserInfo(ctx, &userpb.GetUserInfoRequest{
		BaseRequest: req.BaseRequest,
		Userid:      req.UserID,
	})
	if err != nil {
		return err
	}

	id, err := fs.nodeWeekUsageBillLogic.CreateBill(
		ctx,
		int(userInfoResp.UserInfo.Id),
		int(userInfoResp.UserInfo.GroupId),
		userInfoResp.UserInfo.Username,
		userInfoResp.UserInfo.Name,
		int(req.WallTime),
		int(req.GwallTime),
		req.StartTime,
		req.EndTime,
	)
	if err != nil {
		return err
	}
	resp.Id = int32(id)
	return nil
}

// PaginationGetNodeWeekUsageBillRecords 分页查询机器节点时长账单
func (fs *FeeService) PaginationGetNodeWeekUsageBillRecords(
	ctx context.Context,
	req *feepb.PaginationGetNodeWeekUsageBillRecordsResquest,
	resp *feepb.PaginationGetNodeWeekUsageBillRecordsResponse,
) error {
	logger.Info("PaginationGetNodeWeekUsageBillRecords: ", req.BaseRequest)
	if !verify.Identify(verify.QueryNodeWeekUsageBill, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"QueryNodeWeekUsageBill permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("QueryNodeWeekUsageBill permission forbidden")
	}

	isAdmin := verify.IsAdmin(req.BaseRequest.UserInfo.Levels)
	isTutor := verify.IsTutor(req.BaseRequest.UserInfo.Levels)

	var infos *logic.PaginationGetWeekUsageBillResult
	var err error
	if !isAdmin && !isTutor {
		// 普通学生权限,只能查询自己的所有的账单信息
		infos, err = fs.nodeWeekUsageBillLogic.PaginationGetWithTimeRangeWithUserID(
			ctx,
			int(req.BaseRequest.UserInfo.UserId),
			int(req.PageIndex),
			int(req.PageSize),
			req.StartTimeUnix,
			req.EndTimeUnix,
		)
	} else if !isAdmin && isTutor {
		// 导师权限,能查询自己组的所有的用户的账单的信息
		infos, err = fs.nodeWeekUsageBillLogic.PaginationGetWithTimeRangeWithGroupID(
			ctx,
			int(req.BaseRequest.UserInfo.GroupId),
			int(req.PageIndex),
			int(req.PageSize),
			req.StartTimeUnix,
			req.EndTimeUnix,
		)
	} else {
		// 管理员权限,能查看所有的用户的账单的信息
		infos, err = fs.nodeWeekUsageBillLogic.PaginationGetWithTimeRange(
			ctx,
			int(req.PageIndex),
			int(req.PageSize),
			req.StartTimeUnix,
			req.EndTimeUnix,
		)
	}
	if err != nil {
		return err
	}
	resp.Count = int32(infos.Count)
	resp.Bills = make([]*feepb.NodeWeekUsageBill, len(infos.Data))
	for index := range infos.Data {
		resp.Bills[index] = &feepb.NodeWeekUsageBill{
			Id:          int32(infos.Data[index].ID),
			UserID:      int32(infos.Data[index].UserID),
			Username:    infos.Data[index].Username,
			Name:        infos.Data[index].UserName,
			WallTime:    int32(infos.Data[index].WallTime),
			GwallTime:   int32(infos.Data[index].GWallTime),
			Fee:         infos.Data[index].Fee,
			PayFee:      infos.Data[index].PayFee,
			StartTime:   infos.Data[index].StartTime.Unix(),
			EndTime:     infos.Data[index].EndTime.Unix(),
			PayFlag:     int32(infos.Data[index].PayFlag),
			PayTime:     infos.Data[index].PayTime.Time.Unix(),
			PayType:     int32(infos.Data[index].PayType.Int64),
			PayMessage:  infos.Data[index].PayMessage.String,
			UserGroupID: int32(infos.Data[index].UserGroupID),
			CreateTime:  infos.Data[index].CreateTime.Unix(),
		}
		if infos.Data[index].ExtraAttributes != nil {
			resp.Bills[index].ExtraAttributes = infos.Data[index].ExtraAttributes.String()
		}
	}
	return nil
}

var _ feepb.FeeHandler = (*FeeService)(nil)

// NewFee 创建新的fee服务
func NewFee(
	client client.Client,
	nodeDistributeBillLogic *logic.NodeDistributeBill,
	nodeWeekUsageBillLogic *logic.NodeWeekUsageBill,
) *FeeService {
	nodeService := nodepb.NewNodeService("node", client)
	userService := userpb.NewUserService("user", client)
	userGroupService := userpb.NewGroupService("user", client)
	return &FeeService{
		nodeDistributeBillLogic: nodeDistributeBillLogic,
		nodeWeekUsageBillLogic:  nodeWeekUsageBillLogic,
		nodeService:             nodeService,
		userService:             userService,
		userGroupService:        userGroupService,
	}
}
