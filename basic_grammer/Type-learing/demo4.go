/*
在Go中，我们还可以定义函数类型，即函数的签名。
函数类型允许我们将函数作为参数传递给其他函数，或者将函数作为返回值。
*/
package main

import "fmt"

type AddFunc func(a, b int) int

func add(a, b int) int {
	return a + b
}

func apply(f AddFunc, x, y int) int {
	return f(x, y)
}

func main() {
	result := apply(add, 5, 3)
	fmt.Println(result) // 输出: 8
}

/*
在上述代码中，我们定义了一个AddFunc函数类型，它接受两个int类型的参数并返回一个int类型的结果。然后，我们定义了一个add函数，它符合AddFunc的签名。最后，我们定义了一个apply函数，
它接受一个AddFunc类型的参数和两个int类型的参数，并返回调用该函数的结果。
在Go语言中，type 关键字不仅用于定义基础的数据类型别名、结构体、接口和函数类型，还有一些高级用法和特性。接下来，我们将继续探讨 type 在Go语言中的其他用法。
*/
