package main

import (
	"fmt"
	"time"
)

/*
select选择语句可以用于监听多个Channel的操作，以实现非阻塞的并发控制。
select只能用在通道中，它的语法类似于switch语句，但case分支中处理的是Channel的发送和接收操作。
读取chan数据，无论谁先放入，我们就用谁，抛弃其他的
*/

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Hello"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "World"
	}()

	// 读取chan数据，无论谁先放入，我们就用谁，抛弃其他的.
	// select 和 swtich差不多， 只是select在通道中使用，case表达式需要是一个通道结果
	//如果上面的结果还处于阻塞中，就会先执行default
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
		//default:
		//    fmt.Println("default")
	}
}

/*
select用法总结
1、每一个case必须是一个通道的操作 <-
2、所有chan操作都有要结果（通道表达式都必须会被求值）
3、如果任意的通道拿到了结果。它就会立即执行该case、其他就会被忽略
4、如果有多个case都可以运行，select是随机选取一个执行，其他的就不会执行。
5、如果存在default，执行该语句，如果不存在，阻塞等待 select 直到某个通道可以运行。
*/
/*
Channel的常见使用场景
线程间的数据共享和通信
Channel可以用于在不同Goroutine之间共享和传递数据，实现线程间的通信。
任务的并发执行和结果汇总
可以使用Channel来协调多个Goroutine并发执行任务，并将结果汇总到主Goroutine中。
*/
