package broker

import (
	"bytes"
	"context"
	"encoding/gob"
	"time"

	hpcbroker "github.com/essayZW/hpcmanager/broker"
	"github.com/essayZW/hpcmanager/gateway/proto"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/essayZW/hpcmanager/verify"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/logger"
)

// CheckApplyMessage 审核加入组申请消息
type CheckApplyMessage struct {
	CheckStatus  bool
	CheckMessage string
	ApplyID      int
	TutorCheck   bool
}

// Public 发布该消息
func (message *CheckApplyMessage) Public(
	rabbitmqBroker broker.Broker,
	baseRequest *proto.BaseRequest,
) error {
	m, err := hpcbroker.NewMessage(message, baseRequest)
	if err != nil {
		return err
	}
	if err := rabbitmqBroker.Publish(hpcbroker.Topic("group.apply.check"), m); err != nil {
		return err
	}
	logger.Info("Public message: ", message, " with request: ", baseRequest)
	return nil
}

// checkApplyCustomer 审核加入组申请之后的消息消费者
func checkApplyCustomer(client client.Client) func(broker.Event) error {
	userService := userpb.NewUserService("user", client)
	userGroupService := userpb.NewGroupService("user", client)
	return func(p broker.Event) error {
		message := p.Message()
		logger.Infof(
			"Received message: [%s] from ID: %s",
			message.Header["Time"],
			message.Header["ID"],
		)
		// 解码消息
		var body CheckApplyMessage
		buff := bytes.NewBuffer(message.Body)
		if err := gob.NewDecoder(buff).Decode(&body); err != nil {
			logger.Warn("Message customer: decode error: ", err)
			return err
		}
		// 只关注管理员审核通过的消息
		if body.TutorCheck {
			return nil
		}
		if !body.CheckStatus {
			return nil
		}
		c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()
		baseRequest := &proto.BaseRequest{
			RequestInfo: &proto.RequestInfo{
				Id: message.Header["ID"],
			},
			UserInfo: &proto.UserInfo{},
		}
		// 查询用户申请的组的ID
		queryResp, err := userGroupService.GetApplyInfoByID(c, &userpb.GetApplyInfoByIDRequest{
			ApplyID:     int32(body.ApplyID),
			BaseRequest: baseRequest,
		})
		if err != nil {
			logger.Warn(err)
			return err
		}
		logger.Debug(queryResp)
		// 赋予临时权限
		baseRequest.UserInfo.Levels = append(baseRequest.UserInfo.Levels, int32(verify.SuperAdmin))
		// 添加用户到组
		resp, err := userService.JoinGroup(c, &userpb.JoinGroupRequest{
			UserID:      queryResp.Apply.UserID,
			GroupID:     queryResp.Apply.ApplyGroupID,
			BaseRequest: baseRequest,
		})
		if err != nil || !resp.Success {
			logger.Warn(err)
			return err
		}
		return nil
	}
}
