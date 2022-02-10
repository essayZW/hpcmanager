package service

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/logger"
	publicproto "github.com/essayZW/hpcmanager/proto"
	"github.com/essayZW/hpcmanager/user/logic"
	user "github.com/essayZW/hpcmanager/user/proto"
	"go-micro.dev/v4/client"
)

// UserService 服务
type UserService struct {
	userLogic  *logic.User
	userpLogic *logic.UserPermission
}

// Ping 测试
func (s *UserService) Ping(ctx context.Context, req *publicproto.Empty, resp *publicproto.PingResponse) error {
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	logger.Info("PING ", resp)
	return nil
}

// Login 用户登录
func (s *UserService) Login(ctx context.Context, req *user.LoginRequest, resp *user.LoginResponse) error {
	logger.Infof("User login: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	// 检查登录信息
	success, err := s.userLogic.LoginCheck(req.GetUsername(), req.GetPassword())
	if err != nil {
		logger.Error("login error:", err)
		return errors.New("login error")
	}
	if !success {
		return errors.New("invalid username or password")
	}
	// 查询用户信息
	info, err := s.userLogic.GetByUsername(req.GetUsername())
	if err != nil {
		logger.Error("login error ", err)
		return errors.New("login error")
	}
	resp.UserInfo = &user.UserInfo{
		Id:       int32(info.ID),
		Username: info.Username,
		Name:     info.Name,
	}
	// 创建登录token
	token := s.userLogic.CreateToken(req.GetUsername())
	if token == "" {
		return errors.New("login error")
	}

	resp.Token = token
	return nil
}

// CheckLogin 检查用户登录状态，并返回登录用户的信息以及权限信息
func (s *UserService) CheckLogin(ctx context.Context, req *user.CheckLoginRequest, resp *user.CheckLoginResponse) error {
	logger.Infof("User login check: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	// 通过token查询用户信息
	info, err := s.userLogic.GetUserByToken(req.GetToken())
	if err != nil {
		return errors.New("token invalid")
	}
	resp.Login = true
	resp.UserInfo = &user.UserInfo{
		Id:       int32(info.ID),
		Name:     info.Name,
		Username: info.Username,
		GroupId:  int32(info.GroupID),
	}
	// 查询用户的权限信息
	permissionInfo, err := s.userpLogic.GetUserPermissionByID(info.ID)
	if err != nil {
		logger.Error(err)
		return errors.New("Permission info query error")
	}
	resp.PermissionLevel = make([]int32, len(permissionInfo))
	for index := range permissionInfo {
		resp.PermissionLevel[index] = int32(permissionInfo[index].Level)
	}
	return nil
}

var _ user.UserHandler = (*UserService)(nil)

// NewUser 创建一个新的用户服务实例
func NewUser(client client.Client, userLogic *logic.User, userp *logic.UserPermission) *UserService {
	return &UserService{
		userLogic:  userLogic,
		userpLogic: userp,
	}
}
