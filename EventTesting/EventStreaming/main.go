package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Set up a Kafka producer to send events to a topic
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "my-topic",
	})

	// Send some events to the topic
	writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("key1"),
			Value: []byte("value1"),
		},
		kafka.Message{
			Key:   []byte("key2"),
			Value: []byte("value2"),
		},
	)

	// Set up a Kafka consumer to read events from the topic
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "my-topic",
	})

	// Read events from the topic
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("Received message: key=%s, value=%s\n", string(msg.Key), string(msg.Value))
	}
}
