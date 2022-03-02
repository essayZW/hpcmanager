// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user.proto

package user

import (
	fmt "fmt"
	_ "github.com/essayZW/hpcmanager/gateway/proto"
	proto1 "github.com/essayZW/hpcmanager/proto"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for User service

func NewUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for User service

type UserService interface {
	Ping(ctx context.Context, in *proto1.Empty, opts ...client.CallOption) (*proto1.PingResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	CheckLogin(ctx context.Context, in *CheckLoginRequest, opts ...client.CallOption) (*CheckLoginResponse, error)
	ExistUsername(ctx context.Context, in *ExistUsernameRequest, opts ...client.CallOption) (*ExistUsernameResponse, error)
	AddUser(ctx context.Context, in *AddUserRequest, opts ...client.CallOption) (*AddUserResponse, error)
	CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...client.CallOption) (*CreateTokenResponse, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...client.CallOption) (*GetUserInfoResponse, error)
	PaginationGetUserInfo(ctx context.Context, in *PaginationGetUserInfoRequest, opts ...client.CallOption) (*PaginationGetUserInfoResponse, error)
	JoinGroup(ctx context.Context, in *JoinGroupRequest, opts ...client.CallOption) (*JoinGroupResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Ping(ctx context.Context, in *proto1.Empty, opts ...client.CallOption) (*proto1.PingResponse, error) {
	req := c.c.NewRequest(c.name, "User.Ping", in)
	out := new(proto1.PingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "User.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) CheckLogin(ctx context.Context, in *CheckLoginRequest, opts ...client.CallOption) (*CheckLoginResponse, error) {
	req := c.c.NewRequest(c.name, "User.CheckLogin", in)
	out := new(CheckLoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ExistUsername(ctx context.Context, in *ExistUsernameRequest, opts ...client.CallOption) (*ExistUsernameResponse, error) {
	req := c.c.NewRequest(c.name, "User.ExistUsername", in)
	out := new(ExistUsernameResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) AddUser(ctx context.Context, in *AddUserRequest, opts ...client.CallOption) (*AddUserResponse, error) {
	req := c.c.NewRequest(c.name, "User.AddUser", in)
	out := new(AddUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...client.CallOption) (*CreateTokenResponse, error) {
	req := c.c.NewRequest(c.name, "User.CreateToken", in)
	out := new(CreateTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...client.CallOption) (*GetUserInfoResponse, error) {
	req := c.c.NewRequest(c.name, "User.GetUserInfo", in)
	out := new(GetUserInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) PaginationGetUserInfo(ctx context.Context, in *PaginationGetUserInfoRequest, opts ...client.CallOption) (*PaginationGetUserInfoResponse, error) {
	req := c.c.NewRequest(c.name, "User.PaginationGetUserInfo", in)
	out := new(PaginationGetUserInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) JoinGroup(ctx context.Context, in *JoinGroupRequest, opts ...client.CallOption) (*JoinGroupResponse, error) {
	req := c.c.NewRequest(c.name, "User.JoinGroup", in)
	out := new(JoinGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Ping(context.Context, *proto1.Empty, *proto1.PingResponse) error
	Login(context.Context, *LoginRequest, *LoginResponse) error
	CheckLogin(context.Context, *CheckLoginRequest, *CheckLoginResponse) error
	ExistUsername(context.Context, *ExistUsernameRequest, *ExistUsernameResponse) error
	AddUser(context.Context, *AddUserRequest, *AddUserResponse) error
	CreateToken(context.Context, *CreateTokenRequest, *CreateTokenResponse) error
	GetUserInfo(context.Context, *GetUserInfoRequest, *GetUserInfoResponse) error
	PaginationGetUserInfo(context.Context, *PaginationGetUserInfoRequest, *PaginationGetUserInfoResponse) error
	JoinGroup(context.Context, *JoinGroupRequest, *JoinGroupResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Ping(ctx context.Context, in *proto1.Empty, out *proto1.PingResponse) error
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
		CheckLogin(ctx context.Context, in *CheckLoginRequest, out *CheckLoginResponse) error
		ExistUsername(ctx context.Context, in *ExistUsernameRequest, out *ExistUsernameResponse) error
		AddUser(ctx context.Context, in *AddUserRequest, out *AddUserResponse) error
		CreateToken(ctx context.Context, in *CreateTokenRequest, out *CreateTokenResponse) error
		GetUserInfo(ctx context.Context, in *GetUserInfoRequest, out *GetUserInfoResponse) error
		PaginationGetUserInfo(ctx context.Context, in *PaginationGetUserInfoRequest, out *PaginationGetUserInfoResponse) error
		JoinGroup(ctx context.Context, in *JoinGroupRequest, out *JoinGroupResponse) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Ping(ctx context.Context, in *proto1.Empty, out *proto1.PingResponse) error {
	return h.UserHandler.Ping(ctx, in, out)
}

func (h *userHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.UserHandler.Login(ctx, in, out)
}

func (h *userHandler) CheckLogin(ctx context.Context, in *CheckLoginRequest, out *CheckLoginResponse) error {
	return h.UserHandler.CheckLogin(ctx, in, out)
}

func (h *userHandler) ExistUsername(ctx context.Context, in *ExistUsernameRequest, out *ExistUsernameResponse) error {
	return h.UserHandler.ExistUsername(ctx, in, out)
}

func (h *userHandler) AddUser(ctx context.Context, in *AddUserRequest, out *AddUserResponse) error {
	return h.UserHandler.AddUser(ctx, in, out)
}

func (h *userHandler) CreateToken(ctx context.Context, in *CreateTokenRequest, out *CreateTokenResponse) error {
	return h.UserHandler.CreateToken(ctx, in, out)
}

func (h *userHandler) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, out *GetUserInfoResponse) error {
	return h.UserHandler.GetUserInfo(ctx, in, out)
}

func (h *userHandler) PaginationGetUserInfo(ctx context.Context, in *PaginationGetUserInfoRequest, out *PaginationGetUserInfoResponse) error {
	return h.UserHandler.PaginationGetUserInfo(ctx, in, out)
}

func (h *userHandler) JoinGroup(ctx context.Context, in *JoinGroupRequest, out *JoinGroupResponse) error {
	return h.UserHandler.JoinGroup(ctx, in, out)
}
