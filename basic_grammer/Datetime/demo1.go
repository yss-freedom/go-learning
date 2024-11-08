package main

import (
	"fmt"
	"time"
)

//1. 时间的基本获取
/*
在Go语言中，time.Time 类型用于表示时间。我们可以通过
time.Now() 函数获取当前的时间对象，进而获取年、月、日、时、分、秒等信息。
*/

func main() {
	now := time.Now()
	fmt.Printf("current time: %v\n", now)
	//得到的是time.Time类型的结构体数据，包含 常量：日月年时分秒 周日-周六 方法：获取常量，计算。
	fmt.Printf("current time数据类型: %T\n", now)
	//打印年月日时分秒，得到的都是int类型数据

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	//Printf ： 整数补位--02如果不足两位，左侧用0补齐输出

	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
