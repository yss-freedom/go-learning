package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

/*
Cron库允许在运行时动态调整任务的配置。
以下是一个示例代码，展示如何动态添加、删除和更新定时任务
*/

func main() {
	c := cron.New(cron.WithSeconds())

	// 添加一个每5秒钟执行一次的定时任务
	entryID, err := c.AddFunc("*/5 * * * * *", func() {
		fmt.Println("Task 1: Every 5 seconds", time.Now())
	})
	if err != nil {
		fmt.Println("Error adding task 1:", err)
		return
	}
	// 启动定时器
	c.Start()
	// 等待一段时间，以便观察任务1的执行
	time.Sleep(10 * time.Second)

	// 删除任务1
	c.Remove(entryID)

	// 添加一个每10秒钟执行一次的定时任务
	_, err = c.AddFunc("*/10 * * * * *", func() {
		fmt.Println("Task 2: Every 10 seconds", time.Now())
	})
	if err != nil {
		fmt.Println("Error adding task 2:", err)
		return
	}

	// 为了演示效果，让主程序运行一段时间
	time.Sleep(30 * time.Second)

	c.Stop()
}
