package main

import (
	"github.com/gin-gonic/gin"
	"goPOC/config"
	"goPOC/controller/basic"
	"goPOC/controller/nats"
	"goPOC/model/db"
	"log"
)

func main() {

	//1.加载配置文件.
	config.LoadConfiguration()
	return

	//2.加载配置log.

	//3.加载yaml配置文件进全局结构体:system.Configuration

	//4.创建DB连接,表,索引.
	initDB()

	//5.配置routers.
	router := setRouters()

	//6.配置定时任务.

	//7.启动系统.
	router.Run(":80")

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
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	setBasicRouterGroup(router)
	//setNatsRouterGroup(router)
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
		routerGroup.GET("/test", nats.NatsTest)
		routerGroup.GET("/NatsAsyncConsumer", nats.NatsAsyncConsumer)
	}
}
