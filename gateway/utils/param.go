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
	startDateMicroUnixStr, ok := ctx.GetQuery("startDateMicroUnix")
	if ok {
		return zeroTime, zeroTime, errors.New("缺少startDateMicroUnix参数")
	}
	endDateMicroUnixStr, ok := ctx.GetQuery("endStartMicroUnix")
	if ok {
		return zeroTime, zeroTime, errors.New("缺少endDateMicroUnix参数")
	}

	startUnixMicro, err := strconv.Atoi(startDateMicroUnixStr)
	if err != nil {
		return zeroTime, zeroTime, errors.New("invalid startDateMicroUnix参数")
	}
	endUnixMicro, err := strconv.Atoi(endDateMicroUnixStr)
	if err != nil {
		return zeroTime, zeroTime, errors.New("invalid endDateMicroUnix参数")
	}

	startDate := time.UnixMicro(int64(startUnixMicro))
	if startDate.IsZero() {
		return zeroTime, zeroTime, errors.New("invalid startDateMicroUnix参数")
	}
	endDate := time.UnixMicro(int64(endUnixMicro))
	if endDate.IsZero() {
		return zeroTime, zeroTime, errors.New("invalid endDateMicroUnix参数")
	}
	if endDate.Before(startDate) {
		return zeroTime, zeroTime, errors.New("错误的时间范围")
	}
	return startDate, endDate, nil
}
