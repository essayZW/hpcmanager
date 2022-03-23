package broker

import (
	"bytes"
	"encoding/gob"

	hpcbroker "github.com/essayZW/hpcmanager/broker"
	"github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/logger"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/client"
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
	if err := rabbitmqBroker.Publish(hpcbroker.Topic("node.apply.check"), m); err != nil {
		return err
	}
	logger.Info("Public message: ", message, " with request: ", baseRequest)
	return nil
}

func checkApplyCustomer(client client.Client) func(broker.Event) error {
	return func(p broker.Event) error {
		message := p.Message()
		logger.Infof("Received message: [%s] from ID: %s", message.Header["Time"], message.Header["ID"])
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
		// TODO: 创建机器分配处理工单
		return nil
	}
}
