package main

import (
	"fmt"
	"math"
)

//接口中的指针
/*
在Go语言中，接口是一种类型，它定义了对象的行为。
当接口的实现是大型结构体时，使用指针作为接口的实现可以减少数据拷贝，提高性能
*/

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	circle := Circle{Radius: 5}
	printArea(&circle) // 注意这里传递的是&circle，因为Circle的Area方法接收的是指针
}

//注意：在上面的printArea函数调用中，
//我们传递了&circle而不是circle，
//因为Circle的Area方法接收的是*Circle类型的参数。
