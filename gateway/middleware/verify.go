package middleware

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/gateway/proto"
	gatewaypb "github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/gateway/response"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/logger"
)

type verify struct {
	userService userpb.UserService
	excludeAPI  []string
}

// HandlerFunc 进行初步的权限信息以及用户信息获取
func (v *verify) HandlerFunc(ctx *gin.Context) {
	// 检查当前请求的接口是否需要用户提供token
	if !v.needVerify(ctx.Request.URL.Path) {
		logger.Debug("excludeAPI ", ctx.Request.URL.Path)
		ctx.Next()
		return
	}
	// 获取用户的token信息
	token, ok := ctx.GetQuery("access_token")
	if !ok {
		resp := response.New(403, errors.New("forbidden! need token"), false, "forbidden! need token")
		resp.Send(ctx)
		ctx.Abort()
		return
	}
	baseReq, _ := ctx.Get(BaseRequestKey)
	info, err := v.userService.CheckLogin(context.Background(), &userpb.CheckLoginRequest{
		Token:       token,
		BaseRequest: baseReq.(*gatewaypb.BaseRequest),
	})
	if err != nil {
		resp := response.New(403, err, false, "forbidden! invalid token")
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
