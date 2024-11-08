package main

/*
高阶函数：可以将一个函数作为另外一个函数的参数。
fun1()
fun2(fun1)
fun1 函数作为 fun2 函数的参数
fun2函数，叫做高阶函数，接收了另外一个函数作为参数的函数
fun1函数，叫做回调函数，作为另外一个函数的参数
*/

import "fmt"

// 回调函数
func main() {
	// 函数调用
	r1 := add(1, 2)
	fmt.Println("r1是:", r1)
	// 高阶函数调用
	r2 := oper(1, 2, add)
	fmt.Println("r2是:", r2)
	r3 := oper(1, 2, sub)
	fmt.Println("r3是:", r3)

	// 匿名函数
	fun1 := func(a, b int) int {
		return a * b
	}
	r4 := oper(1, 2, fun1) // 调用匿名函数 *
	fmt.Println("r4是:", r4)

	// 能够直接传递匿名函数
	r5 := oper(8, 0, func(a int, b int) int {
		if b == 0 {
			fmt.Println("除数不能为0")
			return 0
		}
		return a / b
	})
	fmt.Println("r5是:", r5)
}

// 运算 (运算的数字，运算操作)
// 高阶函数，参数是接收另外一个函数
func oper(a, b int, fun func(int, int) int) int {
	//打印出的fun是每次传入的函数内存地址
	fmt.Println(a, b, fun)
	r := fun(a, b)
	return r
}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
