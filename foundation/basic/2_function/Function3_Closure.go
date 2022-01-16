package main

import (
	"fmt"
)

/*
闭包函数.

概念:
	Closure
	引用了外部环境变量的函数,就是闭包函数.
	函数 + 引用环境 = 闭包函数.

 */
func main() {

	//测试简单的闭包
	testSimpleClosure()

	//函数返回值为一个闭包函数. 即返回值是: func()
	testRetureClosure()

	//函数的参数是一个闭包函数.
	callFuncWithClosureParam()
}

/**
结果:
	在函数funcWithClosureParam中开始执行传入的闭包
	这是在调用funcWithClosureParam时传过去的闭包(匿名)函数中的内容.
	在函数funcWithClosureParam中结束执行传入的闭包
 */
func callFuncWithClosureParam()  {
	funcWithClosureParam(func() {
		fmt.Println("这是在调用funcWithClosureParam时传过去的闭包(匿名)函数中的内容.")
	})
}

func funcWithClosureParam(funcParam func()) {
	fmt.Println("在函数funcWithClosureParam中开始执行传入的闭包")
	funcParam()
	fmt.Println("在函数funcWithClosureParam中结束执行传入的闭包")
}

/**
用函数返回的闭包, 来执行该闭包,看传入的参数的值的变化.
执行结果:
	返回闭包后,执行闭包前: 1
	返回闭包后,执行闭包后第一次: 2
	返回闭包后,执行闭包后第二次: 3
	闭包的地址:0xc00009c018
	闭包的类型:int
	11
	0xc00009c028

注意:
	两次返回的闭包的地址是不同的.
 */
func testRetureClosure() {
	number := 1
	// 创建一个累加器, 初始值为1, 由于该函数返回值是个闭包函数,所以这里的accumulator变量是个闭包.
	accumulator := Accumulate(number) //这行执行之后, number的值并没有变量,还是1. 因为真正改变其值的是函数里面的闭包,而这个闭包在这行完后还没执行.
	fmt.Println("返回闭包后,执行闭包前:", number)
	// 累加1并打印
	fmt.Println("返回闭包后,执行闭包后第一次:", accumulator())//这里才可以执行的闭包.
	//注意,这里的结果是3! 但执行的还是不带参数的闭包! 也就是说,闭包是有"记忆效应"的.
	//就是说, "同一个闭包,如果多次执行,那么产生的值是会向下累积影响的".
	fmt.Println("返回闭包后,执行闭包后第二次:",accumulator())
	// 打印累加器闭包的函数地址
	fmt.Printf("闭包的地址:%p\n", &accumulator)
	fmt.Printf("闭包的类型:%T\n", accumulator())
	// 创建一个新的闭包函数.初始值为10
	accumulator2 := Accumulate(10)
	// 累加1并打印
	fmt.Println(accumulator2())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator2)
}

/*
	提供一个值, 每次调用函数会指定对值进行累加
	注意函数的返回值类型:func(), 也是个函数. 由于其调用了外部的变量,所以它是也是个"闭包函数".
*/
func Accumulate(value int) func() int {
	//这是一个匿名函数. 但是其引用了外部定义的变量"value",所以其也是个闭包函数.
	// 返回一个"闭包函数"(引用了外部变量"value"的函数)
	return func() int {
		// 累加
		value++
		// 返回一个累加值
		return value
	}
}

/*
结果:
	simple before closure: hello world
	simple after closure: hello dude
 */
func testSimpleClosure() {
	// 准备一个字符串
	str := "hello world"
	fmt.Println("simple before closure:", str)
	// 创建一个匿名函数
	foo := func() {
		// 匿名函数中访问str
		str = "hello dude"
	}
	// 调用匿名函数
	foo()
	fmt.Println("simple after closure:", str)
}

