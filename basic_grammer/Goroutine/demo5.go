package main

import (
	"fmt"
	"sync"
	"time"
)

/*
互斥锁（sync.Mutex）是最基本的同步机制之一，用于确保同一时间只有一个goroutine能够访问特定的资源。
当一个goroutine持有互斥锁时，其他试图获取该锁的goroutine将会被阻塞，直到锁被释放
*/

// 定义全局变量 票库存为10张
var tickets int = 10

// 定义一个锁  Mutex 锁头
var mutex sync.Mutex

func main() {
	go saleTicket("张三")
	go saleTicket("李四")
	go saleTicket("王五")
	go saleTicket("赵六")

	time.Sleep(time.Second * 5)
}

// 售票函数
func saleTicket(name string) {
	for {
		// 在拿到共享资源之前先上锁
		mutex.Lock()
		if tickets > 0 {
			time.Sleep(time.Millisecond * 1)
			fmt.Println(name, "剩余票的数量为：", tickets)
			tickets--
		} else {
			// 票卖完，解锁
			mutex.Unlock()
			fmt.Println("票已售完")
			break
		}
		// 操作完毕后，解锁
		mutex.Unlock()
	}
}
