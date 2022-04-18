package broker

import (
	"bytes"
	"encoding/gob"
	"errors"

	"github.com/essayZW/hpcmanager/logger"
	groupbroker "github.com/essayZW/hpcmanager/user/broker"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/client"
)

// userJoinGroupCustomer 消费用户成功加入组消息
func userJoinGroupCustomer(client client.Client) func(broker.Event) error {
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
		return errors.New("need implement")
	}
}
