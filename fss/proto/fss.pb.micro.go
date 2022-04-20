// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/fss.proto

package fss

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

// Api Endpoints for FssService service

func NewFssServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FssService service

type FssService interface {
	Ping(ctx context.Context, in *proto1.Empty, opts ...client.CallOption) (*proto1.PingResponse, error)
	StoreFile(ctx context.Context, in *StoreFileRequest, opts ...client.CallOption) (*StoreFileResponse, error)
}

type fssService struct {
	c    client.Client
	name string
}

func NewFssService(name string, c client.Client) FssService {
	return &fssService{
		c:    c,
		name: name,
	}
}

func (c *fssService) Ping(ctx context.Context, in *proto1.Empty, opts ...client.CallOption) (*proto1.PingResponse, error) {
	req := c.c.NewRequest(c.name, "FssService.Ping", in)
	out := new(proto1.PingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fssService) StoreFile(ctx context.Context, in *StoreFileRequest, opts ...client.CallOption) (*StoreFileResponse, error) {
	req := c.c.NewRequest(c.name, "FssService.StoreFile", in)
	out := new(StoreFileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FssService service

type FssServiceHandler interface {
	Ping(context.Context, *proto1.Empty, *proto1.PingResponse) error
	StoreFile(context.Context, *StoreFileRequest, *StoreFileResponse) error
}

func RegisterFssServiceHandler(s server.Server, hdlr FssServiceHandler, opts ...server.HandlerOption) error {
	type fssService interface {
		Ping(ctx context.Context, in *proto1.Empty, out *proto1.PingResponse) error
		StoreFile(ctx context.Context, in *StoreFileRequest, out *StoreFileResponse) error
	}
	type FssService struct {
		fssService
	}
	h := &fssServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FssService{h}, opts...))
}

type fssServiceHandler struct {
	FssServiceHandler
}

func (h *fssServiceHandler) Ping(ctx context.Context, in *proto1.Empty, out *proto1.PingResponse) error {
	return h.FssServiceHandler.Ping(ctx, in, out)
}

func (h *fssServiceHandler) StoreFile(ctx context.Context, in *StoreFileRequest, out *StoreFileResponse) error {
	return h.FssServiceHandler.StoreFile(ctx, in, out)
}