package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go-micro.dev/v4/client"
)

var verifyMiddleware *verify

// Registry 注册中间件
func Registry(router *gin.RouterGroup, client client.Client, redisConn *redis.Client) {
	router.Use(gin.Recovery())

	install := &install{
		redis: redisConn,
	}
	// 注册检查安装状态的中间件
	router.Use(install.check)
	// 注册添加基础请求信息的中间件
	router.Use(baseReq)
	// 注册进行日志输出的中间件
	router.Use(log)
	// 注册初步鉴权的中间件
	verifyMiddleware = newVerify(client)
	router.Use(verifyMiddleware.HandlerFunc)
}

// RegistryExcludeAPIPath 向初步鉴权中间件注册不需要验证的API地址
func RegistryExcludeAPIPath(path string) {
	verifyMiddleware.registryExcludeAPIPath(path)
}
