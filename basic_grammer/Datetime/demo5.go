package main

import (
	"fmt"
	"time"
)

//4. 解析时间字符串
/*
time.Parse 和 time.ParseInLocation 函数可以将符合特定格式的字符串解析为 time.Time 对象。
Parse函数默认使用本地时区，而ParseInLocation允许指定时区。
*/

func main() {

	const layout = "2006-01-02 15:04:05" //模板时间必须是这个
	// 假设这是需要解析的时间字符串

	const timeStr = "2023-09-10 12:34:56"

	t, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println("Parsed time:", t)
}
