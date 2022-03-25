// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/usergroup.proto

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

// Api Endpoints for GroupService service

func NewGroupServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GroupService service

type GroupService interface {
	Ping(ctx context.Context, in *proto1.Empty, opts ...client.CallOption) (*proto1.PingResponse, error)
	GetGroupInfoByID(ctx context.Context, in *GetGroupInfoByIDRequest, opts ...client.CallOption) (*GetGroupInfoByIDResponse, error)
	PaginationGetGroupInfo(ctx context.Context, in *PaginationGetGroupInfoRequest, opts ...client.CallOption) (*PaginationGetGroupInfoResponse, error)
	CreateJoinGroupApply(ctx context.Context, in *CreateJoinGroupApplyRequest, opts ...client.CallOption) (*CreateJoinGroupApplyResponse, error)
	SearchTutorInfo(ctx context.Context, in *SearchTutorInfoRequest, opts ...client.CallOption) (*SearchTutorInfoResponse, error)
	PageGetApplyGroupInfo(ctx context.Context, in *PageGetApplyGroupInfoRequest, opts ...client.CallOption) (*PageGetApplyGroupInfoResponse, error)
	CheckApply(ctx context.Context, in *CheckApplyRequest, opts ...client.CallOption) (*CheckApplyResponse, error)
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...client.CallOption) (*CreateGroupResponse, error)
	GetApplyInfoByID(ctx context.Context, in *GetApplyInfoByIDRequest, opts ...client.CallOption) (*GetApplyInfoByIDResponse, error)
	RevokeUserApplyGroup(ctx context.Context, in *RevokeUserApplyGroupRequest, opts ...client.CallOption) (*RevokeUserApplyGroupResponse, error)
}

type groupService struct {
	c    client.Client
	name string
}

func NewGroupService(name string, c client.Client) GroupService {
	return &groupService{
		c:    c,
		name: name,
	}
}

func (c *groupService) Ping(ctx context.Context, in *proto1.Empty, opts ...client.CallOption) (*proto1.PingResponse, error) {
	req := c.c.NewRequest(c.name, "GroupService.Ping", in)
	out := new(proto1.PingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) GetGroupInfoByID(ctx context.Context, in *GetGroupInfoByIDRequest, opts ...client.CallOption) (*GetGroupInfoByIDResponse, error) {
	req := c.c.NewRequest(c.name, "GroupService.GetGroupInfoByID", in)
	out := new(GetGroupInfoByIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) PaginationGetGroupInfo(ctx context.Context, in *PaginationGetGroupInfoRequest, opts ...client.CallOption) (*PaginationGetGroupInfoResponse, error) {
	req := c.c.NewRequest(c.name, "GroupService.PaginationGetGroupInfo", in)
	out := new(PaginationGetGroupInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) CreateJoinGroupApply(ctx context.Context, in *CreateJoinGroupApplyRequest, opts ...client.CallOption) (*CreateJoinGroupApplyResponse, error) {
	req := c.c.NewRequest(c.name, "GroupService.CreateJoinGroupApply", in)
	out := new(CreateJoinGroupApplyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) SearchTutorInfo(ctx context.Context, in *SearchTutorInfoRequest, opts ...client.CallOption) (*SearchTutorInfoResponse, error) {
	req := c.c.NewRequest(c.name, "GroupService.SearchTutorInfo", in)
	out := new(SearchTutorInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) PageGetApplyGroupInfo(ctx context.Context, in *PageGetApplyGroupInfoRequest, opts ...client.CallOption) (*PageGetApplyGroupInfoResponse, error) {
	req := c.c.NewRequest(c.name, "GroupService.PageGetApplyGroupInfo", in)
	out := new(PageGetApplyGroupInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) CheckApply(ctx context.Context, in *CheckApplyRequest, opts ...client.CallOption) (*CheckApplyResponse, error) {
	req := c.c.NewRequest(c.name, "GroupService.CheckApply", in)
	out := new(CheckApplyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...client.CallOption) (*CreateGroupResponse, error) {
	req := c.c.NewRequest(c.name, "GroupService.CreateGroup", in)
	out := new(CreateGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) GetApplyInfoByID(ctx context.Context, in *GetApplyInfoByIDRequest, opts ...client.CallOption) (*GetApplyInfoByIDResponse, error) {
	req := c.c.NewRequest(c.name, "GroupService.GetApplyInfoByID", in)
	out := new(GetApplyInfoByIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupService) RevokeUserApplyGroup(ctx context.Context, in *RevokeUserApplyGroupRequest, opts ...client.CallOption) (*RevokeUserApplyGroupResponse, error) {
	req := c.c.NewRequest(c.name, "GroupService.RevokeUserApplyGroup", in)
	out := new(RevokeUserApplyGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GroupService service

type GroupServiceHandler interface {
	Ping(context.Context, *proto1.Empty, *proto1.PingResponse) error
	GetGroupInfoByID(context.Context, *GetGroupInfoByIDRequest, *GetGroupInfoByIDResponse) error
	PaginationGetGroupInfo(context.Context, *PaginationGetGroupInfoRequest, *PaginationGetGroupInfoResponse) error
	CreateJoinGroupApply(context.Context, *CreateJoinGroupApplyRequest, *CreateJoinGroupApplyResponse) error
	SearchTutorInfo(context.Context, *SearchTutorInfoRequest, *SearchTutorInfoResponse) error
	PageGetApplyGroupInfo(context.Context, *PageGetApplyGroupInfoRequest, *PageGetApplyGroupInfoResponse) error
	CheckApply(context.Context, *CheckApplyRequest, *CheckApplyResponse) error
	CreateGroup(context.Context, *CreateGroupRequest, *CreateGroupResponse) error
	GetApplyInfoByID(context.Context, *GetApplyInfoByIDRequest, *GetApplyInfoByIDResponse) error
	RevokeUserApplyGroup(context.Context, *RevokeUserApplyGroupRequest, *RevokeUserApplyGroupResponse) error
}

func RegisterGroupServiceHandler(s server.Server, hdlr GroupServiceHandler, opts ...server.HandlerOption) error {
	type groupService interface {
		Ping(ctx context.Context, in *proto1.Empty, out *proto1.PingResponse) error
		GetGroupInfoByID(ctx context.Context, in *GetGroupInfoByIDRequest, out *GetGroupInfoByIDResponse) error
		PaginationGetGroupInfo(ctx context.Context, in *PaginationGetGroupInfoRequest, out *PaginationGetGroupInfoResponse) error
		CreateJoinGroupApply(ctx context.Context, in *CreateJoinGroupApplyRequest, out *CreateJoinGroupApplyResponse) error
		SearchTutorInfo(ctx context.Context, in *SearchTutorInfoRequest, out *SearchTutorInfoResponse) error
		PageGetApplyGroupInfo(ctx context.Context, in *PageGetApplyGroupInfoRequest, out *PageGetApplyGroupInfoResponse) error
		CheckApply(ctx context.Context, in *CheckApplyRequest, out *CheckApplyResponse) error
		CreateGroup(ctx context.Context, in *CreateGroupRequest, out *CreateGroupResponse) error
		GetApplyInfoByID(ctx context.Context, in *GetApplyInfoByIDRequest, out *GetApplyInfoByIDResponse) error
		RevokeUserApplyGroup(ctx context.Context, in *RevokeUserApplyGroupRequest, out *RevokeUserApplyGroupResponse) error
	}
	type GroupService struct {
		groupService
	}
	h := &groupServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&GroupService{h}, opts...))
}

type groupServiceHandler struct {
	GroupServiceHandler
}

func (h *groupServiceHandler) Ping(ctx context.Context, in *proto1.Empty, out *proto1.PingResponse) error {
	return h.GroupServiceHandler.Ping(ctx, in, out)
}

func (h *groupServiceHandler) GetGroupInfoByID(ctx context.Context, in *GetGroupInfoByIDRequest, out *GetGroupInfoByIDResponse) error {
	return h.GroupServiceHandler.GetGroupInfoByID(ctx, in, out)
}

func (h *groupServiceHandler) PaginationGetGroupInfo(ctx context.Context, in *PaginationGetGroupInfoRequest, out *PaginationGetGroupInfoResponse) error {
	return h.GroupServiceHandler.PaginationGetGroupInfo(ctx, in, out)
}

func (h *groupServiceHandler) CreateJoinGroupApply(ctx context.Context, in *CreateJoinGroupApplyRequest, out *CreateJoinGroupApplyResponse) error {
	return h.GroupServiceHandler.CreateJoinGroupApply(ctx, in, out)
}

func (h *groupServiceHandler) SearchTutorInfo(ctx context.Context, in *SearchTutorInfoRequest, out *SearchTutorInfoResponse) error {
	return h.GroupServiceHandler.SearchTutorInfo(ctx, in, out)
}

func (h *groupServiceHandler) PageGetApplyGroupInfo(ctx context.Context, in *PageGetApplyGroupInfoRequest, out *PageGetApplyGroupInfoResponse) error {
	return h.GroupServiceHandler.PageGetApplyGroupInfo(ctx, in, out)
}

func (h *groupServiceHandler) CheckApply(ctx context.Context, in *CheckApplyRequest, out *CheckApplyResponse) error {
	return h.GroupServiceHandler.CheckApply(ctx, in, out)
}

func (h *groupServiceHandler) CreateGroup(ctx context.Context, in *CreateGroupRequest, out *CreateGroupResponse) error {
	return h.GroupServiceHandler.CreateGroup(ctx, in, out)
}

func (h *groupServiceHandler) GetApplyInfoByID(ctx context.Context, in *GetApplyInfoByIDRequest, out *GetApplyInfoByIDResponse) error {
	return h.GroupServiceHandler.GetApplyInfoByID(ctx, in, out)
}

func (h *groupServiceHandler) RevokeUserApplyGroup(ctx context.Context, in *RevokeUserApplyGroupRequest, out *RevokeUserApplyGroupResponse) error {
	return h.GroupServiceHandler.RevokeUserApplyGroup(ctx, in, out)
}
