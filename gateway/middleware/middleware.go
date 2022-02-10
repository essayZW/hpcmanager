package middleware

import (
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/client"
)

var verifyMiddleware *verify

// Registry 注册中间件
func Registry(router *gin.RouterGroup, client client.Client) {
	router.Use(gin.Recovery())
	router.Use(baseReq)

	verifyMiddleware = newVerify(client)
	router.Use(verifyMiddleware.HandlerFunc)

	router.Use(log)
}

// RegistryExcludeAPIPath 向初步鉴权中间件注册不需要验证的API地址
func RegistryExcludeAPIPath(path string) {
	verifyMiddleware.registryExcludeAPIPath(path)
}
