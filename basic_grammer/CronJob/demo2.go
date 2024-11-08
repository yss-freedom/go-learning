package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

/*
Cron库允许使用自定义的Job类型，
实现更加灵活的任务调度。以下是一个示例代码，展示如何使用自定义的Job类型：
*/

// MyJob 定义一个自定义的Job类型
type MyJob struct {
	// 可以根据需要添加其他字段
}

// Run 实现cron.Job接口中的Run方法
func (j *MyJob) Run() {
	fmt.Println("MyJob is running", time.Now())
}

func main() {
	c := cron.New(cron.WithSeconds())

	// 创建一个自定义Job的实例
	myJob := &MyJob{}

	// 添加自定义Job到Cron实例中。这里使用AddJob方法
	_, err := c.AddJob("*/5 * * * * *", myJob)
	if err != nil {
		fmt.Println("Error adding job:", err)
		return
	}

	c.Start()
	time.Sleep(30 * time.Second)
	c.Stop()
}
