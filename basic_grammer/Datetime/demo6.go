package main

import (
	"fmt"
	"time"
)

/*
4.2 带时区的解析
如果时间字符串中包含时区信息，可以使用 time.ParseInLocation 进行解析，并指定时区。
加载时区
func LoadLocation(name string) (*Location, error)
*/

func main() {
	// 假设这是需要解析的时间字符串，包含时区信息
	const layout = "2006-01-02 15:04:05 MST"
	const timeStr = "2023-09-10 12:34:56 CST"

	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	t, err := time.ParseInLocation(layout, timeStr, loc)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println("Parsed time with location:", t)
}
