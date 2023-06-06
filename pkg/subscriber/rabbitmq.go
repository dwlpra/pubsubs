package subscriber

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
)

// NewRabbitMQSubscriber adalah konstruktor untuk WatermillSubscriber.
func NewRabbitMQSubscriber(descriptor, exchange string, log bool) (WatermillSubscriber, error) {
	sub, err := amqp.NewSubscriber(amqp.Config{
		Connection: amqp.ConnectionConfig{
			AmqpURI: descriptor,
		},
		Marshaler: amqp.DefaultMarshaler{},
		Exchange: amqp.ExchangeConfig{
			GenerateName: func(topic string) string {
				return exchange
			},
			Type:    "topic",
			Durable: true,
		},
		Queue: amqp.QueueConfig{
			GenerateName: func(topic string) string {
				return topic
			},
			Durable: true,
		},
		QueueBind: amqp.QueueBindConfig{
			GenerateRoutingKey: func(topic string) string {
				return topic
			},
		},
		Publish: amqp.PublishConfig{
			GenerateRoutingKey: func(topic string) string {
				return topic
			},
			Mandatory: true,
		},
		Consume: amqp.ConsumeConfig{
			Qos: amqp.QosConfig{
				PrefetchCount: 1,
			},
		},
		TopologyBuilder: &amqp.DefaultTopologyBuilder{},
	}, watermill.NewStdLogger(log, log))
	if err != nil {
		return nil, err
	}
	return &watermillSubscriber{
		subscriber: sub,
	}, nil
}
