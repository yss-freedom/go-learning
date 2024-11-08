package main

import "fmt"

//指针和切片
/*
在Go语言中，切片本身就是通过指针引用的，因此修改切片元素的值也会影响到原始数组。
但是，当重新分配切片时（例如使用append函数），切片的底层数组可能会改变，
此时原有的切片指针将不再指向新的底层数组。
*/
func modifySlice(s *[]int) {
	(*s)[0] = 99
	*s = append(*s, 100)
}

func main() {
	slice := []int{1, 2, 3}
	fmt.Println("Original Slice:", slice)
	modifySlice(&slice)
	fmt.Println("Modified Slice:", slice)
}
