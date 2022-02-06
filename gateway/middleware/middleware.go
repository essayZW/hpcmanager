package middleware

import "github.com/gin-gonic/gin"

// Registry 注册中间件
func Registry(router *gin.RouterGroup) {
	router.Use(log)
	router.Use(verify)
}
