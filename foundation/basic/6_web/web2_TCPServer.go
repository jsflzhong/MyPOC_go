package web

import (
	"fmt"
	"log"
	"net"
	"time"
)

/*
net 包中有相应功能的函数，函数定义如下：
	func ListenTCP(net string, laddr *TCPAddr) (l *TCPListener, err os.Error)
	func (l *TCPListener) Accept() (c Conn, err os.Error)

ListenTCP 函数会在本地 TCP 地址 laddr 上声明并返回一个 *TCPListener，net 参数必须是 "tcp"、"tcp4"、"tcp6"，
如果 laddr 的端口字段为 0，函数将选择一个当前可用的端口，可以用 Listener 的 Addr 方法获得该端口。
*/

/*

*/
func echo(conn *net.TCPConn) {
	tick := time.Tick(5 * time.Second) // 五秒的心跳间隔
	for now := range tick {
		//利用连接conn, 可以在server端,向client端写数据.
		n, err := conn.Write([]byte(now.String()))
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}
		fmt.Printf("@@@send %d bytes to %s\n", n, conn.RemoteAddr())
	}
}

/*
关键函数:
address := net.TCPAddr()
listener, err := net.ListenTCP()
conn, err := listener.AcceptTCP()
n, err := conn.Write
 */
func BuildTCPServer() {
	//用net创建address
	address := net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"), // 把字符串IP地址转换为net.IP类型
		Port: 8000,
	}
	log.Printf("@@@创建address成功,address的type是:%T\n", address)

	// 用net创建TCP4服务器端监听器
	listener, err := net.ListenTCP("tcp4", &address)
	if err != nil {
		log.Fatal(err) // Println + os.Exit(1)
	}
	log.Printf("@@@创建TCP4服务器端监听器成功,listener的type是:%T\n", listener)

	// 在无限循环中,接收TCP请求.
	log.Printf("@@@开始在循环中接收TCP请求.")
	for {
		//用listener接收请求, 返回conn连接
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err) // 错误直接退出
		}
		fmt.Println("@@@服务端接收到客户端的请求,remote address:", conn.RemoteAddr())

		go echo(conn)
	}
}
