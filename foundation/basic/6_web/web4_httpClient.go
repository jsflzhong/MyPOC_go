package web

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

/*
具体来说，我们可以通过 net/http 包里面的 Client 类提供的如下方法发起 HTTP 请求：
	func (c *Client) Get(url string) (r *Response, err error)
	func (c *Client) Post(url string, bodyType string, body io.Reader) (r *Response, err error)
	func (c *Client) PostForm(url string, data url.Values) (r *Response, err error)
	func (c *Client) Head(url string) (r *Response, err error)
	func (c *Client) Do(req *Request) (resp *Response, err error)
 */

/*
HTTP GET
	resp, err := http.Get()
	一般我们可以通过 resp.Body 获取响应实体，通过 resp.Header 获取响应头，通过 resp.StatusCode 获取响应状态码
 */
func HttpGet()  {
	resp, err := http.Get("http://c.biancheng.net")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("@@@HttpGet, get result from server:/n",string(body))
}

/*
HTTP POST
只需调用 http.Post() 方法并依次传递下面的 3 个参数即可：
	请求的目标 URL
	将要 POST 数据的资源类型（MIMEType）
	数据的比特流（[]byte 形式）
其中 &buf 为图片的资源。
 */
func HttpPost()  {
	/*resp, err := http.Post("http://c.biancheng.net/upload", "image/jpeg", &buf)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))*/
}

/*
http.PostForm() 方法实现了标准编码格式为“application/x-www-form-urlencoded”的表单提交，
下面的示例代码模拟了 HTML 表单向后台提交信息的过程：

注意：POST 请求参数需要通过 url.Values 方法进行编码和封装。
 */
func postForm()  {
	resp, err := http.PostForm("http://www.baidu.com", url.Values{"wd": {"golang"}})
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}

/*
HTTP 的 Head 请求表示只请求目标 URL 的响应头信息，不返回响应实体。
可以通过 net/http 包的 http.Head() 方法发起 Head 请求，该方法和 http.Get() 方法一样，只需要传入目标 URL 参数即可。

结果:
	Content-Type : [text/html]
	Connection : [keep-alive]
	X-Swift-Savetime : [Wed, 19 Feb 2020 10:36:38 GMT]
	Eagleid : [db934ba615826299356832830e]
	Vary : [Accept-Encoding]
	Last-Modified : [Wed, 19 Feb 2020 10:16:35 GMT]
	Etag : ["d72e-59eeb15527a35"]
	Accept-Ranges : [bytes]
	Ali-Swift-Global-Savetime : [1582108577]
	Via : [cache46.l2cn1829[0,200-0,H], cache30.l2cn1829[0,0], cache22.cn1163[0,200-0,H], cache18.cn1163[1,0]]
	Server : [Tengine]
	Content-Length : [55086]
	Date : [Wed, 19 Feb 2020 10:36:17 GMT]
	Age : [521358]
	X-Cache : [HIT TCP_MEM_HIT dirn:13:395022727]
	X-Swift-Cachetime : [31104000]
	Timing-Allow-Origin : [*]
 */
func HttpHead()  {
	resp, err := http.Head("http://c.biancheng.net")
	if err != nil {
		fmt.Println("Request Failed: ", err.Error())
		return
	}
	defer resp.Body.Close() // 打印头信息
	for key, value := range resp.Header {
		fmt.Println(key, ":", value)
	}
}

/*
在多数情况下，http.Get()、http.Post() 和 http.PostForm() 就可以满足需求，
但是如果我们发起的 HTTP 请求需要更多的自定义请求信息，比如：
	设置自定义 User-Agent，而不是默认的 Go http package；
	传递 Cookie 信息；
	发起其它方式的 HTTP 请求，比如 PUT、PATCH、DELETE 等。

此时可以通过 http.Client 类提供的 Do() 方法来实现，
使用该方法时，就不再是通过缺省的 DefaultClient 对象调用 http.Client 类中的方法了，
而是需要我们手动实例化 Client 对象并传入添加了自定义请求头信息的请求对象来发起 HTTP 请求：

http.NewRequest 方法需要传入三个参数，
	第一个是请求方法，
	第二个是目标 URL，
	第三个是请求实体，只有 POST、PUT、DELETE 之类的请求才需要设置请求实体，对于 HEAD、GET 而言，传入 nil 即可。

http.NewRequest 方法返回的第一个值就是请求对象实例 req，该实例所属的类是 http.Request，
可以调用该类上的公开方法和属性对请求对象进行自定义配置，比如请求方法、URL、请求头等。

设置完成后，就可以将请求对象传入 client.Do() 方法发起 HTTP 请求，
之后的操作和前面四个基本方法一样，
http.Post、http.PostForm、http.Head、http.NewRequest 方法的底层实现及返回值和 http.Get 方法一样。
 */
func HttpDo()  {
	// 初始化客户端请求对象
	req, err := http.NewRequest("GET", "http://c.biancheng.net", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 添加自定义请求头
	req.Header.Add("Custom-Header", "Custom-Value")
	// 其它请求头配置
	client := &http.Client{
		// 设置客户端属性
	}
	//发起请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	//利用os,把response打印到控制台.
	io.Copy(os.Stdout, resp.Body)
}

/*
自定义httpClient
前面我们使用的 http.Get()、http.Post()、http.PostForm() 和 http.Head() 方法其实都是在 http.DefaultClient 的基础上进行调用的，
比如 http.Get() 等价于 http.Default-Client.Get()，依次类推。

http.DefaultClient 在字面上就向我们传达了一个信息，既然存在默认的 Client，那么 HTTP Client 大概是可以自定义的。
实际上确实如此，在 net/http 包中，的确提供了 Client 类型。让我们来看一看 http.Client 类型的结构：

	type Client struct {
		// Transport 用于确定HTTP请求的创建机制。
		// 如果为空，将会使用DefaultTransport
		Transport RoundTripper
		// CheckRedirect定义重定向策略。
		// 如果CheckRedirect不为空，客户端将在跟踪HTTP重定向前调用该函数。
		// 两个参数req和via分别为即将发起的请求和已经发起的所有请求，最早的
		// 已发起请求在最前面。
		// 如果CheckRedirect返回错误，客户端将直接返回错误，不会再发起该请求。
		// 如果CheckRedirect为空，Client将采用一种确认策略，将在10个连续
		// 请求后终止
		CheckRedirect func(req *Request, via []*Request) error
		// 如果Jar为空，Cookie将不会在请求中发送，并会
		// 在响应中被忽略
		Jar CookieJar
	}

其中 Transport 类型必须实现 http.RoundTripper 接口。
Transport 指定了执行一个 HTTP 请求的运行机制，倘若不指定具体的 Transport，默认会使用 http.DefaultTransport，
这意味着 http.Transport 也是可以自定义的。net/http 包中的 http.Transport 类型实现了 http.RoundTripper 接口。

CheckRedirect 函数指定处理重定向的策略。
当使用 HTTP Client 的 Get() 或者是 Head() 方法发送 HTTP 请求时，
若响应返回的状态码为 30x （比如 301 / 302 / 303 / 307），HTTP Client 会在遵循跳转规则之前先调用这个 CheckRedirect 函数。

Jar 可用于在 HTTP Client 中设定 Cookie，Jar 的类型必须实现了 http.CookieJar 接口，
该接口预定义了 SetCookies() 和 Cookies() 两个方法。
如果 HTTP Client 中没有设定 Jar，Cookie 将被忽略而不会发送到客户端。
实际上，我们一般都用 http.SetCookie() 方法来设定 Cookie。

使用自定义的 http.Client 及其 Do() 方法，我们可以非常灵活地控制 HTTP 请求，
比如发送自定义 HTTP Header 或是改写重定向策略等。创建自定义的 HTTP Client 非常简单
 */
func SelfDefineHttpClient()  {
	/*client := &http.Client {
	//		CheckRedirect: redirectPolicyFunc,
	//	}
	//	resp, err := client.Get("http://example.com")
	//	// ...
	//	req, err := http.NewRequest("GET", "http://example.com", nil)
	//	// ...
	//	req.Header.Add("User-Agent", "Our Custom User-Agent")
	//	req.Header.Add("If-None-Match", `W/"TheFileEtag"`)
	//	resp, err := client.Do(req)
	//	// ...*/
}