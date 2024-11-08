package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
非缓冲通道
上面我们讲的通道都是无缓冲通道，只能放一个数据，无缓冲的Channel也称为同步Channel。在无缓冲的Channel中，发送操作和接收操作必须同时准备就绪，否则会被阻塞。
发送和接受都是阻塞的。一次发送对应一个接收。

缓冲通道
有缓冲的Channel也称为异步Channel。它允许在缓冲区未满的情况下发送多个数据，直到缓冲区满为止。
通道带了一个缓冲区，发送的数据直到缓冲区填满为止，才会被阻塞，接收的也是，只有缓冲区清空，才会阻塞。

缓冲区通道，放入数据，不会产生死锁，它不需要等待另外的线程来拿，它可以放多个数据。
如果缓冲区满了，还没有人取，也会产生死锁。
缓冲通道可以在同一个goroutine中发送数据和接收数据
可以通过len来判断缓冲通道中的数据数量
*/
// 缓冲通道 chan，cap
func main() {

	// 非缓冲通道
	ch1 := make(chan int)
	//非缓冲通道默认的大小和容量都为0值
	fmt.Println(cap(ch1), len(ch1)) // 0 0
	//非缓冲通道只能在不同的goroutine中存放和取值，否则报死锁错误
	//ch1 <- 100
	//
	//v := <-ch1
	//fmt.Println(v)

	// 缓冲通道
	// 缓冲区通道，放入数据，不会产生死锁，它不需要等待另外的线程来拿，它可以放多个数据。
	// 如果缓冲区满了，还没有人取，也会产生死锁。
	// 缓冲通道可以在同一个goroutine中发送数据和接收数据
	ch2 := make(chan string, 5)
	fmt.Println(cap(ch2), len(ch2)) // 5 0
	ch2 <- "1"
	fmt.Println(cap(ch2), len(ch2)) // 5 1 , 可以通过len来判断缓冲通道中的数据数量
	ch2 <- "2"
	ch2 <- "3"
	fmt.Println(cap(ch2), len(ch2)) // 5 3
	ch2 <- "4"
	ch2 <- "5"
	fmt.Println(cap(ch2), len(ch2)) // 5 5
	//缓冲通道可以在同一个goroutine中发送数据和接收数据
	data := <-ch2
	ch2 <- "6"                      // 向通道中存数据，如果一直存没取，当存满，继续存时，会报死锁deadlock!
	fmt.Println("缓冲通道取出的数据:", data) //1 先进先出，根据放入数据的先后顺序取出数据

	ch3 := make(chan string, 4)
	go test9(ch3)
	fmt.Println("--------------------------")
	for s := range ch3 {
		time.Sleep(time.Second)
		fmt.Println("main中读取的数据:", s)
	}
	fmt.Println("main-end")
}

func test9(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- "test - " + strconv.Itoa(i)
		fmt.Println("子goroutine放入数据：", "test - "+strconv.Itoa(i))
	}
	close(ch)
}

/*
缓冲通道，可以定义缓冲区的数量
如果缓冲区没有满，可以继续存放，如果满了，也会阻塞等待
如果缓冲区空的，读取也会等待，如果缓冲区中有多个数据，依次按照先进先出的规则进行读取。
如果缓冲区满了，同时有两个线程在读或者写，这个时候和普通的chan一样。一进一出。
*/
