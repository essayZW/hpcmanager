package middleware

import (
	"time"

	"github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/logger"
	"github.com/gin-gonic/gin"
)

// log 记录请求日志
func log(ctx *gin.Context) {
	value, _ := ctx.Get(BaseRequestKey)
	baseReq := value.(*proto.BaseRequest)
	logger.Infof("%v||%s||%v||%v||%v", baseReq.RequestInfo.Id,
		ctx.Request.Method,
		ctx.ClientIP(),
		ctx.Request.URL.Path,
		ctx.Request.UserAgent())
	start := time.Now()
	ctx.Next()
	end := time.Now()
	logger.Infof("%s||%v", baseReq.RequestInfo.Id, end.Sub(start))
}
