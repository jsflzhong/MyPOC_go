package __package

/*
包的习惯用法：
	包名一般是小写的，使用一个简短且有意义的名称。
	包名一般要和所在的目录同名，也可以不同，包名中不能包含- 等特殊符号。
	包一般使用域名作为目录名称，这样能保证包名的唯一性，比如 GitHub 项目的包一般会放到GOPATH/src/github.com/userName/projectName 目录下。
	包名为 main 的包为应用程序的入口包，编译不包含 main 包的源码文件时不会得到可执行文件。
	一个文件夹下的所有源码文件只能属于同一个包，同样属于同一个包的源码文件不能放在多个文件夹下。

包的特性如下：
	一个目录下的同级文件归属一个包。
	包名可以与其目录不同名。
	包名为 main 的包为应用程序的入口包，编译源码没有 main 包时，将无法编译输出可执行的文件。

要在代码中引用其他包的内容，需要使用 import 关键字导入使用的包。具体语法如下：
import "包的路径"

注意事项：
	import 导入语句通常放在源码文件开头包声明语句的下面；
	导入的包名需要使用双引号包裹起来；
	包名是从GOPATH/src/ 后开始计算的，使用/ 进行路径分隔。

包的导入路径
包的引用路径有两种写法，分别是全路径导入和相对路径导入。
	全路径导入
		包的绝对路径就是GOROOT/src/或GOPATH/src/后面包的存放路径，如下所示：
		import "lab/test"
		import "database/sql/driver"
		import "database/sql"

	相对路径导入
		相对路径只能用于导入GOPATH 下的包，标准包的导入只能使用全路径导入。
		例如包 a 的所在路径是GOPATH/src/lab/a，包 b 的所在路径为GOPATH/src/lab/b，如果在包 b 中导入包 a ，则可以使用相对路径导入方式。示例如下：

		// 相对路径导入
		import "../a"

		当然了，也可以使用上面的全路径导入，如下所示：
		// 全路径导入
		import "lab/a"

 */