package main

import (
	"fmt"
	"time"
)

//6.1 时区表示
/*
在Go语言中，时区通过time.Location类型表示。可以使用time.LoadLocation函数加载一个时区，
或者使用time.FixedZone函数创建一个固定偏移量的时区
*/

func main() {
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}
	fmt.Println(loc)

	// 创建固定偏移量的时区
	beijing := time.FixedZone("Beijing Time", 8*3600)
	fmt.Println(beijing)
}
