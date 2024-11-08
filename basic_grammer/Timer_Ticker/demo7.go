package main

import (
	"fmt"
	"time"
)

/*
Ticker定时器可以周期性地不断触发时间事件，不需要额外的Reset操作。
Ticker的结构中也包含一个Time类型的管道C，
每隔固定时间段d就会向该管道发送当前的时间，根据这个管道消息来触发事件。
Ticker定时器是Go标准库time包中的一个重要组件。
它允许你每隔一定的时间间隔执行一次指定的操作。Ticker定时器在创建时会启动一个后台goroutine，
该goroutine会按照指定的时间间隔不断向一个通道（Channel）发送当前的时间值。
*/
/*
要监听Ticker定时器的事件，你可以使用range关键字或者select语句来监听Ticker定时器的通道。
每次时间间隔到达时，Ticker定时器的通道都会接收到一个当前的时间值。
for range ticker.C {
    // 在这里执行周期性任务
}
*/

func main() {
	ticker := time.NewTicker(1 * time.Second)
	//查看定时器数据类型
	fmt.Printf("定时器数据类型%T\n", ticker)
	//启动协程来监听定时器触发事件，通过time.Sleep函数来等待5秒钟，然后调用ticker.Stop()函数来停止定时器。
	//最后，输出"定时器停止"表示定时器已经成功停止。
	go func() {
		for range ticker.C {
			fmt.Println("Ticker ticked")
		}

	}()

	//执行5秒后，让定时器停止
	time.Sleep(5 * time.Second)
	ticker.Stop()
	fmt.Println("定时器停止")
}
