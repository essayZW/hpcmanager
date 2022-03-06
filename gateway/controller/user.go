package controller

import (
	"context"

	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/gateway/middleware"
	gatewaypb "github.com/essayZW/hpcmanager/gateway/proto"
	jsonparam "github.com/essayZW/hpcmanager/gateway/request/json"
	"github.com/essayZW/hpcmanager/gateway/response"
	"github.com/essayZW/hpcmanager/logger"
	"github.com/essayZW/hpcmanager/proto"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/client"
)

// User 控制器
type User struct {
	userService userpb.UserService
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
	// 暂时直接返回中间件处理的信息
	rep := response.New(200, baseReq.(*gatewaypb.BaseRequest).UserInfo, true, "success")
	rep.Send(ctx)
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
	return userRouter
}

// NewUser 创建一个用户控制器
func NewUser(client client.Client, configConn config.DynamicConfig) Controller {
	return &User{
		userService: userpb.NewUserService("user", client),
	}
}
