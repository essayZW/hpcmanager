// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/hpc.proto

package hpc

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

// Api Endpoints for Hpc service

func NewHpcEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Hpc service

type HpcService interface {
	Ping(ctx context.Context, in *proto1.Empty, opts ...client.CallOption) (*proto1.PingResponse, error)
	AddUserWithGroup(ctx context.Context, in *AddUserWithGroupRequest, opts ...client.CallOption) (*AddUserWithGroupResponse, error)
	AddUserToGroup(ctx context.Context, in *AddUserToGroupRequest, opts ...client.CallOption) (*AddUserToGroupResponse, error)
	GetUserInfoByID(ctx context.Context, in *GetUserInfoByIDRequest, opts ...client.CallOption) (*GetUserInfoByIDResponse, error)
	GetGroupInfoByID(ctx context.Context, in *GetGroupInfoByIDRequest, opts ...client.CallOption) (*GetGroupInfoByIDResponse, error)
	GetNodeUsage(ctx context.Context, in *GetNodeUsageRequest, opts ...client.CallOption) (*GetNodeUsageResponse, error)
	GetUserInfoByUsername(ctx context.Context, in *GetUserInfoByUsernameRequest, opts ...client.CallOption) (*GetUserInfoByUsernameResponse, error)
	GetGroupInfoByGroupName(ctx context.Context, in *GetGroupInfoByGroupNameRequest, opts ...client.CallOption) (*GetGroupInfoByGroupNameResponse, error)
	GetQuotaByHpcUserID(ctx context.Context, in *GetQuotaByHpcUserIDRequest, opts ...client.CallOption) (*GetQuotaByHpcUserIDResponse, error)
	SetQuotaByHpcUserID(ctx context.Context, in *SetQuotaByHpcUserIDRequest, opts ...client.CallOption) (*SetQuotaByHpcUserIDResponse, error)
}

type hpcService struct {
	c    client.Client
	name string
}

func NewHpcService(name string, c client.Client) HpcService {
	return &hpcService{
		c:    c,
		name: name,
	}
}

func (c *hpcService) Ping(ctx context.Context, in *proto1.Empty, opts ...client.CallOption) (*proto1.PingResponse, error) {
	req := c.c.NewRequest(c.name, "Hpc.Ping", in)
	out := new(proto1.PingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hpcService) AddUserWithGroup(ctx context.Context, in *AddUserWithGroupRequest, opts ...client.CallOption) (*AddUserWithGroupResponse, error) {
	req := c.c.NewRequest(c.name, "Hpc.AddUserWithGroup", in)
	out := new(AddUserWithGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hpcService) AddUserToGroup(ctx context.Context, in *AddUserToGroupRequest, opts ...client.CallOption) (*AddUserToGroupResponse, error) {
	req := c.c.NewRequest(c.name, "Hpc.AddUserToGroup", in)
	out := new(AddUserToGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hpcService) GetUserInfoByID(ctx context.Context, in *GetUserInfoByIDRequest, opts ...client.CallOption) (*GetUserInfoByIDResponse, error) {
	req := c.c.NewRequest(c.name, "Hpc.GetUserInfoByID", in)
	out := new(GetUserInfoByIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hpcService) GetGroupInfoByID(ctx context.Context, in *GetGroupInfoByIDRequest, opts ...client.CallOption) (*GetGroupInfoByIDResponse, error) {
	req := c.c.NewRequest(c.name, "Hpc.GetGroupInfoByID", in)
	out := new(GetGroupInfoByIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hpcService) GetNodeUsage(ctx context.Context, in *GetNodeUsageRequest, opts ...client.CallOption) (*GetNodeUsageResponse, error) {
	req := c.c.NewRequest(c.name, "Hpc.GetNodeUsage", in)
	out := new(GetNodeUsageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hpcService) GetUserInfoByUsername(ctx context.Context, in *GetUserInfoByUsernameRequest, opts ...client.CallOption) (*GetUserInfoByUsernameResponse, error) {
	req := c.c.NewRequest(c.name, "Hpc.GetUserInfoByUsername", in)
	out := new(GetUserInfoByUsernameResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hpcService) GetGroupInfoByGroupName(ctx context.Context, in *GetGroupInfoByGroupNameRequest, opts ...client.CallOption) (*GetGroupInfoByGroupNameResponse, error) {
	req := c.c.NewRequest(c.name, "Hpc.GetGroupInfoByGroupName", in)
	out := new(GetGroupInfoByGroupNameResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hpcService) GetQuotaByHpcUserID(ctx context.Context, in *GetQuotaByHpcUserIDRequest, opts ...client.CallOption) (*GetQuotaByHpcUserIDResponse, error) {
	req := c.c.NewRequest(c.name, "Hpc.GetQuotaByHpcUserID", in)
	out := new(GetQuotaByHpcUserIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hpcService) SetQuotaByHpcUserID(ctx context.Context, in *SetQuotaByHpcUserIDRequest, opts ...client.CallOption) (*SetQuotaByHpcUserIDResponse, error) {
	req := c.c.NewRequest(c.name, "Hpc.SetQuotaByHpcUserID", in)
	out := new(SetQuotaByHpcUserIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hpc service

type HpcHandler interface {
	Ping(context.Context, *proto1.Empty, *proto1.PingResponse) error
	AddUserWithGroup(context.Context, *AddUserWithGroupRequest, *AddUserWithGroupResponse) error
	AddUserToGroup(context.Context, *AddUserToGroupRequest, *AddUserToGroupResponse) error
	GetUserInfoByID(context.Context, *GetUserInfoByIDRequest, *GetUserInfoByIDResponse) error
	GetGroupInfoByID(context.Context, *GetGroupInfoByIDRequest, *GetGroupInfoByIDResponse) error
	GetNodeUsage(context.Context, *GetNodeUsageRequest, *GetNodeUsageResponse) error
	GetUserInfoByUsername(context.Context, *GetUserInfoByUsernameRequest, *GetUserInfoByUsernameResponse) error
	GetGroupInfoByGroupName(context.Context, *GetGroupInfoByGroupNameRequest, *GetGroupInfoByGroupNameResponse) error
	GetQuotaByHpcUserID(context.Context, *GetQuotaByHpcUserIDRequest, *GetQuotaByHpcUserIDResponse) error
	SetQuotaByHpcUserID(context.Context, *SetQuotaByHpcUserIDRequest, *SetQuotaByHpcUserIDResponse) error
}

func RegisterHpcHandler(s server.Server, hdlr HpcHandler, opts ...server.HandlerOption) error {
	type hpc interface {
		Ping(ctx context.Context, in *proto1.Empty, out *proto1.PingResponse) error
		AddUserWithGroup(ctx context.Context, in *AddUserWithGroupRequest, out *AddUserWithGroupResponse) error
		AddUserToGroup(ctx context.Context, in *AddUserToGroupRequest, out *AddUserToGroupResponse) error
		GetUserInfoByID(ctx context.Context, in *GetUserInfoByIDRequest, out *GetUserInfoByIDResponse) error
		GetGroupInfoByID(ctx context.Context, in *GetGroupInfoByIDRequest, out *GetGroupInfoByIDResponse) error
		GetNodeUsage(ctx context.Context, in *GetNodeUsageRequest, out *GetNodeUsageResponse) error
		GetUserInfoByUsername(ctx context.Context, in *GetUserInfoByUsernameRequest, out *GetUserInfoByUsernameResponse) error
		GetGroupInfoByGroupName(ctx context.Context, in *GetGroupInfoByGroupNameRequest, out *GetGroupInfoByGroupNameResponse) error
		GetQuotaByHpcUserID(ctx context.Context, in *GetQuotaByHpcUserIDRequest, out *GetQuotaByHpcUserIDResponse) error
		SetQuotaByHpcUserID(ctx context.Context, in *SetQuotaByHpcUserIDRequest, out *SetQuotaByHpcUserIDResponse) error
	}
	type Hpc struct {
		hpc
	}
	h := &hpcHandler{hdlr}
	return s.Handle(s.NewHandler(&Hpc{h}, opts...))
}

type hpcHandler struct {
	HpcHandler
}

func (h *hpcHandler) Ping(ctx context.Context, in *proto1.Empty, out *proto1.PingResponse) error {
	return h.HpcHandler.Ping(ctx, in, out)
}

func (h *hpcHandler) AddUserWithGroup(ctx context.Context, in *AddUserWithGroupRequest, out *AddUserWithGroupResponse) error {
	return h.HpcHandler.AddUserWithGroup(ctx, in, out)
}

func (h *hpcHandler) AddUserToGroup(ctx context.Context, in *AddUserToGroupRequest, out *AddUserToGroupResponse) error {
	return h.HpcHandler.AddUserToGroup(ctx, in, out)
}

func (h *hpcHandler) GetUserInfoByID(ctx context.Context, in *GetUserInfoByIDRequest, out *GetUserInfoByIDResponse) error {
	return h.HpcHandler.GetUserInfoByID(ctx, in, out)
}

func (h *hpcHandler) GetGroupInfoByID(ctx context.Context, in *GetGroupInfoByIDRequest, out *GetGroupInfoByIDResponse) error {
	return h.HpcHandler.GetGroupInfoByID(ctx, in, out)
}

func (h *hpcHandler) GetNodeUsage(ctx context.Context, in *GetNodeUsageRequest, out *GetNodeUsageResponse) error {
	return h.HpcHandler.GetNodeUsage(ctx, in, out)
}

func (h *hpcHandler) GetUserInfoByUsername(ctx context.Context, in *GetUserInfoByUsernameRequest, out *GetUserInfoByUsernameResponse) error {
	return h.HpcHandler.GetUserInfoByUsername(ctx, in, out)
}

func (h *hpcHandler) GetGroupInfoByGroupName(ctx context.Context, in *GetGroupInfoByGroupNameRequest, out *GetGroupInfoByGroupNameResponse) error {
	return h.HpcHandler.GetGroupInfoByGroupName(ctx, in, out)
}

func (h *hpcHandler) GetQuotaByHpcUserID(ctx context.Context, in *GetQuotaByHpcUserIDRequest, out *GetQuotaByHpcUserIDResponse) error {
	return h.HpcHandler.GetQuotaByHpcUserID(ctx, in, out)
}

func (h *hpcHandler) SetQuotaByHpcUserID(ctx context.Context, in *SetQuotaByHpcUserIDRequest, out *SetQuotaByHpcUserIDResponse) error {
	return h.HpcHandler.SetQuotaByHpcUserID(ctx, in, out)
}
