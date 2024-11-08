package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

/*
除了使用闭包外，我们还可以定义一个带参数的Job类型。
这需要实现cron.Job接口，该接口包含一个Run方法
*/

// ParamJob 定义带参数的Job类型
type ParamJob struct {
	param string
}

// Run 实现cron.Job接口的Run方法
func (j *ParamJob) Run() {
	fmt.Println("ParamJob executed with parameter:", j.param, time.Now())
}

func main() {
	c := cron.New(cron.WithSeconds())

	// 创建一个ParamJob实例并设置参数
	jobParam := "Hello, Cron with ParamJob!"
	//注意，这里定义对象的时候使用指针
	paramJob := &ParamJob{param: jobParam}

	// 将ParamJob实例添加到Cron实例中
	// 注意：由于AddJob方法期望的是一个cron.Job接口，因此我们需要将ParamJob实例的指针转换为cron.Job接口
	c.AddJob("@every 2s", paramJob)

	c.Start()
	// 使用一个空的select语句来阻塞主goroutine
	select {}
}
