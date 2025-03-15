package main

import (
	"fmt"
	"runtime"
	"testing"
)

/*
*
目的:

	用recover()来抓住底层抛上来的panic宕机错误, 和运行时异常(例如空指针).
	该机制类似try...catch.

测试结果:

	运行前
	测试1:手动宕机前
	@@@error: &{这是手动触发panic时抛向上层的结构体信息}
	测试2:赋值宕机前
	@@@runtime error: runtime error: invalid memory address or nil pointer dereference
	运行后
*/
func TestMainRecoverAndPanic(t *testing.T) {
	testRecoverAndPanic()
}

/*
*
测试用recovery在上层抓住: panic 和 运行时异常(例如空指针)
*/
func testRecoverAndPanic() {
	fmt.Println("运行前")

	// 测试1: 调用下面的ProtectRun()函数, 传过去带有panic的匿名函数, 触发那边的recover.
	// 用panic关键字触发的是 非runtime的error.
	ProtectRun(func() {
		fmt.Println("测试1:手动宕机前")
		// 使用panic传递上下文
		// 使用 panic 手动触发一个错误，并将一个结构体附带信息传递过去，此时，recover 就会获取到这个结构体信息，并打印出来。
		panic(&panicContext{
			"这是手动触发panic时抛向上层的结构体信息",
		})
		fmt.Println("测试1:手动宕机后, 本行不会被执行,因为在上一行panic就返回了.") //这行不会被执行到.
	})

	// 测试2: 调用下面的ProtectRun()函数, 传过去带有空指针的匿名函数, 触发那边的recover.
	// 这里没用panic关键字来触发error, 而是写了个空指针, 这样的error类型就会是runtime error了.
	ProtectRun(func() {
		fmt.Println("测试2:赋值宕机前")
		var a *int
		//模拟代码中空指针赋值造成的错误，此时会由 Runtime 层抛出错误，被 ProtectRun() 函数的 recover() 函数捕获到。
		*a = 1
		fmt.Println("测试2:赋值宕机后,本行不会被执行,因为在上一行发生空指针时就返回了.") //这行不会被执行到.
	})

	// 测试3: 调用下面的ProtectRun()函数, 传过去带有if判断的手动panic.
	ProtectRun(func() {
		fmt.Println("测试3:赋值宕机前")
		if 1 == 1 {
			//只是一个字符串的panic, 上面测试1是一个结构体. 大同小异了.
			panic("@@@手动panic.")
		}
		fmt.Println("测试3:赋值宕机后,本行不会被执行,因为在上一行手动panic就返回了.") //这行不会被执行到.
	})

	fmt.Println("运行后") //这行可以被执行到.
}

/*
声明描述错误的结构体，保存执行错误的函数。
崩溃时需要传递的上下文信息
*/
type panicContext struct {
	function string // 所在函数
}

/*
保护方式允许一个函数
*/
func ProtectRun(entry func()) {
	// 定义一个延迟处理的闭包函数,该函数会在外层函数执行到最后一行后开始执行.
	//用于: 当 下面执行:传入的entry()函数中的panic()触发崩溃时，本外层ProtectRun() 函数将结束运行，此时点 defer 后的本闭包将会发生调用。
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		// recover() 会获取到 panic 传入的参数(上面传入的是一个结构体)
		// 注意,不但可以抓到从panic抛上来的宕机错误, 也可以抓到类似空指针的runtime运行时错误.
		err := recover()

		//使用 switch 对 err 变量进行类型断言。
		switch err.(type) {
		//如果错误是有 Runtime 层,即运行时错误抛出的运行时错误，如空指针访问、除数为 0 等情况，打印运行时错误。
		case runtime.Error:
			fmt.Println("@@@runtime error:", err)
		default: // 非运行时错误
			fmt.Println("@@@error:", err)
		}
	}()
	//开始调用传入的函数. 该函数中可能会发生panic,导致宕机. 不过上面已经用自定义并defer的func()函数中用recover来回复这里可能发生的宕机了.
	entry()
}
