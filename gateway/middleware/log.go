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
	logger.Infof("%v||%v||%v||%v||%s||%v||%v||%v", baseReq.RequestInfo.Id,
		baseReq.UserInfo.UserId,
		baseReq.UserInfo.UserId,
		baseReq.UserInfo.Levels,
		ctx.Request.Method,
		ctx.Request.RemoteAddr,
		ctx.Request.RequestURI,
		ctx.Request.UserAgent())
	start := time.Now()
	ctx.Next()
	end := time.Now()
	logger.Infof("%s||%v", baseReq.RequestInfo.Id, end.Sub(start))
}
