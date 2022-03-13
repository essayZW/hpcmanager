package broker

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strconv"
	"time"

	"github.com/asim/go-micro/plugins/broker/rabbitmq/v4"
	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/gateway/proto"
	"go-micro.dev/v4/broker"
)

// NewRabbitmq 创建rabbitmq队列连接
func NewRabbitmq() (broker.Broker, error) {
	// 创建rabbitmq队列连接
	rabbitmqConf, err := config.LoadRabbitmq()
	if err != nil {
		return nil, fmt.Errorf("Rabbitmq conn error: %s", err)
	}
	rabbitmqBroker := rabbitmq.NewBroker(
		broker.Addrs("amqp://"+rabbitmqConf.Address),
		rabbitmq.DurableExchange(),
	)
	if err := rabbitmqBroker.Connect(); err != nil {
		return nil, fmt.Errorf("Rabbitmq conn error: %s", err)
	}
	return rabbitmqBroker, nil
}

// NewMessage 创建新的消息体
func NewMessage(data interface{}, baseRequest *proto.BaseRequest) (*broker.Message, error) {
	// 序列化数据
	var buff bytes.Buffer
	if err := gob.NewEncoder(&buff).Encode(data); err != nil {
		return nil, err
	}
	return &broker.Message{
		Header: NewMessageHeader(baseRequest),
		Body:   buff.Bytes(),
	}, nil
}

// NewMessageHeader 创建新的消息体的头部
func NewMessageHeader(baseRequest *proto.BaseRequest) map[string]string {
	return map[string]string{
		"ID":     baseRequest.RequestInfo.Id,
		"UserID": strconv.Itoa(int(baseRequest.UserInfo.UserId)),
		"Time":   strconv.FormatInt(time.Now().Unix(), 10),
	}
}

// TopicPrefix exchange的topic的前缀
const topicPrefix = "hpcmanager.micro"

// Topic 拼接topic
func Topic(topic string) string {
	return topicPrefix + "." + topic
}
