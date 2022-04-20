package broker

import (
	hpcbroker "github.com/essayZW/hpcmanager/broker"
	"github.com/essayZW/hpcmanager/gateway/proto"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/logger"
)

// UserJoinGroupMessage用户成功加入用户组的消息
type UserJoinGroupMessage struct {
	UserID  int
	GroupID int
}

// Public 发布消费消息
func (message *UserJoinGroupMessage) Public(
	rabbitmqBroker broker.Broker,
	baseRequest *proto.BaseRequest,
) error {
	m, err := hpcbroker.NewMessage(message, baseRequest)
	if err != nil {
		return err
	}
	if err := rabbitmqBroker.Publish(hpcbroker.Topic("group.user.join"), m); err != nil {
		return err
	}
	logger.Info("Public message: ", message, " with request: ", baseRequest)
	return nil
}
