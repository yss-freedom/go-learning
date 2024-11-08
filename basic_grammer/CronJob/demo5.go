package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

/*
如果我们需要在任务函数中使用参数，可以使用闭包或者定义一个带参数的函数类型。
使用闭包传递参数
闭包是一种捕获并存储其外部作用域的引用的函数。
利用闭包，我们可以轻松地将参数传递给定时任务函数。
*/
func main() {
	c := cron.New(cron.WithSeconds())

	// 定义一个带参数的外部函数 外函数带有参数，返回一个函数
	executeTask := func(param string) func() {
		return func() {
			fmt.Println("Task executed with parameter:", param)
		}
	}

	// 使用闭包传递参数
	taskParam := "Hello, Cron with Closure!"
	c.AddFunc("@every 1s", executeTask(taskParam))

	c.Start()
	// 为了让程序运行足够长的时间以观察定时任务的执行，我们使用一个空的select语句来阻塞主goroutine
	select {}
}
