package main

import (
	"fmt"
	"time"
)

/*
5. 时间的计算
time 包提供了丰富的函数和方法来支持时间的计算，如时间的加减、时间间隔的计算等。
*/
//5.1 时间的加减
//使用 time.Time 的 Add 方法可以在原有时间的基础上加上或减去一定的时间间隔。

func main() {
	now := time.Now()
	// 加上一小时
	oneHourLater := now.Add(time.Hour)
	fmt.Println("One hour later:", oneHourLater)

	// 减去一天
	oneDayAgo := now.Add(-24 * time.Hour)
	fmt.Println("One day ago:", oneDayAgo)
}
