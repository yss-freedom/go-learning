package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	// 创建一个新的Cron实例，默认是支持分钟级别的调度，加上cron.WithSeconds() 支持秒级别调度
	c := cron.New(cron.WithSeconds()) //精确到秒级

	// 添加一个每3秒钟执行一次的定时任务 也可以使用预定义时间格式
	spec := "@every 3s"
	// func (c *Cron) AddFunc(spec string, cmd func()) (EntryID, error)
	// AddFunc第一个参数是一个Cron表达式，表示任务的执行时间和频率；第二个参数是一个无参的函数
	c.AddFunc(spec, func() {
		fmt.Println("Task executed every 3 seconds", time.Now())
	})

	// 启动Cron实例，开始执行定时任务
	c.Start()

	// 为了演示效果，让主程序运行一段时间
	time.Sleep(30 * time.Second)

	// 停止Cron实例（在实际应用中，通常不需要手动停止Cron实例，除非程序需要退出）
	c.Stop()
}
