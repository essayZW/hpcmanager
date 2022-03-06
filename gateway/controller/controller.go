package controller

import "github.com/gin-gonic/gin"

// Controller 控制器接口
type Controller interface {
	// Registry 控制器自己注册指定的接口地址到路由上
	Registry(*gin.RouterGroup) *gin.RouterGroup
}
