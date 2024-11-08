package main

/*
在Go语言中，值类型的变量在赋值时默认进行深拷贝（因为值类型的数据直接存储在变量中），
而引用类型的变量在赋值时默认进行浅拷贝（因为引用类型的数据存储在堆上，变量中只存储指向数据的指针）。
*/

//引用类型的浅拷贝

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := slice1 // 浅拷贝
	slice2[0] = 10
	fmt.Println(slice1) // 输出: [10 2 3]
	//slice2是slice1的浅拷贝，它们共享同一个底层数组。因此，修改slice2的第一个元素也会影响到slice1。
}
