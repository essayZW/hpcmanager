package middleware

import (
	"github.com/essayZW/hpcmanager/gateway/response"
	"github.com/essayZW/hpcmanager/gateway/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type sys struct {
	redis *redis.Client
}

func (i *sys) checkInstall(ctx *gin.Context) {
	status := utils.IsInstall(i.redis)
	if ctx.Request.URL.Path != "/api/sys/install" && !status {
		resp := response.New(403, nil, false, "系统还未初始化,需要进行系统初始化")
		resp.Send(ctx)
		ctx.Abort()
		return
	}
	ctx.Next()
}
