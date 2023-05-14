package main

import (
	"fmt"

	"github.com/dwlpra/pubsubs/pkg/publisher"
)

type Model struct {
	Nama string
	Umur int
}

func main() {
	rabbitmqURL := "amqp://guest:guest@localhost:5672/"

	// Initialize WatermillPublisher
	abc, _ := publisher.NewRabbitMQPublisher(rabbitmqURL, "exchange_name", "routing_key", 5, true)

	// Prepare a message
	data := Model{
		Nama: "ade dwi putra",
		Umur: 18,
	}
	// Publish a message using WatermillPublisher
	for i := 0; i < 99999; i++ {
		err := abc.Publish("like_tt", data)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("success")

}
