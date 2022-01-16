package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("hello world!")

	//Array
	var a [10]int //way1
	a[0] = 1
	a[1] = 2
	var b = [...]int{1, 2, 3} //way2

	//loop
	for index, value := range a {
		fmt.Printf("index: %d, value:%d, \n", index, value)
	}
	fmt.Println("==============")
	for index, value := range b {
		fmt.Printf("index: %d, value:%d, \n", index, value)
	}

	//judge -> if
	var c = 1
	if c == 1 {
		fmt.Println("c==1")
	}
	fmt.Println("==============")

	//list1
	list1 := list.New() //方式2: var list2 list1.List
	list1.PushFront(1)
	handle := list1.PushBack(2)
	list1.InsertBefore("before", handle)
	list1.InsertAfter("after", handle)
	//打印结果: [1, before, 2, after]
	for i := list1.Front(); i != nil; i = i.Next() {
		fmt.Println("value:", i.Value)
	}
	fmt.Println("==============")

	//map
	map1 := map[string]int{"key1": 1, "key2": 2} //way1, do initialization
	map2 := make(map[string]int)                 //way2, no initialization
	map2["key1"] = 3
	map2["key2"] = 4
	fmt.Printf("map1 -1: %d \n", map1["key1"])
	fmt.Printf("map1 -2: %d \n", map1["key2"])
	fmt.Printf("map2 -1: %d \n", map2["key1"])
	fmt.Printf("map2 -2: %d \n", map2["key2"])
	for k, v := range map1 {
		fmt.Printf("key:%v, value:%d \n", k, v)
	}
	fmt.Println("==============")

	/*
		用&符号在变量的前面,就是取地址.
		用*符号在变量的前面,就是取值.
		指针变量是用来接地址的.
	*/
	//pointer and address
	var cat = "cat"
	var dog = "dog"
	//使用 fmt.Printf 的动词%p打印 cat 和 str 变量的内存地址，指针的值是带有0x十六进制前缀的一组数据。
	//注意在变量前加"&", 表示拿的不是变量的值,而是其内存地址.
	fmt.Printf("取地址: %p \n", &cat)  //%p 用来打印格式化的地址
	pointer := &dog //这里用来接&地址的pointer是个指针类型的变量, 而不是普通的变量
	//注意在变量前加"*", 表示拿的不是变量的内存地址,而是其值.
	fmt.Printf("用指针取值: %s \n", *pointer) //%s 用来打印格式化的值
	fmt.Printf("用来接地址的变量pointer的类型是指针! : %T \n", pointer)

}
