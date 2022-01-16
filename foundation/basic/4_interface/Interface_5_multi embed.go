package main

import "io"

/*
接口嵌套组合

在Go语言中，不仅结构体与结构体之间可以嵌套，接口与接口间也可以通过嵌套创造出新的接口。

一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。
只要接口的所有方法被实现，则这个接口中的所有嵌套接口的方法均可以被调用。



*/
func main() {
	//接口内嵌接口
	embedInterface()
}

func embedInterface() {
	// 创建自定义结构体的实例, 由于该接口体实现了interface3的两个内嵌接口:interface1和2, 所以它可以被interface3所引用.
	var wc io.WriteCloser = new(device)
	// 调用interface1的方法.
	wc.Write(nil)
	// 调用interface2的方法.
	wc.Close()

	// 也可以把实例交给interface1或2的引用.
	var writeOnly io.Writer = new(device)
	// 写入数据
	writeOnly.Write(nil)
}

/*
interface1
 */
type Writer interface {
	Write(p []byte) (n int, err error)
}

/*
interface2
*/
type Closer interface {
	Close() error
}

/*
interface3, 内部嵌套了上面的两个接口.
*/
type WriteCloser interface {
	Writer
	Closer
}

/*
声明一个结构体去实现上面接口1和接口2的函数.
实现后, 它的实例就可以被接口3所引用.
 */
type device struct {
}

// 实现io.Writer的Write()方法
func (d *device) Write(p []byte) (n int, err error) {
	return 0, nil
}
// 实现io.Closer的Close()方法
func (d *device) Close() error {
	return nil
}



