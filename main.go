package main

import (
	"github.com/gin-gonic/gin"
	"goPOC/controller/basic"
)

func main() {

	//1.加载配置文件.

	//2.加载配置log.

	//3.加载yaml配置文件进全局结构体:system.Configuration

	//4.创建DB连接,表,索引.

	//5.配置routers.
	router := setRouters()

	//6.配置定时任务.

	//7.启动系统.
	router.Run(":80")

}

func setRouters() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	setBasicRouterGroup(router)
	return router
}

func setBasicRouterGroup(router *gin.Engine) {
	routerGroup := router.Group("/basic")
	{
		routerGroup.GET("/simpleTest", basic.SimpleTest)
	}
}
