// api 是专门提供RESTful接口的路由请求处理handler
// 除了GET方法，也支持POST,PUT,DELETE,OPTION等常用的restful方法

package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"mcgo/foundation/basic/6_web/Gin/website_1/common"
	"mcgo/foundation/basic/6_web/Gin/website_1/model"
	"mcgo/foundation/basic/6_web/Gin/website_1/model/VO/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/**
这里主要是开发REST endpoints的文件
关于后端endpoint接口的细节, 请见README中的第六条.
*/

/*
1.endpoint: GET, 查询单个user.

	url: v1/user/:id

	测试: http://localhost/v1/user/1

	注意, handler(endpoint)的参数是需要符合接口规范的.
*/
func GetUser(c *gin.Context) {
	//1.获取request中的入参
	id := c.Param("id")
	//string 转 int64
	intID, _ := strconv.ParseInt(id, 10, 64)

	//2.调用结构体中的API, 来实现对DB的操作逻辑.查询user.
	user, err := new(model.User).GetUser(intID)
	switch {
	case err == sql.ErrNoRows:
		//对于查不到对应的数据这种error,是正常的情况,一般还是会返回200, 然后用自定义的状态码来约定前后端.
		c.JSON(http.StatusOK, new(response.BaseResponse).NoData("No data can be queried."))
		return
	case err != nil:
		c.JSON(http.StatusInternalServerError, new(response.BaseResponse).Error(err.Error()))
		return
	}

	//3.把查询到的entity进行Json化,并处理标准的response返回.
	baseResponse := new(response.BaseResponse)
	userJson, err := json.Marshal(user)
	if common.CheckError(err) {
		c.JSON(http.StatusInternalServerError, baseResponse.Error(err.Error()))
	}
	c.JSON(http.StatusOK, baseResponse.Ok(string(userJson)))
}

/*
   2.endpoint: PUT, 根据id修改user的其他字段的值.

   	url: v1/user/insert

   	测试:
   		http://localhost/v1/user/insert
   		PUT  {"nick":"Michael4","Password":"123","Age":"14"}

*/

func InsertUser(c *gin.Context) {
	user := new(model.User)
	err := c.Bind(user)
	if common.CheckError(err) {
		//这个API会终止本次请求.
		log.Fatal("@@@[InsertUser]", err)
	}
	rowNum, err := new(model.User).InsertUser(*user)
	if common.CheckError(err) {
		c.JSON(http.StatusInternalServerError, new(response.BaseResponse).Error(err.Error()))
	}
	c.JSON(http.StatusOK, new(response.BaseResponse).Ok(strconv.Itoa(rowNum)))
}

/*
3.endpoint: POST, 插入单个user.

	url: v1/user/update

	测试:
		http://localhost/v1/user/update
		PUT  {"uid":"8", "nick":"Michael5","Password":"125","Age":"15"}
*/
func UpdateUser(c *gin.Context) {
	//把入参的JSON转换到结构体上. 入参中JSON的key,需要与User结构体中的字段后的'json'值一致.
	user := new(model.User)
	err := c.Bind(user)
	if common.CheckError(err) {
		//这个API会终止本次请求.
		//log.Fatal("@@@[UpdateUser]",err)
		c.JSON(http.StatusInternalServerError, new(response.BaseResponse).Error(err.Error()))
	}
	rowNum, err := new(model.User).UpdateUser(*user)
	if common.CheckError(err) {
		c.JSON(http.StatusInternalServerError, new(response.BaseResponse).Error(err.Error()))
	}
	c.JSON(http.StatusOK, new(response.BaseResponse).Ok(strconv.Itoa(rowNum)))
}

/*
4.endpoint: DELETE, 删除单个user.

	url: v1/user/delete/:id

	测试:
		DELETE	 http://localhost/v1/user/delete/8
*/
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idString, err := strconv.Atoi(id)
	if common.CheckError(err) {
		c.JSON(http.StatusInternalServerError, new(response.BaseResponse).Error(err.Error()))
	}
	rowNum, err := new(model.User).DeleteUser(idString)
	if common.CheckError(err) {
		c.JSON(http.StatusInternalServerError, new(response.BaseResponse).Error(err.Error()))
	}
	c.JSON(http.StatusOK, new(response.BaseResponse).Ok(strconv.Itoa(rowNum)))
}

/*
   string转成int：
   int, err := strconv.Atoi(string)
   string转成int64：
   int64, err := strconv.ParseInt(string, 10, 64)
   int转成string：
   string := strconv.Itoa(int)
   int64转成string：
   string := strconv.FormatInt(int64,10)
*/
