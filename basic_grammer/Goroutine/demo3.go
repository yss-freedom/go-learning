package main

import (
	"fmt"
	"runtime"
)

// 并发控制
// 控制并发顺序
/*
正常情况下，main主函数里面的协程一般是先于go func()执行的，
但是我们在main函数里面加上了runtime.Gosched()，
这样main函数里面的代码执行就让出CPU时间片，让其他Goroutine先执行
*/
func main() {
	// goroutine是竞争cpu的  ，调度
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine", i)
		}
	}()

	for i := 0; i < 5; i++ {
		// gosched:礼让, 让出时间片，让其他的 goroutine 先执行
		// cpu是随机，相对来说，可以让一下，但是不一定能够成功
		// schedule
		runtime.Gosched()
		fmt.Println("main-", i)
	}
}
