package main

import "fmt"

// 指针的高级用法
// 指针数组和指针切片
// 指针数组和指针切片允许我们存储多个指针变量，并通过这些指针访问和修改对应的值。
// 指针数组
func main() {

	a := 1
	b := 2
	c := 3
	d := 4

	// 创建一个指针数组
	arr1 := [4]*int{&a, &b, &c, &d}
	fmt.Println(arr1)

	// 通过指针修改a的值
	// arr1[0] 0xc00000e0a8
	*arr1[0] = 100
	fmt.Println(a)

	a = 200
	fmt.Println(*arr1[0])

	//指针切片
	slice := []int{4, 5, 6}
	var pSlice *[]int = &slice
	fmt.Println("Pointer to Slice:", *pSlice)
}
