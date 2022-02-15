package service

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/logger"
	"github.com/essayZW/hpcmanager/permission/db"
	"github.com/essayZW/hpcmanager/permission/logic"
	permissionpb "github.com/essayZW/hpcmanager/permission/proto"
	publicpb "github.com/essayZW/hpcmanager/proto"
	"github.com/essayZW/hpcmanager/verify"
	"go-micro.dev/v4/client"
)

// PermissionService 权限服务
type PermissionService struct {
	userpLogic *logic.UserPermission
}

// Ping ping测试
func (permission *PermissionService) Ping(ctx context.Context, req *publicpb.Empty, resp *publicpb.PingResponse) error {
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	logger.Info("PING ", resp)
	return nil
}

// GetUserPermission 查询用户拥有的权限信息
func (permission *PermissionService) GetUserPermission(ctx context.Context, req *permissionpb.GetUserPermissionRequest, resp *permissionpb.GetUserPermissionResponse) error {
	logger.Infof("GetUserPermission %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	permissionInfo, err := permission.userpLogic.GetUserPermissionByID(ctx, int(req.GetId()))
	if err != nil {
		return errors.New("no user permission info")
	}
	resp.Info = make([]*permissionpb.PermissionInfo, len(permissionInfo))
	for index, singleInfo := range permissionInfo {
		resp.Info[index] = &permissionpb.PermissionInfo{
			Id:              int32(singleInfo.ID),
			Name:            singleInfo.Name,
			Description:     singleInfo.Description,
			Level:           int32(singleInfo.Level),
			CreateTime:      singleInfo.CreateTime.Unix(),
			ExtraAttributes: singleInfo.ExtraAttributes.String(),
		}
	}
	return nil
}

// AddUserPermission 添加用户权限
func (permission *PermissionService) AddUserPermission(ctx context.Context, req *permissionpb.AddUserPermissionRequest, resp *permissionpb.AddUserPermissionResponse) error {
	logger.Infof("AddUserPermission %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	// 鉴权，只有管理员才可以进行此操作
	if !verify.Identify(verify.AddUserPermissionAction, req.BaseRequest.UserInfo.Levels) {
		logger.Info("AdduserPermission permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("AdduserPermission permission forbidden")
	}
	err := permission.userpLogic.AddUserPermission(ctx, &db.UserPermission{
		UserID:      int(req.GetUserid()),
		UserGroupID: int(req.GetUserGroupID()),
	}, verify.Level(req.Level))
	if err != nil {
		resp.Success = false
		logger.Warn("AddUserPermission logic error: ", err)
		return errors.New("add userpermission info error")
	}
	resp.Success = true
	return nil
}

// RemoveUserPermission 删除用户的某个权限
func (permission *PermissionService) RemoveUserPermission(ctx context.Context, req *permissionpb.RemoveUserPermissionRequest, resp *permissionpb.RemoveUserPermissionResponse) error {
	logger.Infof("RemoveUserPermission %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	// 鉴权，只有管理员才可以进行此项操作
	if !verify.Identify(verify.RemoveUserPermissionAction, req.BaseRequest.UserInfo.Levels) {
		logger.Info("RemoveUserPermission permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("RemoveUserPermission permission forbidden")
	}
	err := permission.userpLogic.RemoveUserPermission(ctx, int(req.GetUserid()), verify.Level(req.GetLevel()))
	if err != nil {
		return errors.New("Remove user permission error")
	}
	resp.Success = true
	return nil
}

var _ permissionpb.PermissionHandler = (*PermissionService)(nil)

// NewPermission 创建一个新的Permission服务
func NewPermission(client client.Client, userpermissionLogic *logic.UserPermission) *PermissionService {
	return &PermissionService{
		userpLogic: userpermissionLogic,
	}
}
