package middleware

import (
	"github.com/essayZW/hpcmanager/gateway/request"
	"github.com/gin-gonic/gin"
)

// BaseRequestKey 基础请求体存放在context中的键名
const BaseRequestKey = "__BASE_REQUEST__"

func baseReq(ctx *gin.Context) {
	ctx.Set(BaseRequestKey, request.NewBaseRequest(ctx))
}
