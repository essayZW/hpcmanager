package broker

import (
	"bytes"
	"encoding/gob"

	hpcbroker "github.com/essayZW/hpcmanager/broker"
	"github.com/essayZW/hpcmanager/gateway/proto"
	"go-micro.dev/v4/broker"
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
func (message *CheckApplyMessage) Public(rabbitmqBroker broker.Broker, baseRequest *proto.BaseRequest) error {
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
func checkApplyCustomer(p broker.Event) error {
	message := p.Message()
	logger.Infof("Received message: [%s] from ID: %s", message.Header["Time"], message.Header["ID"])
	// 解码消息
	var body CheckApplyMessage
	buff := bytes.NewBuffer(message.Body)
	if err := gob.NewDecoder(buff).Decode(&body); err != nil {
		logger.Warn("Message customer: decode error: ", err)
	}
	// 只关注管理员审核通过的消息
	if body.TutorCheck {
		return nil
	}
	if !body.CheckStatus {
		return nil
	}
	// 添加用户到组
	return nil
}
