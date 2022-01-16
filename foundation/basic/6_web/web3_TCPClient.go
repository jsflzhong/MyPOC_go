package web

import (
	"log"
	"net"
)

/*
Go语言可以通过 net 包中的 DialTCP 函数来建立一个 TCP 连接，并返回一个 TCPConn 类型的对象，
当连接建立时服务器端也会同时创建一个同类型的对象，此时客户端和服务器段通过各自拥有的 TCPConn 对象来进行数据交换。

一般而言，客户端通过 TCPConn 对象将请求信息发送到服务器端，读取服务器端响应的信息；
服务器端读取并解析来自客户端的请求，并返回应答信息。
这个连接会在客户端或服务端任何一端关闭之后失效，不然这连接可以一直使用。

建立连接的函数定义如下：
	func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error)
	参数说明如下：
		net 参数是 "tcp4"、"tcp6"、"tcp" 中的任意一个，分别表示 TCP(IPv4-only)、TCP(IPv6-only) 或者 TCP(IPv4,IPv6) 的任意一个；
		laddr 表示本机地址，一般设置为 nil；
		raddr 表示远程的服务地址。

*/

func TCPClinet() {
	service := "127.0.0.1:8000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Fatal(err)
	}
	//创建与server的连接
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	log.Println("@@@客户端创建连接成功")
	if err != nil {
		log.Fatal(err)
	}
	//向server发送数据
	log.Println("@@@客户端开始向server发送数据.")
	n, err := conn.Write([]byte("HEAD / HTTP/1.1\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(n)
}
