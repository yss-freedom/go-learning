package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
2. 周期性检查系统状态
Ticker定时器还可以用于周期性检查系统状态。
例如，你可以每隔一段时间检查一次服务器的负载、内存使用情况或数据库连接数等关键指标，
并在发现异常时采取相应的措施。
*/
// 模拟检查系统状态的函数
func checkSystemStatus() {
	// 这里可以添加实际的检查逻辑
	// 例如：检查CPU使用率、内存使用情况等
	// 这里我们随机生成一个0到100之间的数作为模拟结果
	status := rand.Intn(101)
	fmt.Printf("System status: %d\n", status)

	// 假设状态大于80表示系统异常
	if status > 80 {
		fmt.Println("Warning: System status is above normal!")
		// 这里可以添加处理异常的逻辑
		// 例如：发送警报、重启服务等
	}
}

func main() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop() // 确保程序结束时停止Ticker定时器

	for range ticker.C {
		checkSystemStatus()
	}
}
