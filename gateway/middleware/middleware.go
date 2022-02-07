package middleware

import (
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/client"
)

// Registry 注册中间件
func Registry(router *gin.RouterGroup, client client.Client) {
	router.Use(gin.Recovery())
	router.Use(baseReq)

	v := newVerify(client)
	router.Use(v.HandlerFunc)

	router.Use(log)
}
