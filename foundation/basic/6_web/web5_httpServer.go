package web

/*
使用 net/http 包提供的 http.ListenAndServe() 方法，可以对指定的地址进行监听，开启一个 HTTP，服务端该方法的原型如下：
	func ListenAndServe(addr string, handler Handler) error

该方法用于在指定的 TCP 网络地址 addr 进行监听，然后调用服务端处理程序来处理传入的连接请求。

ListenAndServe 方法有两个参数，其中第一个参数 addr 即监听地址，第二个参数表示服务端处理程序，通常为空。
第二个参数为空时，意味着服务端调用 http.DefaultServeMux 进行处理，
而服务端编写的业务逻辑处理程序 http.Handle() 或 http.HandleFunc() 默认注入 http.DefaultServeMux 中

	http.Handle("/foo", fooHandler)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))

	如果想更多地控制服务端的行为，可以自定义 http.Server，代码如下所示：
	s := &http.Server{
		Addr: ":8080",
		Handler: myHandler,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

简单的server参考: we1_simple.go
 */