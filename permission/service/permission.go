package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	hpcdb "github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
	"github.com/essayZW/hpcmanager/permission/db"
	"github.com/essayZW/hpcmanager/permission/logic"
	permissionpb "github.com/essayZW/hpcmanager/permission/proto"
	publicpb "github.com/essayZW/hpcmanager/proto"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/essayZW/hpcmanager/verify"
	"go-micro.dev/v4/client"
)

// PermissionService 权限服务
type PermissionService struct {
	userpLogic      *logic.UserPermission
	permissionLogic *logic.Permission

	userService userpb.UserService
}

// Ping ping测试
func (permission *PermissionService) Ping(ctx context.Context, req *publicpb.Empty, resp *publicpb.PingResponse) error {
	logger.Infof("Ping %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	return nil
}

// GetUserPermission 查询用户拥有的权限信息
func (permission *PermissionService) GetUserPermission(ctx context.Context, req *permissionpb.GetUserPermissionRequest, resp *permissionpb.GetUserPermissionResponse) error {
	logger.Infof("GetUserPermission %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)

	isAdmin := verify.IsAdmin(req.BaseRequest.UserInfo.Levels)
	isTutor := verify.IsTutor(req.BaseRequest.UserInfo.Levels)

	if !isAdmin && !isTutor {
		// 是学生用户
		if req.BaseRequest.UserInfo.UserId != req.Id {
			return errors.New("common user only can query self permission info")
		}
	}

	if !isAdmin && isTutor {
		// 是导师用户
		rpcResp, err := permission.userService.GetUserInfo(ctx, &userpb.GetUserInfoRequest{
			BaseRequest: req.BaseRequest,
			Userid:      req.Id,
		})
		if err != nil {
			// 说明目标用户的信息自己没有权限进行查询,或者目标用户信息有误
			return errors.New("user permission info query permission forbidden")
		}

		if rpcResp.UserInfo.GroupId != req.BaseRequest.UserInfo.GroupId {
			return errors.New("user permission info query permission forbidden")
		}
	}
	permissionInfo, err := permission.userpLogic.GetUserPermissionByID(ctx, int(req.GetId()))
	if err != nil {
		return errors.New("no user permission info")
	}
	resp.Info = make([]*permissionpb.PermissionInfo, len(permissionInfo))
	for index, singleInfo := range permissionInfo {
		resp.Info[index] = &permissionpb.PermissionInfo{
			Id:          int32(singleInfo.ID),
			Name:        singleInfo.Name,
			Description: singleInfo.Description,
			Level:       int32(singleInfo.Level),
			CreateTime:  singleInfo.CreateTime.Unix(),
		}
		if singleInfo.ExtraAttributes != nil {
			resp.Info[index].ExtraAttributes = singleInfo.ExtraAttributes.String()
		}
	}
	return nil
}

// AddUserPermission 添加用户权限
func (permission *PermissionService) AddUserPermission(ctx context.Context, req *permissionpb.AddUserPermissionRequest, resp *permissionpb.AddUserPermissionResponse) error {
	logger.Infof("AddUserPermission %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	// 鉴权，只有超级管理员才可以进行此操作
	if !verify.Identify(verify.AddUserPermissionAction, req.BaseRequest.UserInfo.Levels) {
		logger.Info("AdduserPermission permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("AdduserPermission permission forbidden")
	}
	// 由于SuperAdmin权限只能有一个用户拥有,因此需要进行判断
	if req.Level == int32(verify.SuperAdmin) {
		return errors.New("SuperAdmin user has exists")
	}
	err := permission.userpLogic.AddUserPermission(ctx, &db.UserPermission{
		UserID: int(req.GetUserid()),
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
	// 鉴权，只有超级管理员才可以进行此操作
	if !verify.Identify(verify.RemoveUserPermissionAction, req.BaseRequest.UserInfo.Levels) {
		logger.Info("RemoveUserPermission permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("RemoveUserPermission permission forbidden")
	}
	// 由于SuperAdmin权限只能有一个用户拥有,因此需要进行判断
	if req.Level == int32(verify.SuperAdmin) {
		return errors.New("SuperAdmin user permission can't remove")
	}
	_, err := hpcdb.Transaction(context.Background(), func(c context.Context, i ...interface{}) (interface{}, error) {

		err := permission.userpLogic.RemoveUserPermission(c, int(req.GetUserid()), verify.Level(req.GetLevel()))
		if err != nil {
			return nil, errors.New("Remove user permission error")
		}
		permissions, err := permission.userpLogic.GetUserPermissionByID(c, int(req.Userid))
		if err != nil {
			return nil, errors.New("remove user permission error")
		}
		if len(permissions) == 0 {
			// 删除权限之后若没有了任何的权限则添加Guest权限
			if err := permission.userpLogic.AddUserPermission(c, &db.UserPermission{
				UserID: int(req.Userid),
			}, verify.Guest); err != nil {
				return nil, errors.New("remove user permission error")
			}
		}
		resp.Success = true
		return nil, nil
	})
	return err
}

// AddPermission 添加新的权限等级
func (permission *PermissionService) AddPermission(ctx context.Context, req *permissionpb.AddPermissionRequest, resp *permissionpb.AddPermissionResponse) error {
	logger.Infof("AddPermission %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	// 鉴权，只有管理员才可以进行此项操作
	if !verify.Identify(verify.AddPermission, req.BaseRequest.UserInfo.Levels) {
		logger.Info("AddPermission permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("AddPermission permission forbidden")
	}
	extraAttributesJSONStr, err := hpcdb.NewJSON(req.GetInfo().GetExtraAttributes())
	if err != nil {
		return fmt.Errorf("ExtraAttributes json parse error: %v", err)
	}
	id, err := permission.permissionLogic.Add(ctx, &db.Permission{
		Name:            req.GetInfo().GetName(),
		Level:           int8(req.GetInfo().GetLevel()),
		Description:     req.GetInfo().GetDescription(),
		CreateTime:      time.Now(),
		ExtraAttributes: extraAttributesJSONStr,
	})
	if err != nil {
		return errors.New("add permission info error")
	}
	resp.PermissionID = int32(id)
	return nil
}

var _ permissionpb.PermissionHandler = (*PermissionService)(nil)

// NewPermission 创建一个新的Permission服务
func NewPermission(client client.Client, userpermissionLogic *logic.UserPermission, permissionLogic *logic.Permission) *PermissionService {
	userService := userpb.NewUserService("user", client)
	return &PermissionService{
		userpLogic:      userpermissionLogic,
		permissionLogic: permissionLogic,
		userService:     userService,
	}
}
