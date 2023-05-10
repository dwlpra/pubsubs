package subscriber

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

// WatermillSubscriber adalah interface yang harus diimplementasikan oleh jenis subscriber apa pun.
type WatermillSubscriber interface {
	Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error)
	Close() error
}

// watermillSubscriber adalah tipe struct yang mengimplementasikan interface WatermillSubscriber, digunakan untuk subscribe pesan dengan Watermill.
type watermillSubscriber struct {
	subscriber message.Subscriber
}

// Subscribe melakukan subscribe ke topic yang ditentukan dan mengembalikan channel pesan.
func (s *watermillSubscriber) Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error) {
	return s.subscriber.Subscribe(ctx, topic)
}

// Close menutup koneksi ke message broker.
func (s *watermillSubscriber) Close() error {
	return s.subscriber.Close()
}
