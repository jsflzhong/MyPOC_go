package main

import (
	"fmt"
	"io"
)

/*
一个strut可以实现多个接口


*/
func main() {
	//测试用一个struct实现2个interface.
	multiImplement()

	//多个类型可以实现相同的接口.
	//测试用一个struct实现一个接口的多个方法, 但是其中的部分方法是由该结构体的内嵌结构体实现的.
	implementMltiInterface()
}

type Socket struct {
}

/*
结构体Socket实现了io.Writer接口的Write方法.
 */
func (s *Socket) Write(p []byte) (n int, err error) {
	return 0, nil
}

/*
结构体Socket也同时实现了第二个接口: io.Closer接口的Close方法.
*/
func (s *Socket) Close() error {
	return nil
}

/*
使用io.Writer的代码, 并不知道Socket和io.Closer的存在
*/
func myWriter( writer io.Writer){
	fmt.Println("@@@myWriter is calling...")
	writer.Write( nil )
}


/*
使用io.Closer, 并不知道Socket和io.Writer的存在
 */
func myCloser( closer io.Closer) {
	fmt.Println("@@@myCloser is calling...")
	closer.Close()
}

/**
测试用一个struct实现2个interface.

结果:
	@@@myWriter is calling...
	@@@myCloser is calling...
 */
func multiImplement()  {
	socket := new(Socket)
	//由于Socket结构体分别实现了两个接口,所以下面可以传参.
	myWriter(socket)
	myCloser(socket)
}

////////////////////////////////////////////2//////////////////////////////////////////////////

/*
接口: 一个服务需要满足能够开启和写日志的功能
 */
type Service interface {
	Start()  // 开启服务
	Log(string)  // 日志输出
}

/*
struct1: 日志器, 实现了上面接口中的Log方法.
 */
type Logger struct {
}

/*
实现Service的Log()方法
 */
func (g *Logger) Log(l string) {
	fmt.Println("@@@Log......")
}

/*
struct1: 游戏服务, 实现了上面接口中的Start方法.
但是内嵌了Logger结构体,该结构体实现了Service接口的第二个方法.
所以GameService这个结构体也算是实现了Service接口的全部两个方法. 内嵌的结构体中实现的也算.
 */
type GameService struct {
	Logger  // 嵌入日志器
}

/*
实现Service的Start()方法
 */
func (g *GameService) Start() {
}

/*
多个类型可以实现相同的接口.
测试用一个struct实现一个接口的多个方法, 但是其中的部分方法是由该结构体的内嵌结构体实现的.
 */
func implementMltiInterface() {
	var s Service = new(GameService)
	s.Start()
	s.Log("@@@log...")
}