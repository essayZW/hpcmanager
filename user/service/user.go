package service

import (
	"context"

	"github.com/essayZW/hpcmanager/logger"
	publicproto "github.com/essayZW/hpcmanager/proto"
	"github.com/essayZW/hpcmanager/user/db"
	user "github.com/essayZW/hpcmanager/user/proto"
	"go-micro.dev/v4/client"
)

// UserService 服务
type UserService struct {
	db *db.UserDB
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
func (s *UserService) Login(ctx context.Context, req *user.LoginRequest, resp *user.CreateSessionResponse) error {
	return nil
}

// CreateToken 创建用户的身份认证token
func (s *UserService) CreateToken(ctx context.Context, req *user.CreateSessionRequest, resp *user.CreateSessionResponse) error {
	return nil
}

var _ user.UserHandler = (*UserService)(nil)

// NewUser 创建一个新的用户服务实例
func NewUser(client client.Client, conn *db.UserDB) *UserService {
	return &UserService{
		db: conn,
	}
}
