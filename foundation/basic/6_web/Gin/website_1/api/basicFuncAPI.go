package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"goPOC/foundation/basic/6_web/Gin/website_1/common"
	"goPOC/foundation/basic/6_web/Gin/website_1/model"
	"goPOC/foundation/basic/6_web/Gin/website_1/model/VO/request"
	"goPOC/foundation/basic/6_web/Gin/website_1/model/VO/response"
	"log"
	"net/http"
)

/*
1.POST:绑定POST请求体中的json字符串,到一个VO.
	用 c.Bind() 这个函数, 来取出请求中的jason并绑定到一个VO上.

endpoint:http://127.0.0.1:8080/v2/login

测试:#curl -v -X POST http://127.0.0.1/v2/login -H 'content-type:application/json' -d '{"user":"Michael","password":"123"}'
*/
func BindJSON2VO(c *gin.Context) {
	var loginVO request.Login
	// 其实就是将request中的Body中的数据按照JSON格式解析到json变量中
	// Login 这个VO那边有做 字段对应jason key名字 的处理, 去看.
	// 基础知识参考: 'Struct_4_JSON.go'
	if err := c.Bind(&loginVO); err != nil {
		log.Println("@@@bind error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if loginVO.User != "Michael" || loginVO.Password != "123" {
		log.Println("@@@401, unauthorized")
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

/*
2.GET: 返回:由结构体转成的JSON.
	用c.JSON()这个函数来返回response, 注意, 很多后端对前端的交互处理, 都是Context这个对象的函数来完成的.

endpoint:http://127.0.0.1/v2/people/struct2json
	注意:此时的key,要在路由配置那边, 配置的与vo中的binding的key相匹配. content-type是URI.

测试
*/
func ReturnStruct(c *gin.Context) {
	baseResponse := response.BaseResponse{}
	//用defer和recover来抓本层和底层抛出的异常并做处理和返回到FE.
	defer func() {
		if err := recover(); err != nil {
			log.Println("@@@[ReturnStruct]error:", err) //打印的err,是底层调用panic("...")自定义的字符串.
			errString := err.(string)
			c.JSON(http.StatusBadRequest, baseResponse.Error(errString))
		}
	}()
	people := model.People{Id: 1, Name: "name", Age: 11}
	peopleJson, err := json.Marshal(people)
	if common.CheckError(err) {
		c.JSON(http.StatusBadRequest, baseResponse.Error(err.Error()))
	}
	c.JSON(http.StatusOK, baseResponse.Ok(string(peopleJson)))

}

/*
3.GET. URI绑定入参到VO.
	用函数:c.ShouldBindUri() 来完成. 依然是操作 Context 这个关键对象.

endpoint:http://127.0.0.1/v2/user/:user/:password
	注意:此时的key,要在路由配置那边, 配置的与vo中的binding的key相匹配. content-type是URI.

测试:#curl -v http://127.0.0.1/v2/user/Michael/123
*/
func BindURI(c *gin.Context) {
	var login request.Login
	if err := c.ShouldBindUri(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"username": login.User, "password": login.Password})
}

/*
4.GET: 认证

Endpoint:http://127.0.0.1/auth/check
result:{"secret":{"email":"hanru@163.com","phone":"123433"},"user":"hanru"}
*/
func Authenticate(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	if secret, ok := secrets[user]; ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}
}

//模拟从DB取出三个用户数据, 用来为上面的登录功能提供用户数据.
var secrets = gin.H{
	"hanru":     gin.H{"email": "hanru@163.com", "phone": "123433"},
	"wangergou": gin.H{"email": "wangergou@example.com", "phone": "666"},
	"ruby":      gin.H{"email": "ruby@guapa.com", "phone": "523443"},
}
