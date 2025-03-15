package main

import (
	"fmt"
	"testing"
)

/*
*
函数类型实现接口——把函数作为接口来调用

本节将对:"用结构体实现接口"与"用函数实现接口"的过程进行对比。
*/
func TestMainFuncInterface(t *testing.T) {

	invoker := structImplementInterface()

	funcImplementInterface(invoker)
}

/*
*
实现接口的方式二: 用结构体实现接口, 可以多态调用.

用该函数体实现接口的方式, 看下面.
*/
func structImplementInterface() Invoker {
	// 声明接口变量
	var invoker Invoker
	// 实例化自定义的结构体
	s := new(Struct)
	// 将实例化的结构体赋值到接口
	//**注意: s 类型为 *Struct，由于该结构体已经实现了Invoker 接口类型，因此赋值给 invoker 时是成功的**.
	//如果把下面的"func (s *Struct) Call(p interface{}) {"方法注释掉, 即表示结构体Struct并没有实现接口Invoker,所以这一行赋值会直接红线报错.
	invoker = s
	// 使用接口变量, 可以直接调用上面已经实例化了的结构体Struct里的方法Call.
	invoker.Call("hello")
	// 当然, 不用父接口调用, 结构体自己直接调用也是可以的.
	s.Call("hello2")
	return invoker
}

/*
*
实现接口的方式二: 用函数实现接口.
上面是用结构体实现的接口. 这里是用一个函数而非结构体来直接实现接口.

##这里涉及到了一个新的类型: "函数类型". 定义方式: type FuncCaller func(interface{}). 即: 用关键字type + func

问题:

	用函数实现接口的好处是什么?
	答案:
	1.参考自己的文件:Exception.go， 只要用结构体实现了接口,就可以在其他方法中,例如返回上层接口, 以类似多态的形式,返回多种"实现类"(实现结构体)
*/
func funcImplementInterface(invoker Invoker) {
	// 调用自定义的"函数类型".
	// 用函数实现接口的方式: 将匿名函数强转为自定义的FuncCaller类型，再赋值给接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("用函数实现接口:", v)
	})
	// 使用接口, 就可以直接调用FuncCaller.Call，内部会调用函数本体,即上面的匿名函数.
	invoker.Call("hello")
}

/*
定义一个接口, 当做调用器.
*/
type Invoker interface {
	//该方法类似JAVA中的抽象方法, 需要在其他地方被实现后使用.
	//调用时会传入一个 interface{} 类型的变量，*这种类型的变量表示任意类型的值*
	Call(interface{})
}

/*
定义一个结构体类型
*/
type Struct struct {
}

/*
注意这里的写法: func(s *Struct) Call(p interface{}) {...)
作用: 用上面自定义的结构体"Struct", 来实现上面自定义的接口"Invoker"中的方法Call.

可理解为:func(用这个结构体) 实现Call方法(是这个接口里的Call方法) {..具体实现..}
*/
func (s *Struct) Call(p interface{}) {
	fmt.Println("用结构体实现接口:", p)
}

/*
上面用结构体实现接口的函数已经完事. 这里开始用函数类型实现接口的函数..

先自定义一个"函数类型".
注意写法: type (函数类型的)名字 func(参数)
写法和结构体差不多, 只是把最后的"关键字struct"替换为关键字"func(参数)"而已.
*/
type FuncCaller func(interface{})

/*
与上面同理, 用自定义的函数类型, 来实现自定义的接口Invoker中的Call方法.

注意写法, 与上面的用结构体实现的写法大同小异.
*/
func (f FuncCaller) Call(p interface{}) {
	// 调用f函数本体
	f(p)
}
