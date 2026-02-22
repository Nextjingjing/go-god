package main

import (
	"context"
	"log"

	"github.com/Nextjingjing/go-god/13-kafka/internal/utils"
	"github.com/segmentio/kafka-go"
)

func main() {
	broker := []string{"localhost:29092", "localhost:39092"}
	topic := "my-topic"

	// Ensures that a Kafka topic exists.
	// if not, then create topic via controller.
	err := utils.EnsureTopic(broker, topic, 3, 2)

	w := &kafka.Writer{
		Addr:     kafka.TCP(broker...),
		Topic:    topic,
		Balancer: &kafka.Hash{}, //
	}

	err = w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Three!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Four!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	println("Producer Success!")
}
