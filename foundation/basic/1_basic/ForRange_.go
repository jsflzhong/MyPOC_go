package main

import "fmt"

/*
概念:
	for range: 键值循环

	for range 结构是Go语言特有的一种的迭代结构，
	在许多情况下都非常有用，for range 可以遍历数组、切片、字符串、map 及通道（channel）

形式:
	一般形式为：
	for key, val := range coll {
		...
	}

注意:
	val 始终为集合中对应索引的"值拷贝"，
	因此它一般只具有"只读"性质，***"对它所做的任何修改都不会影响到集合中原有的值"***。

返回值:
	数组、切片、字符串返回索引和值。
	map 返回键和值。
	通道（channel）只返回通道内的值。
*/
func main() {
	testSimpleForRange()

}

/*

 */
func testSimpleForRange() {
	/*
		1.遍历切片.
		切片其实就是多个相同类型元素的连续集合，既然切片是一个集合，那么我们就可以迭代其中的元素，
		Go语言有个特殊的关键字 range，它可以配合关键字 for 来迭代切片里的每一个元素，如下所示：

		执行结果:
		Index: 0 Value: 10
		Index: 1 Value: 20
		Index: 2 Value: 30
		Index: 3 Value: 40

		注意:
			切片是有index的.
	*/
	// 创建一个整型切片，并赋值
	slice := []int{10, 20, 30, 40}
	// 迭代每一个元素，并显示其值
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}

	/*
		2.遍历数组
		结果:
			key:0  value:1
			key:1  value:2
			key:2  value:3
			key:3  value:4
	*/
	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d  value:%d\n", key, value)
	}

	/*
		3.遍历字符串.
		结果:
			key:0 value:0x68
			key:1 value:0x65
			key:2 value:0x6c
			key:3 value:0x6c
			key:4 value:0x6f
			key:5 value:0x20
			key:6 value:0x4f60
			key:9 value:0x597d
		注意:
			代码中的变量 value，实际类型是 rune 类型，以十六进制打印出来就是字符的编码

	*/
	var str = "hello 你好"
	for key, value := range str {
		fmt.Printf("key:%d value:0x%x\n", key, value)
	}

	/*
		4.遍历map
		结果:
			hello 100
			world 200
	*/
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}

	/*
		5.遍历通道（channel）——接收通道数据
		结果:
			channel: 1
			channel: 2
			channel: 4
	*/
	//创建了一个整型类型的通道。
	c := make(chan int)
	//启动了一个 goroutine，其逻辑的实现体现在第 5～8 行，在通道中推送数据 1、2、3，然后结束并关闭通道。
	go func() {
		c <- 1
		c <- 2
		c <- 4
		close(c)
	}() //这段 goroutine 在声明结束后，在本行马上被执行
	for v := range c {
		fmt.Println("channel:", v)
	}
}
