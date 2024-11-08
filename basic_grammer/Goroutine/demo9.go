package main

import "fmt"

func main() {
	//在主Goroutine中定义通道
	ch := make(chan int)
	go func() {
		ch <- 42 // 在另一个Goroutine中发送数据
	}()

	value := <-ch      // 在主Goroutine中 从Channel中接收数据
	fmt.Println(value) // 输出42

	/*
		关闭后的Channel仍然可以从其中接收数据，但不能再向其发送数据。
		如果向一个已关闭的Channel发送数据，会引发panic。
		从这个关闭的channel中不但可以读取出已发送的数据，还可以不断的读取零值:
	*/
	//c := make(chan int, 10)
	//c <- 1
	//c <- 2
	//close(c)
	//fmt.Println(<-c) //1
	//fmt.Println(<-c) //2
	//fmt.Println(<-c) //0
	//fmt.Println(<-c) //0
	//c := make(chan int, 10)
	//c <- 1
	//c <- 2
	//close(c)
	//for i := range c {
	//	fmt.Println(i)
	//}
	/*
		通过i, ok := <-c可以查看Channel的状态，判断值是零值还是正常读取的值。
		ok 判断chan的状态是否是关闭，如果是关闭，不会再取值了。
		ok, 如果是true，就代表我们还在读数据
		ok, 如果是fasle，就说明该通道已关闭
	*/
	c := make(chan int, 10)
	close(c)
	i, ok := <-c
	fmt.Printf("%d, %t", i, ok) //0, false

}
