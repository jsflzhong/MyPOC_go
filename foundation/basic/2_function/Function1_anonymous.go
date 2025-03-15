package main

import "fmt"

/*
Go 语言的函数属于“一等公民”（first-class），也就是说：
	函数本身可以作为值进行传递。
	支持匿名函数和闭包（closure）。
	函数可以满足接口。

Go语言里面拥三种类型的函数：
	普通的带有名字的函数
	匿名函数或者 lambda 函数
	方法

###重点:
	在函数中，实参通过值传递的方式进行传递，因此函数的形参是实参的拷贝，对形参进行修改不会影响实参，
	但是，如果实参包括引用类型，如指针、slice(切片)、map、2_function、channel 等类型，实参可能会由于函数的间接引用被修改。

*/
func mainAnoymous() {
	NormalFunction()

	testAnonymousFunction()
}

/*
定义匿名函数的两种方式.
*/
func testAnonymousFunction() {
	//方式1:定义一个匿名函数
	func(data int) {
		fmt.Println("anonymous1:", data)
	}(100) //在这里立即调用!

	//方式2:定义匿名函数体, 并把其保存到f()中
	f := func(data int) {
		fmt.Println("anonymous2:", data)
	}
	// 使用f()调用
	f(101)
}

func NormalFunction() int {
	testMultiReturnValue()
	return 1
}

/*
函数声明包括:
	函数名、形式参数列表、返回值列表（可省略）以及函数体。

func 函数名(形式参数列表)(返回值列表){
    函数体
}

测试多返回值的函数
测试结果:1 2
*/
func testMultiReturnValue() {
	a, b := multiReturnWithoutVariableName()
	fmt.Println("multiReturnWithoutVariableName:", a, b) //1 2

	c, d := multiReturnWithVariableName()
	fmt.Println("multiReturnWithVariableName:", c, d) //1 2
}

/*
多个返回值, 但是定义返回值时没写变量名,只写了返回的两个类型.
*/
func multiReturnWithoutVariableName() (int, int) {
	return 1, 2
}

/*
多个返回值, 定义返回值时既写了返回的两个类型, 也写了变量名.
*/
func multiReturnWithVariableName() (a int, b int) {
	a = 1
	b = 2
	return //如果定义了带变量名的返回值, 则return时可以不写这俩变量名.
}
