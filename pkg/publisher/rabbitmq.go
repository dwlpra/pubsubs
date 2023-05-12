package publisher

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
)

// NewWatermillPublisher adalah konstruktor untuk WatermillPublisher.
func NewRabbitMQPublisher(descriptor, exchange string, maxRetries int, log bool) (WatermillPublisher, error) {
	pub, err := amqp.NewPublisher(amqp.Config{
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
				return fmt.Sprintf("%s.%s", exchange, topic)
			},
		},
		Publish: amqp.PublishConfig{
			GenerateRoutingKey: func(topic string) string {
				return fmt.Sprintf("%s.%s", exchange, topic)
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
	return &watermillPublisher{
		publisher:  pub,
		maxRetries: maxRetries,
	}, nil
}
