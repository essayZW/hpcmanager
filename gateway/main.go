package main

import (
	"flag"
	"strconv"

	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/gateway/middleware"
	"github.com/essayZW/hpcmanager/logger"
	"github.com/gin-gonic/gin"
)

func init() {
	logger.SetName("gateway")
}

func main() {
	var port int
	var debug bool
	flag.IntVar(&port, "port", 80, "port to listen")
	flag.BoolVar(&debug, "debug", true, "debug mode")
	hpcmanager.LoadCommonArgs()
	flag.Parse()
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.Default()

	// 注册V1接口的相关方法
	v1 := server.Group("/v1")
	// 添加日志中间件
	v1.Use(middleware.Log)
	// 添加鉴权中间件
	v1.Use(middleware.Verify)
	server.Run(":" + strconv.Itoa(port))
}
