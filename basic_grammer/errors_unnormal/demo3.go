package main

/*
fmt.Errorf()函数比errors.New()更灵活，它允许你使用格式化字符串来创建错误。
这对于需要动态构建错误信息的场景非常有用
*/

import (
	"fmt"
)

func main() {
	err := fmt.Errorf("错误代码: %d, 错误信息: %s", 1001, "操作失败")
	if err != nil {
		fmt.Println(err.Error())
	}
}
