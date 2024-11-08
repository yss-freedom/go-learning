package main

import (
	"fmt"
	"time"
)

//在实际应用中，你可能需要在某个条件下提前停止Ticker定时器。
//这时，你可以使用一个额外的通道来发送停止信号：

func main() {
	ticker := time.NewTicker(5 * time.Second)
	stopChan := make(chan struct{})

	go func() {
		// 模拟一个运行一段时间的任务
		time.Sleep(15 * time.Second)
		// 发送停止信号
		stopChan <- struct{}{}
	}()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
		case <-stopChan:
			fmt.Println("Ticker stopped")
			ticker.Stop()
			return
		}
	}
}
