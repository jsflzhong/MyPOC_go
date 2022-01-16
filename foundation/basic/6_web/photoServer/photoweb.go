package photoServer

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
)

/**********
文件上传和浏览的服务器

注册了三个endpoint:
	1.注册上传图片的endpoint.
	2.注册根据id查询图片的endpoint
	3.注册列出所有已上传的图片的endpoint

注意, 每个handler相当于java中的一个controller中的handler,
但其本身不能作为网络endpoint存在, 需要在main方法中用:http.HandleFunc("/xxx", funcName) 来将其"注册"成一个endpoint,之后才可以被网络访问.
*/

/*
注册所有endpoints, 即网络端点.
只有用http.HandleFunc()这样的函数, 把我们自定义的函数注册后, 那些自定义的函数才能被网络访问到. 类似Controller中的handler.

同时, 注释了一部分可用的调用io.WriteString() 返回html字符串给FE的代码, 虽然可用, 但不建议.
应该将业务逻辑程序和表现层分离开来，各自单独处理。这时候，就需要使用网页模板技术了。
Go 标准库中的 html/template 包对网页模板有着良好的支持。

首大, 给main包中的main方法调用.
*/
const (
	//UPLOAD_DIR = "./uploadFiles"
	UPLOAD_DIR   = "D:\\michael.cui\\workspace_go\\src\\awesomeProject\\src\\com.jsflzhong\\6_web\\photoServer\\uploadFiles"
	TEMPLATE_DIR = "D:\\michael.cui\\workspace_go\\src\\awesomeProject\\src\\com.jsflzhong\\6_web\\photoServer\\views"
	STATIC_DIR = "D:\\michael.cui\\workspace_go\\src\\awesomeProject\\src\\com.jsflzhong\\6_web\\photoServer\\public"

	HTML_LIST   = "list.html"
	HTML_UPLOAD = "upload.html"

	ListDir = 0x0001
)

//模板全局缓存
var templates = make(map[string]*template.Template)

/*
全局初始化函数.

init()会在main函数之前被执行
在innit中预加载模板进内存.
否则每次都会在handler代码中读取本地的html模板文件.
*/
func init() {
	//迭代html模板目录, 之后在目录下自由添加新模板,这里即可自动加载.
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil {
		panic(err)
		return
	}
	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		//判断模板文件的扩展名. 只取.html文件.
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("@@@Pre-Loading template:", templatePath)
		//字符串转template对象.
		//template.Must() 确保了模板在不能解析成功时，一定会触发错误处理流程.
		templateHtml := template.Must(template.ParseFiles(templatePath))
		templates[templateName] = templateHtml
	}
}

/**
注册所有的endpoint.
注册时,在写明handler名时,外层包装一层用defer处理的异常处理器.
*/
func RegistEndpoints() {
	//静态文件服务
	//mux := http.NewServeMux()
	//staticDirHandler(mux, "/assets/", STATIC_DIR, 0)

	//注册上传图片的endpoint.
	http.HandleFunc("/uploads", handleError(uploadHandler))
	//注册根据id查询图片的endpoint
	http.HandleFunc("/view", handleError(viewHandler))
	//注册列出所有已上传的图片的endpoint
	http.HandleFunc("/", handleError(listHandler))

	//监听8080端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

/*
Handler1: 上传图片.
endpoint: /uploads
*/
func uploadHandler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("@@@[uploadHandler]new request, request method is:", request.Method)
	//如果是get请求, 返回一个html表单.
	if request.Method == "GET" {
		//注掉直接HTML的方式,改用html/template包下的ParseFiles()函数来读取html文件.
		//注意, 调用可以返回error的方法时, 可以与if在一行内套用.
		if err := renderHtml(responseWriter, HTML_UPLOAD, nil); err != nil {
			checkError(err)
		}
	}

	//如果是POST请求,则开始处理图片上传流程
	if request.Method == "POST" {
		//从request中读取表单用的文件
		formFile, fileHeader, err := request.FormFile("image") //name="image"
		if err != nil {
			checkError(err)
		}
		filename := fileHeader.Filename
		//注意formFile也需要关闭资源.
		defer formFile.Close()
		localFile, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			checkError(err)
		}
		defer localFile.Close()
		//注意if的条件. 交换数据也写进if条件中了.
		if _, err := io.Copy(localFile, formFile); err != nil {
			//注意http.Error的使用方式.(常用)
			http.Error(responseWriter, err.Error(),
				http.StatusInternalServerError)
			return
		}
		fmt.Println("@@@[uploadHandler], upload is done, fileName:", filename)

		//重定向到另外的endpoint. 需自定义该endpoint, 用于根据id查询图片.
		//注意, id的值,是从这里直接带过去的.
		http.Redirect(responseWriter, request, "/view?id="+filename, http.StatusFound)
	}
}

/*
Handler2: 根据id查询指定图片
endpoint: /view

net/http 包提供的这个 ServeFile() 函数可以将服务端的一个文件内容读写到 http.Response-Writer,
并返回给请求来源的 *http.Request 客户端。
*/
func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	fmt.Println("@@@[viewHandler]New request, imageId is:", imageId)
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists {
		//注意http.NotFound的使用方式.(常用)
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	fmt.Println("@@@[viewHandler]Done, imageId is:", imageId)
	//注意写文件回客户端的API.
	//net/http 包提供的这个 ServeFile() 函数可以将服务端的一个文件内容读写到 http.Response-Writer 并返回给请求来源的 *http.Request 客户端。
	http.ServeFile(w, r, imagePath)
}

/*
检查目标路径是否存在.
*/
func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

/**
handler3: 列出所有已上传的图片.
endpoint: /
*/
func listHandler(w http.ResponseWriter, r *http.Request) {
	//遍历目标目录下的所有文件
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	if err != nil {
		checkError(err)
	}
	//html模板中是用双大括号+range来迭代后端这里传过去的集合的.
	imageNameMap := make(map[string]interface{})
	imageNameArray := []string{}
	for _, fileInfo := range fileInfoArr {
		imageNameArray = append(imageNameArray, fileInfo.Name())
	}
	//注意,该map的key"imageNameArray", 需要与html模板中的"$.imageNameArray"匹配, 然后模板中才能取到该map的key的值.
	imageNameMap["imageNameArray"] = imageNameArray

	//注意, 调用可以返回error的方法时, 可以与if在一行内套用.
	if err := renderHtml(w, HTML_LIST, imageNameMap); err != nil {
		checkError(err)
	}
}

/*
DRY
从html模板缓存中,根据传入的模板name,调用缓存中的模板.
*/
func renderHtml(w http.ResponseWriter, htmlName string, dataForHtml map[string]interface{}) error {
	templates[htmlName].Execute(w, dataForHtml)
	return nil
}

/*
DRY
check error是否为空的函数.
当不为空时, 利用panic向上层抛出异常. 会被注册endpoint时的外层异常处理器safeHandler()处理.
*/
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

/*
DRY
异常处理器.

***defer + panic 合并运用的函数.

传入的参数和返回值都是一个"函数"，且都是http.HandlerFunc类型，
这种类型的函数有两个参数：http.ResponseWriter 和 *http.Request。

传入的函数: 是业务处理函数, 里面有可能抛出panic.
defer的函数: 是处理上面函数中可能抛上来的panic的函数.
返回的函数: 之所以要返回一个函数,是因为这个异常处理器,是要在注册endpoint那里调用的,而那里的第二个入参,要求是个函数.
	例如:http.HandleFunc("/uploads", handleError(uploadHandler)) //第二个入参要求是个函数. 所以本异常处理器要求返回一个函数.所以才要用闭包!(使用了外层变量的(匿名)函数)

使用了 defer 关键字搭配 recover() 方法可以处理业务函数里抛上来的panic.
倘若业务逻辑处理函数(InFunction)里边引发了 panic，
则调用 recover() 对其进行检测，若为一般性的错误，则输出 HTTP 50x 出错信息并记录日志，而程序将继续良好运行。
*/
func handleError(InFunction http.HandlerFunc) http.HandlerFunc {
	//返回闭包函数
	//之所以要用一个(匿名)闭包函数, 是因为返回值要求是一个函数,见上面函数的注释.
	return func(w http.ResponseWriter, r *http.Request) {
		//defer匿名函数, 该函数会在下面运行传入的函数后被执行兜底.
		defer func() {
			//如果InFunction发生异常,则会被这里recover住.
			if e, ok := recover().(error); ok {
				//把500写回. 或利用下面注释的代码,返回自定义错误码.
				http.Error(w, e.Error(), http.StatusInternalServerError)
				// 或者输出自定义的 50x 错误页面
				// w.WriteHeader(http.StatusInternalServerError)
				// renderHtml(w, "error", e)
				// logging
				log.Println("@@@WARN: panic in %V - %V", InFunction, e)
				log.Println("@@@Debug Stack is :", string(debug.Stack()))
			}
		}()

		//注意,本闭包函数其实相当于一层代理,在里面可以做前后绕, 或抓异常(像上面), 也可以调用被代理的函数,例如:
		// next handler, 调用被包装(被代理)的目标函数. 可以非常轻松地实现了业务与非业务之间的剥离. 这里写非业务代码. 例如调用时间计算等.
		//InFunction.ServeHTTP(w, r)

		//运行传入的函数
		InFunction(w, r)
	}
}

/*
处理静态文件的函数.
例如CSS JS等.
 */
func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(responseWriter http.ResponseWriter, reqeust *http.Request) {
		file := staticDir + reqeust.URL.Path[len(prefix)-1:]
		if (flags & ListDir) == 0 {
			if exists := isExists(file); !exists {
				http.NotFound(responseWriter, reqeust)
				return
			}
		}
		http.ServeFile(responseWriter, reqeust, file)
	})
}
