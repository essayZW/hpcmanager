package middleware

import "github.com/gin-gonic/gin"

// Verify 进行初步的权限信息获取
func Verify(ctx *gin.Context) {
	ctx.Next()
}
