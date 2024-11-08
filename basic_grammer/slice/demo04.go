package main

import "fmt"

// 查找切片中的最大值
func findMax(numbers []int) int {
	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func main() {
	numbers := []int{3, 5, 2, 8, 1}
	fmt.Println("Max:", findMax(numbers)) // 输出: Max: 8
}
