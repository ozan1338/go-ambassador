package events

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var Producer *kafka.Producer

func SetupProducer() {
	var err error
	Producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "pkc-ew3qg.asia-southeast2.gcp.confluent.cloud:9092",
		"security.protocol": "SASL_SSL",
		"sasl.username": "ZRITSDHTM4YORCX3",
		"sasl.password": "iOJVSZ5sHVRnmunF7VvCw+lC1iADXyNZGeYuVlZZfUlvcvUn4fotwbsxRoW2WY2W",
		"sasl.mechanism": "PLAIN",
	})
	if err != nil {
		panic(err)
	}

	// defer Producer.Close()
}


func Produce(topic string,message interface{}, key string) {

	value, _ := json.Marshal(message)

	Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key: []byte(key),
		Value:          value,
	}, nil)

	Producer.Flush(15000)
}