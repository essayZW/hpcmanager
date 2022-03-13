package broker

import (
	"github.com/asim/go-micro/plugins/broker/rabbitmq/v4"
	hpcbroker "github.com/essayZW/hpcmanager/broker"
	"go-micro.dev/v4/broker"
)

// RegistryCustomer 注册broker消费者
func RegistryCustomer(rabbitmqBroker broker.Broker) {
	rabbitmqBroker.Subscribe(hpcbroker.Topic("group.apply.check"),
		checkApplyCustomer,
		broker.Queue(hpcbroker.Topic("group.apply.check")),
		rabbitmq.AckOnSuccess())
}
