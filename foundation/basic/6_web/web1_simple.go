package web

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*func main() {
	//创建一个8080端口的简单服务,返回字符串
	//test1_simleServer()

	//创建一个8080端口的简单服务,返回html
	test2_readHtml()
}*/

func test1_simleServer() {
	// myFunc 为向 url发送请求时，调用的函数
	http.HandleFunc("/", simleServer)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func simleServer(w http.ResponseWriter, r *http.Request)  {
	//返回给FE, 并显示在浏览器上.
	fmt.Fprint(w, "@@@")
}

func test2_readHtml()  {
	http.HandleFunc("/", readHtml)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

//todo path error
func readHtml(w http.ResponseWriter, r *http.Request) {
	content, _ := ioutil.ReadFile("index.html")
	w.Write(content)
}