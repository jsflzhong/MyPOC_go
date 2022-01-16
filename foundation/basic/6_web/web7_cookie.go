package web

import (
	"fmt"
	"log"
	"net/http"
	"time"
)


func Regist() {
	// myFunc 为向 url发送请求时，调用的函数
	http.HandleFunc("/cookieSet", cookieSet)
	http.HandleFunc("/cookieRead", cookieRead)

	//监听8080端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

/*
设置cookie
 */
func cookieSet(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{Name: "username", Value: "zuolan", Expires: expiration}
	cookie = http.Cookie{Name: "id", Value: "1", Expires: expiration}
	fmt.Println("@@@设置cookie的过期时间为1天, 准备写回客户端.")
	http.SetCookie(w,&cookie)
}

/*
读取cookie
 */
func cookieRead(w http.ResponseWriter, r *http.Request)  {
	cookie, err := r.Cookie("username")
	CheckErr(err)
	fmt.Println("@@@已经取到key为username的cookie的值, 是:",cookie)

	fmt.Println("@@@准备迭代取出所有的cookie键值对...")
	for _, cookie := range r.Cookies() {
		//使用fmt.Fprint(w,...)可以把字符串写回客户端
		fmt.Fprint(w,cookie.Name,":", cookie.Value,"\n")
	}
}
