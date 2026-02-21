package utils

import (
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
)

// EnsureTopic ensures that a Kafka topic exists with the specified configuration.
// It first dials a bootstrap broker to discover the current cluster controller,
// then establishes a direct connection to the controller to request topic creation.
//
// This function is idempotent: if the topic already exists, it returns nil
// instead of an error.
//
// Parameters:
//   - brokers: A slice of broker addresses used for initial discovery.
//   - topic: The name of the topic to ensure.
//   - partitions: The desired number of partitions.
//   - replication: The replication factor (should not exceed the number of available brokers).
//
// Returns nil if the topic is created or already exists; otherwise, returns the connection
// or creation error.
func EnsureTopic(brokers []string, topic string, partitions int, replication int) error {
	conn, err := kafka.Dial("tcp", brokers[0])
	if err != nil {
		return err
	}
	defer conn.Close()

	// To create a topic manually when auto.create.topics.enable='false',
	// we must communicate directly with the Cluster Controller.
	controller, err := conn.Controller()
	if err != nil {
		return err
	}

	// Connect to the Controller node
	controllerAddr := net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port))
	controllerConn, err := kafka.Dial("tcp", controllerAddr)
	if err != nil {
		return err
	}
	defer controllerConn.Close()

	topicConfig := kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     partitions,
		ReplicationFactor: replication,
	}

	err = controllerConn.CreateTopics(topicConfig)

	// if Topic is already exists, then no error
	if err == kafka.TopicAlreadyExists {
		return nil
	}

	if err != nil {
		return err
	}

	return nil
}
