package main

import "fmt"

// 定义函数
// 1、无参、没有返回值函数
//func add() {
//	fmt.Println("计算加法函数")
//}

// 无参数有返回值函数
// 需要指定函数返回值类型，并且用return将返回值返回
// 在调用处用变量接收
//func add() int {
//	return 3
//}

// 有参，有一个返回值函数
// 指定参数类型，和返回值类型
//func add(x int, y int) int {
//	return x + y
//}

// 4、有参，有多个返回值函数
// 需要指明参数类型，和多个返回值类型，多个返回值返回类型需要用括号括起来
// 多个返回值用逗号隔开
// 调用的时候，用多个变量接收返回值
func add(x int, y int) (int, int) {
	return x + y, y
}

func main() {
	//add()
	//a := add()
	//fmt.Println(a)

	//a := add(2, 5)
	//fmt.Println(a)

	a, b := add(3, 5)
	fmt.Println(a, b)
}
