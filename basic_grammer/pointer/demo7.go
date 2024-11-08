package main

import "fmt"

//指针和方法
/*
在Go语言中，方法的接收者可以是值类型或指针类型。
使用指针作为接收者可以减少数据拷贝，提高性能，特别是当处理大型结构体时
*/

type Rectangle struct {
	Width  int
	Height int
}

func (r *Rectangle) Area() int {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{Width: 5, Height: 10}
	fmt.Println("Area:", rect.Area())
}
