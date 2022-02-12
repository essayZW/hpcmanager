package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	hpcDB "github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/logger"
	publicproto "github.com/essayZW/hpcmanager/proto"
	"github.com/essayZW/hpcmanager/user/db"
	"github.com/essayZW/hpcmanager/user/logic"
	user "github.com/essayZW/hpcmanager/user/proto"
	"github.com/essayZW/hpcmanager/verify"
	"go-micro.dev/v4/client"
)

// UserService 服务
type UserService struct {
	userLogic  *logic.User
	userpLogic *logic.UserPermission
}

// Ping 测试
func (s *UserService) Ping(ctx context.Context, req *publicproto.Empty, resp *publicproto.PingResponse) error {
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	logger.Info("PING ", resp)
	return nil
}

// Login 用户登录
func (s *UserService) Login(ctx context.Context, req *user.LoginRequest, resp *user.LoginResponse) error {
	logger.Infof("User login: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	// 检查登录信息
	success, err := s.userLogic.LoginCheck(req.GetUsername(), req.GetPassword())
	if err != nil {
		logger.Error("login error:", err)
		return errors.New("login error")
	}
	if !success {
		return errors.New("invalid username or password")
	}
	// 查询用户信息
	info, err := s.userLogic.GetByUsername(req.GetUsername())
	if err != nil {
		logger.Error("login error ", err)
		return errors.New("login error")
	}
	resp.UserInfo = &user.UserInfo{
		Id:       int32(info.ID),
		Username: info.Username,
		Name:     info.Name,
	}
	// 创建登录token
	token := s.userLogic.CreateToken(req.GetUsername())
	if token == "" {
		return errors.New("login error")
	}

	resp.Token = token
	return nil
}

// CheckLogin 检查用户登录状态，并返回登录用户的信息以及权限信息
func (s *UserService) CheckLogin(ctx context.Context, req *user.CheckLoginRequest, resp *user.CheckLoginResponse) error {
	logger.Infof("User login check: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	// 通过token查询用户信息
	info, err := s.userLogic.GetUserByToken(req.GetToken())
	if err != nil {
		return errors.New("token invalid")
	}
	resp.Login = true
	resp.UserInfo = &user.UserInfo{
		Id:       int32(info.ID),
		Name:     info.Name,
		Username: info.Username,
		GroupId:  int32(info.GroupID),
	}
	// 查询用户的权限信息
	permissionInfo, err := s.userpLogic.GetUserPermissionByID(info.ID)
	if err != nil {
		logger.Error(err)
		return errors.New("Permission info query error")
	}
	resp.PermissionLevel = make([]int32, len(permissionInfo))
	for index := range permissionInfo {
		resp.PermissionLevel[index] = int32(permissionInfo[index].Level)
	}
	return nil
}

// ExistUsername 检查是否存在某个用户名的用户
func (s *UserService) ExistUsername(ctx context.Context, req *user.ExistUsernameRequest, resp *user.ExistUsernameResponse) error {
	logger.Infof("ExistUsername: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	// 直接通过用户名查询用户信息
	resp.Exist = true
	if _, err := s.userLogic.GetByUsername(req.GetUsername()); err != nil {
		resp.Exist = false
	}
	return nil
}

// AddUser 添加一个新的用户,返回新用户的用户ID信息
func (s *UserService) AddUser(ctx context.Context, req *user.AddUserRequest, resp *user.AddUserResponse) error {
	logger.Infof("AddUser: %s||%v", req.BaseRequest.RequestInfo.Id, req.BaseRequest.UserInfo.UserId)
	if !verify.Identify(verify.AddUserAction, req.BaseRequest.UserInfo.Levels) {
		logger.Info("Adduser permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("Adduser permission forbidden")
	}
	extraAttributes, err := hpcDB.NewJSON(req.UserInfo.GetExtraAttributes())
	if err != nil {
		return fmt.Errorf("Parse extraAttributes error: %v", err)
	}
	addedUserInfo := &db.User{
		Username:        req.UserInfo.GetUsername(),
		Password:        req.UserInfo.GetPassword(),
		Tel:             req.UserInfo.GetTel(),
		Email:           req.UserInfo.GetEmail(),
		Name:            req.UserInfo.GetName(),
		PinyinName:      req.UserInfo.GetPyName(),
		CollegeName:     req.UserInfo.GetCollege(),
		GroupID:         int(req.UserInfo.GetGroupId()),
		CreateTime:      time.Now(),
		ExtraAttributes: extraAttributes,
	}
	id, err := s.userLogic.AddUser(addedUserInfo)
	if err != nil {
		return fmt.Errorf("Adduser error: %v", err)
	}
	resp.Userid = int32(id)
	// TODO 调用命令调度系统添加计算节点上的用户
	// 添加新用户默认权限信息
	err = s.userpLogic.AddUserPermission(&db.UserPermission{
		UserID:      id,
		UserGroupID: int(req.UserInfo.GetGroupId()),
	}, verify.Common)
	if err != nil {
		return fmt.Errorf("Init user permission info error: %v", err)
	}
	return nil
}

var _ user.UserHandler = (*UserService)(nil)

// NewUser 创建一个新的用户服务实例
func NewUser(client client.Client, userLogic *logic.User, userp *logic.UserPermission) *UserService {
	return &UserService{
		userLogic:  userLogic,
		userpLogic: userp,
	}
}
