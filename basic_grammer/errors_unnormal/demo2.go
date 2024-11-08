package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("这是一个错误")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("错误类型%T\n", err)
		//可以使用Error()方法，将err转化为字符串
		// 注意，如果err为空，调用Error()方法会报错，所以该方法一定要在if err != nil条件中执行

		fmt.Printf("错误类型%T\n", err.Error())
	}
}
