package main

/*
go 语言是追求简洁的，他没有使用 try…catch 语句
如果有些情况，必须要处理异常，就需要使用panic，抛出了异常，recover，接收这个异常来处理。
recover结合defer处理 panic 恐慌
recover必须在defer语句中调用，才能捕获到panic。defer语句会延迟函数的执行，直到包含它的函数即将返回时，才执行defer语句中的函数。
一般我们在某个函数抛出的panic，就在那个函数里面解决了
recover会返回panic的参数
*/

import (
	"fmt"
)

func main() {
	//defer语句绑定的匿名函数
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("捕获到异常:", r)
		}
	}()
	panic("发生了异常")
	fmt.Println("这行代码不会被执行")
}
