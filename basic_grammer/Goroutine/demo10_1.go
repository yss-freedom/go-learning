package main

import (
	"fmt"
	"time"
)

// 通过for range简化
// 读取chan中的数据, for 一个个取，并且会自动判断chan是否close 迭代器
// 通过for range来遍历通道，返回只有一个数据，就是每次循环读取的通道中的数据
// 关闭通道
// 告诉接收方，我不会再有其他数据发送到chan了。
func main() {
	// 在main线程中定义的通道
	ch1 := make(chan int)
	go test8(ch1)
	// 读取chan中的数据, for 一个个取，并且会自动判断chan是否close 迭代器
	//通过for range来遍历通道，返回只有一个数据，就是每次循环读取的通道中的数据

	for data := range ch1 {
		time.Sleep(time.Second)
		fmt.Println(data)
	}
	fmt.Println("end")
}

// 通道可以参数传递
func test8(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// 关闭通道，告诉接收方，不会在往ch中放入数据
	close(ch)
}
