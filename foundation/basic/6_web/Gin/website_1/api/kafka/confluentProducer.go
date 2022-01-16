package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

func ProducerConfluent() {
	config := &kafka.ConfigMap{
		"metadata.broker.list": GetBrokersString(),
		"security.protocol":    "SASL_SSL",
		"sasl.mechanisms":      "SCRAM-SHA-256",
		"sasl.username":        GetUsername(),
		"sasl.password":        GetPassword(),
		"group.id":             GetGroupId(),
		"default.topic.config": kafka.ConfigMap{"auto.offset.reset": "earliest"},
		//"debug":                           "generic,broker,security",
	}
	topic := GetTopicPrefix() + "test"
	p, err := kafka.NewProducer(config)
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created Producer %v\n", p)
	deliveryChan := make(chan kafka.Event)

	for i := 0; i < 10; i++ {
		value := fmt.Sprintf("[%d] Hello Go!", i+1)
		log.Println("@@@Producer sending the msg......")
		err = p.Produce(
			&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: []byte(value)},
			deliveryChan)
		log.Println("@@@Producer receiving the ack......")
		e := <-deliveryChan
		m := e.(*kafka.Message)
		if m.TopicPartition.Error != nil {
			fmt.Printf("@@@Delivery failed: %v\n", m.TopicPartition.Error)
		} else {
			fmt.Printf("@@@Delivered message to topic %s [%d] at offset %v\n",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}
	}
	close(deliveryChan)
}
