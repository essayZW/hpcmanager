// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/project.proto

package project

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

// Api Endpoints for Project service

func NewProjectEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Project service

type ProjectService interface {
	Ping(ctx context.Context, in *proto1.Empty, opts ...client.CallOption) (*proto1.PingResponse, error)
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...client.CallOption) (*CreateProjectResponse, error)
	GetProjectInfoByID(ctx context.Context, in *GetProjectInfoByIDRequest, opts ...client.CallOption) (*GetProjectInfoByIDResponse, error)
	PaginationGetProjectInfos(ctx context.Context, in *PaginationGetProjectInfosRequest, opts ...client.CallOption) (*PaginationGetProjectInfosResponse, error)
}

type projectService struct {
	c    client.Client
	name string
}

func NewProjectService(name string, c client.Client) ProjectService {
	return &projectService{
		c:    c,
		name: name,
	}
}

func (c *projectService) Ping(ctx context.Context, in *proto1.Empty, opts ...client.CallOption) (*proto1.PingResponse, error) {
	req := c.c.NewRequest(c.name, "Project.Ping", in)
	out := new(proto1.PingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectService) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...client.CallOption) (*CreateProjectResponse, error) {
	req := c.c.NewRequest(c.name, "Project.CreateProject", in)
	out := new(CreateProjectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectService) GetProjectInfoByID(ctx context.Context, in *GetProjectInfoByIDRequest, opts ...client.CallOption) (*GetProjectInfoByIDResponse, error) {
	req := c.c.NewRequest(c.name, "Project.GetProjectInfoByID", in)
	out := new(GetProjectInfoByIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectService) PaginationGetProjectInfos(ctx context.Context, in *PaginationGetProjectInfosRequest, opts ...client.CallOption) (*PaginationGetProjectInfosResponse, error) {
	req := c.c.NewRequest(c.name, "Project.PaginationGetProjectInfos", in)
	out := new(PaginationGetProjectInfosResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Project service

type ProjectHandler interface {
	Ping(context.Context, *proto1.Empty, *proto1.PingResponse) error
	CreateProject(context.Context, *CreateProjectRequest, *CreateProjectResponse) error
	GetProjectInfoByID(context.Context, *GetProjectInfoByIDRequest, *GetProjectInfoByIDResponse) error
	PaginationGetProjectInfos(context.Context, *PaginationGetProjectInfosRequest, *PaginationGetProjectInfosResponse) error
}

func RegisterProjectHandler(s server.Server, hdlr ProjectHandler, opts ...server.HandlerOption) error {
	type project interface {
		Ping(ctx context.Context, in *proto1.Empty, out *proto1.PingResponse) error
		CreateProject(ctx context.Context, in *CreateProjectRequest, out *CreateProjectResponse) error
		GetProjectInfoByID(ctx context.Context, in *GetProjectInfoByIDRequest, out *GetProjectInfoByIDResponse) error
		PaginationGetProjectInfos(ctx context.Context, in *PaginationGetProjectInfosRequest, out *PaginationGetProjectInfosResponse) error
	}
	type Project struct {
		project
	}
	h := &projectHandler{hdlr}
	return s.Handle(s.NewHandler(&Project{h}, opts...))
}

type projectHandler struct {
	ProjectHandler
}

func (h *projectHandler) Ping(ctx context.Context, in *proto1.Empty, out *proto1.PingResponse) error {
	return h.ProjectHandler.Ping(ctx, in, out)
}

func (h *projectHandler) CreateProject(ctx context.Context, in *CreateProjectRequest, out *CreateProjectResponse) error {
	return h.ProjectHandler.CreateProject(ctx, in, out)
}

func (h *projectHandler) GetProjectInfoByID(ctx context.Context, in *GetProjectInfoByIDRequest, out *GetProjectInfoByIDResponse) error {
	return h.ProjectHandler.GetProjectInfoByID(ctx, in, out)
}

func (h *projectHandler) PaginationGetProjectInfos(ctx context.Context, in *PaginationGetProjectInfosRequest, out *PaginationGetProjectInfosResponse) error {
	return h.ProjectHandler.PaginationGetProjectInfos(ctx, in, out)
}
