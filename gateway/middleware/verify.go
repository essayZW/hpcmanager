package middleware

import "github.com/gin-gonic/gin"

// verify 进行初步的权限信息获取
func verify(ctx *gin.Context) {
	ctx.Next()
}
