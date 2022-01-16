package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

/**
这是一个基于Sarama库的对接Kafka的客户端.

Apache Kafka 的 Golang 客户端库 Sarama。Sarama 是 MIT 许可的 Apache Kafka 0.8 及更高版本的 Golang 客户端库。

1.1 安装依赖库sarama
go get github.com/Shopify/sarama
该库要求kafka版本在0.8及以上，支持kafka定义的high-level API和low-level API，但不支持常用的consumer自动rebalance和offset追踪，
所以一般得结合cluster版本使用。


1.2 sarama-cluster依赖库
go get github.com/bsm/sarama-cluster
需要kafka 0.9及以上版本。代码示例来自官网，可到官网查看更多信息。

生产者:
	我们可以使用 Sarama 库的 AsyncProducer 或 SyncProducer 生产消息。
	在大多数情况下首选使用 AsyncProducer 生产消息。它通过一个 channel 接收消息，并在后台尽可能高效的异步生产消息。

	SyncProducer 发送 Kafka 消息后阻塞，直到接收到 ACK 确认。
	SyncProducer 有两个警告：它通常效率较低，并且实际的耐用性保证取决于 Producer.RequiredAcks 的配置值。
	在某些配置中，有时仍会丢失由 SyncProducer 确认的消息，但是使用比较简单。

消费者:

*/

/**
Configuration
*/

var configProducer *sarama.Config
var configConsumer *cluster.Config
var consumerConfig2 *sarama.Config
var groupID string
var topic string

/**
请求: localhost/kafka/sendMsg
*/
func SendMsg2Kafka(context *gin.Context) {
	//AsyncProducer()
	//kafkaClient()
	ProducerConfluent()
	context.JSON(http.StatusOK, gin.H{"status": "U r ok"})
}

// 初始配置. 需要被main调用.
func InitConfiguration() {
	initProducerConfig()
	initConsumerConfigWithCluster()
}

func initProducerConfig() {
	//设置配置
	configProducer = sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	configProducer.Producer.RequiredAcks = sarama.WaitForAll
	//随机的分区类型
	configProducer.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	configProducer.Producer.Return.Successes = true
	configProducer.Producer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	configProducer.Version = sarama.V0_11_0_0
}

func initConsumerConfigWithCluster() {
	groupID = "group-1"
	topic = "ptxqyt3u-"
	configConsumer = cluster.NewConfig()
	configConsumer.Group.Return.Notifications = true
	configConsumer.Consumer.Offsets.CommitInterval = 1 * time.Second
	configConsumer.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始
}

func initConsumerConfig() {
	groupID = "group-1"
	topic = "ptxqyt3u-"
	consumerConfig2 = sarama.NewConfig()
	consumerConfig2.Version = sarama.V2_8_0_0 // specify appropriate version
	consumerConfig2.Consumer.Return.Errors = false
	//consumerConfig.Consumer.Offsets.AutoCommit.Enable = true      // 禁用自动提交，改为手动
	//consumerConfig.Consumer.Offsets.AutoCommit.Interval = time.Second * 1 // 测试3秒自动提交
	consumerConfig2.Consumer.Offsets.Initial = sarama.OffsetNewest
}

//异步的生产者
func AsyncProducer() {
	log.Println("@@@asyncProducer is running...")
	//使用配置,新建一个异步生产者
	producer, e := sarama.NewAsyncProducer(
		[]string{"sulky-01.srvs.cloudkafka.com:9094", "sulky-02.srvs.cloudkafka.com:9094", "sulky-03.srvs.cloudkafka.com:9094"},
		configProducer)
	if e != nil {
		panic(e)
	}
	defer producer.AsyncClose()

	//设置要发送的消息的: 主题, key
	msg := &sarama.ProducerMessage{
		Topic: "ptxqyt3u-",
		Key:   sarama.StringEncoder("myKey1"),
	}

	//设置要发送的消息的: value
	randomString := getRandomString()
	//将字符串转化为字节数组
	msg.Value = sarama.ByteEncoder(randomString)

	//使用通道发送消息
	producer.Input() <- msg

	//接收ack
	select {
	case suc := <-producer.Successes():
		fmt.Println("@@@offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
	case fail := <-producer.Errors():
		fmt.Println("@@@err: ", fail.Err)
	}
}

func syncProducer() {

}

//同步的消费者. 需要被main调用.
func Consumer() {
	log.Println("@@@consumer is running...")
	c, err := cluster.NewConsumer(
		GetBrokers(),
		groupID,
		strings.Split(topic, ","),
		configConsumer)
	if err != nil {
		panic(err)
		return
	}
	defer c.Close()
	go func(c *cluster.Consumer) {
		errors := c.Errors()
		notification := c.Notifications()
		for {
			select {
			case err := <-errors:
				panic(err)
			case <-notification:
			}
		}
	}(c)

	for msg := range c.Messages() {
		fmt.Fprintf(os.Stdout, "@@@Consumer get msg: %s/%d/%d\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
		c.MarkOffset(msg, "") //MarkOffset 并不是实时写入kafka，有可能在程序crash时丢掉未提交的offset
	}
}

func Consumer2() {
	consumer, err := sarama.NewConsumer(GetBrokers(), consumerConfig2)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = consumer.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}()

}

func kafkaClient() {
	config := sarama.NewConfig()
	config.Version = sarama.V0_10_0_0
	client, err := sarama.NewClient(GetBrokers(), config)
	if err != nil {
		panic("client create error")
	}
	defer client.Close()
	//获取主题的名称集合
	topics, err := client.Topics()
	if err != nil {
		panic("@@@get topics err")
	}
	for _, e := range topics {
		fmt.Println("@@@topic:", e)
	}
	//获取broker集合
	brokers := client.Brokers()
	//输出每个机器的地址
	for _, broker := range brokers {
		fmt.Println("@@@broker:", broker.Addr())
	}
}

func getRandomString() string {
	unixTime := time.Now().Unix()
	randInt := rand.Intn(100)
	randomString := fmt.Sprintf("%d%d", unixTime, randInt)
	log.Println("@@@randomString:", randomString)
	return randomString
}
