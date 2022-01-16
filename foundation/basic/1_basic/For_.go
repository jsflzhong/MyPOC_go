package main

import "fmt"


func main() {
	testSimpleFor()

	testInfiniteFor()

	testInitialFor()

	testBreak()
}

func testBreak() {
	var i int

	//无尽for循环,但在for体内有带条件的break.
	for ; ; i++ {
		if i > 10 {
			fmt.Println("i:",i)
			//结束for循环,跳转到for外面的下一行.
			break
		}
	}
	fmt.Println("break后会跳转到这一行.")


	//上面的无尽for循环代码, 等同于下面的简易写法:
	for {
		if i > 10 {
			fmt.Println("i:",i)
			//结束for循环,跳转到for外面的下一行.
			break
		}
		i ++
	}
	fmt.Println("第二个break后会跳转到这一行.")
}

/**
结果:
2
1
 */
func testInitialFor() {
	step := 2
	for ; step > 0; step-- { //第一个分号必须写,因为初始化语句被移动到了上一行,这里虽然省略了,但分号要有.
		fmt.Println(step)
	}
}

func testInfiniteFor() {
	sum := 0
	for {
		sum++
		if sum > 100 {
			break
		}
	}
}

/**
测试循环.
结果:
a: 1
a: 2
a: 3
a: 4
a: 5
a: 6
a: 7
a: 8
a: 9
*/
func testSimpleFor() {
	for a := 0; a < 10; a++ {
		fmt.Println("a:", a)
	}
}
