package service

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
	publicpb "github.com/essayZW/hpcmanager/proto"
	"github.com/essayZW/hpcmanager/user/logic"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/essayZW/hpcmanager/verify"
	"go-micro.dev/v4/client"
)

// UserGroupService 提供关于用户组方面的接口
type UserGroupService struct {
	userGroupLogic *logic.UserGroup
	userLogic      *logic.User
}

// Ping 用户组服务ping测试
func (group *UserGroupService) Ping(ctx context.Context, req *publicpb.Empty, resp *publicpb.PingResponse) error {
	logger.Info("UserGroup PING ", req)
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	return nil
}

// GetGroupInfoByID 查询用户组信息
func (group *UserGroupService) GetGroupInfoByID(ctx context.Context, req *userpb.GetGroupInfoByIDRequest, resp *userpb.GetGroupInfoByIDResponse) error {
	logger.Infof("GetGrouoInfo: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	if !verify.Identify(verify.GetGroupInfo, req.BaseRequest.UserInfo.Levels) {
		logger.Info("GetGroupInfo permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("GetGroupInfo permission forbidden")
	}
	// 只有组管理员或者系统管理员才可以查看组信息
	isAdmin := verify.IsAdmin(req.BaseRequest.UserInfo.Levels)
	isTutor := verify.IsTutor(req.BaseRequest.UserInfo.Levels)
	if isTutor && !isAdmin && req.GroupID != req.BaseRequest.UserInfo.GroupId {
		return errors.New("Tutor can only view group information for his managed group")
	}
	info, err := group.userGroupLogic.GetGroupInfoByID(ctx, int(req.GetGroupID()))
	if err != nil {
		return errors.New("group info query error")
	}
	resp.GroupInfo = &userpb.GroupInfo{
		Id:              int32(info.ID),
		Name:            info.Name,
		QueueName:       info.QueueName,
		NodeGroupName:   info.NodeUserGroupName,
		CreateTime:      info.CreateTime.Unix(),
		CreaterID:       int32(info.CreaterID),
		CreaterUsername: info.CreaterUsername,
		CreaterName:     info.CreaterName,
		TutorID:         int32(info.TutorID),
		TutorUsername:   info.TutorUsername,
		TutorName:       info.TutorName,
		Balance:         info.Balance,
	}
	if info.ExtraAttributes != nil {
		resp.GroupInfo.ExtraAttributes = info.ExtraAttributes.String()
	}
	return nil
}

// PaginationGetGroupInfo 分页查询用户组基本信息
func (group *UserGroupService) PaginationGetGroupInfo(ctx context.Context, req *userpb.PaginationGetGroupInfoRequest, resp *userpb.PaginationGetGroupInfoResponse) error {
	logger.Infof("PaginationGetGroupInfo: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	if !verify.Identify(verify.GetGroupInfo, req.BaseRequest.UserInfo.Levels) {
		logger.Info("PaginationGetGroupInfo permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("PaginationGetGroupInfo permission forbidden")
	}
	// 只有管理员才可以分页查询组信息
	if !verify.IsAdmin(req.BaseRequest.UserInfo.Levels) {
		logger.Info("PaginationGetGroupInfo permission forbidden: not admin BaseRequest: ", req.BaseRequest)
		return errors.New("Only admin can query all group's info")
	}

	infos, err := group.userGroupLogic.PaginationGetGroupInfo(ctx, int(req.PageIndex), int(req.PageSize))
	if err != nil {
		return errors.New("Pagination query group info error")
	}
	resp.Count = int32(infos.Count)
	resp.GroupInfos = make([]*userpb.GroupInfo, len(infos.Infos))
	for index, info := range infos.Infos {
		resp.GroupInfos[index] = &userpb.GroupInfo{
			Id:              int32(info.ID),
			Name:            info.Name,
			QueueName:       info.QueueName,
			NodeGroupName:   info.NodeUserGroupName,
			CreateTime:      info.CreateTime.Unix(),
			CreaterID:       int32(info.CreaterID),
			CreaterUsername: info.CreaterUsername,
			CreaterName:     info.CreaterName,
			TutorID:         int32(info.TutorID),
			TutorUsername:   info.TutorUsername,
			TutorName:       info.TutorName,
			Balance:         info.Balance,
		}
		if info.ExtraAttributes != nil {
			resp.GroupInfos[index].ExtraAttributes = info.ExtraAttributes.String()
		}
	}
	return nil
}

// CreateJoinGroupApply 创建用户加入组申请单
func (group *UserGroupService) CreateJoinGroupApply(ctx context.Context, req *userpb.CreateJoinGroupApplyRequest, resp *userpb.CreateJoinGroupApplyResponse) error {
	logger.Infof("CreateJoinGroupApply: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	if !verify.Identify(verify.ApplyJoinGroup, req.BaseRequest.UserInfo.Levels) {
		logger.Info("CreateJoinGroupApply permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("CreateJoinGroupApply permission forbidden")
	}
	// 查询申请人的信息
	userInfo, err := group.userLogic.GetUserInfoByID(ctx, int(req.BaseRequest.UserInfo.UserId))
	if err != nil {
		return errors.New("apply fail: error userid")
	}
	if userInfo.GroupID != 0 {
		return errors.New("apply fail: user have a group")
	}
	id, err := db.Transication(context.Background(), func(c context.Context, i ...interface{}) (interface{}, error) {
		return group.userGroupLogic.CreateUserJoinGroupApply(c, userInfo, int(req.ApplyGroupID))
	})
	if err != nil {
		return err
	}
	resp.ApplyID = int32(id.(int64))
	logger.Info("Create new user join group apply: ", id)
	resp.Success = true
	// TODO 发送异步消息，表明申请已经创建
	return nil
}

// SearchTutorInfo 通过用户名查询导师信息
func (group *UserGroupService) SearchTutorInfo(ctx context.Context, req *userpb.SearchTutorInfoRequest, resp *userpb.SearchTutorInfoResponse) error {
	logger.Infof("SearchTutorInfo: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	if !verify.Identify(verify.SearchTutorInfo, req.BaseRequest.UserInfo.Levels) {
		logger.Info("SearchTutorInfo permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("SearchTutorInfo permission forbidden")
	}
	info, err := group.userGroupLogic.GetByTutorUsername(ctx, req.Username)
	if err != nil {
		return errors.New("tutor dones not exists")
	}
	resp.GroupID = int32(info.ID)
	resp.TutorID = int32(info.TutorID)
	resp.TutorUsername = info.TutorUsername
	resp.TutorName = info.TutorName
	resp.GroupName = info.Name
	return nil
}

var _ userpb.GroupServiceHandler = (*UserGroupService)(nil)

// NewGroup 创建一个新的group服务
func NewGroup(client client.Client, userGroupLogic *logic.UserGroup, userLogic *logic.User) *UserGroupService {
	return &UserGroupService{
		userGroupLogic: userGroupLogic,
		userLogic:      userLogic,
	}
}
