package main

import "fmt"

/*
New 和 Make 的概念, 作用, 和区别.

New:
	new是一个go内置的"函数". 接收一个Type, 返回指针. 用来分配内存地址.

Make:
	make 函数只用于 map，slice 和 channel，并且不返回指针,而是返回所创建的类型本身.

###### Go语言中的 new 和 make 主要区别如下 ######：
	make 只能用来分配及初始化类型为 slice、map、chan 的数据。new 可以分配任意类型的数据；
	new 分配返回的是指针，即类型 "*Type"。make 返回引用，即 "Type"；
	new 分配的空间被清零。make 分配空间后，会进行初始化；
*/
func main1() {

	//new函数的应用
	testNew()

	//make函数的应用.
	testMake()

}

/*
what:
	new是一个go内置的"函数". 接收一个Type, 返回指针. 用来分配内存地址.

why:
	new 函数用来为Type( 系统默认的数据类型 或 自定义的Type类型), 分配空间, 返回指向该内存地址的指针.

结论:
	"new"返回的永远是类型(Type)的"指针"，指针指向分配类型的内存地址。

内置new函数的定义:
	func new(Type) *Type
	可以看出，new 函数只接受一个参数，这个参数是一个类型，并且返回一个指向该类型内存地址的指针。
	同时 new 函数会把分配的内存置为零，也就是类型的零值。
*/
func testNew() {
	/*
		使用 new 函数为变量分配内存空间。
	*/
	//先定义一个指针. ##注意, Go无法在下一行new时返回一个变量的同时还定义该变量的类型.
	//需要在上面定义其类型, 或用 := 符号来让程序推测其类型.
	var sum *int
	sum = new(int) //分配内存空间
	*sum = 98
	fmt.Println(*sum)

	/*
		new 函数不仅仅能够为系统默认的数据类型，分配空间，自定义类型也可以使用 new 函数来分配空间
	*/
	type Student struct {
		name string
		age  int
	}

	var s *Student
	s = new(Student)
	s.name = "name1"

	fmt.Println("My Student is:", s)
	fmt.Println()
}

/*
what:
	make 也是用于内存分配的，但是和 new 不同，它只用于 chan、map 以及 slice 的内存创建
	它返回的类型就是这三个类型本身，而不是他们的指针类型，
	因为这三种类型就是引用类型，所以就没有必要返回他们的指针了

why:
	用于 chan、map 以及 slice 的内存创建

内置make函数的定义:
	func make(t Type, size ...IntegerType) Type
	注意:make 函数的 t 参数必须是 chan（通道）、map（字典）、slice（切片）中的一个，并且返回值也是类型本身。

注意：
	make 函数只用于 map，slice 和 channel，并且不返回指针。
	如果想要获得一个显式的指针，可以使用 new 函数进行分配，或者显式地使用一个变量的地址。

*/
func testMake() {
	//make map
	map1 := make(map[string]int)
	fmt.Printf("map1的类型是:%T", map1) //map1的类型是:map[string]int
	fmt.Println()
	fmt.Println("map:", map1) //map: map[]

	//make slice
	slice1 := make([]int, 2)
	fmt.Printf("slice1的类型是:%T", slice1) //slice1的类型是:[]int
	fmt.Println()
	fmt.Println("切片:", slice1) //切片: [0 0]
}
