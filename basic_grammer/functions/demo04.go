package main

import "fmt"

//给出两个数，得到其中较大者

func main() {
	num := max(3, 5)
	fmt.Println(num)
}

// 定义比较大小函数
func max(x, y int) int {
	var result int
	if x > y {
		result = x
	} else {
		result = y
	}
	return result

}
