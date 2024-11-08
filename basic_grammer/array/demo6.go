package main

/*
一维数组： 线性的，一组数
二维数组： 表格性的，数组套数组
三维数组： 立体空间性的，数组套数组套数组
xxxx维数组：xxx，数组套数组套数组…
*/

import "fmt"

func main() {

	// 定义一个多维数组  二维

	arr := [3][4]int{
		{0, 1, 2, 3},   // arr[0]  //数组
		{4, 5, 6, 7},   // arr[1]
		{8, 9, 10, 11}, // arr[2]
	}
	// 二维数组，一维数组存放的是一个数组
	fmt.Println(arr[0])
	// 要获取这个二维数组中的某个值，找到对应一维数组的坐标，arr[0] 当做一个整体
	fmt.Println(arr[0][1])
	fmt.Println("------------------")
	// 如何遍历二维数组
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Println(arr[i][j])
		}
	}
	// for range
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
