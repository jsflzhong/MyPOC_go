package main

import (
	"fmt"
	"time"
)

/*
UTC 标准时间是以 GMT（Greenwich Mean Time，格林尼治时间）这个时区为主，所以本地时间与 UTC 时间的时差就是本地时间与 GMT 时间的时差。
UTC + 时区差 ＝ 本地时间

国内一般使用的是北京时间，与 UTC 的时间关系如下：
UTC + 8 个小时 = 北京时间

在Go语言的 time 包里面有两个时区变量，如下：
time.UTC：UTC 时间
time.Local：本地时间
*/

func main() {
	//获取简单的当前时间, 年月日等.
	getSimpleTime()

	//获取"时间戳": 从1970 年 1 月 1 日（08:00:00GMT）至当前时间的总毫秒数
	getTimeStamp()

	//时间戳与时间之间的转换
	transferTimeStamp2Time()

	//时间加法
	addTime()

	//格式化时间
	formatTime()

	//字符转时间
	string2Time()
}

/*
Parse 函数可以解析一个格式化的时间字符串并返回它代表的时间。
func Parse(layout, value string) (Time, error)

与 Parse 函数类似的还有 ParseInLocation 函数。
func ParseInLocation(layout, value string, loc *Location) (Time, error)

ParseInLocation 与 Parse 函数类似，但有两个重要的不同之处：
第一，当缺少时区信息时，Parse 将时间解释为 UTC 时间，而 ParseInLocation 将返回值的 Location 设置为 loc；
第二，当时间字符串提供了时区偏移量信息时，Parse 会尝试去匹配本地时区，而 ParseInLocation 会去匹配 loc。

结果:
	2019-12-12 15:22:12 +0000 UTC
	2019-12-12 15:22:12 +0800 CST
*/
func string2Time() {
	var layout string = "2006-01-02 15:04:05"
	var timeStr string = "2019-12-12 15:22:12"
	timeObj1, _ := time.Parse(layout, timeStr)
	fmt.Println(timeObj1) //2019-12-12 15:22:12 +0000 UTC
	timeObj2, _ := time.ParseInLocation(layout, timeStr, time.Local)
	fmt.Println(timeObj2) //2019-12-12 15:22:12 +0800 CST
}

/*
时间类型有一个自带的 Format 方法进行格式化，需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S
而是使用Go语言的诞生时间 2006 年 1 月 2 号 15 点 04 分 05 秒。

结果:
	2020-02-23 17:47:14.571 Sun Feb
	2020-02-23 05:47:14.571 PM Sun Feb
	2020/02/23 17:47
	17:47 2020/02/23
	2020/02/23
*/
func formatTime() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan")) //2022-01-14 10:14:25.957 Fri Jan
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan")) //2022-01-14 10:14:25.957 AM Fri Jan
	fmt.Println(now.Format("2006/01/02 15:04"))                   //2022/01/14 10:14
	fmt.Println(now.Format("15:04 2006/01/02"))                   //10:14 2022/01/14
	fmt.Println(now.Format("2006/01/02"))                         //2022/01/14
}

/*
Sub
	求两个时间之间的差值：
	func (t Time) Sub(u Time) Duration

Equal
	判断两个时间是否相同：
	func (t Time) Equal(u Time) bool

Before
	判断一个时间点是否在另一个时间点之前：
	func (t Time) Before(u Time) bool

After
判断一个时间点是否在另一个时间点之后：
func (t Time) After(u Time) bool
*/
func addTime() {
	now := time.Now()
	hourLater := now.Add(time.Hour)
	fmt.Println("一小时之后:", hourLater)
}

func transferTimeStamp2Time() {
	now := time.Now()                                                                   //获取当前时间
	timestamp := now.Unix()                                                             //时间戳
	fmt.Println("当前时间戳:", timestamp)                                                    //当前时间戳: 1642126465
	timeObj := time.Unix(timestamp, 0)                                                  //将时间戳转为时间格式
	fmt.Println("当前时间戳转时间:", timeObj)                                                   //当前时间戳转时间: 2022-01-14 10:14:25 +0800 CST
	year := timeObj.Year()                                                              //年
	month := timeObj.Month()                                                            //月
	day := timeObj.Day()                                                                //日
	hour := timeObj.Hour()                                                              //小时
	minute := timeObj.Minute()                                                          //分钟
	second := timeObj.Second()                                                          //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second) //2022-01-14 10:14:25
}

/*
时间戳是自 1970 年 1 月 1 日（08:00:00GMT）至当前时间的总毫秒数，它也被称为 Unix 时间戳（UnixTimestamp）。
*/
func getTimeStamp() {
	now := time.Now()                       //获取当前时间
	timestamp1 := now.Unix()                //时间戳
	timestamp2 := now.UnixNano()            //纳秒时间戳
	fmt.Printf("现在的时间戳：%v\n", timestamp1)   //当前时间戳: 1642126465
	fmt.Printf("现在的纳秒时间戳：%v\n", timestamp2) //现在的纳秒时间戳：1642126465956779700
	fmt.Println()
}

func getSimpleTime() {
	now := time.Now() //获取当前时间
	//current time:2022-01-14 10:14:25.9294167 +0800 CST m=+0.002209801
	fmt.Printf("current time:%v\n", now)
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	//@@@getSimpleTime: 2022-01-14 10:14:25
	fmt.Printf("@@@getSimpleTime: %d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	weekday := now.Weekday().String()
	//当前星期几: Friday
	fmt.Println("当前星期几:", weekday)

	fmt.Println()
}
