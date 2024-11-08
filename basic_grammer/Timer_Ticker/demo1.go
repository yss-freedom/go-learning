package main

import (
	"fmt"
	"time"
)

/*
Timer定时器是一种一次性定时器，即在未来某个时刻触发的事件只会执行一次。
Timer的结构中包含一个Time类型的管道C，主要用于事件通知。
在未到达设定时间时，管道内没有数据写入，一直处于阻塞状态；到达设定时间后，会向管道内写入一个系统时间，触发事件。
*/

func main() {
	// func NewTimer(d Duration) *Timer
	timer := time.NewTimer(2 * time.Second) // 设置超时时间2秒
	// 先打印下当前时间
	fmt.Println("当前时间:", time.Now())

	//我们可以打印下这个只读通道的值
	//timer.C就是我们在定义定时器的时候，存放的时间，等待对应的时间。现在这个就是根据当前时间加2秒

	fmt.Println("通道里面的值", <-timer.C)
}
