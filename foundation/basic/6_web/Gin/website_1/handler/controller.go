// 处理website站点路由请求的handler，考虑这里使用了模板，结合MVC使用的习惯，故以controller来命名。
// 当然，API也可以都放在这里处理，但为了结构清晰，不提倡这么做。
package handler

import (
	"net/http"

	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const pathPhoto = "src/com.jsflzhong/6_web/Gin/website_1/website/photo/"

/*
这里的handler函数们,相当于controller里的handler, 只不过是在main文件那边, 被注册进routerGroup的.

这里主要是开发 非REST endpoints的文件. 关于REST API的开发 ,请见api/目录下.
 */

/*
由于main那边用router.LoadHTMLGlob(..)预先加载了模板tpl文件的路径.
所以这里可以直接写tpl的文件名.

routerGroup.GET("/index.html", handler.IndexHandler)
 */
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "作品欣赏",
	})
}

/*
跳转到添加页面
routerGroup.GET("/add.html", handler.AddHandler)
 */
func AddHandler(c *gin.Context) {
	time := time.Now().Unix()
	h := md5.New()
	h.Write([]byte(strconv.FormatInt(time, 10)))
	token := hex.EncodeToString(h.Sum(nil))
	c.HTML(http.StatusOK, "add.html", gin.H{
		"Title": "添加作品",
		"token": token,
	})
}

/*
上传文件的功能.

routerGroup.POST("/postme.html", handler.PostmeHandler)
 */
func PostmeHandler(c *gin.Context) {
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	// err := c.Request.ParseMultipartForm(200000)

	// Multipart form
	form, _ := c.MultipartForm()

	//intro := c.PostForm("intro")
	files := form.File["uploadImg[]"]

	if len(files) == 0 {
		c.String(http.StatusOK, "没有上传文件！")
		return
	}

	for _, file := range files {
		outputFilePath := pathPhoto + file.Filename
		fmt.Println("@@@已取得要上传的文件的全限定名,准备上传:",outputFilePath)
		// Upload the file to specific dst.
		//log.Println(outputFilePath)
		if err := c.SaveUploadedFile(file, outputFilePath); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

	}

	// 存储到IPFS
	log.Println("IPFS")
	c.String(http.StatusOK, "上传成功！")

}
