package publisher

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

// Publisher adalah interface yang harus diimplementasikan oleh jenis publisher apa pun.
type WatermillPublisher interface {
	Publish(topic string, messages interface{}) error
	Close() error
}

// WatermillPublisher adalah tipe struct yang mengimplementasikan interface Publisher, digunakan untuk publish pesan dengan Watermill.
type watermillPublisher struct {
	publisher  message.Publisher
	maxRetries int
}

// PublishWithRetry melakukan publish pesan dengan Watermill dan melakukan retry jika gagal.
// Retry dilakukan dengan waktu tunggu yang dinamis berdasarkan jumlah retry yang sudah dilakukan.
func (p *watermillPublisher) PublishWithRetry(topic string, msgs interface{}) error {

	var err error
	retries := 0

	dataBytes, err := json.Marshal(msgs)
	if err != nil {
		return err
	}
	msg := message.NewMessage(watermill.NewUUID(), dataBytes)

	for retries <= p.maxRetries {
		err = p.publisher.Publish(topic, msg)
		if err == nil {
			return nil
		}

		// calculate dynamic retry timeout
		retryTimeout := time.Duration(retries+1) * time.Second

		// wait before retry
		time.Sleep(retryTimeout)

		retries++
	}

	fmt.Printf("Failed to publish message after %d retries: %v\n", p.maxRetries, err)
	return fmt.Errorf("failed to publish message after %d retries: %v", p.maxRetries, err)
}

// Publish melakukan publish pesan dengan Watermill dan melakukan retry jika gagal.
func (p *watermillPublisher) Publish(topic string, msgs interface{}) error {
	err := p.PublishWithRetry(topic, msgs)
	if err != nil {
		return err
	}
	return nil
}

// Close menutup koneksi ke message broker.
func (p *watermillPublisher) Close() error {
	return p.publisher.Close()
}
