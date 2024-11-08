package main

import (
	"fmt"
	"time"
)

/*
time.AfterFunc函数可以接受一个Duration类型的参数和一个函数f，
返回一个*Timer类型的指针。在创建Timer之后，等待一段时间d，
然后执行函数f。
*/

func main() {
	duration := time.Duration(1) * time.Second
	f := func() {
		fmt.Println("f has been called after 1s by time.AfterFunc")
	}
	// func AfterFunc(d Duration, f func()) *Timer
	//等待duration时间后，执行f函数
	//这里是经过1秒后。执行f函数
	timer := time.AfterFunc(duration, f)
	defer timer.Stop()
	time.Sleep(2 * time.Second)
}
