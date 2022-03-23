package controller

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/gateway/middleware"
	gatewaypb "github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/gateway/request/json"
	jsonparam "github.com/essayZW/hpcmanager/gateway/request/json"
	"github.com/essayZW/hpcmanager/gateway/response"
	"github.com/essayZW/hpcmanager/gateway/utils"
	hpcpb "github.com/essayZW/hpcmanager/hpc/proto"
	"github.com/essayZW/hpcmanager/logger"
	"github.com/essayZW/hpcmanager/proto"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/client"
)

// User 控制器
type User struct {
	userService      userpb.UserService
	hpcService       hpcpb.HpcService
	userGroupService userpb.GroupService
}

// ping /api/user/ping GET ping!
func (user *User) ping(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	res, err := user.userService.Ping(context.Background(), &proto.Empty{
		BaseRequest: baseReq.(*gatewaypb.BaseRequest),
	})
	var resp *response.Response
	if err != nil {
		resp = response.New(500, err, false, "ping fail!")
	} else {
		resp = response.New(200, res, true, "success")
	}
	resp.Send(ctx)
}

// login /api/user/token POST create a user login token
func (user *User) login(ctx *gin.Context) {
	var params jsonparam.Login
	if err := ctx.ShouldBindJSON(&params); err != nil {
		rep := response.New(500, err.Error(), false, "用户名或者密码格式错误")
		rep.Send(ctx)
		return
	}
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	res, err := user.userService.Login(context.Background(), &userpb.LoginRequest{
		Username:    params.Username,
		Password:    params.Password,
		BaseRequest: baseReq.(*gatewaypb.BaseRequest),
	})
	if err != nil {
		rep := response.New(500, err, false, "用户名或者密码错误")
		rep.Send(ctx)
		return
	}
	rep := response.New(200, res, true, "login success")
	rep.Send(ctx)
}

// loginValid /api/user/token GET query token info of user
func (user *User) loginValid(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)
	// 暂时直接返回中间件处理的信息
	rep := response.New(200, map[string]interface{}{
		"Username": baseRequest.UserInfo.Username,
		"Name":     baseRequest.UserInfo.Name,
		"UserId":   baseRequest.UserInfo.UserId,
		"GroupId":  baseRequest.UserInfo.GroupId,
		"Levels":   baseRequest.UserInfo.Levels,
	}, true, "success")
	rep.Send(ctx)
}

// logout /api/user/token DELETE 删除用户token,退出登录
func (user *User) logout(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	_, err := user.userService.Logout(context.Background(), &userpb.LogoutRequest{
		BaseRequest: baseRequest,
		Username:    baseRequest.UserInfo.Username,
	})
	if err != nil {
		res := response.New(500, nil, false, err.Error())
		res.Send(ctx)
		return
	}
	res := response.New(200, nil, true, "退出登录成功")
	res.Send(ctx)
	return

}

// getByUserID 通过用户ID查询用户信息
func (user *User) getByUserID(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		resp := response.New(200, nil, false, "id参数错误")
		resp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	resp, err := user.userService.GetUserInfo(c, &userpb.GetUserInfoRequest{
		BaseRequest: baseRequest,
		Userid:      int32(id),
	})
	if err != nil {
		resp := response.New(200, nil, false, "用户信息查询失败")
		resp.Send(ctx)
		return
	}
	httpResponse := response.New(200, resp.UserInfo, true, "success")
	httpResponse.Send(ctx)
}

// getByUsername 通过用户帐号查询用户基础信息,主要用于判断用户是否存在以及获取其用户ID
func (user *User) getIDByUsername(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	username := ctx.Param("username")
	if username == "" {
		resp := response.New(200, nil, false, "用户帐号错误")
		resp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	resp, err := user.userService.ExistUsername(c, &userpb.ExistUsernameRequest{
		Username:    username,
		BaseRequest: baseRequest,
	})
	if err != nil {
		logger.Warn(err)
		httpResp := response.New(200, nil, false, "用户信息查询失败")
		httpResp.Send(ctx)
		return
	}
	var httpResp *response.Response
	if resp.Exist {
		httpResp = response.New(200, map[string]interface{}{
			"id":      resp.UserInfo.Id,
			"groupID": resp.UserInfo.GroupId,
			"name":    resp.UserInfo.Name,
		}, true, "success")
	} else {
		httpResp = response.New(200, nil, false, "用户不存在")
	}
	httpResp.Send(ctx)
}

// paginationGetUserInfo /api/group GET 分页查询用户信息
func (user *User) paginationGetUserInfo(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	pageIndex, pageSize, err := utils.ParsePagination(ctx)
	if err != nil {
		httpResp := response.New(200, nil, false, err.Error())
		httpResp.Send(ctx)
		return
	}
	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	resp, err := user.userService.PaginationGetUserInfo(c, &userpb.PaginationGetUserInfoRequest{
		BaseRequest: baseRequest,
		PageIndex:   int32(pageIndex),
		PageSize:    int32(pageSize),
	})
	if err != nil {
		httpResp := response.New(200, nil, false, "查询用户信息失败")
		httpResp.Send(ctx)
		return
	}
	responseData := response.PaginationQueryResponse{
		Count: int(resp.Count),
		Data:  resp.UserInfos,
	}
	if resp.UserInfos == nil {
		responseData.Data = make([]*userpb.UserInfo, 0)
	}
	httpResp := response.New(200, responseData, true, "success")
	httpResp.Send(ctx)
}

// updateUserInfo /api/user PATCH 更新用户信息
func (user *User) updateUserInfo(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	var param json.UpdateUserInfoParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		httpResp := response.New(200, nil, false, err.Error())
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	resp, err := user.userService.UpdateUserInfo(c, &userpb.UpdateUserInfoRequest{
		BaseRequest: baseRequest,
		NewInfos: &userpb.UserInfo{
			Id:      int32(param.ID),
			Tel:     param.Tel,
			Email:   param.Email,
			College: param.College,
		},
	})
	if err != nil || !resp.Success {
		logger.Warn(err)
		httpResp := response.New(200, nil, false, "更改用户信息失败")
		httpResp.Send(ctx)
		return
	}
	httpResp := response.New(200, nil, true, "success")
	httpResp.Send(ctx)
}

// createUser /api/user POST 管理员添加新用户
func (user *User) createUserByAdmin(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	var param json.CreateUserWithGroup
	if err := ctx.ShouldBindJSON(&param); err != nil {
		httpResp := response.New(200, nil, false, err.Error())
		httpResp.Send(ctx)
		return
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	// 先创建对应的用户
	addUserResp, err := user.userService.AddUser(c, &userpb.AddUserRequest{
		UserInfo: &userpb.UserInfo{
			Name:     param.Name,
			Username: param.Username,
			Password: param.Password,
			Tel:      param.Tel,
			Email:    param.Email,
			College:  param.CollegeName,
		},
		BaseRequest: baseRequest,
	})
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("添加用户失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}

	// 如果param.GroupID存在代表需要将其添加到对应的用户组并创建相应的hpc计算用户
	if param.GroupID == 0 {
		httpResp := response.New(200, map[string]interface{}{
			"userID": addUserResp.Userid,
		}, true, "创建Guest用户成功")
		httpResp.Send(ctx)
		return
	}
	// 添加用户组
	joinGroupResp, err := user.userService.JoinGroup(c, &userpb.JoinGroupRequest{
		BaseRequest: baseRequest,
		UserID:      addUserResp.Userid,
		GroupID:     int32(param.GroupID),
	})
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("添加用户到用户组失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}

	if !joinGroupResp.Success {
		httpResp := response.New(200, nil, false, "添加用户到用户组失败")
		httpResp.Send(ctx)
		return
	}
	httpResp := response.New(200, map[string]interface{}{
		"userID": addUserResp.Userid,
	}, true, "操作成功")
	httpResp.Send(ctx)
}

// Registry 为用户控制器注册相应的处理函数
func (user *User) Registry(router *gin.RouterGroup) *gin.RouterGroup {
	logger.Info("registry gateway controller User")
	userRouter := router.Group("/user")
	userRouter.GET("/ping", user.ping)
	middleware.RegistryExcludeAPIPath("GET:/api/user/ping")

	userRouter.POST("/token", user.login)
	middleware.RegistryExcludeAPIPath("POST:/api/user/token")

	userRouter.GET("/token", user.loginValid)

	userRouter.DELETE("/token", user.logout)
	userRouter.GET("/:id", user.getByUserID)
	userRouter.GET("/name/:username", user.getIDByUsername)
	userRouter.GET("", user.paginationGetUserInfo)
	userRouter.PATCH("", user.updateUserInfo)
	userRouter.POST("", user.createUserByAdmin)
	return userRouter
}

// NewUser 创建一个用户控制器
func NewUser(client client.Client, configConn config.DynamicConfig) Controller {
	return &User{
		userService:      userpb.NewUserService("user", client),
		hpcService:       hpcpb.NewHpcService("hpc", client),
		userGroupService: userpb.NewGroupService("user", client),
	}
}
