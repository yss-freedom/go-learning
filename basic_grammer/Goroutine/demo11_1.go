package main

import "fmt"

//存放值，不取值，造成死锁

// 定义通道 chan
// 这个 goroutine 希望告诉 main 线程，我还没结束。（通信）
func main() {
	// 定一个bool的通道
	var ch chan bool
	ch = make(chan bool)

	//在一个goroutine中去往通道中放入数据
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("goroutine-", i)
		}
		//time.Sleep(time.Second * 3)
		ch <- true
	}()

	// 定义好通道之后，如果没有 goroutine来使用（必须在两个及以上goroutine），那么就会产生死锁
	// deadlock!
	data := <-ch
	fmt.Println("ch data:", data)

	// 死锁的产生，没有goroutine来消耗通道（存取）
	ch2 := make(chan int)
	ch2 <- 10
}
