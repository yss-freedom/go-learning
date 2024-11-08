package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
Goroutine的终止和清理
使用runtime.Goexit()终止Goroutine
runtime.Goexit()函数用于终止当前Goroutine的执行。当调用Goexit()时，当前Goroutine会立即停止执行，并把控制权交还给调度器。
这意味着当前Goroutine不会继续执行后续的代码，也不会返回到调用它的地方。同时，其他仍在运行的Goroutine将继续执行。例如
*/

func main() {
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			//return // 只是终止了函数

			//这里设置，终止当前的 goroutine，该Goroutine中下面的代码不再执行
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	time.Sleep(1 * time.Second)
}
