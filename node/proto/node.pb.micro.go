// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/node.proto

package node

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

// Api Endpoints for Node service

func NewNodeEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Node service

type NodeService interface {
	Ping(ctx context.Context, in *proto1.Empty, opts ...client.CallOption) (*proto1.PingResponse, error)
	CreateNodeApply(ctx context.Context, in *CreateNodeApplyRequest, opts ...client.CallOption) (*CreateNodeApplyResponse, error)
}

type nodeService struct {
	c    client.Client
	name string
}

func NewNodeService(name string, c client.Client) NodeService {
	return &nodeService{
		c:    c,
		name: name,
	}
}

func (c *nodeService) Ping(ctx context.Context, in *proto1.Empty, opts ...client.CallOption) (*proto1.PingResponse, error) {
	req := c.c.NewRequest(c.name, "Node.Ping", in)
	out := new(proto1.PingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeService) CreateNodeApply(ctx context.Context, in *CreateNodeApplyRequest, opts ...client.CallOption) (*CreateNodeApplyResponse, error) {
	req := c.c.NewRequest(c.name, "Node.CreateNodeApply", in)
	out := new(CreateNodeApplyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Node service

type NodeHandler interface {
	Ping(context.Context, *proto1.Empty, *proto1.PingResponse) error
	CreateNodeApply(context.Context, *CreateNodeApplyRequest, *CreateNodeApplyResponse) error
}

func RegisterNodeHandler(s server.Server, hdlr NodeHandler, opts ...server.HandlerOption) error {
	type node interface {
		Ping(ctx context.Context, in *proto1.Empty, out *proto1.PingResponse) error
		CreateNodeApply(ctx context.Context, in *CreateNodeApplyRequest, out *CreateNodeApplyResponse) error
	}
	type Node struct {
		node
	}
	h := &nodeHandler{hdlr}
	return s.Handle(s.NewHandler(&Node{h}, opts...))
}

type nodeHandler struct {
	NodeHandler
}

func (h *nodeHandler) Ping(ctx context.Context, in *proto1.Empty, out *proto1.PingResponse) error {
	return h.NodeHandler.Ping(ctx, in, out)
}

func (h *nodeHandler) CreateNodeApply(ctx context.Context, in *CreateNodeApplyRequest, out *CreateNodeApplyResponse) error {
	return h.NodeHandler.CreateNodeApply(ctx, in, out)
}
