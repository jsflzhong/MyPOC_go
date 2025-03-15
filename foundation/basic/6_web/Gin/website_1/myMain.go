package main

import (
	"context"
	"log"
	"mcgo/foundation/basic/6_web/Gin/website_1/api" //第一段"mpoc/“要与哪里一致？ 与根目录下的文件"go.mod"里的第一行定义一致
	"mcgo/foundation/basic/6_web/Gin/website_1/api/kafka"
	"mcgo/foundation/basic/6_web/Gin/website_1/handler"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const pathStatic = "D:/michael.cui/workspace_go/src/awesomeProject/src/com.jsflzhong/6_web/Gin/website_1/website/static"
const pathFavicon = "D:/michael.cui/workspace_go/src/awesomeProject/src/com.jsflzhong/6_web/Gin/website_1/website/photo/1.jpg"

/*
test:http://localhost/index.html
*/
func main() {
	//1.生成engine
	router := gin.Default()

	//2.加载静态资源
	loadStaticResources(router)

	//3.加载所有页面模板文件
	loadTpl(router)

	//4.配置endpoints(***website分组***)
	configEndpoints(router)

	//4.2 加载kafka的配置信息, 然后启动消费者.
	//kafka.InitConfiguration()
	//kafka.Consumer()
	kafka.ConsumerConfluent()

	//5.***配置server参数,并启动server***.
	server := configAndStartServer(router)

	//6.***优雅Shutdown（或重启）服务***
	shutDownServer(server)

}

/*
定义中间件1: CORS中间件

这里的中间件, 是针对一个路由组全局生效的.
如果想针对单个路由生效,则:

	第二种: 针对单个endpoint生效的中间件的用法:
	v1.GET("/user/:id/*action",setCors(), api.GetUser) //只针对一个目标endpoint.
*/
func setCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("@@@[Middleware1:setCors] is running...")
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

/*
定义中间件2: 在filter中set一个key-value进request.
*/
func setRequestParam() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("@@@[Middleware2:setRequestParam] is running...")
		c.Set("middlewareKey1", "middlewareValue2")
		c.Next()
	}
}

/*
优雅Shutdown（或重启）服务
*/
func shutDownServer(server *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt) // syscall.SIGKILL
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
	}
	log.Println("Server exiting")
}

/*
配置server参数,并启动server.

router.Run(":80")
简单的话,像上行这样写就可以启动了，但下面所有代码（go1.8+）是为了优雅处理重启等动作。
*/
func configAndStartServer(router *gin.Engine) *http.Server {
	log.Println("@@@configAndStartServer......")
	//初始化GO内置的结构体:Server. 传入上面的engine. 并配置端口,超时时间等参数.
	server := &http.Server{
		Addr:         ":80",
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	go func() {
		// 监听请求
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	return server
}

/*
配置endpoints
*/
func configEndpoints(router *gin.Engine) {

	log.Println("@@@4.配置endpoints(***website分组***)")

	endpointGroup1StaticFiles(router)

	endpointGroup2Rest(router)

	endpointGroup3BasicFunc(router)

	endpointGroup4Auth(router)

	endpointGroup5kafka(router)
}

/*
定义endpoint分组5, kafka相关的接口
*/
func endpointGroup5kafka(router *gin.Engine) {
	routerGroup := router.Group("/kafka")
	{
		routerGroup.GET("/sendMsg", kafka.SendMsg2Kafka)
	}
}

/*
endpoints group4.  Authentication.
*/
func endpointGroup4Auth(router *gin.Engine) {
	authGroup := router.Group("/auth", gin.BasicAuth(gin.Accounts{
		"hanru":     "hanru123",
		"wangergou": "1234",
		"ruby":      "hello2",
		"lucy":      "4321",
	}))
	{
		authGroup.GET("/check", api.Authenticate)
	}
}

/*
定义endpoint分组1, engine定义的根目录是:/
*/
func endpointGroup1StaticFiles(router *gin.Engine) {
	routerGroup := router.Group("/")
	//下面几个handler, 是定义在controller.go文件那边里的.
	{
		routerGroup.GET("/index.html", handler.IndexHandler)
		routerGroup.GET("/add.html", handler.AddHandler)
		routerGroup.POST("/postme.html", handler.PostmeHandler)
	}
}

/*
定义endpoint分组2, engine定义的根目录是:/v1
此处定义的是REST风格的endpoints
*/
func endpointGroup2Rest(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		//使用中间件. 分为针对engine全局, 或针对单个目标endpoint两种.
		//中间件, 就是拦截器. 会在handler之前被调用.可做类似Cors跨域, 或rate limit限流等.
		//这个位置, 表示该组中间件只针对/v1这个根endpoints下的所有endpoint生效.
		useMiddleWare(v1)

		//配置REST风格的endpoint.
		v1.GET("/user/:id", api.GetUser)
		v1.POST("/user/insert", api.InsertUser)
		v1.PUT("/user/update", api.UpdateUser)
		v1.DELETE("/user/delete/:id", api.DeleteUser)

		//其他几种路由方式.
		routerExample(v1)

		//配置请求方式. 例如:POST,PUT, DELETE,PATCH等.
		configRequestType()
	}
}

func endpointGroup3BasicFunc(router *gin.Engine) {
	v2 := router.Group("/v2")
	{
		//使用中间件. 分为针对engine全局, 或针对单个目标endpoint两种.
		//中间件, 就是拦截器. 会在handler之前被调用.可做类似Cors跨域, 或rate limit限流等.
		//这个位置, 表示该组中间件只针对/v2这个根endpoints下的所有endpoint生效.
		useMiddleWare(v2)

		//配置REST风格的endpoint. POST.
		v2.POST("/login", api.BindJSON2VO)
		v2.GET("/user/:user/:password", api.BindURI)
		v2.GET("/people/struct2json", api.ReturnStruct)
	}
}

func routerExample(v1 *gin.RouterGroup) {
	/*v1.GET("/someGet", getting)
	v1.POST("/somePost", posting)
	v1.PUT("/somePut", putting)
	v1.DELETE("/someDelete", deleting)
	v1.PATCH("/somePatch", patching)
	v1.HEAD("/someHead", head)
	v1.OPTIONS("/someOptions", options)*/
}

/*
配置请求方式. 例如:POST,PUT, DELETE,PATCH等.
*/
func configRequestType() {
	// AJAX OPTIONS ，下面是有关OPTIONS用法的示例
	// v1.OPTIONS("/users", OptionsUser)      // POST
	// v1.OPTIONS("/users/:id", OptionsUser)  // PUT, DELETE
	/*
		// 对应的endpoint那边的API(handler)里面,要增加处理:
		func OptionsUser(c *gin.Context) {
		    c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST,PUT") //配置请求的方式.
		    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		    c.Next()
			...
		}

	*/
}

/*
使用各个自定义的中间件(类似filter的东西).

顺序问题: 中间件的调用顺序, 会按照这里的排列顺序进行.

中间件实践:

	中间件最大的作用，莫过于用于一些记录log，错误handler，还有就是对部分接口的鉴权, 验签等.
*/
func useMiddleWare(routerGroup *gin.RouterGroup) {

	routerGroup.Use(setCors())

	routerGroup.Use(setRequestParam())
}

/*
使用自定义中间件:setCors()
*/
func useCors(routerGroup *gin.RouterGroup) {
	// 第一种: 针对一整个engine组生效的中间件的用法:
	routerGroup.Use(setCors()) //针对v1的engine组
	// 第二种: 针对单个endpoint生效的中间件的用法:
	// v1.GET("/user/:id/*action",setCors(), api.GetUser) //只针对一个目标endpoint.
}

/*
加载所有页面模板，多级目录结构需要这样写
*/
func loadTpl(router *gin.Engine) {
	log.Println("@@@3.加载所有页面模板文件")
	//router.LoadHTMLGlob("src/com.jsflzhong/6_web/Gin/website_1/website/tpl/*/*")
	//currentPath := getCurrentDirectory()

	//打印: @@@currentPath: D:\code\go\home\workspace_go\workspace_go;D:\env\gopath
	//解释: 注意是两个值, 因为不但设置了全局GOPATH, 而且这个项目的IDE里还设置了project GOPATH, 也就是第一个值.
	//currentPath := os.Getenv("GOPATH")

	currentPath := os.Getenv("GOPATH")
	firstPath := strings.Split(currentPath, ";")[0]
	tplPath := filepath.Join(firstPath, "src/awesomeProject/src/com.jsflzhong/6_web/Gin/website_1/website/tpl/*/*")
	log.Println("@@@currentPath:", currentPath)
	log.Println("@@@firstPath:", firstPath)
	log.Println("@@@tplPath:", tplPath)
	router.LoadHTMLGlob(tplPath)
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

/*
静态资源加载，本例为css,js以及资源图片.这些目录下资源是可以随时更新，而不用重新启动程序.
注意:router.StaticFS("/xxx,xxx) 这是个创建endpoint的API.第二参指定了当访问第一参时,映射到项目下的哪个路径下.

测试:

	http://localhost/public/images/1.png (等于访问pathStatic/images/1.png , 而pathStatic是上面定义的常量.)
	http://localhost/favicon.ico (等于访问pathFavicon这个变量指定的文件)

每次请求响应都会在服务端有日志产生，包括响应时间，加载资源名称，响应状态值等等。
*/
func loadStaticResources(router *gin.Engine) {
	log.Println("@@@2.加载静态资源")
	// StaticFS 是加载一个完整的目录资源：
	router.StaticFS("/public", http.Dir(pathStatic))
	//StaticFile 是加载单个文件
	router.StaticFile("/favicon.ico", pathFavicon)
}
