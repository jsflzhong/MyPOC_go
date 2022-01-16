package main

import (
	"errors"
	"fmt"
)

/*
Go语言的错误处理思想及设计包含以下特征：
一个"可能造成错误的函数"，需要返回值中返回一个"错误接口（error）"，如果调用是成功的，错误接口将返回 "nil"，否则返回错误。
在函数调用后需要检查错误，如果发生错误，则进行必要的错误处理。

Go语言没有类似 Java 或 .NET 中的异常处理机制，虽然可以使用 defer、panic、recover 模拟，但官方并不主张这样做，
Go语言的设计者认为其他语言的异常机制已被过度使用，上层逻辑需要为函数发生的异常付出太多的资源，
同时，如果函数使用者觉得错误处理很麻烦而忽略错误，那么程序将在不可预知的时刻崩溃。

Go语言希望开发者将错误处理视为正常开发"必须实现的环节"，正确地处理每一个可能发生错误的函数，
同时，Go语言使用"返回值返回错误的机制"，也能大幅降低编译器、运行时处理错误的复杂度，让开发者真正地掌握错误的处理。

例1:
	net.Dial() 是Go语言系统包 net 即中的一个函数，一般用于创建一个 Socket 连接。
	net.Dial 拥有两个返回值，即 Conn 和 error，这个函数是阻塞的，
	因此在 Socket 操作后，会返回 Conn 连接对象和 error，如果发生错误，error 会告知错误的类型，Conn 会返回空
	func Dial(network, address string) (Conn, error) {
		var d Dialer
		return d.Dial(network, address)
	}

例2:
	在 io 包中的 Writer 接口也拥有错误返回，代码如下：
	type Writer interface {
		Write(p []byte) (n int, err error) //第二个返回值的类型是:error接口(注意小写)
	}

例3:
	io 包中还有 Closer 接口，只有一个错误返回，代码如下：
	type Closer interface {
		Close() error //还是error接口
	}

"error"接口:
	这个内置接口,是 Go 系统声明的接口类型，代码如下：
	type error interface {
		Error() string
	}
	注意,Go中不叫Exception, 叫Error, 中文译为"错误".



 */
func main() {
	//使用error方式一: 自定义结构体实现error. 模仿errors包的实现原理: 自定义new方法 + 自定义结构体去实现上述error接口里的Error方法.
	selfDefinedError("@@@test")

	//使用error方式二: 使用Go语言内置的errors包来创建error.
	useDefaultError()

	//使用error方式三: 自定义结构体并含有多个字段,可以丰富错误信息. 类似上述方式一, 只是结构体中有多个字段,用来描述多种错误信息.
	testMultiFieldsSelfDefineError()

}

/*
使用error方式一: 自定义error. 模仿errors包的实现原理: 自定义new方法 + 自定义结构体去实现上述error接口里的Error方法.
 */
func selfDefinedError(text string) {
	err := New(text)
	if err != nil {
		//上面被调用的New方法,肯定会返回error,所以这里肯定会被执行.
		fmt.Println("error is not nil:", err)
		return
	}
	fmt.Println("error is nil")
}

/*
	根据传入的字符串, 返回自创建的error对象.
	注意: 方法返回的是接口:error, 类似多态: 所有实现这个接口的结构体,例如下面的errorString, 都能在这里被返回.
 */
func New(text string) error {
	//返回的是包含了被text初始化字段s的自定义结构体:errorString.
	//由于errorString这个结构体在下面实现了error接口(的Error方法),所以这里可以用多态返回.s
	return &errorString{text}
}

/*
	创建结构体: 用来描述错误字符串
 */
type errorString struct {
	s string
}

/*
	用上述结构体来"errorString",来实现error接口的Error方法. 返回发生何种错误.
	作用:可以在其他方法(例如上面的New())中,利用接口和多态返回这个结构体.
 */
func (e *errorString) Error() string {
	fmt.Println("@@@errorString's Error is running...")
	return e.s //e.s表示取结构体e的字段s.
}

/**
使用error方式二: 使用Go语言内置的errors包来创建error.
Go语言有个叫做:errors 的包,内置了对New函数的定义, 类似上面自己的实现.
Go语言的 errors 中对 New 的定义非常简单，类似上面自己的实现, 代码如下：

// 创建错误对象
func New(text string) error {
    return &errorString{text}
}
// 错误字符串
type errorString 3_struct {
    s string
}
// 返回发生何种错误
func (e *errorString) Error() string {
    return e.s
}

 */
func useDefaultError()  {
	i, err := div(1, 0)
	if err != nil {
		fmt.Println("@@@divisor is zero!")
		return
	}
	fmt.Println("Successfully call div(), result is :", i)
}

// 定义除数为0的错误
var errDivisionByZero = errors.New("division by zero")

/**
在代码中使用错误定义: 定义一个会返回error的函数
 */
func div(dividend, divisor int) (int, error) {
	// 判断除数为0的情况并返回
	if divisor == 0 {
		return 0, errDivisionByZero
	}
	// 正常计算，返回空错误
	return dividend / divisor, nil
}

/*
自定义一个含有多个字段的结构体
 */
type ParseError struct {
	Filename string // 文件名
	Line     int    // 行号
}
/*
用上述自定义结构体, 来实现error接口，返回错误不止一个错误描述, 因为ParseError这个结构体中定义了多个字段.
 */
func (e *ParseError) Error() string {
	return fmt.Sprintf("@@@selfDeifnedError(),%s:%d", e.Filename, e.Line)
}

/*
定义一个函数,可以返回不止一个错误信息
 */
func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}

func testMultiFieldsSelfDefineError()  {
	err := newParseError("@@@testMultiFieldsSelfDefineError", 1)
	fmt.Println("@@@error:",err.Error())
}




