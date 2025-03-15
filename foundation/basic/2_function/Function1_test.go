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

	默认情况下，t.Log 只有在 go test -v 模式下才会显示。
	fmt.Println 在 go test 运行时，默认被吞掉，不会显示，因为 Go 运行测试时会捕获标准输出，只有失败时才可能打印。
	VS Code 运行 go test 时默认不加 -v，导致 t.Log 也不会输出。

	方案 1：VS Code 里启用 -v
		VS Code 默认不会加 -v，可以手动修改：

		打开 VS Code 设置 (Ctrl + ,)
		搜索 Go Test Flags

		添加 -v 选项：
		["-v"]
		(那个设置里有中括号了, 不要再输入中括号)

	方案2: 直接指定目标单文件的全路径名:
		go test -v -run TestJusttest D:\code\MyCode\MC_go\foundation\basic\2_function\test_test.go

	方案2: 直接指定目标单文件夹, 应该是一次性把该路径下的所有test文件全执行了:
		在根目录: go test -v -run TestJusttest ./foundation/basic/2_function

/*
1.单元（功能）测试
*/
func TestNormalFunction(t *testing.T) {
	result := NormalFunction()
	if result != 1 {
		t.Error("@@@Test failed.")
	}
}
