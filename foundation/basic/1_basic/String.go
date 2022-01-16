package main

import (
	"fmt"
	"log"
)

func main() {
	Concat("name", "pswd")
}

/*
拼接字符串. 连接字符串.
测试结果:
	@@@After concat: Concat string,param1='name',param2='pswd'
*/
func Concat(param1 string, param2 string) {
	//注意API, 不是Printf.
	result := fmt.Sprintf("Concat string,param1='%s',param2='%s'", param1, param2)
	log.Println("@@@After concat:", result)
}


