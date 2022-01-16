package main

import (
	"fmt"
	"os"
	"sync"
)

/*
defer
概念:
	defer 语句会将其后面跟随的语句进行延迟处理，
	在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 的逆序进行执行，(在函数的最后一行后执行)
	也就是说，先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行。

作用:
	关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源，
	典型的例子就是对一个互斥解锁，或者关闭一个文件。
*/
func main() {
	//Use defer to release lock
	//releaseLock("a")

	releaseFileHandle("D:\\install.log")


}

var (
	// 一个演示用的映射
	valueByKey = make(map[string]int)
	// 保证使用映射时的并发安全的互斥锁
	valueByKeyGuard sync.Mutex
)

/*
	演示释放锁的函数.
 */
func releaseLock(key string) int {
	valueByKeyGuard.Lock()

	// defer后面的语句不会马上调用, 而是延迟到函数结束时调用
	defer valueByKeyGuard.Unlock()

	//在这行执行完后, 上面的defer语句会自动执行,用来释放资源或锁.
	return valueByKey[key]
}

/*
	读取文件操作中, 用defer释放文件句柄. 写一次释放, 多处return后会触发.
 */
func releaseFileHandle(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		return 0
	}
	// 延迟调用Close, 此时Close不会被调用
	defer f.Close()
	info, err := f.Stat()
	fmt.Println("info:",info,"error:",err)
	if err != nil {
		// defer机制触发, 调用Close关闭文件
		return 0
	}
	size := info.Size()
	fmt.Println("size:",size) //单位是byte
	// defer机制触发, 调用Close关闭文件
	return size
}
