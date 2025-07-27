package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"strings"
	"sd/internal/model"
	"sd/internal/store"
)

func StartConsumer(broker string, topic string, store *store.Store) error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		GroupID: "message-group",
		Topic:  topic,
		MaxAttempts: 3,
	})
	defer r.Close()

	for{
		message,err := r.ReadMessage(context.Background())
		if err != nil{
			return err
		}

		parts := strings.Split(string(message.Value), "|")
		if len(parts) != 3 {
			continue // Invalid message format
		}

		msg := model.Message{
			ID:      parts[0],
			Content: parts[1],
			State:   parts[2],
		}
		store.AddMessage(msg)
	}
}