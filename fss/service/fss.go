package service

import (
	"context"
	"errors"
	"time"

	"github.com/essayZW/hpcmanager/fss/logic"
	fss "github.com/essayZW/hpcmanager/fss/proto"
	fsspb "github.com/essayZW/hpcmanager/fss/proto"
	"github.com/essayZW/hpcmanager/logger"
	publicproto "github.com/essayZW/hpcmanager/proto"
	"github.com/essayZW/hpcmanager/verify"
	"go-micro.dev/v4/client"
)

type FssService struct {
	fssLogic *logic.Fss
}

// Ping ping测试
func (fss *FssService) Ping(
	ctx context.Context,
	req *publicproto.Empty,
	resp *publicproto.PingResponse,
) error {
	logger.Info("Ping: ", req.BaseRequest)
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	return nil
}

// StoreFile 存储文件到指定位置
func (fss *FssService) StoreFile(ctx context.Context, req *fss.StoreFileRequest, resp *fss.StoreFileResponse) error {
	logger.Info("StoreFile: ", req.BaseRequest)
	if !verify.Identify(verify.StoreFile, req.GetBaseRequest().GetUserInfo().GetLevels()) {
		logger.Info(
			"StoreFile permission forbidden: ",
			req.BaseRequest.RequestInfo.Id,
			", fromUserId: ",
			req.BaseRequest.UserInfo.UserId,
			", withLevels: ",
			req.BaseRequest.UserInfo.Levels,
		)
		return errors.New("StoreFile permission forbidden")
	}
	// 检查文件大小,拒绝存储大于30MB的文件
	if len(req.File) >= 31457280 {
		return errors.New("file size can't larger than 30MB")
	}
	c, cancel := context.WithTimeout(context.Background(), time.Duration(1)*time.Second)
	defer cancel()
	fileName, err := fss.fssLogic.StoreFile(c, req.FileName, req.File)
	if err != nil {
		return err
	}
	resp.FilePath = fileName
	return nil
}

var _ fsspb.FssServiceHandler = (*FssService)(nil)

func NewFss(client client.Client, fssLogic *logic.Fss) *FssService {
	return &FssService{
		fssLogic: fssLogic,
	}
}
