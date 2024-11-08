package main

import (
	"fmt"
	"time"
)

/*
多线程会遇到的问题
1. 临界资源的安全问题
临界资源：指并发环境中多个进程、线程、协程共享的资源

在并发编程中对临界资源的处理不当，往往会导致数据不一致的问题。
*/

func main() {
	// 此时的a就是临界资源：多个协程共享的变量，会导致程序结果未知
	a := 1

	go func() {
		a = 2
		fmt.Println("goroutine a:", a)
	}()

	a = 3
	time.Sleep(3 * time.Second)
	fmt.Println("main a:", a)
}
