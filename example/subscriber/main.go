package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dwlpra/pubsubs/pkg/subscriber"
)

func main() {
	// Ganti "amqp://guest:guest@localhost:5672/" dengan string koneksi RabbitMQ Anda
	sub, err := subscriber.NewRabbitMQSubscriber("amqp://guest:guest@localhost:5672/", "abc", true)
	if err != nil {
		log.Fatalf("Failed to create subscriber: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Ganti "your_topic" dengan topik yang ingin Anda subscribe
	msgs, err := sub.Subscribe(ctx, "sss")
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}

	// Menerima sinyal keluar
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		a := 0
		for msg := range msgs {
			// Logika pemrosesan pesan Anda
			a++
			fmt.Println(string(msg.Payload))
			// Acknowledge pesan setelah diproses
			msg.Ack()
		}
	}()

	// Menunggu sinyal keluar
	<-sigChan

	// Membersihkan sumber daya saat keluar
	cancel()
	time.Sleep(1 * time.Second)

	if err := sub.Close(); err != nil {
		log.Fatalf("Failed to close subscriber: %v", err)
	}

	log.Println("Subscriber stopped")
}
