package middleware

import (
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/client"
)

type verify struct {
	userService userpb.UserService
}

// HandlerFunc 进行初步的权限信息以及用户信息获取
func (v *verify) HandlerFunc(ctx *gin.Context) {
	ctx.Next()
}

func newVerify(client client.Client) *verify {
	return &verify{
		userService: userpb.NewUserService("user", client),
	}
}
