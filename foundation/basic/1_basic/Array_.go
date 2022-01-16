package main

import "fmt"

func main() {
	var a [3]int             // 定义三个整数的数组
	fmt.Println(a[0])        // 打印第一个元素
	fmt.Println(a[len(a)-1]) // 打印最后一个元素

	// 打印索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	// 仅打印元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	//默认情况下，数组的每个元素都会被初始化为元素类型对应的零值，对于数字类型来说就是 0，同时也可以使用数组字面值语法，用一组值来初始化数组：
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // "0"

	//var q [3]int = [3]int{1, 2, 3}
	//在数组的定义中，如果在数组长度的位置出现“...”省略号，则表示数组的长度是根据初始化值的个数来计算，因此，上面数组 q 的定义可以简化为：
	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"

	//遍历数组也和遍历切片类似，代码如下所示：
	var team [3]string
	team[0] = "hammer"
	team[1] = "soldier"
	team[2] = "mum"
	for k, v := range team { //注意关键字range的应用
		fmt.Println(k, v)
	}
}