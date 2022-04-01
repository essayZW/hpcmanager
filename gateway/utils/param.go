package utils

import (
	"errors"
	"strconv"
	"time"

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

// ParseDateRange 解析时间范围参数
func ParseDateRange(ctx *gin.Context) (time.Time, time.Time, error) {
	zeroTime := time.Time{}
	startDateMilliUnixStr, ok := ctx.GetQuery("startDateMilliUnix")
	if !ok {
		return zeroTime, zeroTime, errors.New("缺少startDateMilliUnix参数")
	}
	endDateMilliUnixStr, ok := ctx.GetQuery("endDateMilliUnix")
	if !ok {
		return zeroTime, zeroTime, errors.New("缺少endDateMilliUnix参数")
	}

	startUnixMilli, err := strconv.Atoi(startDateMilliUnixStr)
	if err != nil {
		return zeroTime, zeroTime, errors.New("invalid startDateMilliUnix参数")
	}
	endUnixMilli, err := strconv.Atoi(endDateMilliUnixStr)
	if err != nil {
		return zeroTime, zeroTime, errors.New("invalid endDateMilliUnix参数")
	}

	startDate := time.UnixMilli(int64(startUnixMilli))
	if startDate.IsZero() {
		return zeroTime, zeroTime, errors.New("invalid startDateMilliUnix参数")
	}
	endDate := time.UnixMilli(int64(endUnixMilli))
	if endDate.IsZero() {
		return zeroTime, zeroTime, errors.New("invalid endDateMilliUnix参数")
	}
	if endDate.Before(startDate) {
		return zeroTime, zeroTime, errors.New("错误的时间范围")
	}
	return startDate, endDate, nil
}
