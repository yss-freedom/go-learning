package main

import (
	"fmt"
	"time"
)

/*
双向通道
channel 是用来实现 goroutine 通信的。一个写、一个读、这是双向通道，上面我们讲的都是双向通道。
单向Channel
在并发编程中，有时需要在不同的函数中对Channel进行限制，例如只允许发送或只允许接收。这时可以使用单向Channel。
*/
// 只发送的通道
func send(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("发送的值:", i)
	}
	close(ch)
}

// 只接收的通道
func receive(ch <-chan int) {
	for x := range ch {
		fmt.Println("接收到的值:", x)
	}

}

func main() {
	ch := make(chan int, 2)
	go send(ch)
	go receive(ch)
	time.Sleep(time.Second)
}
