package main

//自定义类型的方法
/*
在Go中，你可以为任何自定义类型（结构体、类型别名等）定义方法。
方法是附加到类型上的函数，其第一个参数是接收者（receiver），接收者指定了方法所属的类型。
*/

import "fmt"

type MyInt1 int

// 为 MyInt 类型定义方法

func (m MyInt1) SayHello() {
	fmt.Println("Hello from MyInt1:", m)
}

func main() {
	var x MyInt1 = 10
	x.SayHello() // 输出: Hello from MyInt: 10
}
