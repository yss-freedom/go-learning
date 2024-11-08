package main

import (
	"fmt"
	"sync"
	"time"
)

//sync.WaitGroup类型可以用来等待一组Goroutine完成。例如：
// waitgroup

var wg sync.WaitGroup

func main() {
	// 公司最后关门的人   0
	// wg.Add(2) wg.Add(2)来告诉WaitGroup我们要等待两个Goroutine完成  开启几个协程，就add几个
	// wg.Done() 我告知我已经结束了  defer wg.Done()来在Goroutine完成时通知WaitGroup
	// 开启几个协程，就add几个
	wg.Add(2)

	go test1()
	go test2()

	fmt.Println("main等待ing")
	wg.Wait() // 等待 wg 归零，wg.Wait()来等待所有Goroutine完成 代码才会继续向下执行
	fmt.Println("end")

	// 理想状态：所有协程执行完毕之后，自动停止。
	//如果每次都强制设置个等待时间。那么协程代码也可能在这个时间内还没跑完，也可能提前就跑完了，所以设置死的等待时间不合理。此时就需要用到了等待组WaitGroup
	//time.Sleep(1 * time.Second)

}
func test1() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("test1--", i)
	}
	wg.Done() //这里就将该代码块放在了其他逻辑之后
}
func test2() {
	defer wg.Done() // defer wg.Done()来在Goroutine完成时通知WaitGroup  如果不用defer就得把该方法放在其他代码之后
	for i := 0; i < 10; i++ {
		fmt.Println("test2--", i)
	}
}

/*
避免死锁：确保在获取锁之后，无论发生什么情况（包括panic），都能够释放锁。可以使用defer语句来确保锁的释放。
减少锁的持有时间：锁的持有时间越长，其他goroutine被阻塞的时间就越长，系统的并发性能就越差。因此，应该尽量减少锁的持有时间，只在必要的代码段中持有锁。
避免嵌套锁：尽量避免在一个锁已经持有的情况下再尝试获取另一个锁，这可能会导致死锁。
避免忘记调用Done：如果忘记调用Done方法，WaitGroup将会永远等待下去，导致程序无法正常结束。
避免负数计数器：调用Add方法时，如果传入的参数为负数，或者导致计数器变为负数，将会导致panic。
*/
