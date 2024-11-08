package main

import (
	"fmt"
	"time"
)

/*
1. 定时打印日志
Ticker定时器的一个常见应用是定时打印日志。
通过设置一个Ticker定时器，你可以每隔固定的时间间隔输出一次日志信息，
从而监控程序的运行状态。
*/

func main() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop() // 确保程序结束时停止Ticker定时器

	for range ticker.C {
		// 打印当前时间作为日志
		fmt.Println("Current time:", time.Now())
	}
}
