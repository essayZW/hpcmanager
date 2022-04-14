package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/essayZW/hpcmanager/logger"
	projectdb "github.com/essayZW/hpcmanager/project/db"
	"github.com/essayZW/hpcmanager/project/logic"
	projectpb "github.com/essayZW/hpcmanager/project/proto"
	publicproto "github.com/essayZW/hpcmanager/proto"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/essayZW/hpcmanager/verify"
	"go-micro.dev/v4/client"
)

// ProjectService 用户服务
type ProjectService struct {
	projectLogic *logic.Project
	userService  userpb.UserService
}

// Ping ping测试
func (ps *ProjectService) Ping(
	ctx context.Context,
	req *publicproto.Empty,
	resp *publicproto.PingResponse,
) error {
	logger.Info("Project PING ", req)
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	return nil
}

// CreateProject 创建新的项目记录
func (ps *ProjectService) CreateProject(
	ctx context.Context,
	req *projectpb.CreateProjectRequest,
	resp *projectpb.CreateProjectResponse,
) error {
	logger.Info("CreateProject: ", req.BaseRequest)
	if !verify.Identify(verify.CreateProject, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"CreateProject permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("CreateProject permission forbidden")
	}
	id, err := ps.projectLogic.Create(
		context.Background(),
		int(req.BaseRequest.UserInfo.UserId),
		req.BaseRequest.UserInfo.Name,
		req.BaseRequest.UserInfo.Username,
		&projectdb.Project{
			Name:        req.ProjectInfo.Name,
			From:        req.ProjectInfo.From,
			Numbering:   req.ProjectInfo.Numbering,
			Expenses:    req.ProjectInfo.Expenses,
			Description: req.ProjectInfo.Description,
		},
	)
	if err != nil {
		return err
	}
	resp.ProjectID = int32(id)
	return nil
}

// GetProjectInfoByID 通过ID查询项目信息
func (ps *ProjectService) GetProjectInfoByID(
	ctx context.Context,
	req *projectpb.GetProjectInfoByIDRequest,
	resp *projectpb.GetProjectInfoByIDResponse,
) error {
	logger.Info("GetProjectInfoByID: ", req.BaseRequest)
	if !verify.Identify(verify.GetProjectInfo, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"GetProjectInfoByID permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("GetProjectInfoByID permission forbidden")
	}
	info, err := ps.projectLogic.GetByID(context.Background(), int(req.Id))
	if err != nil {
		return err
	}
	// 进行用户查询范围权限鉴定,普通用户只可以查询自己的,导师可以查询自己组的,管理员可以查询所有人的
	isTutor := verify.IsTutor(req.BaseRequest.UserInfo.Levels)
	isAdmin := verify.IsAdmin(req.BaseRequest.UserInfo.Levels)
	if !isAdmin && !isTutor {
		// 普通用户
		if info.CreaterUserID != int(req.BaseRequest.UserInfo.UserId) {
			return errors.New("user only can query self project info")
		}
	}
	if !isAdmin && isTutor {
		// 导师用户
		// 查询项目所属用户的组
		resp, err := ps.userService.GetUserInfo(ctx, &userpb.GetUserInfoRequest{
			Userid:      int32(info.CreaterUserID),
			BaseRequest: req.BaseRequest,
		})
		if err != nil || resp.UserInfo.GroupId != req.BaseRequest.UserInfo.GroupId {
			return errors.New("tutor only can query self group's user's project info")
		}
	}
	resp.Data = &projectpb.ProjectInfo{
		Id:              int32(info.ID),
		Name:            info.Name,
		From:            info.From,
		Numbering:       info.Numbering,
		Expenses:        info.Expenses,
		Description:     info.Description,
		CreaterUserID:   int32(info.CreaterUserID),
		CreaterUsername: info.CreaterUsername,
		CreaterName:     info.CreaterUserName,
		CreateTime:      info.CreateTime.Unix(),
		ModifyUserID:    int32(info.ModifyUserID),
		ModifyUsername:  info.ModifyUsername,
		ModifyName:      info.ModifyUserName,
		ModifyTime:      info.ModifyTime.Unix(),
	}
	if info.ExtraAttributes != nil {
		resp.Data.ExtraAttributes = info.ExtraAttributes.String()
	}
	return nil
}

// PaginationGetProjectInfos 分页查询项目信息
func (ps *ProjectService) PaginationGetProjectInfos(
	ctx context.Context,
	req *projectpb.PaginationGetProjectInfosRequest,
	resp *projectpb.PaginationGetProjectInfosResponse,
) error {
	logger.Info("PaginationGetProjectInfos: ", req.BaseRequest)
	if !verify.Identify(verify.GetProjectInfo, req.BaseRequest.UserInfo.Levels) {
		logger.Info(
			"PaginationGetProjectInfos permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("PaginationGetProjectInfos permission forbidden")
	}
	isAdmin := verify.IsAdmin(req.BaseRequest.UserInfo.Levels)
	isTutor := verify.IsTutor(req.BaseRequest.UserInfo.Levels)
	var err error
	var res *logic.PaginationProjectResult
	// 验证查询范围
	if !isAdmin && !isTutor {
		// 普通用户只可以查询自己创建的项目信息
		res, err = ps.projectLogic.PaginationGetByCreaterUserID(
			context.Background(),
			int(req.PageIndex),
			int(req.PageSize),
			int(req.BaseRequest.UserInfo.UserId),
		)
	} else if !isAdmin && isTutor {
		// 导师用户只可以查询自己组用户的所有的项目信息
		ids, err := ps.userService.ListGroupUser(ctx, &userpb.ListGroupUserRequest{
			BaseRequest: req.BaseRequest,
			GroupID:     req.BaseRequest.UserInfo.GroupId,
		})
		if err != nil {
			return err
		}
		intIds := make([]int, len(ids.Ids))
		for i := range intIds {
			intIds[i] = int(ids.Ids[i])
		}
		res, err = ps.projectLogic.PaginationGetByCreaterUserID(context.Background(), int(req.PageIndex), int(req.PageSize), intIds...)
	} else {
		// 管理员用户可以查看所有的用户创建的项目信息
		res, err = ps.projectLogic.PaginationGet(context.Background(), int(req.PageIndex), int(req.PageSize))
	}
	if err != nil {
		return fmt.Errorf("PaginationGetProjectInfos error: %s", err.Error())
	}
	resp.Count = int32(res.Count)
	resp.Infos = make([]*projectpb.ProjectInfo, 0)
	for _, info := range res.Data {
		tempInfo := &projectpb.ProjectInfo{
			Id:              int32(info.ID),
			Name:            info.Name,
			From:            info.From,
			Numbering:       info.Numbering,
			Expenses:        info.Expenses,
			Description:     info.Description,
			CreaterUserID:   int32(info.CreaterUserID),
			CreaterUsername: info.CreaterUsername,
			CreaterName:     info.CreaterUserName,
			CreateTime:      info.CreateTime.Unix(),
			ModifyUserID:    int32(info.ModifyUserID),
			ModifyUsername:  info.ModifyUsername,
			ModifyName:      info.ModifyUserName,
			ModifyTime:      info.ModifyTime.Unix(),
		}
		if info.ExtraAttributes != nil {
			tempInfo.ExtraAttributes = info.ExtraAttributes.String()
		}
		resp.Infos = append(resp.Infos, tempInfo)
	}
	return nil
}

var _ projectpb.ProjectHandler = (*ProjectService)(nil)

// NewProject 创建用户服务
func NewProject(client client.Client, projectLogic *logic.Project) *ProjectService {
	userService := userpb.NewUserService("user", client)
	return &ProjectService{
		projectLogic: projectLogic,
		userService:  userService,
	}
}
