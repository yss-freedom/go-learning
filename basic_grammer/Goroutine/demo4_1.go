package main

import (
	"fmt"
	"time"
)

/*
2. 经典案例：售票问题
并发本身并不复杂，但是因为有了资源竞争的问题，就使得我们开发出好的并发程序变得复杂起来，因为会引起很多莫名其妙的问题。
如果多个goroutine在访问同一个数据资源的时候，其中一个线程修改了数据，那么这个数值就被修改了，对于其他的goroutine来讲，这个数值可能是不对的。
*/

// 定义全局变量 票库存为10张
var ticket int = 10

func main() {
	// 单线程不存在问题，多线程资源争抢就出现了问题
	//多人抢票
	go saleTickets("张三")
	go saleTickets("李四")
	go saleTickets("王五")
	go saleTickets("赵六")

	time.Sleep(time.Second * 5)
}

// 售票函数
func saleTickets(name string) {
	for {
		if ticket > 0 {
			time.Sleep(time.Millisecond * 1)
			fmt.Println(name, "剩余票的数量为：", ticket)
			//每卖出一张票，票的数量减一
			ticket--
		} else {
			fmt.Println("票已售完")
			break
		}
	}
}
