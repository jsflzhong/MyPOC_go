package confluent

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"goPOC/config"
	"log"
	"net/http"
	"os"
)

// ConfluentProducer localhost/kafka/ConfluentProducer
func ConfluentProducer(context *gin.Context) {
	confluentProducer()
	context.JSON(http.StatusOK, gin.H{"status": "U R OK !"})
}

func confluentProducer() {
	log.Println("@@@ConfluentProducer is running ......")
	kafkaConfig := &kafka.ConfigMap{
		"metadata.broker.list": config.GetConfig().Kafka.Broker,
		"security.protocol":    "SASL_SSL",
		"sasl.mechanisms":      "SCRAM-SHA-256",
		"sasl.username":        config.GetConfig().Kafka.UserName,
		"sasl.password":        config.GetConfig().Kafka.Pswd,
		"group.id":             config.GetConfig().Kafka.GroupId,
		"default.topic.config": kafka.ConfigMap{"auto.offset.reset": "earliest"},
		//"debug":                           "generic,broker,security",
	}
	topic := config.GetConfig().Kafka.TopicPrefix + "test"
	p, err := kafka.NewProducer(kafkaConfig)
	if err != nil {
		log.Println("@@@Failed to create producer", err)
		os.Exit(1)
	}
	fmt.Printf("Created Producer %v\n", p)
	deliveryChan := make(chan kafka.Event)

	for i := 0; i < 10; i++ {
		value := fmt.Sprintf("[%d] Hello Go!", i+1)
		log.Println("@@@ConfluentProducer sending the msg......")
		err = p.Produce(
			&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: []byte(value)},
			deliveryChan)
		log.Println("@@@ConfluentProducer receiving the ack......")
		e := <-deliveryChan
		m := e.(*kafka.Message)
		if m.TopicPartition.Error != nil {
			fmt.Printf("@@@ConfluentProducer delivery failed: %v\n", m.TopicPartition.Error)
		} else {
			fmt.Printf("@@@ConfluentProducer delivered message to topic %s [%d] at offset %v\n",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}
	}
	close(deliveryChan)
}
