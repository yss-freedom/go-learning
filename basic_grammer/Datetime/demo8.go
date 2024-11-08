package main

import (
	"fmt"
	"time"
)

/*
5.2 时间间隔的计算
time.Duration 类型代表两个时间点之间经过的时间，以纳秒为单位。
可以使用 time.Time 的 Sub 方法来计算两个时间点之间的时间间隔。
*/

func main() {
	start := time.Now()
	// 假设这里有一些耗时操作
	time.Sleep(2 * time.Second)
	end := time.Now()

	duration := end.Sub(start)
	fmt.Println("Duration:", duration)

	// 转换为不同单位
	hours := duration.Hours()
	minutes := duration.Minutes()
	seconds := duration.Seconds()
	fmt.Printf("Duration in hours: %v\n", hours)
	fmt.Printf("Duration in minutes: %v\n", minutes)
	fmt.Printf("Duration in seconds: %v\n", seconds)
	//使用time.ParseDuration解析时间间隔字符串
	d, _ := time.ParseDuration("1.2h")
	fmt.Println("Parsed Duration:", d)

}
