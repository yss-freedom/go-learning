package main

import (
	"fmt"
	"time"
)

//6.2 时区转换
/*
在处理跨时区的时间时，可能需要将时间从一个时区转换到另一个时区。
这通常涉及到创建两个时间对象，
一个代表原始时间，另一个代表转换后的时间，并指定不同的时区。
*/

func main() {
	// 原始时间（UTC）
	utcTime := time.Date(2023, 9, 10, 12, 0, 0, 0, time.UTC)
	fmt.Println("UTC Time:", utcTime)

	// 转换为北京时间
	beijingLoc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}
	beijingTime := utcTime.In(beijingLoc)
	fmt.Println("Beijing Time:", beijingTime)
}
