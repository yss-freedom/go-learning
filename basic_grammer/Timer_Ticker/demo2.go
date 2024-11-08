package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second) // 设置超时时间2秒
	<-timer.C
	//经过两秒后执行下面代码
	fmt.Println("after 2s Time out!")
}
