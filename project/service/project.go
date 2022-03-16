package service

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/logger"
	projectdb "github.com/essayZW/hpcmanager/project/db"
	"github.com/essayZW/hpcmanager/project/logic"
	projectpb "github.com/essayZW/hpcmanager/project/proto"
	publicproto "github.com/essayZW/hpcmanager/proto"
	"github.com/essayZW/hpcmanager/verify"
	"go-micro.dev/v4/client"
)

// ProjectService 用户服务
type ProjectService struct {
	projectLogic *logic.Project
}

// Ping ping测试
func (ps *ProjectService) Ping(ctx context.Context, req *publicproto.Empty, resp *publicproto.PingResponse) error {
	logger.Info("Project PING ", req)
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	return nil
}

// CreateProject 创建新的项目记录
func (ps *ProjectService) CreateProject(ctx context.Context, req *projectpb.CreateProjectRequest, resp *projectpb.CreateProjectResponse) error {
	logger.Info("CreateProject: ", req.BaseRequest)
	if !verify.Identify(verify.CreateProject, req.BaseRequest.UserInfo.Levels) {
		logger.Info("CreateProject permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("CreateProject permission forbidden")
	}
	id, err := ps.projectLogic.Create(context.Background(), int(req.ProjectInfo.CreaterUserID), req.ProjectInfo.CreaterName, req.ProjectInfo.CreaterUsername, &projectdb.Project{
		Name:        req.ProjectInfo.Name,
		From:        req.ProjectInfo.From,
		Numbering:   req.ProjectInfo.Numbering,
		Expenses:    req.ProjectInfo.Expenses,
		Description: req.ProjectInfo.Description,
	})
	if err != nil {
		return err
	}
	resp.ProjectID = int32(id)
	return nil
}

var _ projectpb.ProjectHandler = (*ProjectService)(nil)

// NewProject 创建用户服务
func NewProject(client client.Client, projectLogic *logic.Project) *ProjectService {
	return &ProjectService{
		projectLogic: projectLogic,
	}
}
