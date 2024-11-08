package main

import "fmt"

func main() {
	// 初始化数组的几种方式
	// 1.在定义数组的时候就直接初始化，用大括号包裹数组的值
	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	//2.短变量快速初始化
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	//自动根据数组的长度来给 … 赋值，自动推导长度
	//数据如果来自用户，我不知道用户给我多少个数据，数组
	//… 代表数组的长度
	//Go的编译器会自动根据数组的长度来给 … 赋值，自动推导长度
	//注意点：这里的数组不是无限长的，也是固定的大小，大小取决于数组元素个数。
	var arr3 = [...]int{1, 2, 3, 4, 5, 5, 6, 7, 8, 8, 8}
	fmt.Println(len(arr3))
	fmt.Println(arr3)

	//数组默认值,我只想给其中的某几个index位置赋值。
	//{index:值}
	//给赋值的下标有值，没赋值的下标是默认值
	var arr4 [10]int
	arr4 = [10]int{1: 100, 5: 500}
	fmt.Println(arr4) // [0 100 0 0 0 500 0 0 0 0]

	//遍历数组元素
	//遍历数组可以通过下标获取元素 arr[index]、使用for循环或者for-range结构。
	//直接通过下标获取元素 arr[index]
	//使用for循环遍历：
	//for i := 0; i < len(arr); i++ {
	//	fmt.Println(arr[i])
	//}
	//
	////使用for-range遍历：
	//for index, value := range arr {
	//	fmt.Printf("Index: %d, Value: %d\n", index, value)
	//}

}
