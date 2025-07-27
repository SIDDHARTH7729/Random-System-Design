package kafka

import (
	"context"
	"fmt"
	"math/rand"


	"sd/internal/model"

	"github.com/segmentio/kafka-go"
)

var states = []string{"completed", "failed", "pending"}

func ProduceMessage(broker string,topic string,count int) error{
	w := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
		RequiredAcks: kafka.RequireAll,
		MaxAttempts: 3,
	}

	defer w.Close()

	messages := make([]kafka.Message,0,count)
	for i := 0;i<count;i++{
		msg := model.Message{
			ID:      fmt.Sprintf("msg-%d", i),
			Content: fmt.Sprintf("This is message %d", i),
			State:   states[rand.Intn(len(states))],
		}
		kmsg := kafka.Message{
			Key:   []byte(msg.ID),
			Value: []byte(fmt.Sprintf("%s|%s|%s", msg.ID, msg.Content, msg.State)),
		}
		messages = append(messages, kmsg)


	}

	return w.WriteMessages(context.Background(), messages...)
}
