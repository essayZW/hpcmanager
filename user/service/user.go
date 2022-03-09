package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	hpcDB "github.com/essayZW/hpcmanager/db"
	hpcpb "github.com/essayZW/hpcmanager/hpc/proto"
	"github.com/essayZW/hpcmanager/logger"
	permissionpb "github.com/essayZW/hpcmanager/permission/proto"
	publicproto "github.com/essayZW/hpcmanager/proto"
	"github.com/essayZW/hpcmanager/user/db"
	"github.com/essayZW/hpcmanager/user/logic"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/essayZW/hpcmanager/verify"
	"go-micro.dev/v4/client"
)

// UserService 服务
type UserService struct {
	userLogic         *logic.User
	groupLogic        *logic.UserGroup
	permissionService permissionpb.PermissionService
	hpcService        hpcpb.HpcService
}

// Ping 测试
func (s *UserService) Ping(ctx context.Context, req *publicproto.Empty, resp *publicproto.PingResponse) error {
	logger.Info("User PING ", req)
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	return nil
}

// Login 用户登录
func (s *UserService) Login(ctx context.Context, req *userpb.LoginRequest, resp *userpb.LoginResponse) error {
	logger.Infof("User login: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	// 检查登录信息
	success, err := s.userLogic.LoginCheck(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		logger.Error("login error:", err)
		return errors.New("login error")
	}
	if !success {
		return errors.New("invalid username or password")
	}
	// 查询用户信息
	info, err := s.userLogic.GetByUsername(ctx, req.GetUsername())
	if err != nil {
		logger.Error("login error ", err)
		return errors.New("login error")
	}
	resp.UserInfo = &userpb.UserInfo{
		Id:       int32(info.ID),
		Username: info.Username,
		Name:     info.Name,
		GroupId:  int32(info.GroupID),
	}
	// 创建登录token
	token := s.userLogic.CreateToken(ctx, req.GetUsername())
	if token == "" {
		return errors.New("login error")
	}

	resp.Token = token
	return nil
}

// CheckLogin 检查用户登录状态，并返回登录用户的信息以及权限信息
func (s *UserService) CheckLogin(ctx context.Context, req *userpb.CheckLoginRequest, resp *userpb.CheckLoginResponse) error {
	logger.Infof("CheckLogin: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	// 通过token查询用户信息
	info, err := s.userLogic.GetUserByToken(ctx, req.GetToken())
	if err != nil {
		return errors.New("token invalid")
	}
	resp.Login = true
	resp.UserInfo = &userpb.UserInfo{
		Id:       int32(info.ID),
		Name:     info.Name,
		Username: info.Username,
		GroupId:  int32(info.GroupID),
	}
	// 查询用户的权限信息
	permissionInfo, err := s.permissionService.GetUserPermission(ctx, &permissionpb.GetUserPermissionRequest{
		BaseRequest: req.BaseRequest,
		Id:          int32(info.ID),
	})
	if err != nil {
		logger.Error(err)
		return errors.New("Permission info query error")
	}
	resp.PermissionLevel = make([]int32, len(permissionInfo.Info))
	for index := range permissionInfo.Info {
		resp.PermissionLevel[index] = int32(permissionInfo.Info[index].Level)
	}
	return nil
}

// ExistUsername 检查是否存在某个用户名的用户
func (s *UserService) ExistUsername(ctx context.Context, req *userpb.ExistUsernameRequest, resp *userpb.ExistUsernameResponse) error {
	logger.Infof("ExistUsername: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	thisCtx := context.Background()
	// 直接通过用户名查询用户信息
	info, err := s.userLogic.GetByUsername(thisCtx, req.GetUsername())
	if err != nil {
		return nil
	}
	resp.Exist = true
	resp.UserInfo = &userpb.UserInfo{
		Id:       int32(info.ID),
		GroupId:  int32(info.GroupID),
		Username: info.Username,
		Name:     info.Name,
	}
	return nil
}

// AddUser 添加一个新的用户,返回新用户的用户ID信息
func (s *UserService) AddUser(ctx context.Context, req *userpb.AddUserRequest, resp *userpb.AddUserResponse) error {
	logger.Infof("AddUser: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	if !verify.Identify(verify.AddUserAction, req.BaseRequest.UserInfo.Levels) {
		logger.Info("Adduser permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("Adduser permission forbidden")
	}
	_, err := hpcDB.Transication(context.Background(), func(c context.Context, i ...interface{}) (interface{}, error) {
		extraAttributes, err := hpcDB.NewJSON(req.UserInfo.GetExtraAttributes())
		if err != nil {
			return nil, fmt.Errorf("Parse extraAttributes error: %v", err)
		}
		addedUserInfo := &db.User{
			Username:        req.UserInfo.GetUsername(),
			Password:        req.UserInfo.GetPassword(),
			Tel:             req.UserInfo.GetTel(),
			Email:           req.UserInfo.GetEmail(),
			Name:            req.UserInfo.GetName(),
			PinyinName:      req.UserInfo.GetPyName(),
			CollegeName:     req.UserInfo.GetCollege(),
			GroupID:         int(req.UserInfo.GetGroupId()),
			HpcUserID:       0,
			CreateTime:      time.Now(),
			ExtraAttributes: extraAttributes,
		}
		id, err := s.userLogic.AddUser(c, addedUserInfo)
		if err != nil {
			return nil, fmt.Errorf("Adduser error: %v", err)
		}
		resp.Userid = int32(id)
		// 添加新用户默认权限信息
		// 新建立的用户默认没有组因此只有Guest权限
		addResp, err := s.permissionService.AddUserPermission(ctx, &permissionpb.AddUserPermissionRequest{
			Userid:      int32(id),
			Level:       int32(verify.Guest),
			BaseRequest: req.BaseRequest,
		})
		if err != nil || !addResp.Success {
			return nil, fmt.Errorf("Init user permission info error: %v", err)
		}
		return nil, nil
	})
	return err
}

// CreateToken 创建用户token接口
func (s *UserService) CreateToken(ctx context.Context, req *userpb.CreateTokenRequest, resp *userpb.CreateTokenResponse) error {
	logger.Infof("CreateToken: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	// 根据用户名查询用户信息，确保用户存在
	userInfo, err := s.userLogic.GetByUsername(ctx, req.GetUsername())
	if err != nil {
		return errors.New("user don't exists")
	}
	token := s.userLogic.CreateToken(ctx, req.GetUsername())
	resp.Token = token
	resp.UserInfo = &userpb.UserInfo{
		Id:       int32(userInfo.ID),
		Username: userInfo.Username,
		Name:     userInfo.Name,
		GroupId:  int32(userInfo.GroupID),
	}
	return nil
}

// GetUserInfo 查询用户详细信息
func (s *UserService) GetUserInfo(ctx context.Context, req *userpb.GetUserInfoRequest, resp *userpb.GetUserInfoResponse) error {
	logger.Infof("GetUserInfo: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	if !verify.Identify(verify.GetUserInfo, req.GetBaseRequest().GetUserInfo().GetLevels()) {
		logger.Info("GetUserInfo permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("GetUserInfo permission forbidden")
	}
	isAdmin := verify.IsAdmin(req.BaseRequest.UserInfo.Levels)
	isTutor := verify.IsTutor(req.BaseRequest.UserInfo.Levels)
	if !isAdmin && !isTutor {
		// 普通用户，判断是否是本人
		if req.BaseRequest.UserInfo.UserId != req.Userid {
			return errors.New("you can only query your own user information")
		}
	}
	userInfo, err := s.userLogic.GetUserInfoByID(ctx, int(req.Userid))
	if err != nil {
		return errors.New("User information query error")
	}
	if isTutor && !isAdmin {
		// 是导师且不是管理员，判断其是否属于对应的组
		if req.BaseRequest.UserInfo.GroupId != int32(userInfo.GroupID) {
			return errors.New("Cannot query user information of other groups")
		}
	}
	resp.UserInfo = &userpb.UserInfo{
		Id:         int32(userInfo.ID),
		GroupId:    int32(userInfo.GroupID),
		Username:   userInfo.Username,
		Name:       userInfo.Name,
		Tel:        userInfo.Tel,
		Email:      userInfo.Email,
		PyName:     userInfo.PinyinName,
		College:    userInfo.CollegeName,
		CreateTime: userInfo.CreateTime.Unix(),
		HpcUserID:  int32(userInfo.HpcUserID),
	}
	if userInfo.ExtraAttributes != nil {
		resp.UserInfo.ExtraAttributes = userInfo.ExtraAttributes.String()
	}
	return nil
}

// PaginationGetUserInfo 分页查询用户信息
func (s *UserService) PaginationGetUserInfo(ctx context.Context, req *userpb.PaginationGetUserInfoRequest, resp *userpb.PaginationGetUserInfoResponse) error {
	logger.Infof("PaginationGetUserInfo: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	if !verify.Identify(verify.GetUserInfo, req.GetBaseRequest().GetUserInfo().GetLevels()) {
		logger.Info("PaginationGetUserInfo permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("PaginationGetUserInfo permission forbidden")
	}
	// 只能导师和管理员有查询多个用户的权限
	isAdmin := verify.IsAdmin(req.BaseRequest.UserInfo.Levels)
	isTutor := verify.IsTutor(req.BaseRequest.UserInfo.Levels)
	if !isAdmin && !isTutor {
		logger.Info("PaginationGetUserInfo permission forbidden: must be tutor or admin, baseRequest: ", req.BaseRequest)
		return errors.New("PaginationGetUserInfo permission forbidden: must be tutor or admin")
	}
	var infos *logic.PaginationUserResult
	var err error
	if isAdmin {
		infos, err = s.userLogic.PaginationGetUserInfo(ctx, int(req.PageIndex), int(req.PageSize), 0)
	} else {
		// 导师只可以查看自己组的用户信息
		infos, err = s.userLogic.PaginationGetUserInfo(ctx, int(req.PageIndex), int(req.PageSize), int(req.BaseRequest.UserInfo.GroupId))
	}
	if err != nil {
		return errors.New("User infos query error")
	}
	resp.Count = int32(infos.Count)
	resp.UserInfos = make([]*userpb.UserInfo, len(infos.Infos))
	for index, userInfo := range infos.Infos {
		resp.UserInfos[index] = &userpb.UserInfo{
			Id:         int32(userInfo.ID),
			GroupId:    int32(userInfo.GroupID),
			Username:   userInfo.Username,
			Name:       userInfo.Name,
			Tel:        userInfo.Tel,
			Email:      userInfo.Email,
			PyName:     userInfo.PinyinName,
			College:    userInfo.CollegeName,
			CreateTime: userInfo.CreateTime.Unix(),
			HpcUserID:  int32(userInfo.HpcUserID),
		}
		if userInfo.ExtraAttributes != nil {
			resp.UserInfos[index].ExtraAttributes = userInfo.ExtraAttributes.String()
		}
	}
	return nil
}

// JoinGroup 将一个没有组的用户加入到一个已经存在的组中,并提升其权限为Common
func (s *UserService) JoinGroup(ctx context.Context, req *userpb.JoinGroupRequest, resp *userpb.JoinGroupResponse) error {
	logger.Infof("JoinGroup: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	if !verify.Identify(verify.GetUserInfo, req.GetBaseRequest().GetUserInfo().GetLevels()) {
		logger.Info("JoinGroup permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("JoinGroup permission forbidden")
	}
	_, err := hpcDB.Transication(ctx, func(c context.Context, i ...interface{}) (interface{}, error) {
		// 验证用户不属于任何一个组
		userInfo, err := s.userLogic.GetUserInfoByID(ctx, int(req.UserID))
		if err != nil {
			return nil, errors.New("user not exists")
		}
		if userInfo.GroupID != 0 {
			return nil, errors.New("user already has a group")
		}
		// 查询组信息
		groupInfo, err := s.groupLogic.GetGroupInfoByID(ctx, int(req.GroupID))
		if err != nil {
			return nil, errors.New("group not exists")
		}

		hpcResp, err := s.hpcService.AddUserToGroup(ctx, &hpcpb.AddUserToGroupRequest{
			UserName:   userInfo.Username,
			HpcGroupID: int32(groupInfo.HpcGroupID),
		})
		if err != nil {
			return nil, err
		}
		err = s.userLogic.SetHpcUserID(ctx, userInfo.ID, int(hpcResp.HpcUserID))
		if err != nil {
			return nil, err
		}
		err = s.userLogic.ChangeUserGroup(ctx, userInfo.ID, groupInfo.ID)
		if err != nil {
			return nil, err
		}

		// 删除原来的Guest权限
		s.permissionService.RemoveUserPermission(ctx, &permissionpb.RemoveUserPermissionRequest{
			Userid: int32(userInfo.ID),
			Level:  int32(verify.Guest),
		})
		// 添加Common权限
		addResp, err := s.permissionService.AddUserPermission(ctx, &permissionpb.AddUserPermissionRequest{
			Userid: int32(userInfo.ID),
			Level:  int32(verify.Common),
		})
		if err != nil {
			return nil, err
		}
		if !addResp.Success {
			return nil, errors.New("add user permission error")
		}
		resp.Success = true
		return nil, nil
	})
	return err
}

// Logout 用户退出登录
func (s *UserService) Logout(ctx context.Context, req *userpb.LogoutRequest, resp *userpb.LogoutResponse) error {
	logger.Infof("Logout: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	s.userLogic.DeleteToken(ctx, req.Username)
	return nil
}

var _ userpb.UserHandler = (*UserService)(nil)

// NewUser 创建一个新的用户服务实例
func NewUser(client client.Client, userLogic *logic.User, groupLogic *logic.UserGroup) *UserService {
	permissionService := permissionpb.NewPermissionService("permission", client)
	hpcService := hpcpb.NewHpcService("hpc", client)
	return &UserService{
		userLogic:         userLogic,
		permissionService: permissionService,
		hpcService:        hpcService,
	}
}
