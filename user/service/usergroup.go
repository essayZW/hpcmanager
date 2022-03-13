package service

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/db"
	hpcpb "github.com/essayZW/hpcmanager/hpc/proto"
	"github.com/essayZW/hpcmanager/logger"
	permissionpb "github.com/essayZW/hpcmanager/permission/proto"
	publicpb "github.com/essayZW/hpcmanager/proto"
	userdb "github.com/essayZW/hpcmanager/user/db"
	"github.com/essayZW/hpcmanager/user/logic"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/essayZW/hpcmanager/verify"
	"go-micro.dev/v4/client"
)

// UserGroupService 提供关于用户组方面的接口
type UserGroupService struct {
	userGroupLogic *logic.UserGroup
	userLogic      *logic.User

	permissionService permissionpb.PermissionService
	hpcService        hpcpb.HpcService
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
	logger.Infof("GetGroupInfo: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
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
		CreateTime:      info.CreateTime.Unix(),
		CreaterID:       int32(info.CreaterID),
		CreaterUsername: info.CreaterUsername,
		CreaterName:     info.CreaterName,
		TutorID:         int32(info.TutorID),
		TutorUsername:   info.TutorUsername,
		TutorName:       info.TutorName,
		Balance:         info.Balance,
		HpcGroupID:      int32(info.HpcGroupID),
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
			CreateTime:      info.CreateTime.Unix(),
			CreaterID:       int32(info.CreaterID),
			CreaterUsername: info.CreaterUsername,
			CreaterName:     info.CreaterName,
			TutorID:         int32(info.TutorID),
			TutorUsername:   info.TutorUsername,
			TutorName:       info.TutorName,
			Balance:         info.Balance,
			HpcGroupID:      int32(info.HpcGroupID),
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
	id, err := db.Transaction(context.Background(), func(c context.Context, i ...interface{}) (interface{}, error) {
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
		return errors.New("tutor does not exists")
	}
	resp.GroupID = int32(info.ID)
	resp.TutorID = int32(info.TutorID)
	resp.TutorUsername = info.TutorUsername
	resp.TutorName = info.TutorName
	resp.GroupName = info.Name
	return nil
}

// PageGetApplyGroupInfo 分页查询用户加入组申请记录,对于不同的权限的角色，所查询的范围不同,根据用户的最高权限来进行查询范围判断
func (group *UserGroupService) PageGetApplyGroupInfo(ctx context.Context, req *userpb.PageGetApplyGroupInfoRequest, resp *userpb.PageGetApplyGroupInfoResponse) error {
	logger.Infof("PageGetApplyGroupInfo: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	isAdmin := verify.IsAdmin(req.BaseRequest.UserInfo.Levels)
	isTutor := verify.IsTutor(req.BaseRequest.UserInfo.Levels)

	var result *logic.PaginationApplyResult
	var err error
	if isAdmin {
		// 管理员只可以查看已经经过导师审核的所有的申请信息
		result, err = group.userGroupLogic.AdminPageGetApplyInfo(ctx, int(req.PageIndex), int(req.PageSize))
	} else if isTutor {
		// 导师可以查看自己组的所有申请信息
		result, err = group.userGroupLogic.TutorPageGetApplyInfo(ctx, int(req.PageIndex), int(req.PageSize), int(req.BaseRequest.UserInfo.GroupId))
	} else {
		// 普通用户只可以查看自己发起的申请信息
		result, err = group.userGroupLogic.CommonPageGetApplyInfo(ctx, int(req.PageIndex), int(req.PageSize), int(req.BaseRequest.UserInfo.UserId))
	}
	if err != nil {
		return err
	}
	resp.Applies = make([]*userpb.UserGroupApply, len(result.Applies))
	for index := range result.Applies {
		resp.Applies[index] = &userpb.UserGroupApply{
			Id:                     int32(result.Applies[index].ID),
			UserID:                 int32(result.Applies[index].ApplyGroupID),
			UserUsername:           result.Applies[index].UserUsername,
			UserName:               result.Applies[index].UserName,
			ApplyGroupID:           int32(result.Applies[index].ApplyGroupID),
			TutorID:                int32(result.Applies[index].TutorID),
			TutorUsername:          result.Applies[index].TutorUsername,
			TutorName:              result.Applies[index].TutorName,
			TutorCheckStatus:       int32(result.Applies[index].TutorCheckStatus),
			ManagerCheckStatus:     int32(result.Applies[index].ManagerCheckStatus),
			Status:                 int32(result.Applies[index].Status),
			MessageTutor:           result.Applies[index].MessageTutor.String,
			MessageManager:         result.Applies[index].MessageManager.String,
			TutorCheckTime:         result.Applies[index].TutorCheckTime.Time.Unix(),
			ManagerCheckTime:       result.Applies[index].ManagerCheckTime.Time.Unix(),
			ManagerCheckerID:       int32(result.Applies[index].ManagerCheckerID.Int64),
			ManagerCheckerUsername: result.Applies[index].ManagerCheckerUsername.String,
			ManagerCheckerName:     result.Applies[index].ManagerCheckerName.String,
			CreateTime:             result.Applies[index].CreateTime.Time.Unix(),
		}
		if result.Applies[index].ExtraAttributes != nil {
			resp.Applies[index].ExtraAttributes = result.Applies[index].ExtraAttributes.String()
		}
	}
	resp.Count = int32(result.Count)
	return nil
}

// CheckApply 审核用户加入组申请
func (group *UserGroupService) CheckApply(ctx context.Context, req *userpb.CheckApplyRequest, resp *userpb.CheckApplyResponse) error {
	logger.Infof("CheckApply: %v||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	if !verify.Identify(verify.CheckJoinGroupApply, req.BaseRequest.UserInfo.Levels) {
		logger.Info("CheckJoinGroupApply permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("CheckJoinGroupApply permission forbidden")
	}
	isAdmin := verify.IsAdmin(req.BaseRequest.UserInfo.Levels)
	isTutor := verify.IsTutor(req.BaseRequest.UserInfo.Levels)
	var status bool
	var err error
	// 需要考虑身份同时是导师以及是管理员的情况下审核的冲突情况
	if req.TutorCheck {
		if isTutor {
			status, err = group.userGroupLogic.TutorCheckApply(ctx, int(req.BaseRequest.UserInfo.UserId), int(req.ApplyID), req.CheckStatus, req.CheckMessage)
		} else {
			return errors.New("must be tutor user")
		}
	} else {
		if isAdmin {
			status, err = group.userGroupLogic.AdminCheckApply(ctx, int(req.ApplyID), int(req.BaseRequest.UserInfo.UserId),
				req.BaseRequest.UserInfo.Username, req.BaseRequest.UserInfo.Name, req.CheckStatus, req.CheckMessage)
		} else {
			return errors.New("must be admin")
		}
	}
	if err != nil {
		return err
	}
	resp.Success = status
	return nil
}

// CreateGroup 创建新的用户组并将用户设置为新组的导师
func (group *UserGroupService) CreateGroup(ctx context.Context, req *userpb.CreateGroupRequest, resp *userpb.CreateGroupResponse) error {
	logger.Infof("CreateGroup: %v||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	if !verify.Identify(verify.CreateGroup, req.BaseRequest.UserInfo.Levels) {
		logger.Info("CreateGroup permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("CreateGroup permission forbidden")
	}
	idInterface, err := db.Transaction(ctx, func(c context.Context, i ...interface{}) (interface{}, error) {
		// 查询新组的新导师的信息
		tutorInfo, err := group.userLogic.GetUserInfoByID(c, int(req.TutorID))
		if err != nil {
			return nil, err
		}
		// 新的导师必须是不属于任何组的一个新用户
		if tutorInfo.GroupID != 0 {
			return nil, errors.New("this user has in a group")
		}
		// 调用hpc服务添加对应的计算节点组
		hpcResp, err := group.hpcService.AddUserWithGroup(c, &hpcpb.AddUserWithGroupRequest{
			// 使用导师的英语姓名作为计算节点上的用户组名
			TutorUsername: tutorInfo.PinyinName,
			GroupName:     req.Name,
			QueueName:     req.QueueName,
			BaseRequest:   req.BaseRequest,
		})
		if err != nil {
			return nil, err
		}
		// 创建组信息
		id, err := group.userGroupLogic.CreateGroup(c, &userdb.User{
			ID:       int(req.BaseRequest.UserInfo.UserId),
			Username: req.BaseRequest.UserInfo.Username,
			Name:     req.BaseRequest.UserInfo.Name,
		}, tutorInfo, req.Name, int(hpcResp.HpcGroupID))
		if err != nil {
			return nil, err
		}
		// 修改导师用户关联的hpc_user表ID
		err = group.userLogic.SetHpcUserID(c, tutorInfo.ID, int(hpcResp.HpcUserID))
		if err != nil {
			return nil, err
		}
		// 修改导师所属的组ID
		err = group.userLogic.ChangeUserGroup(c, tutorInfo.ID, int(id))
		if err != nil {
			return nil, err
		}
		// 删除该用户原来的Guest权限
		// NOTE 删除不管是否删除成功,若存在则会删除成功,不存在则忽略删除失败的错误消息
		group.permissionService.RemoveUserPermission(ctx, &permissionpb.RemoveUserPermissionRequest{
			Userid:      int32(tutorInfo.ID),
			Level:       int32(verify.Guest),
			BaseRequest: req.BaseRequest,
		})
		// 添加权限记录
		addresp, err := group.permissionService.AddUserPermission(ctx, &permissionpb.AddUserPermissionRequest{
			Userid:      int32(tutorInfo.ID),
			Level:       int32(verify.Tutor),
			BaseRequest: req.BaseRequest,
		})
		if err != nil {
			return nil, err
		}
		if !addresp.Success {
			return nil, errors.New("user permission create error")
		}
		return id, nil
	})
	if err != nil {
		return err
	}
	id, ok := idInterface.(int64)
	if !ok {
		return errors.New("error result for create group")
	}
	resp.GroupID = int32(id)
	resp.Success = true
	return nil
}

var _ userpb.GroupServiceHandler = (*UserGroupService)(nil)

// NewGroup 创建一个新的group服务
func NewGroup(client client.Client, userGroupLogic *logic.UserGroup, userLogic *logic.User) *UserGroupService {
	return &UserGroupService{
		userGroupLogic:    userGroupLogic,
		userLogic:         userLogic,
		permissionService: permissionpb.NewPermissionService("permission", client),
		hpcService:        hpcpb.NewHpcService("hpc", client),
	}
}
