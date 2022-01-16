package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
)

/*
处理panic.

注意,上层不能在最后一行不用defer来调用本方法, 因为会被异常部分截断,从而调用不到.即: 必须用defer.

注意: todo 暂时无法解决在defer中返回response给客户端的问题.
 */
func RecoverPanic(c *gin.Context) bool{
	// 发生宕机时，获取panic传递的上下文并打印
	// recover() 会获取到 panic 传入的参数(上面传入的是一个结构体)
	// 注意,不但可以抓到从panic抛上来的宕机错误, 也可以抓到类似空指针的runtime运行时错误.
	err := recover()
	//使用 switch 对 err 变量进行类型断言。
	switch err.(type) {
	//如果错误是有 Runtime 层,即运行时错误抛出的运行时错误，如空指针访问、除数为 0 等情况，打印运行时错误。
	case runtime.Error:
		fmt.Println("@@@[recoverPanic]runtime error:", err)
	default: // 非运行时错误
		fmt.Println("@@@[recoverPanic]error:", err)
	}
	return true
}

func TestPanic() {
	if 1==1 {
		panic("@@@[TestPanic], just for causing panic.")
	}
}
