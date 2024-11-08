package main

import (
	"fmt"
	"time"
)

/*
使用Stop方法可以停止一个Timer定时器。该方法返回一个布尔值，表示定时器是否在超时前被停止。
如果返回true，表示定时器在超时前被成功停止；如果返回false，表示定时器已经超时或已经被停止过。
*/

func main() {
	timer := time.NewTimer(2 * time.Second) // 设置超时时间2秒
	//这里，创建定时器后，里面执行了停止方法，肯定是在定时器超时前停止了，返回true
	res := timer.Stop()
	fmt.Println(res) // 输出：true
}
