package messenger

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/online.scheduling-api/config"
	"github.com/segmentio/kafka-go"
)

func Produce(ctx context.Context, topic string, message interface{}) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{config.GetMessengerBroker()},
		Topic:   topic,
	})

	content, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Unable to produce message")
	}

	w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(uuid.New().String()),
		Value: content,
	})
}
