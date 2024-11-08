package main

import "fmt"

func main() {
	//var a []int
	//a = append(a, 1)                 // 追加1个元素
	//a = append(a, 1, 2, 3)           // 追加多个元素, 手写解包方式
	//a = append(a, []int{1, 2, 3}...) // 追加一个切片, 切片需要解包
	//fmt.Println(a)

	var a = []int{1, 2, 3}
	a = append([]int{0}, a...)          // 在开头位置添加1个元素
	a = append([]int{-3, -2, -1}, a...) // 在开头添加1个切片
	fmt.Println(a)

	//var a []int
	//a = append(a[:i], append([]int{x}, a[i:]...)...)     // 在第i个位置插入x
	//a = append(a[:i], append([]int{1,2,3}, a[i:]...)...) // 在第i个位置插入切片

	//a = []int{1, 2, 3, ...}
	//a = a[1:]                       // 删除开头1个元素
	//a = a[N:]                       // 删除开头N个元素

	//a = []int{1, 2, 3, ...}
	//a = append(a[:0], a[1:]...) // 删除开头1个元素
	//a = append(a[:0], a[N:]...) // 删除开头N个元素

	//a = []int{1, 2, 3}
	//a = a[:copy(a, a[1:])] // 删除开头1个元素
	//a = a[:copy(a, a[N:])] // 删除开头N个元素

	//a = []int{1, 2, 3, ...}
	//a = append(a[:i], a[i+1], ...)
	//a = append(a[:i], a[i+N:], ...)

	//a = []int{1, 2, 3}
	//a = a[:copy(a[:i], a[i+1:])] // 删除中间1个元素
	//a = a[:copy(a[:i], a[i+N:])] // 删除中间N个元素

	//a = []int{1, 2, 3, ...}
	//
	//a = a[:len(a)-1]   // 删除尾部1个元素
	//a = a[:len(a)-N]   // 删除尾部N个元素

}
