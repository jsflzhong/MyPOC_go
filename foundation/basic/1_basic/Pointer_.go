package main

import "fmt"

/**
演示拿变量的内存地址.
*/
func main() {
	//用指针拿内存地址(而不是值)
	getAddressWithPointer()

	//用指针拿值(而不是地址)
	getValueWithPointer()

	//通过指针修改其保存的值(而不是地址)
	updateValueWithPointer()

}

func updateValueWithPointer() {
	x, y := 1, 2
	fmt.Println("swap之前的值:", x, y)
	//注意,传过去的应该是变量的地址, 而不是值. 因为对面的形参类型都是指针,也没用*, 所以接不了值.
	swap(&x, &y)
	fmt.Println("swap之后的值:", x, y)
}

/*
注意: 这种声明变量的方式!
说明参数为 a、b的类型都为 *int 指针类型!
所以在调用这个函数时,传的参数应该是地址, 是这样的: &x, &y.
*/
func swap(a, b *int) {
	//取a指针保存的"值"(不是地址,因为用*了), 赋给临时变量t(注意,此时变量t是int类型,不是指针,因为左边用*取值后才赋值给它的)
	t := *a

	// 取b指针的值, 赋给a指针指向的变量. 注意,在进行指针之间"值"的交换时, 左右都得写*号.
	*a = *b

	// 将a指针的值赋给b指针指向的变量. 注意, 左边的变量是*号时, 赋值号不能写成:=, 会报错.
	*b = t

}

/*
用&符号在变量的前面,就是取地址.
用*符号在变量的前面,就是取值.
*/
func getValueWithPointer() {
	//准备一个字符串
	var house = "123abc"

	//取该字符串的"地址", 用指针存. 注意: 此时变量ptr的类型其实为: *string的指针!
	ptr := &house
	//打印:变量ptr的类型 (是个指针,而不是普通变量)
	fmt.Printf("ptr's type: %T\n", ptr) //ptr's type: *string
	//打印:ptr里保存的目标内存地址
	fmt.Printf("address: %p\n", ptr) //address: 0xc000042240

	//对"指针"进行"取值"操作(而不是地址), 用*
	value := *ptr
	//打印:取值后,保存值的变量value的类型
	fmt.Printf("value's type: %T\n", value)
	//打印:取值后,保存值的变量value的类型
	fmt.Printf("value's value: %s\n", value)
}

/*
用&符号在变量的前面,就是取地址.
用*符号在变量的前面,就是取值.
*/
func getAddressWithPointer() {
	var cat = 1
	var dog = "banana"
	//使用 fmt.Printf 的动词%p打印 cat 和 str 变量的内存地址，指针的值是带有0x十六进制前缀的一组数据。
	fmt.Printf("测试取地址: %p %p \n", &cat, &dog) //注意在变量前加"&", 表示拿的不是变量的值,而是其内存地址.
}
