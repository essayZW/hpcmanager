package middleware

import (
	"github.com/essayZW/hpcmanager/logger"
	"github.com/gin-gonic/gin"
)

// Log 记录请求日志
func Log(ctx *gin.Context) {
	logger.Infof("%s||%v||%v||%v", ctx.Request.Method, ctx.Request.RemoteAddr, ctx.Request.RequestURI, ctx.Request.UserAgent())
	ctx.Next()
}
