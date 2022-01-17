package nats

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"log"
	"net/http"
)

const natsURL = "nats://42.192.156.142:4222"

// NatsPublisher localhost/nats/NatsPublisher
// Should start the consumer first through the REST API below.
func NatsPublisher(context *gin.Context) {
	log.Println("@@@NatsPublisher is running......")
	testPublisher()
	context.JSON(http.StatusOK, gin.H{"status": "U r ok"})
}

func testPublisher() {
	// Connect to a server
	log.Println("@@@Publisher connecting to nats server......")
	//nc, _ := nats.Connect(nats.DefaultURL)
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal("@@@ Nats connection is not success:", err)
	}

	// Simple Publisher
	log.Println("@@@Publishing message to nats server......")
	nc.Publish("foo", []byte("Hello World"))

	// Simple Async Subscriber
	//log.Println("@@@Subscribing nats server......")
	//nc.Subscribe("foo", func(m *nats.Msg) {
	//	fmt.Printf("@@@Async Received a message: %s\n", string(m.Data))
	//})

	// Responding to a request message
	//nc.Subscribe("request", func(m *nats.Msg) {
	//	m.Respond([]byte("answer is 42"))
	//})

	// Simple Sync Subscriber
	//sub, err := nc.SubscribeSync("foo")
	//m, err := sub.NextMsg(10)
	//fmt.Printf("@@@Sync Received a message: %s\n", string(m.Data))

	//// Channel Subscriber
	//ch := make(chan *nats.Msg, 64)
	//sub, err := nc.ChanSubscribe("foo", ch)
	//msg := <-ch
	//fmt.Printf("@@@Sync Received a message: %s\n", string(msg.Data))
	//
	//// Unsubscribe
	//sub.Unsubscribe()
	//
	//// Drain
	//sub.Drain()
	//
	//// Requests
	//msg, err := nc.Request("help", []byte("help me"), 10*time.Millisecond)
	//
	//// Replies
	//nc.Subscribe("help", func(m *nats.Msg) {
	//	nc.Publish(m.Reply, []byte("I can help!"))
	//})

	// Drain connection (Preferred for responders)
	// Close() not needed if this is called.
	nc.Drain()

	// Close connection
	nc.Close()
}

// NatsAsyncConsumer localhost/nats/NatsAsyncConsumer
func NatsAsyncConsumer(context *gin.Context) {
	log.Println("@@@NatsAsyncConsumer is running......")
	natsConsumer()
	context.JSON(http.StatusOK, gin.H{"status": "U r ok"})
}

func natsConsumer() {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal("@@@ Nats connection is not success:", err)
	}

	// Simple Async Subscriber
	log.Println("@@@Subscribing nats server......")
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("@@@Async Received a message: %s\n", string(m.Data))
	})

}
