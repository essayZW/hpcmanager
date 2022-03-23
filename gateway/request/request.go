package request

import (
	"github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// NewBaseRequest 构建RPC请求基本信息
func NewBaseRequest(ctx *gin.Context) *proto.BaseRequest {
	v4 := uuid.New()
	res := &proto.BaseRequest{
		RequestInfo: &proto.RequestInfo{
			RemoteIP: ctx.ClientIP(),
			Id:       v4.String(),
		},
		UserInfo: &proto.UserInfo{
			Levels: make([]int32, 0),
		},
	}
	return res
}
