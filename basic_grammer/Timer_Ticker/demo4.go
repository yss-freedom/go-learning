package main

import (
	"fmt"
	"time"
)

/*
对于已经过期或者是已经停止的Timer，可以通过Reset方法重新激活它，并设置新的超时时间。
Reset方法也返回一个布尔值，表示定时器是否在重置前已经停止或过期。
*/

func main() {
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("time out1")
	//经过两秒后，定时器超时了
	res1 := timer.Stop()
	//此时再stop，得到的是false
	fmt.Printf("res1 is %t\n", res1) // 输出：false
	//然后我们重置定时器。重置成3秒后超时
	timer.Reset(3 * time.Second)
	res2 := timer.Stop()
	//此时再stop，由于定时器没超时，得到的是true
	fmt.Printf("res2 is %t\n", res2) // 输出：true
}
