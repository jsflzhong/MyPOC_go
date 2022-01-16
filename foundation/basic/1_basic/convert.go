package main

import (
	"fmt"
	"strconv"
)

func main() {
	//int to str
	i := 30
	str := strconv.Itoa(i)
	fmt.Printf("type:%T, value:%#v\n", str, str)

	//str to int
	str2 := "20"
	//注意这里必须加一个返回值.
	i2, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Printf("%v 转换失败",str2)
	} else {
		fmt.Printf("type:%T, value:%#v\n", i2, i2)
	}
}


//其他类型转字符串
/*
	string转成int：
		int, err := strconv.Atoi(string)
	int转成string：
		string := strconv.Itoa(int)
		//也可以:s := string(k)

	string转成int64：
		int64, err := strconv.ParseInt(string, 10, 64)
	int64转成string：
		string := strconv.FormatInt(int64,10)
*/