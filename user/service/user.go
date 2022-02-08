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
	userLogic *logic.User
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
	info, err := s.userLogic.QueryByUsername(req.GetUsername())
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

var _ user.UserHandler = (*UserService)(nil)

// NewUser 创建一个新的用户服务实例
func NewUser(client client.Client, userLogic *logic.User) *UserService {
	return &UserService{
		userLogic: userLogic,
	}
}
