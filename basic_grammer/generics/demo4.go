package main

import "fmt"

func main() {
	// 特殊的泛型类型，泛型的参数是多样的，但是实际类型定义就是int
	type AAA[T int | string] int

	var a AAA[int] = 123
	var b AAA[string] = 123
	fmt.Println(a)
	fmt.Println(b)

	//查看类型
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	//var c AAA[string] = "hello" //不能这样赋值，因为AAA的值约束的类型是int

}
