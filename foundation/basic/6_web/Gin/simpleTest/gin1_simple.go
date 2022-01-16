package simpleTest

//需要提前在CMD中下载gin包:#go get -u github.com/gin-gonic/gin (前面的部分:#go get -u 是固定的, 后面只需要添加想下载的包即可.)
import "github.com/gin-gonic/gin"

/*
编译运行程序，打开浏览器，访问http://localhost:8080/ping页面显示：
{"message":"pong"}
 */
func SimpleTest()  {
	//Default返回一个 默认的路由引擎
	engine := gin.Default()

	//定义GET请求,
	//定义endpoint:/ping,
	//定义一个用于处理请求的匿名函数.
	engine.GET("/ping", func(context *gin.Context) {
		//输出json结果给调用方. H is a shortcut for map[string]interface{}
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// listen and serve on 0.0.0.0:8080
	engine.Run()
}
