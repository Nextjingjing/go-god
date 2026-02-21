package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// to create topics when auto.create.topics.enable='true'
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:29092", "my-topic", 0)
	if err != nil {
		panic(err.Error())
	}
	conn.Close()

	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:29092", "localhost:39092"),
		Topic:    "my-topic",
		Balancer: &kafka.LeastBytes{},
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
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	println("Producer Success!")
}
