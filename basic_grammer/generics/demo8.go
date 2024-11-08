package main

import "fmt"

/*
搜索算法：
泛型也可以用于实现通用的搜索算法。例如，定义一个泛型函数来查找切片中的元素
*/

func Find[T comparable](slice []T, value T) (int, bool) {
	for i, v := range slice {
		if v == value {
			return i, true
		}
	}
	return -1, false
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	index, found := Find(numbers, 4)
	if found {
		fmt.Println("Element found at index:", index) // 输出: Element found at index: 2
	} else {
		fmt.Println("Element not found")
	}
}
