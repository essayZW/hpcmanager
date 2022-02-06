package controller

import (
	"context"

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

func (user *User) ping(ctx *gin.Context) {
	res, err := user.userService.Ping(context.TODO(), &proto.Empty{})
	var resp *response.Response
	if err != nil {
		resp = response.New(500, err, false, "ping fail!")
	} else {
		resp = response.New(200, res, true, "success")
	}
	resp.Send(ctx)
}

// Registry 为用户控制器注册相应的处理函数
func (user *User) Registry(router *gin.RouterGroup) *gin.RouterGroup {
	logger.Info("registry gateway controller User")
	userRouter := router.Group("/user")
	userRouter.GET("/ping", user.ping)
	return userRouter
}

// NewUser 创建一个用户控制器
func NewUser(client client.Client) *User {
	return &User{
		userService: userpb.NewUserService("user", client),
	}
}
