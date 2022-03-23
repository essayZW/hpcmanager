package broker

import (
	"github.com/asim/go-micro/plugins/broker/rabbitmq/v4"
	hpcbroker "github.com/essayZW/hpcmanager/broker"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/client"
)

// RegistryCustomer 注册消息队列消费者
func RegistryCustomer(rabbitmqBroker broker.Broker, client client.Client) {
	rabbitmqBroker.Subscribe(hpcbroker.Topic("node.apply.check"),
		checkApplyCustomer(client),
		broker.Queue(hpcbroker.Topic("node.apply.check")),
		rabbitmq.AckOnSuccess())
}
