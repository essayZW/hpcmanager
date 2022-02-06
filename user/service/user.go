package service

import (
	"context"

	publicproto "github.com/essayZW/hpcmanager/proto"
	user "github.com/essayZW/hpcmanager/user/proto"
)

// UserService 服务
type UserService struct {
}

// Ping 测试
func (s *UserService) Ping(ctx context.Context, _ *publicproto.Empty, req *publicproto.PingResponse) error {
	req.Msg = "PONG"
	return nil
}

var _ user.UserHandler = (*UserService)(nil)

// NewUser 创建一个新的用户服务实例
func NewUser() *UserService {
	return &UserService{}
}
