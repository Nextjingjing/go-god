# บทที่ 13 Kafka

## สิ่งที่ต้องรู้มาก่อน
- พื้นฐานภาษา GO

## ปัญหาก่อนมี Kafka
- การสื่อสารแบบ Synchronization (เช่น REST API, gRPC) ทำให้เกิด High Coupling
  - หากมี Service ไหนตาย จะพากันตายหมด
  - หากมี Service ไหนช้า จะทำให้ช้าทั้งระบบ
- การคุยแบบเป็นคู่ๆ ทำให้ช้า ถ้ามีคนรอฟังหลายคนแล้วเราส่งข้อความได้แค่ทีละคนจะทำให้ประสิทธิภาพช้า
  
เหล่านี้คือการที่ Service สื่อสารกันตรงๆ ไม่ผ่านตัวกลางที่เรียกว่า `Broker`

## Kafka คืออะไร?
Kafka เป็นตัวกลางในการสตรีมข้อมูล ทำให้ Service ไม่ต้องสื่อสารกันตรงๆ เพียงแต่สื่อสารผ่าน `Broker`

![ภาพ message broker](/docs/images/kafka-1.png)

จากรูปเราจะเรียกผู้ส่งว่า `Producer` และผู้รับว่า `Consumer`

## องค์ประกอบของ Kafka
- `Cluster` ใน Kafka เพื่อให้เกิด High available เขาจึงสร้างหลายๆ Node เพื่อให้พอมีตัวไหนล่มก็สามารถหาสัก Node มาแทนที่
แต่ละ Node มีอยู่สอง Role ได้แก่
  1. `Broker` คือทำหน้าที่เป็นตัวกลางการสื่อสาร
  2. `Controller` คือหัวหน้าของ Cluster ไม่ได้ใช้ในการสื่อสาร แต่คอยจัดการใครตายเอาใครมาแทน และสร้าง `Topic` (เดี่ยวอธิบาย)
```
จริงๆ คือมีแค่ 1 Controller และ 1 Broker ทำงานจริงๆ ที่เหลือแค่สำรองข้อมูลและรอเสียบแทน
```

- `Topic` คือกลุ่มของข้อความโดย `Producer` จะเผยแพร่ข้อความผ่าน Topic โดย Message ที่อยู่ใน Topic เดียวกันจะเป็นข้อความประเภทเดียวกัน และ `Consumer` สามารถเลือกฟังเฉพาะ Topic ที่ตนเองอยากรู้ได้

![ภาพ message broker](/docs/images/kafka2.png)

- `Message` คือข้อความ/ข้อมูลที่สื่อสารกันภายใน Service เป็นลักษณะ `Key-Value` และใช้ Format เป็น Byte (หากอยากส่ง JSON ต้องแปลงเป็น Byte แล้วให้ Consumer แปลงเป็น JSON)
  - `Key` ***ขออธิบายทีหลัง***
  - `Value` คือข้อมูลที่อยากจะส่ง

- `Partition` คือท่อส่งข้อความของ `Topic` ซึ่ง Partition จะเชื่อมท่อกับ 1 Instance ของ Consumer เท่านั้น (ในหนึ่งช่วงเวลา อาจจะมีการเปลี่ยนท่อกันบ้าง)
 
  - ซึ่งเราควรเขียนให้ใช้ `Key` ให้ส่งไปให้ Partition เดิมหากเป็น Key เดียวกัน วิธีนี้จะรับประกันว่าขอความ เช่น จากรูปเราส่ง orderId = 1 ไปให้ partition 1 ซึ่งมันจะส่งให้ Service A Instance 1 การทำแบบนี้จะทำให้ข้อมูลไม่ผิดลำดับ

![ภาพ message broker](/docs/images/kafka3.png)

- `Group Id` คือ กลุ่มของ `Consumer` เนื่องจาก Consumer ไม่จำเป็นต้องมีแค่ 1 Instance เท่านั้น อาจจะเปิดหลายๆ ตัวก็ได้ และ Message จะถูกส่งไปทุกๆ Group Id ที่กำลังฟัง Topic อยู่ และจะได้ข้อความเดียวกัน เช่น จากรูป Message A ถูกส่งไปทุกๆ Group Id ที่กำลังฟัง Topic Order

![ภาพ message broker](/docs/images/kafka4.png)

## การติดตั้ง Kafka
### ไปอ่านคำอธิบายก่อนค่อยดูโค้ด !

สร้างไฟล์ docker-compose.yml

```yml
services:
  controller-1:
    image: apache/kafka:4.2.0
    container_name: controller-1
    restart: always
    environment:
      CLUSTER_ID: '4L6g3nShT-eMCtK--X86sw'
      KAFKA_NODE_ID: 1
      KAFKA_PROCESS_ROLES: controller
      KAFKA_LISTENERS: CONTROLLER://:9093
      KAFKA_CONTROLLER_LISTENER_NAMES: CONTROLLER
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: CONTROLLER:PLAINTEXT
      KAFKA_CONTROLLER_QUORUM_VOTERS: 1@controller-1:9093,2@controller-2:9093
      KAFKA_LOG_DIRS: '/var/lib/kafka/data'
    volumes:
      - controller-1-data:/var/lib/kafka/data

  controller-2:
    image: apache/kafka:4.2.0
    container_name: controller-2
    restart: always
    environment:
      CLUSTER_ID: '4L6g3nShT-eMCtK--X86sw'
      KAFKA_NODE_ID: 2
      KAFKA_PROCESS_ROLES: controller
      KAFKA_LISTENERS: CONTROLLER://:9093
      KAFKA_CONTROLLER_LISTENER_NAMES: CONTROLLER
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: CONTROLLER:PLAINTEXT
      KAFKA_CONTROLLER_QUORUM_VOTERS: 1@controller-1:9093,2@controller-2:9093
      KAFKA_LOG_DIRS: '/var/lib/kafka/data'
    volumes:
      - controller-2-data:/var/lib/kafka/data

  broker-1:
    image: apache/kafka:4.2.0
    container_name: broker-1
    restart: always
    ports:
      - 29092:9092
    depends_on:
      - controller-1
      - controller-2
    environment:
      CLUSTER_ID: '4L6g3nShT-eMCtK--X86sw'
      KAFKA_NODE_ID: 3
      KAFKA_PROCESS_ROLES: broker

      KAFKA_LISTENERS: 'PLAINTEXT://:19092,PLAINTEXT_HOST://:9092'
      KAFKA_ADVERTISED_LISTENERS: 'PLAINTEXT://broker-1:19092,PLAINTEXT_HOST://localhost:29092'
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT

      KAFKA_CONTROLLER_LISTENER_NAMES: CONTROLLER
      KAFKA_CONTROLLER_QUORUM_VOTERS: 1@controller-1:9093,2@controller-2:9093

      KAFKA_LOG_DIRS: '/var/lib/kafka/data'

      # topic defaults
      KAFKA_AUTO_CREATE_TOPICS_ENABLE: "false"
      KAFKA_NUM_PARTITIONS: 3
      KAFKA_DEFAULT_REPLICATION_FACTOR: 2

      # internal topics
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 2
      KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR: 2
      KAFKA_TRANSACTION_STATE_LOG_MIN_ISR: 1

    volumes:
      - broker-1-data:/var/lib/kafka/data

  broker-2:
    image: apache/kafka:4.2.0
    container_name: broker-2
    restart: always
    ports:
      - 39092:9092
    depends_on:
      - controller-1
      - controller-2
    environment:
      CLUSTER_ID: '4L6g3nShT-eMCtK--X86sw'
      KAFKA_NODE_ID: 4
      KAFKA_PROCESS_ROLES: broker

      KAFKA_LISTENERS: 'PLAINTEXT://:19092,PLAINTEXT_HOST://:9092'
      KAFKA_ADVERTISED_LISTENERS: 'PLAINTEXT://broker-2:19092,PLAINTEXT_HOST://localhost:39092'
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT

      KAFKA_CONTROLLER_LISTENER_NAMES: CONTROLLER
      KAFKA_CONTROLLER_QUORUM_VOTERS: 1@controller-1:9093,2@controller-2:9093

      KAFKA_LOG_DIRS: '/var/lib/kafka/data'

      # topic defaults
      KAFKA_AUTO_CREATE_TOPICS_ENABLE: "false"
      KAFKA_NUM_PARTITIONS: 3
      KAFKA_DEFAULT_REPLICATION_FACTOR: 2

      # internal topics
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 2
      KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR: 2
      KAFKA_TRANSACTION_STATE_LOG_MIN_ISR: 1

    volumes:
      - broker-2-data:/var/lib/kafka/data
  
  kafka-ui:
    image: provectuslabs/kafka-ui:latest
    ports:
      - "8080:8080"
    depends_on:
      - broker-1
      - broker-2
    environment:
      - KAFKA_CLUSTERS_0_NAME=local-cluster
      - KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS=broker-1:19092,broker-2:19092
      - KAFKA_CLUSTERS_0_METRICS_PORT=9997

volumes:
  controller-1-data:
  controller-2-data:
  broker-1-data:
  broker-2-data:
```

- `container_name: controller-1` และ `container_name: controller-2`
  - `KAFKA_PROCESS_ROLES: controller` มี Role เป็น `Controller`
  - `KAFKA_LISTENERS: CONTROLLER://:9093` ช่องทางที่ Controller คุยกัน
  - `KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: CONTROLLER:PLAINTEXT` คุยกันแบบ Plain text ไม้เข้ารหัส
  - `KAFKA_CONTROLLER_QUORUM_VOTERS: 1@controller-1:9093,2@controller-2:9093` หาก 1 ตาย 2 จะเป็นหัวหน้าแทน
- `container_name: broker-1` และ `container_name: broker-2`
  - `KAFKA_AUTO_CREATE_TOPICS_ENABLE: "false"` ไม่ให้ Producer สร้าง Topic ได้เองโดยไม่ผ่าน Controller
  - `KAFKA_NUM_PARTITIONS: 3` จำนวน Partition ใน 1 Topic
  - `KAFKA_DEFAULT_REPLICATION_FACTOR: 2` จำนวนการสำรองข้อมูลใช้ 2 เพราะมี 2 node
  - `# internal topics` อันนี้คือการสื่อสารกันภายใน Cluster เท่านั้น

## เขียนโค้ด Go

### Producer

การสร้าง Topic

```go
package utils

import (
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
)


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
```

- เราเชื่อมต่อไปที่สัก Broker ก่อน
- จากนั้นเราจะติดต่อกับ Controller ผ่าน Broker ตัวนั้น
- จากนั้นสร้าง Topic
```go
topicConfig := kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     partitions,
		ReplicationFactor: replication,
	}
```
- หากมี Topic แล้วจะไม่สร้าง Topic
  - หากไม่มี Topic จะสร้าง Topic

การส่งข้อมูล

```go
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
```

ให้สังเกตที่
```go
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
```
ผมส่ง Key เป็น Key-C เหมือนกัน หากส่งไปควรจะได้ Partition เดียวกัน

## Consumer

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:29092", "localhost:39092"},
		Topic:    "my-topic",
		GroupID:  "demo-group",
		MaxBytes: 10e6, // 10MB
	})
	r.SetOffset(42)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
```

- `GroupID:  "demo-group",` ผมสร้าง Group Id ขึ้นมา

## ลองรันดู
ผมได้สร้าง `Makefile` เพื่อให้ทุกคนรันโค้ดได้อย่างสะดวก ให้เปิดสอง Terminal นะ

### Producer's terminal
```bash
make up
go run ./cmd/producer/producer.go
```

### Consumer's terminal
```bash
go run ./cmd/consumer/consumer.go
```

ที่ Consumer เรามีผลลัพธ์ดังนี้
```bash
message at offset 0: Key-C = Two!
message at offset 0: Key-A = Hello World!
message at offset 0: Key-B = One!
```

และเรามี `kafka-ui` โดยเปิดที่ `http://localhost:8080/`

![kafka-ui](/docs/images/kafka5.png)

ลอง Produce หลายๆ ครั้ง

![kafka-ui](/docs/images/kafka6.png)

สังเกตว่า key-c จะอยู่ใน partition เดียวกัน (ขอย้ำนะ แค่ในหนึ่งหน่วยเวลา)