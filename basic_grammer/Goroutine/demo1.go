package main

import (
	"fmt"
	"time"
)

/*
Goroutine的创建和使用
Goroutine是由Go的运行时调度和管理的轻量级线程。每个goroutine都有自己独立的栈空间，并且由Go的运行时环境进行调度。
与线程不同，goroutine的创建和切换成本更低，因为它只是函数入口和堆栈的封装，不需要像线程那样进行复杂的上下文切换。
*/
func main() {
	//使用go关键字来创建goroutine协程
	go Hello()
	time.Sleep(1 * time.Second) // 等待一秒钟，确保Hello函数有足够的时间执行。如果不加等待，main函数如果结束了，所有的 goroutine也会瞬间销毁
}

func Hello() {
	fmt.Println("hello world")
}
