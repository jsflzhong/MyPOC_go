package main

import (
	"fmt"
	"sort"
	"sync"
)

/*
map 概念
map 是引用类型，可以使用如下方式声明：
var mapname map[keytype]valuetype

其中：
mapname 为 map 的变量名。
keytype 为键类型。
valuetype 是键对应的值类型。
提示：[keytype] 和 valuetype 之间允许有空格。

在声明的时候不需要知道 map 的长度，因为 map 是可以动态增长的，未初始化的 map 的值是 nil，使用函数 len() 可以获取 map 中 pair 的数目。
*/
func main() {
	//1.创建map
	createMap()

	//2.迭代map
	iterateMap()

	//3.迭代和排序
	iterateAndSortMap()

	//4.删除元素
	deleteElement()

	//5.线程安全map
	threadSafeMap()
}

/*
线程安全的map.
需要并发读写时，一般的做法是加锁，但这样性能并不高，
Go语言在 1.9 版本中提供了一种效率较高的并发安全的 sync.Map，
sync.Map 和 map 不同，不是以语言原生形态提供，而是在 sync 包下的特殊结构。

sync.Map 有以下特性：

无须初始化，直接声明即可。

sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，
Store 表示存储，Load 表示获取，Delete 表示删除。

使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，
Range 参数中回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false。
 */
func threadSafeMap() {
	var scene sync.Map

	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	// 从sync.Map中根据键取值
	value, ok := scene.Load("london")
	fmt.Println("value:", value)
	fmt.Println("ok? :", ok)

	// 根据键删除对应的键值对
	scene.Delete("london")

	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate,", k, v)
		return true
	})
}

func deleteElement() {
	scene := make(map[string]int)

	// 准备map数据
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	//删除指定key的键值对.
	delete(scene, "brazil")
	for k, v := range scene {
		fmt.Println(k, v)
	}

	//清空 map 中的所有元素
	/*
		Go语言中并没有为 map 提供任何清空所有元素的函数、方法，
		清空 map 的唯一办法就是重新 make 一个新的 map，
		不用担心垃圾回收的效率，Go语言中的并行垃圾回收效率比写一个清空函数要高效的多。
	*/
}

/*
遍历map并排序
*/
func iterateAndSortMap() {
	scene := make(map[string]int)
	// 准备map数据
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	// 声明一个切片保存map数据
	var sceneList []int

	// 将map数据遍历复制到切片中
	for _, v := range scene {
		sceneList = append(sceneList, v)
	}

	// 对切片进行排序
	sort.Ints(sceneList)

	fmt.Println(sceneList)
	//[4 66 960]
}

/*
遍历map.
*/
func iterateMap() {
	scene := make(map[string]int)

	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	//取key和value.
	for k, v := range scene {
		fmt.Println(k, v)
	}

	//只取key,不取value
	for k := range scene {
		fmt.Println(k)
	}
	//和上面意思一样, 这两种方式都行.
	for k, _ := range scene {
		fmt.Println(k)
	}

	//只取value, 不取key
	for _, v := range scene {
		fmt.Println(v)
	}
}

/*
	创建map和简单赋值元素的两种方式.
*/
func createMap() {
	var mapLit map[string]int
	//var mapCreated map[string]float32
	var mapAssigned map[string]int

	/*
		创建map的方式1.
		连创建再同时初始化几个键值对.
	*/
	mapLit = map[string]int{"one": 1, "two": 2}
	mapAssigned = mapLit
	mapAssigned["two"] = 3 //这里, mapLit中的value也被一并改变了!

	/*
		创建map的方式2.
		只创建, 不同时初始化.
	*/
	mapCreated := make(map[string]float32) //创建map的方式2
	mapCreated["key1"] = 4.5               //向map里添加key和value.
	mapCreated["key2xxx"] = 3.14159        //再向map里添加第二组key和value.

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
	fmt.Println("mapLit:", mapLit)         //mapLit: map[one:1 two:3]
	fmt.Println("mapCreated:", mapCreated) //mapCreated: map[key1:4.5 key2xxx:3.14159]
}
