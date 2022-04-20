// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/award.proto

package award

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

// Api Endpoints for AwardService service

func NewAwardServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AwardService service

type AwardService interface {
	Ping(ctx context.Context, in *proto1.Empty, opts ...client.CallOption) (*proto1.PingResponse, error)
	CreatePaperAward(ctx context.Context, in *CreatePaperAwardRequest, opts ...client.CallOption) (*CreatePaperAwardResponse, error)
}

type awardService struct {
	c    client.Client
	name string
}

func NewAwardService(name string, c client.Client) AwardService {
	return &awardService{
		c:    c,
		name: name,
	}
}

func (c *awardService) Ping(ctx context.Context, in *proto1.Empty, opts ...client.CallOption) (*proto1.PingResponse, error) {
	req := c.c.NewRequest(c.name, "AwardService.Ping", in)
	out := new(proto1.PingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *awardService) CreatePaperAward(ctx context.Context, in *CreatePaperAwardRequest, opts ...client.CallOption) (*CreatePaperAwardResponse, error) {
	req := c.c.NewRequest(c.name, "AwardService.CreatePaperAward", in)
	out := new(CreatePaperAwardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AwardService service

type AwardServiceHandler interface {
	Ping(context.Context, *proto1.Empty, *proto1.PingResponse) error
	CreatePaperAward(context.Context, *CreatePaperAwardRequest, *CreatePaperAwardResponse) error
}

func RegisterAwardServiceHandler(s server.Server, hdlr AwardServiceHandler, opts ...server.HandlerOption) error {
	type awardService interface {
		Ping(ctx context.Context, in *proto1.Empty, out *proto1.PingResponse) error
		CreatePaperAward(ctx context.Context, in *CreatePaperAwardRequest, out *CreatePaperAwardResponse) error
	}
	type AwardService struct {
		awardService
	}
	h := &awardServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AwardService{h}, opts...))
}

type awardServiceHandler struct {
	AwardServiceHandler
}

func (h *awardServiceHandler) Ping(ctx context.Context, in *proto1.Empty, out *proto1.PingResponse) error {
	return h.AwardServiceHandler.Ping(ctx, in, out)
}

func (h *awardServiceHandler) CreatePaperAward(ctx context.Context, in *CreatePaperAwardRequest, out *CreatePaperAwardResponse) error {
	return h.AwardServiceHandler.CreatePaperAward(ctx, in, out)
}
