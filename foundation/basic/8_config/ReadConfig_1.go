package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/fwhezfwhez/errorx"
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
	"time"
)

/*
解决没有第三方包的问题:
	有两种方式. 按顺序尝试:
	1.把包的import,从第三方包的引用位置,拷贝到自己的go文件中, 点击该行左边的红灯泡--ge get....
		如果该资源被墙,则会报错, 提示没有权限之类的. 那么请尝试下面第二种, 即,从github拉取.
	2.先在github找到该仓库.
		2.1.https://github.com/golang  在这里搜索下一级目录的名字. 例如:想要下载:golang.org/x/text/transform, 则在该页面搜索"text".
		2.2.点击进入该仓库,例如text仓库,此时观察右边有"clone or download"的绿色按钮.则该资源可以用git clone下载.
		2.3.新建本地目录. 例如要下载包:golang.org/x/text/transform, 则新建:{GOPATH}\src\golang.org\x\text (也可以从报错信息中观察路径)
		2.4.执行命令:#git clone https://github.com/golang/text.git D:\michael.cui\workspace_go\src\golang.org\x\text
			注意:左半部分的url,是从git页面上的绿色按钮得出的; 右半部分的本地路径,是从控制台报错信息得出的. 两部分都是有根据的.
			执行完后,会发现x\text目录下多了很多子目录(包,即一次性下载了很多子包.). 这些子目录,与git上text仓库里的子目录都是匹配的.
*/

// ~/workspace_go/src/github.com/spf13/viper/viper.go:328
var SupportedExts = []string{"json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "dotenv", "env", "ini"}

const ConfigJson  = "json"
const ConfigYaml  = "yaml"
const ConfigYml  = "yml"
const ConfigProps  = "props"
const ConfigProp  = "prop"

var config *viper.Viper
var m  sync.Mutex

func main() {

	testReadConfig()
}

/*
Init 初始化配置
 */
func init() {
	log.Println("@@@[init]Init config loader...")
	var env string
	if env = os.Getenv("ENV"); env=="" {
		env = "dev"
	}
	v := viper.New()
	v.SetConfigType(ConfigProps)
	v.SetConfigName(env)
	v.AddConfigPath("../config/")
	v.AddConfigPath("config/")
	v.AddConfigPath("src/com.jsflzhong/8_config")
	ReadConfig(v)
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	config = v
	log.Println("@@@[init]Config loader initialize done.")
}

/*
GetConfig 获取配置
 */
func GetConfig() *viper.Viper {
	return config
}

/*
读取配置
 */
func ReadConfig(v *viper.Viper) error{
	m.Lock()
	defer m.Unlock()
	err := v.ReadInConfig()
	if err != nil {
		return errorx.NewFromString("Error on parsing config file!")
	}
	return nil
}

/*
测试入口. 读取配置文件.

热加载: 热加载是生效的. 但是如果是在golang里修改配置文件,需要ctrl+s保存一下,才能触发该监听事件.

测试结果:
	@@@配置文件中,key为addr的value为: 8090
	@@@配置文件中,key为db.host的value为: localhost
	@@@开始测试热加载,准备睡眠10秒...
	Config file changed: D:\michael.cui\workspace_go\src\awesomeProject\src\com.jsflzhong\8_config\dev.properties
	Config file changed: D:\michael.cui\workspace_go\src\awesomeProject\src\com.jsflzhong\8_config\dev.properties
	Config file changed: D:\michael.cui\workspace_go\src\awesomeProject\src\com.jsflzhong\8_config\dev.properties
	@@@开始测试热加载,结束睡眠10秒,开始读取新值...
	@@@配置文件中,key为db.host的[热修改后的]value为: localhost1

Process finished with exit code 0

*/
func testReadConfig() {
	c := GetConfig()
	addr := c.GetString("addr")
	fmt.Println("@@@配置文件中,key为addr的value为:", addr)

	host := c.GetString("db.host")
	fmt.Println("@@@配置文件中,key为db.host的value为:", host)

	//注意,如果想要读取启动参数中的环境变量, 则需要用os的API.
	Env1 := os.Getenv("Env-1")
	fmt.Println("@@@环境变量中,key为Env1的value为:", Env1)
	Env2 := os.Getenv("Env-2")
	fmt.Println("@@@环境变量中,key为Env2的value为:", Env2)

	fmt.Println("@@@开始测试热加载,准备睡眠10秒...")
	time.Sleep(10 * time.Second)
	fmt.Println("@@@开始测试热加载,结束睡眠10秒,开始读取新值...")
	// 这时候去修改 dev.properties.测试热加载.
	hostHot := c.GetString("db.host")
	fmt.Println("@@@配置文件中,key为db.host的[热修改后的]value为:", hostHot)
}