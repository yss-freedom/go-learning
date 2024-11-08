package main

import "fmt"

/*
2. 自定义约束
由于约束有时候很多，我们可以定义一些自己的泛型约束（本质是一个接口）
除了内置约束外，还可以使用自定义接口作为约束条件。例如，定义一个支持加法操作的泛型函数：
*/

// MyInt 泛型的约束提取定义
type MyInt interface {
	int | float64 | int8 | int32 // 作用域泛型的，而不是一个接口方法  这样泛型约束可以用着集中类型都可以
}

// 自定义泛型
func main() {
	var a int = 10
	var b int = 20

	fmt.Println(GetMaxNum(a, b))
}

// GetMaxNum 函数里面泛型T 用自己定义的约束MyInt
func GetMaxNum[T MyInt](a, b T) T {
	if a > b {
		return a
	}
	return b
}
