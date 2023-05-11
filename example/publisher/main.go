package main

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/dwlpra/pubsubs/pkg/publisher"
)

func main() {
	rabbitmqURL := "amqp://guest:guest@localhost:5672/"

	// Initialize WatermillPublisher
	abc, _ := publisher.NewRabbitMQPublisher(rabbitmqURL, "exchange_name", 5, false)

	// Prepare a message

	msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, world!"))

	// Publish a message using WatermillPublisher
	for i := 0; i < 99999; i++ {
		err := abc.Publish("like_tt", msg)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("success")

}
