package broker

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"time"

	"github.com/essayZW/hpcmanager/gateway/proto"
	hpcproto "github.com/essayZW/hpcmanager/hpc/proto"
	"github.com/essayZW/hpcmanager/logger"
	groupbroker "github.com/essayZW/hpcmanager/user/broker"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/essayZW/hpcmanager/verify"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/client"
)

// userJoinGroupCustomer 消费用户成功加入组消息
func userJoinGroupCustomer(client client.Client) func(broker.Event) error {
	hpcService := hpcproto.NewHpcService("hpc", client)
	userService := userpb.NewUserService("user", client)
	return func(b broker.Event) error {
		message := b.Message()
		logger.Infof(
			"Received message: [%s] from ID: %s",
			message.Header["Time"],
			message.Header["ID"],
		)
		// 解码消息
		var body groupbroker.UserJoinGroupMessage
		buff := bytes.NewBuffer(message.Body)
		if err := gob.NewDecoder(buff).Decode(&body); err != nil {
			logger.Warn("Message customer: decode error: ", err)
			return err
		}
		// TODO: 调用hpc服务初始化用户的存储信息
		c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()
		baseRequest := &proto.BaseRequest{
			RequestInfo: &proto.RequestInfo{
				Id: message.Header["ID"],
			},
			UserInfo: &proto.UserInfo{},
		}
		// 赋予临时权限
		baseRequest.UserInfo.Levels = append(baseRequest.UserInfo.Levels, int32(verify.SuperAdmin))

		// 查询用户的信息
		userInfo, err := userService.GetUserInfo(c, &userpb.GetUserInfoRequest{
			BaseRequest: baseRequest,
			Userid:      int32(body.UserID),
		})
		if err != nil {
			logger.Warn(err)
			return err
		}
		// 创建存储信息,默认1TB大小,期限为一年
		hpcResp, err := hpcService.SetQuotaByHpcUserID(c, &hpcproto.SetQuotaByHpcUserIDRequest{
			BaseRequest:    baseRequest,
			HpcUserID:      userInfo.UserInfo.HpcUserID,
			NewMaxQuotaTB:  1,
			NewEndTimeUnix: time.Now().Add(time.Duration(8760) * time.Hour).Unix(),
		})
		if err != nil {
			logger.Warn(err)
			return err
		}
		if !hpcResp.Success {
			logger.Warn("update quota error")
			return errors.New("update quota error")
		}
		return errors.New("need implement")
	}
}
