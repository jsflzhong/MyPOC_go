package main

import (
	"testing"
)

/*
要开始一个"单元测试"，需要准备一个 go 源码文件，在命名文件时文件名必须以"_test.go"结尾，
单元测试源码文件可以由多个测试用例（可以理解为函数）组成，每个测试用例的名称需要以 Test 为前缀，例如：
	func TestXxx( t *testing.T ){
		//......
	}

编写测试用例有以下几点需要注意：
	测试用例文件不会参与正常源码的编译，不会被包含到可执行文件中；
	测试用例的文件名必须以_test.go结尾；
	需要使用 import 导入 "testing" 包；
	"测试函数"的名称要以"Test或Benchmark"开头，后面可以跟任意字母组成的字符串，但"Test后面的第一个字母必须大写"，
		例如 TestAbc()，一个测试用例文件中可以包含多个测试函数；
	"单元测"试则以"(t *testing.T)"作为参数，"性能测试"以(t *testing.B)做为参数；
	测试用例文件使用"go test"命令来执行，源码中不需要 main() 函数作为入口，所有以_test.go结尾的源码文件内以Test开头的函数都会自动执行。

Go语言的 testing 包提供了三种测试方式，分别是:
	单元（功能）测试、性能（压力）测试和覆盖率测试。


打印日志：
	go test 默认 只打印失败的测试信息，不会显示 fmt.Println 的输出。
	你的 TestMainException 没有使用 t.Fail 或 t.Errorf 触发失败，因此 go test 认为测试通过，最终不会显示 fmt.Println 的输出。
	go test 运行时会捕获标准输出，除非使用 -v 参数，否则不会显示 fmt.Println 的内容。

	方案1： (加 -v)
	go test -v -run ^TestMainException$

	方案2：（使用 t.Log） (Log首字母是大写的！)
	t.Log("@@@@@@@1111111") // 使用 t.Log 代替 fmt.Println

	你的 t.Log("@@@@@@@1111111") 没有打印，可能是因为 Go 测试框架的日志默认不会在 成功的测试 中显示。
	Go 的 testing.T.Log 只有在 测试失败 或者 使用 go test -v 运行时 才会输出日志。
*/

/*
1.单元（功能）测试
*/
func TestNormalFunction(t *testing.T) {
	result := NormalFunction()
	if result != 1 {
		t.Error("@@@Test failed.")
	}
}
