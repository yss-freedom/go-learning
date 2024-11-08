package main

import "fmt"

// 切片作为函数参数时，传递的是切片的引用，
// 因此在函数内部对切片的修改会影响到原切片
func modifySlice(s []int) {
	s[0] = 100
}

func main() {
	mySlice := []int{1, 2, 3}
	modifySlice(mySlice)
	fmt.Println(mySlice) // 输出: [100 2 3]
}

/*
切片是Go语言中一种非常强大且灵活的数据结构，它通过封装数组提供了动态调整大小的能力。
切片的操作包括声明、初始化、访问元素、动态变化、遍历、复制等。
切片作为函数参数时，传递的是切片的引用，因此对切片的修改会影响到原切片。
*/
