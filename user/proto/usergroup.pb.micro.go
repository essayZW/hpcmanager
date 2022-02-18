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

// Server API for GroupService service

type GroupServiceHandler interface {
	Ping(context.Context, *proto1.Empty, *proto1.PingResponse) error
	GetGroupInfoByID(context.Context, *GetGroupInfoByIDRequest, *GetGroupInfoByIDResponse) error
}

func RegisterGroupServiceHandler(s server.Server, hdlr GroupServiceHandler, opts ...server.HandlerOption) error {
	type groupService interface {
		Ping(ctx context.Context, in *proto1.Empty, out *proto1.PingResponse) error
		GetGroupInfoByID(ctx context.Context, in *GetGroupInfoByIDRequest, out *GetGroupInfoByIDResponse) error
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
