package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

type SensorData struct {
	SensorID    string  `json:"sensor_id"`
	Timestamp   int64   `json:"timestamp"`
	Poluente   string  `json:"poluente"`
	Valor float64 `json:"valor"`
}

func produceData(producer *kafka.Producer, topic string) {
	sensorData := SensorData{
		SensorID:    "123",
		Timestamp:   time.Now().Unix(),
		Poluente:   "PM2.5",
		Valor: 35.4,
	}

	dataBytes, err := json.Marshal(sensorData)
	if err != nil {
		panic(err)
	}

	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value: dataBytes,
	}, nil)

	producer.Flush(15 * 1000)
}

func consumeData(consumer *kafka.Consumer, topic string) {
	consumer.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			var sensorData SensorData
			if err := json.Unmarshal(msg.Value, &sensorData); err != nil {
				fmt.Printf("Error decoding sensor data: %s\n", err)
				continue
			}
			fmt.Printf("Received sensor data: %+v\n", sensorData)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}
}

func main() {
	topic := "qualidadeAr"

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092,localhost:39092",
		"client.id": "go-producer",
	})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	produceData(producer, topic)

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092,localhost:39092",
		"group.id": "go-consumer-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	consumeData(consumer, topic)
}
