package service

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/award/logic"
	awardpb "github.com/essayZW/hpcmanager/award/proto"
	"github.com/essayZW/hpcmanager/logger"
	publicproto "github.com/essayZW/hpcmanager/proto"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/essayZW/hpcmanager/verify"
	"go-micro.dev/v4/client"
)

type AwardService struct {
	paperAwardLogic *logic.Paper

	userService      userpb.UserService
	userGroupService userpb.GroupService
}

// Ping ping测试
func (as *AwardService) Ping(ctx context.Context, req *publicproto.Empty, resp *publicproto.PingResponse) error {
	logger.Info("AwardService PING ", req)
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	return nil
}

// CreatePaperAward 创建论文奖励申请
func (as *AwardService) CreatePaperAward(
	ctx context.Context,
	req *awardpb.CreatePaperAwardRequest,
	resp *awardpb.CreatePaperAwardResponse,
) error {
	logger.Info("CreatePaperAward: ", req.BaseRequest)
	if !verify.Identify(verify.CreatePaperAward, req.GetBaseRequest().GetUserInfo().GetLevels()) {
		logger.Info(
			"CreatePaperAward permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("CreatePaperAward permission forbidden")
	}
	// 查询用户所在组的信息
	req.BaseRequest.UserInfo.Levels = append(
		req.BaseRequest.UserInfo.Levels,
		int32(verify.SuperAdmin),
	)
	groupResp, err := as.userGroupService.GetGroupInfoByID(ctx, &userpb.GetGroupInfoByIDRequest{
		BaseRequest: req.BaseRequest,
		GroupID:     req.BaseRequest.UserInfo.GroupId,
	})
	// 取消其临时赋予的管理员权限
	req.BaseRequest.UserInfo.Levels = req.BaseRequest.UserInfo.Levels[:len(req.BaseRequest.UserInfo.Levels)-1]
	if err != nil {
		return errors.New("query group info error")
	}
	id, err := as.paperAwardLogic.CreateNew(ctx, &logic.UserInfoParam{
		ID:       int(req.BaseRequest.UserInfo.UserId),
		Username: req.BaseRequest.UserInfo.Username,
		Name:     req.BaseRequest.UserInfo.Name,
	}, &logic.UserInfoParam{
		ID:       int(groupResp.GroupInfo.TutorID),
		Username: groupResp.GroupInfo.TutorUsername,
		Name:     groupResp.GroupInfo.TutorName,
	}, int(req.BaseRequest.UserInfo.GroupId), &logic.PaperInfoParam{
		Title:               req.Title,
		Category:            req.Category,
		Partition:           req.Partition,
		FirstPageImageName:  req.FirstPageImageName,
		ThanksPageImageName: req.ThanksPageImageName,
		RemarkMessage:       req.RemarkMessage,
	})
	if err != nil {
		return err
	}
	resp.Id = int32(id)
	return nil
}

// PaginationGetPaperApply 分页查询论文奖励申请记录
func (as *AwardService) PaginationGetPaperApply(
	ctx context.Context,
	req *awardpb.PaginationGetPaperApplyRequest,
	resp *awardpb.PaginationGetPaperApplyResponse,
) error {
	logger.Info("PaginationGetPaperApply: ", req.BaseRequest)
	if !verify.Identify(verify.QueryPaperAwardApply, req.GetBaseRequest().GetUserInfo().GetLevels()) {
		logger.Info(
			"QueryPaperAwardApply permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("QueryPaperAwardApply permission forbidden")
	}
	var infos *logic.PaginationGetPaperApplyResult
	var err error
	isAdmin := verify.IsAdmin(req.BaseRequest.UserInfo.Levels)
	isTutor := verify.IsTutor(req.BaseRequest.UserInfo.Levels)
	if !isAdmin && !isTutor {
		// 普通学生用户
		infos, err = as.paperAwardLogic.PaginationGetByCreaterID(
			ctx,
			int(req.BaseRequest.UserInfo.UserId),
			int(req.PageIndex),
			int(req.PageSize),
		)
	} else if !isAdmin && isTutor {
		// 导师用户
		infos, err = as.paperAwardLogic.PaginationGetByGroupID(ctx, int(req.BaseRequest.UserInfo.GroupId), int(req.PageIndex), int(req.PageSize))
	} else {
		// 管理员用户
		infos, err = as.paperAwardLogic.PaginationGetAll(ctx, int(req.PageIndex), int(req.PageSize))
	}
	if err != nil {
		return err
	}
	resp.Applies = make([]*awardpb.PaperApply, len(infos.Data))
	resp.Count = int32(infos.Count)
	for i := range infos.Data {
		resp.Applies[i] = &awardpb.PaperApply{
			Id:                       int32(infos.Data[i].ID),
			CreaterID:                int32(infos.Data[i].CreaterID),
			CreaterUsername:          infos.Data[i].CreaterUsername,
			CreaterName:              infos.Data[i].CreaterName,
			UserGroupID:              int32(infos.Data[i].UserGroupID),
			TutorID:                  int32(infos.Data[i].TutorID),
			TutorUsername:            infos.Data[i].TutorUsername,
			TutorName:                infos.Data[i].TutorName,
			PaperTitle:               infos.Data[i].PaperTitle,
			PaperCategory:            infos.Data[i].PaperCategory,
			PaperPartition:           infos.Data[i].PaperPartition,
			PaperFirstPageImageName:  infos.Data[i].PaperFirstPageImageName,
			PaperThanksPageImageName: infos.Data[i].PaperThanksPageImageName,
			RemarkMessage:            infos.Data[i].RemarkMessage,
			CheckStatus:              int32(infos.Data[i].CheckStatus.Int64),
			CheckerID:                int32(infos.Data[i].CheckerID.Int64),
			CheckerUsername:          infos.Data[i].CheckerUsername.String,
			CheckerName:              infos.Data[i].CheckerName.String,
			CheckMoney:               infos.Data[i].CheckMoney,
			CheckMessage:             infos.Data[i].CheckMessage.String,
			CheckTimeUnix:            infos.Data[i].CheckTime.Time.Unix(),
		}
		if infos.Data[i].ExtraAttributes != nil {
			resp.Applies[i].ExtraAttributes = infos.Data[i].ExtraAttributes.String()
		}
	}
	return nil
}

var _ awardpb.AwardServiceHandler = (*AwardService)(nil)

func NewAward(client client.Client, paperAwardLogic *logic.Paper) *AwardService {
	userService := userpb.NewUserService("user", client)
	userGroupService := userpb.NewGroupService("user", client)
	return &AwardService{
		paperAwardLogic:  paperAwardLogic,
		userService:      userService,
		userGroupService: userGroupService,
	}
}
