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
	"测试函数"的名称要以"Test或Benchmark"开头，后面可以跟任意字母组成的字符串，但"第一个字母必须大写"，
		例如 TestAbc()，一个测试用例文件中可以包含多个测试函数；
	"单元测"试则以"(t *testing.T)"作为参数，"性能测试"以(t *testing.B)做为参数；
	测试用例文件使用"go test"命令来执行，源码中不需要 main() 函数作为入口，所有以_test.go结尾的源码文件内以Test开头的函数都会自动执行。

Go语言的 testing 包提供了三种测试方式，分别是:
	单元（功能）测试、性能（压力）测试和覆盖率测试。
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
