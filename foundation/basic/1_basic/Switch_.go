package main

import "fmt"

/*
表达式不需要为常量，甚至不需要为整数

Go语言改进了 switch 的语法设计，case 与 case 之间是独立的代码块，不需要通过 break 语句跳出当前 case 代码块以避免执行

 */
func main() {
	testSimpelBreak()

	testMultiConstant()

	testExpression()

	testFallthrough()
}

/*
在Go语言中 case 是一个独立的代码块，执行完毕后不会像C语言那样紧接着执行下一个 case，等于是默认在每个case中都有个break.
但是为了兼容一些移植代码，依然加入了 fallthrough 关键字来实现:无默认break的功能.
 */
func testFallthrough() {
	var s = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough //这里等于是取消掉了默认的break.所以程序还可能会进入到下面的case块中.
	case s != "world":
		fmt.Println("world")
	}
}

/*
case后不仅仅只可以添加常量，还可以添加表达式.
 */
func testExpression() {
	var r int = 11

	switch {
	case r > 10 && r < 20: //还可以添加表达式.
		fmt.Println(r)
	}
}

/**
当出现多个 case 要放在一起的时候，可以在常量中用逗号分开, 等同于"或".
 */
func testMultiConstant() {
	var a = "mum"
	switch a {
	case "mum", "daddy": //可以在常量中用逗号分开, 等同于"或".
		fmt.Println("family")
	}
}

/*
Go语言改进了 switch 的语法设计，case 与 case 之间是独立的代码块，不需要通过 break 语句跳出当前 case 代码块以避免执行
 */
func testSimpelBreak() {
	var a = "hello"
	switch a {
	case "hello":
		fmt.Println(1)
		//这一行不用写break,程序也不会执行下面的default.等同于默认有个break.
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}
}
