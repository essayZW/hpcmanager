package utils

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ParsePagination 解析分页参数
func ParsePagination(ctx *gin.Context) (int, int, error) {
	var pageSize int
	var pageIndex int
	var err error
	var param string
	var ok bool
	if param, ok = ctx.GetQuery("pageSize"); !ok {
		return 0, 0, errors.New("缺少pageSize参数")
	}
	pageSize, err = strconv.Atoi(param)
	if err != nil {
		return 0, 0, errors.New("pageSize参数错误")
	}
	if param, ok = ctx.GetQuery("pageIndex"); !ok {
		return 0, 0, errors.New("缺少pageIndex参数")
	}
	pageIndex, err = strconv.Atoi(param)
	if err != nil {
		return 0, 0, errors.New("pageIndex参数错误")
	}
	return pageIndex, pageSize, nil
}
