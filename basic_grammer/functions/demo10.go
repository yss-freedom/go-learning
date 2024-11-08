package main

import "fmt"

func main() {
	fmt.Println(feibo(5))
}

// 使用递归实现斐波那契数列,计算出第n个值是多少
// 斐波那契数列 1 1 2 3 5 8 13 .....
func feibo(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	//第n个数是第n-1个数与第n-2个数的和
	return feibo(n-1) + feibo(n-2)
}
