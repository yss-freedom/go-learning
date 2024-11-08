package main

import (
	"fmt"
	"math"
)

//3. 判断数值类型
//math包还提供了一些用于判断数值类型的函数，
//如IsNaN、IsInf等。这些函数在错误处理和数值验证中非常有用。

func main() {
	// 判断是否为NaN
	isNaN := math.IsNaN(math.Sqrt(-1))
	fmt.Println(isNaN) // 输出: true

	// 判断是否为正无穷大
	isInf := math.IsInf(math.Inf(1), 1)
	fmt.Println(isInf) // 输出: true

	// 判断是否为负无穷大
	isNegInf := math.IsInf(math.Inf(-1), -1)
	fmt.Println(isNegInf) // 输出: true
}
