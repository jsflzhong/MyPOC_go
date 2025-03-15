package main

import (
	"log"
	"mcgo/config"
	"mcgo/controller/basic"
	"mcgo/controller/kafka/confluent"
	"mcgo/controller/nats"
	"mcgo/model/db"

	"github.com/gin-gonic/gin"
)

func main() {

	//1.加载配置文件.
	config.LoadConfiguration()

	//2.加载配置log.

	//3.加载yaml配置文件进全局结构体:system.Configuration

	//4.创建DB连接,表,索引.
	//initDB()

	//5.配置routers.
	router := setRouters()

	//6.配置定时任务.

	//7.启动系统.
	startupSystem(router)

}

func startupSystem(router *gin.Engine) {
	port := ":80"
	log.Println("@@@Starting the system with port:", port)
	router.Run(port)
}

func initDB() {
	db, err := db.InitMysql()
	if err != nil {
		log.Fatal("@@@There is an error when open connection to mysql!")
		return
	}
	defer db.Close()
}

func setRouters() *gin.Engine {
	log.Println("@@@Start setting routers......")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	setBasicRouterGroup(router)
	setNatsRouterGroup(router)
	setKafkaRouterGroup(router)
	log.Println("@@@Finish setting routers")
	return router
}

func setBasicRouterGroup(router *gin.Engine) {
	routerGroup := router.Group("/basic")
	{
		routerGroup.GET("/simpleTest", basic.SimpleTest)
	}
}

func setNatsRouterGroup(router *gin.Engine) {
	routerGroup := router.Group("/nats")
	{
		routerGroup.GET("/NatsPublisher", nats.NatsPublisher)
		routerGroup.GET("/NatsAsyncConsumer", nats.NatsAsyncConsumer)
	}
}

func setKafkaRouterGroup(router *gin.Engine) {
	routerGroup := router.Group("/kafka")
	{
		routerGroup.GET("/ConfluentProducer", confluent.ConfluentProducer)
		routerGroup.GET("/ConfluentConsumer", confluent.ConfluentConsumer)
	}
}
