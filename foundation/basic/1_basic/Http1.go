package main

import "net/http"

/**
启动一个"文件服务器"
然后访问:http://127.0.0.1:8080/ 即可.
 */
func main() {
	//使用 http.FileServer 文件服务器将当前目录作为根目录（/目录）的处理器，访问根目录，就会进入当前目录。
	http.Handle("/", http.FileServer(http.Dir(".")))
	//默认的 HTTP 服务侦听在本机 8080 端口
	http.ListenAndServe(":8080",nil)
}
