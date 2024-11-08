package main

import "fmt"

// 返回一个“返回类型为int的函数” ，将闭包函数作为值返回
func fibonacci() func() int {
	a, b := 1, 1 //定义初始的两个值
	return func() int {
		c := a
		a, b = b, a+b
		return c
	}
}

func main() {
	f := fibonacci()
	//计算前10个斐波那契数列
	for i := 0; i < 30; i++ {
		fmt.Println(f())
	}
}
