package main

import (
	"fmt"
	"time"
)

/*
当你不再需要Ticker定时器时，你应该调用它的Stop方法来停止它。
停止Ticker定时器可以释放与之关联的资源，并防止不必要的goroutine继续运行。
ticker.Stop()
停止Ticker定时器后，它的通道将不再接收任何事件。
*/

func main() {
	ticker := time.NewTicker(500 * time.Millisecond) // 创建一个每500毫秒触发一次的Ticker定时器
	timeEnd := make(chan bool)                       // 用于停止Ticker定时器的通道

	go func() {
		for {
			select {
			//当达到设置的停止条件式，停止循环
			case <-timeEnd:
				fmt.Println("===结束任务")
				break
			case t := <-ticker.C:
				fmt.Println("==500毫秒响应一次:", t)
			}
		}
	}()

	time.Sleep(5 * time.Second) // 主线程等待5秒钟
	ticker.Stop()               // 停止Ticker定时器
	timeEnd <- true             // 发送结束信号

	fmt.Println("===定时任务结束===")
}
