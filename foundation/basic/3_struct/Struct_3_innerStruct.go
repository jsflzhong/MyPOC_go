package main

import "fmt"

/*
结构体可以包含一个或多个匿名（或内嵌）字段，即这些字段没有显式的名字，只有字段的类型是必须的，此时类型也就是字段的名字。
匿名字段本身可以是一个结构体类型，即结构体可以包含内嵌结构体。
 */
func main() {
	//用new初始化组合结构体.
	innerStruct()

	//直接用大括号初始化: "组合结构体"以及"内嵌匿名结构体"的初始化.
	innerStruct2()
}

/*
定义内层结构体
 */
type innerS struct {
	in1 int
	in2 int
}

/*
定义外层结构体
*/
type outerS struct {
	b int
	c float32
	int // anonymous field
	innerS //anonymous field
}

func innerStruct() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5 //注意,外层结构体可以直接点出来使用内层结构体的字段.
	outer.in2 = 10
	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)
	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Printf("outer2 is:", outer2)

	in1_2 := outer.innerS.in1 //注意,外层结构体也可以先点出来内层结构体,再点出来使用内层结构体的字段,以免多个内层结构体中的字段名重名.
	fmt.Println("使用外层.内层.内层字段也可以访问内层的字段:",in1_2)
}

/*
结构体:车轮
 */
type Wheel struct {
	Size int
}

/*
结构体:引擎
 */
type Engine struct {
	Power int    // 功率
	Type  string // 类型
}

/*
结构体:车.
内部组合了三个结构体, 两个是在上面(即外部)定义好的, 一个是在本结构体中内嵌的结构体(在内部现定义的,外部没有)
 */
type Car struct {
	//has a
	Wheel
	Engine
	//内嵌匿名结构体
	Dashboard struct{
		OilQuantity int
		WaterQuantity int
	}
}

/*
"组合结构体"以及"内嵌匿名结构体"的初始化.

结果:
	innerStruct2()组合和内嵌结构体的初始化,car: {{1} {2 3} {4 5}}
*/
func innerStruct2()  {
	car := Car{
		//组合结构体的初始化
		Wheel: Wheel{
			Size: 1,
		},
		//组合结构体的初始化
		Engine: Engine{
			Power: 2,
			Type:  "3",
		},
		//内嵌结构体的初始化, 需要用struct关键字,
		//由于 Dashboard 字段的类型并没有被单独定义，
		//因此在初始化其字段时需要先填写 3_struct{…} 声明其类型, 并且还要在这里再声明一次字段.
		Dashboard: struct {
			OilQuantity   int
			WaterQuantity int
		}{
			OilQuantity:   4,
			WaterQuantity: 5,
		},
	}

	fmt.Println("innerStruct2()组合和内嵌结构体的初始化,car:",car)
}









