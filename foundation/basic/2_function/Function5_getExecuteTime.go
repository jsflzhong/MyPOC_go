package main

import (
	"fmt"
	"time"
)

/*
获取函数的执行时间.
有两种方式.


 */
func main() {
	test1()

	test2()
}

/**
方式一
 */
func test1() {
	start := time.Now() // 获取当前时间
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}
	elapsed := time.Since(start) //获取执行时间
	fmt.Println("该函数执行完成耗时：", elapsed)
}

/**
方式二
*/
func test2() {
	start := time.Now() // 获取当前时间
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}
	elapsed := time.Now().Sub(start)//获取执行时间
	fmt.Println("该函数执行完成耗时：", elapsed)
}