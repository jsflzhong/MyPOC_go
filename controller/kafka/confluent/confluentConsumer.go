package confluent

import (
	"fmt"
	"log"
	"mcgo/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
)

// ConfluentConsumer localhost/kafka/ConfluentConsumer
// When test, starting up the kafka consumer by calling this REST API.
func ConfluentConsumer(context *gin.Context) {
	confluentConsumer()
	context.JSON(http.StatusOK, gin.H{"status": "U R Good !"})
}

func confluentConsumer() {
	log.Println("@@@ConfluentConsumer is running ......")
	kafkaConfig := &kafka.ConfigMap{
		"metadata.broker.list":            config.GetConfig().Kafka.Broker,
		"security.protocol":               "SASL_SSL",
		"sasl.mechanisms":                 "SCRAM-SHA-256",
		"sasl.username":                   config.GetConfig().Kafka.UserName,
		"sasl.password":                   config.GetConfig().Kafka.Pswd,
		"group.id":                        config.GetConfig().Kafka.GroupId,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"default.topic.config":            kafka.ConfigMap{"auto.offset.reset": "earliest"},
		//"debug":                           "generic,broker,security",
	}
	topic := config.GetConfig().Kafka.TopicPrefix + "test"

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	log.Println("@@@ConfluentConsumer setting up......")
	kafkaConsumer, err := kafka.NewConsumer(kafkaConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "@@@Failed to create consumer: %s\n", err)
		//os.Exit(1)
	}
	fmt.Printf("@@@Created ConfluentConsumer %v\n", kafkaConsumer)
	err = kafkaConsumer.Subscribe(topic, nil)
	run := true
	counter := 0
	commitAfter := 1000
	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("@@@Caught signal %v: terminating\n", sig)
			run = false
		case event := <-kafkaConsumer.Events():
			switch e := event.(type) {
			case kafka.AssignedPartitions:
				fmt.Printf("@@@ConfluentConsumer get event : AssignedPartitions")
				kafkaConsumer.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				fmt.Printf("@@@@@@ConfluentConsumer get event :RevokedPartitions")
				kafkaConsumer.Unassign()
			case *kafka.Message:
				fmt.Printf("%% @@@@@@ConfluentConsumer get event : Message on %s: %s\n", e.TopicPartition, string(e.Value))
				counter++
				if counter > commitAfter {
					kafkaConsumer.Commit()
					counter = 0
				}
			case kafka.PartitionEOF:
				fmt.Printf("%% @@@ConfluentConsumer get event: PartitionEOF : Reached %v\n", e)
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% @@@ConfluentConsumer Error: %v\n", e)
				run = false
			}
		}
	}
	fmt.Printf("@@@ConfluentConsumer Closing consumer\n")
	kafkaConsumer.Close()
}
