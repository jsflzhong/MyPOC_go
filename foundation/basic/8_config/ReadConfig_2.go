package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
)

//全局存放通过命令行输入的参数的map.
var arguments = make(map[string]string)

func main() {

	//1.获取环境变量(并非nohup启动传参那种)
	//在linux中, 可以配置在/etc/profile中,或$HOME/.profile中. 然后#source一下.
	//在goland中,添加到启动配置的Environment中.
	getenv := os.Getenv("db_password")
	fmt.Println("@@@获取环境变量:", getenv)

	//2.获取启动命令传参(nohup那种)
	//在linux中,添加到命令行后.例如#nohup -key=value
	//在goland中,添加到启动配置的Program arguments中.
	getInputArguments()
}

func getInputArguments() {

	//方式一:自己分割,不推荐.
	//PrintArgs1(os.Args)

	//方式二:用flag库
	useFlag2GetArguments()
}

/*
方式二:用flag库获取输入参数.(推荐)

在Linux中的输入参数的方式:
	在启动命令的后面,有4种方式，效果是相同的
	-word opt
	-word=opt
	--word opt
	--word=opt

在goLand中的输入参数的方式:
	在启动配置--Program arguments中:
	输入:-username=12
 */
func useFlag2GetArguments() {

	//API1:flag.String(),返回的值指针.
	pswd := flag.String("pswd", "default value", "just description")

	//API2:flag.StringVar(),返回的是值
	var username string
	flag.StringVar(&username, "username", "default value", "账号，默认为root")

	//[必须调用,否则取不到值] 从 arguments 中解析注册的 flag
	flag.Parse()

	fmt.Println("@@@pswd:", *pswd) //指针.
	fmt.Println("@@@username:", username)
}

/*
方式一.不推荐
	这种自己分割的方式比较麻烦. 推荐用下面的flag库的成熟方式.
 */
func PrintArgs1(args ...interface{}) {
	//测试结果:
	//@@@获取启动传参: 0  = C:\Users\michael.cui\AppData\Local\Temp\___go_build_ReadConfig_2_go.exe 类型: string
	//@@@获取启动传参: 1  = arg1=12;arg2=23 类型: string
	//注意:k不是key,是次序. 第0个参数是系统的执行文件.默认的. 第1个开始才是自己传入的,但是没有自动分割.
	//所以下面才要取第1个参数,舍弃第0个.然后做分号分割.
	for k, v := range args[0].([]string) {
		fmt.Println("@@@获取启动传参:", k, " =", v, "类型:", reflect.TypeOf(v))
	}
	argArray := args[0].([]string)
	//结果:2
	fmt.Println("数组长度:", len(argArray))
}
