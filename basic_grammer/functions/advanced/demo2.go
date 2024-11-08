package main

//func （xxxx,xxx） xxx,xxxx
//函数也是一种数据类型，可以定义函数类型的变量

import "fmt"

// 函数是什么（数据类型）
func main() {
	a := 10.01
	fmt.Printf("%T\n", a) // 查看变量的类型
	b := [4]int{1, 2, 3, 4}
	fmt.Printf("%T\n", b) // 查看变量的类型
	c := true
	fmt.Printf("%T\n", c) // 查看变量的类型

	// 函数的类型
	func1()                   // 带了括号是函数的调用
	fmt.Printf("%T\n", func1) // 查看变量的类型 func()
	fmt.Printf("%T\n", func2) // 查看变量的类型 func(int, int, ...string) (int, int)
	// func(int, int) (int, int)
	// func(int, int, ...string) (int, int)
	//var fun3 func(int, int, ...string) (int, int)
	fun3 := func2
	r1, r2 := fun3(1, 2, "111")
	fmt.Println(r1, r2)
	// 函数在Go语言中本身也是一个数据类型，加了（） 是调用函数，不加（）, 函数也是一个变量，可以赋值给别人。

	// 函数的类型就等于该函数创建的类型，他也可以赋值给
}

// 无参无返回值的函数
func func1() {

}

// 有参有返回值的函数
func func2(a, b int, c ...string) (int, int) {
	return 0, 0
}

/*
函数的本质
函数在Go语言中不是一个简单的调用或者接收结果的。
函数在Go中是一个符合类型，可以看做是一个特殊的变量。var 定义吗，赋值。类型相同即可
函数类型的样子 ：var f1 函数名(参数) 结果
变量名：指向一段内存 （num --> 0x11111aaaa）
函数名：指向一段函数体的内存地址，是一种特殊类型的变量。我们可以将一个函数赋值给另外一个类型相同的函数
*/
