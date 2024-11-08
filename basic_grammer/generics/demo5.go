package main

import "fmt"

/*
1. 定义泛型函数：
单纯的泛型没啥意义。和函数结合使用， 可以使用调用者（调用者的类型可以自定义，就可以实现泛型。）
带了类型形参的函数就叫做泛型函数，极大的提高代码的灵活心，降低阅读性！
泛型函数可以处理多种类型的数据。例如，定义一个交换两个值的泛型函数：
*/

func Swap[T any](a, b T) (T, T) {
	return b, a
}

func main() {
	a, b := Swap[int](1, 2) // 显式指定类型
	fmt.Println(a, b)       // 输出: 2 1

	c, d := Swap("hello", "world") // 隐式类型推断,不用写类型
	fmt.Println(c, d)              // 输出: world hello
}
