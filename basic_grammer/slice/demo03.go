package main

import "fmt"

// 计算切片中整数的和
func sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum:", sum(numbers)) // 输出: Sum: 15
}
