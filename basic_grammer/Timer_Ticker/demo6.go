package main

import (
	"fmt"
	"time"
)

/*
time.After函数会返回一个*Timer类型的管道，
该管道会在经过指定时间段d后写入数据。调用这个函数相当于实现了一个定时器。
*/

func main() {
	ch := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "test"
	}()

	//使用select，哪个先到来，耗时时间短，执行哪个
	select {
	case val := <-ch:
		fmt.Printf("val is %s\n", val)
	// 这个case是两秒后执行
	// func After(d Duration) <-chan Time
	case <-time.After(2 * time.Second):
		fmt.Println("timeout!!!")
	}
}
