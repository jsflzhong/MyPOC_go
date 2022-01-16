package main

import (
	"container/list"
	"fmt"
)

/*
在Go语言中，列表使用 container/list 包来实现，
内部的实现原理是"双链表"，**列表能够高效地进行任意位置的元素插入和删除操作。**

*/
func main() {

	//1.创建
	list1 := createList()

	//2.初始化 (返回元素句柄,便于下面的删除操作)
	element := initiateList(list1)

	//3.迭代list. (无法用range)
	iterateList(list1)

	//4.删除元素
	delteElement(list1, element)

}

func delteElement(list1 *list.List, element *list.Element) {
	//删除list中的元素
	list1.Remove(element)
	iterateList(list1)
	/*
		只有element那一个元素会被删除.
		此时打印的结果是:
		element: 1
		element: 把我插入句柄前面
		element: 把我插入句柄后面
		element: tail2
		element: tail1

		元素"2"被删除了.
	*/
}

func iterateList(list1 *list.List) {
	fmt.Println("开始打印...")
	for i := list1.Front(); i != nil; i = i.Next() {
		fmt.Println("element:", i.Value)
	}
	fmt.Println("结束打印...")
	fmt.Println("")
}

/**
API:
InsertAfter(v interface {}, mark * Element) * Element	在 mark 点之后插入元素，mark 点由其他插入函数提供
InsertBefore(v interface {}, mark * Element) *Element	在 mark 点之前插入元素，mark 点由其他插入函数提供
PushBackList(other *List)	添加 other 列表元素到尾部
PushFrontList(other *List)

元素句柄(element):
列表插入函数的返回值会提供一个 *list.Element 结构，
这个结构记录着列表元素的值以及与其他节点之间的关系等信息，
从列表中删除元素时，需要用到这个结构进行快速删除。
 */
func initiateList(l *list.List) *list.Element{
	l.PushBack("tail2")
	element := l.PushFront(2)
	l.PushFront(1)
	l.PushBack("tail1")
	//到此为止,list中的元素是: 1,2(元素句柄保存在element变量中),tail2,tail1

	//使用"element"作为"元素句柄". 在该句柄的前后插入元素s
	l.InsertAfter("把我插入句柄后面",element)
	l.InsertBefore("把我插入句柄前面",element)

	/*
	打印后,此时list中的元素是:
	element: 1
	element: 把我插入句柄前面
	element: 2
	element: 把我插入句柄后面
	element: tail2
	element: tail1
	 */

	//返回句柄, 便于之后的删除元素操作.
	return element
}

/*
创建list的2种方式
 */
func createList() *list.List {
	//方式1:通过 container/list 包的 New() 函数初始化 list
	list1 := list.New()
	fmt.Println("list1:", list1)

	//方式2:通过 var 关键字声明初始化 list
	var list2 list.List
	list2.PushFront(111)
	fmt.Println("list2:", list2)

	return list1
}
