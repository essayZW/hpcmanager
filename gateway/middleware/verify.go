package middleware

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/gateway/response"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/logger"
)

// TokeiCookieName  token存储在cookie中的键值
const TokeiCookieName = "TOKEN"

type verify struct {
	userService userpb.UserService
	excludeAPI  []string
}

// HandlerFunc 进行初步的权限信息以及用户信息获取
func (v *verify) HandlerFunc(ctx *gin.Context) {
	// TODO 检查当前请求的接口是否需要用户提供token
	if !v.needVerify(ctx.Request.RequestURI) {
		ctx.Next()
		return
	}
	ctx.Abort()
	// 获取用户的token信息
	token, err := ctx.Cookie(TokeiCookieName)
	if err != nil {
		resp := response.New(403, errors.New("forbidden! need token"), false, "forbidden! need token")
		resp.Send(ctx)
		ctx.Abort()
		return
	}
	info, err := v.userService.CheckLogin(context.Background(), &userpb.CheckLoginRequest{
		Token: token,
	})
	if err != nil {
		logger.Error(err)
		resp := response.New(500, err, false, err.Error())
		resp.Send(ctx)
		ctx.Abort()
		return
	}
	if !info.Login {
		resp := response.New(403, errors.New("forbidden! need token"), false, "forbidden! need token")
		resp.Send(ctx)
		ctx.Abort()
		return
	}
	value, _ := ctx.Get(BaseRequestKey)
	breq := value.(*proto.BaseRequest)
	breq.UserInfo = &proto.UserInfo{
		Levels:  info.GetPermissionLevel(),
		UserId:  info.GetUserInfo().GetId(),
		GroupId: info.GetUserInfo().GetGroupId(),
	}
	ctx.Next()
}

func (v *verify) registryExcludeAPIPath(path string) {
	v.excludeAPI = append(v.excludeAPI, path)
}

func (v *verify) needVerify(path string) bool {
	for index := range v.excludeAPI {
		if v.excludeAPI[index] == path {
			return false
		}
	}
	return true
}

func newVerify(client client.Client) *verify {
	return &verify{
		userService: userpb.NewUserService("user", client),
		excludeAPI:  make([]string, 0),
	}
}
