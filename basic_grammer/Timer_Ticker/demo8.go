package main

import (
	"fmt"
	"time"
)

/*
select {
case t := <-ticker.C:
    // 处理Ticker定时器事件
case <-stopChan:
    // 处理停止信号
}
*/

func main() {
	ticker := time.NewTicker(1 * time.Second) // 创建一个每秒触发一次的Ticker定时器
	defer ticker.Stop()                       // 确保在main函数结束时停止定时器

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
		}
	}
}
