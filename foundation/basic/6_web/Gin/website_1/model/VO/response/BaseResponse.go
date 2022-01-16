package response

import (
	"goPOC/foundation/basic/6_web/Gin/website_1/common"
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
	//如果想让json.Marshal(Pojo)生效, 则字段名要首大.
	Code    int    `json:"Code"`
	Message string `json:"Message"`
	Data    string `json:"Data"`
}

//自定义状态码. 10001:未找到要查询的数据.
var noData = 10001

func (baseResponse BaseResponse)Ok(data string) string {
	response := BaseResponse{
		Code:    http.StatusOK,
		Message: "success",
	}
	if data != "" {
		response.Data = data
	}
	//common.TestPanic()
	marshal, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	return string(marshal)
}

func (baseResponse BaseResponse)Error(errString string) string{
	response := BaseResponse{
		Code:    http.StatusBadRequest,
		Message: "Error",
	}
	if errString != "" {
		response.Data = errString
	}
	responseJson, err := json.Marshal(response)
	if common.CheckError(err) {
		panic(err)
	}
	return string(responseJson)
}

/*
为找到要查询的资源. 外层用c.json(http.StatusOK, ...)返回了状态码200, 内层这里用前后端约定的状态码表示即可.
 */
func (baseResponse BaseResponse) NoData(data string) string {
	response := BaseResponse{
		Code:    noData, //这里是约定的状态码.
		Message: "success",
	}
	if data != "" {
		response.Data = data
	}
	//common.TestPanic()
	marshal, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	return string(marshal)
}