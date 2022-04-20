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
