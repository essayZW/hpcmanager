package controller

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	fsspb "github.com/essayZW/hpcmanager/fss/proto"
	"github.com/essayZW/hpcmanager/gateway/middleware"
	gatewaypb "github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/gateway/response"
	"github.com/essayZW/hpcmanager/logger"
	"github.com/essayZW/hpcmanager/proto"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/client"
)

type fss struct {
	fssService fsspb.FssService
}

// ping /api/fss/ping GET ping测试
func (f *fss) ping(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)
	res, err := f.fssService.Ping(context.Background(), &proto.Empty{
		BaseRequest: baseRequest,
	})
	var resp *response.Response
	if err != nil {
		resp = response.New(500, err, false, "ping fail!")
	} else {
		resp = response.New(200, res, true, "success")
	}
	resp.Send(ctx)
}

// uploadFile /api/fss/file POST 上传文件
func (f *fss) uploadFile(ctx *gin.Context) {
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)

	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("文件上传失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}
	logger.Debug("upload file: ", fileHeader.Filename)
	uploadedFile, err := fileHeader.Open()
	if err != nil {
		logger.Warn(err)
		httpResp := response.New(200, nil, false, "文件上传失败")
		httpResp.Send(ctx)
		return
	}

	bytesBuffer := make([]byte, 1024)
	fileDataBuffer := bytes.NewBuffer(make([]byte, 0))
	for {
		_, err := uploadedFile.Read(bytesBuffer)
		if err == io.EOF {
			break
		}
		_, err = fileDataBuffer.Write(bytesBuffer)
		if err != nil {
			logger.Warn(err)
			httpResp := response.New(200, nil, false, "文件上传失败")
			httpResp.Send(ctx)
			return
		}
	}

	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	resp, err := f.fssService.StoreFile(c, &fsspb.StoreFileRequest{
		BaseRequest: baseRequest,
		FileName:    fileHeader.Filename,
		File:        fileDataBuffer.Bytes(),
	})
	if err != nil {
		httpResp := response.New(200, nil, false, fmt.Sprintf("文件上传失败: %s", err.Error()))
		httpResp.Send(ctx)
		return
	}

	httpResp := response.New(200, map[string]interface{}{
		"filename": resp.FilePath,
	}, true, "success")
	httpResp.Send(ctx)
}

func (f *fss) Registry(router *gin.RouterGroup) *gin.RouterGroup {
	fssRouter := router.Group("/fss")
	fssRouter.GET("/ping", f.ping)
	middleware.RegistryExcludeAPIPath("GET:/api/fss/ping")

	fssRouter.POST("/file", f.uploadFile)

	return fssRouter
}

func NewFss(client client.Client) Controller {
	fssService := fsspb.NewFssService("fss", client)
	return &fss{
		fssService: fssService,
	}
}
