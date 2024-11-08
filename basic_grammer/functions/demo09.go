package main

/*
在Go语言中，可以使用递归函数，也就是一个函数在自身内部调用自身。
递归函数通常用于解决需要重复执行相同操作的问题，比如：计算阶乘、斐波那契数列等。
需要注意的是，递归函数必须有一个终止条件。否则，会无限递归下去，导致栈溢出甚至程序崩溃。
*/

import "fmt"

/*
定义：一个函数自己调用自己，就叫做递归函数
注意：递归函数需要有一个出口，逐渐向出口靠近，没有出口就会形成死循环。
*/
func main() {
	// overflows 栈溢出
	sum := getSum2(100)
	fmt.Println(sum)
}
func getSum2(n int) int {
	//fmt.Println(n)
	//递归出口
	if n == 1 {
		return 1
	}
	//将自己返回，自己调用自己
	return getSum2(n-1) + n
}

// 假设 getSum(5)

// 求和 sum  1 + 2 + 3 + 4 + 5
// getSum(5)☟=15
//      getSum(4)☟=10 + 5
//          getSum(3)☟=6 + 4
//               getSum(2)☟=3 + 3
// //                getSum(1)=1 + 2
