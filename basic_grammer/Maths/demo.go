package main

import "math"
import "fmt"

func main() {
	//1. 绝对值函数
	//绝对值函数Abs用于获取一个数的绝对值。
	value := -10.5
	absValue := math.Abs(value)
	fmt.Println(absValue)

	//2. 平方根函数
	//平方根函数Sqrt用于计算一个数的平方根。如果参数是负数，函数将返回NaN（非数字）。
	sqrtOfNine := math.Sqrt(9)
	fmt.Println(sqrtOfNine) // 输出: 3

	//3. 对数函数
	//math包提供了两种对数函数：自然对数函数Log（以e为底）和常用对数函数Log10（以10为底）。
	logE := math.Log(math.E)
	fmt.Println(logE) // 输出: 1

	log10Of100 := math.Log10(100)
	fmt.Println(log10Of100) // 输出: 2

	//4. 取整函数
	//math包提供了三个取整函数：向上取整函数Ceil、向下取整函数Floor和取整函数Trunc。
	ceilValue := math.Ceil(3.14)
	fmt.Println(ceilValue) // 输出: 4

	floorValue := math.Floor(3.8)
	fmt.Println(floorValue) // 输出: 3

	truncValue := math.Trunc(3.8)
	fmt.Println(truncValue) // 输出: 3

	//5. 幂函数
	//幂函数Pow用于计算一个数的指数次幂。它接受两个参数，第一个参数是底数，第二个参数是指数。
	result := math.Pow(2, 3)
	fmt.Println(result) // 输出: 8

	//6. 三角函数
	//math包提供了一系列三角函数，包括正弦函数Sin、余弦函数Cos、正切函数Tan及其反函数和双曲函数。
	sinValue := math.Sin(math.Pi / 2)
	fmt.Println(sinValue) // 输出: 1

	cosValue := math.Cos(0)
	fmt.Println(cosValue) // 输出: 1

	tanValue := math.Tan(math.Pi / 4)
	fmt.Println(tanValue) // 输出: 1

	//7. 指数函数
	//指数函数Exp用于计算e（自然对数的底数）的x次方。
	expValue := math.Exp(1)
	fmt.Println(expValue) // 输出: 2.718281828459045

	//8. 其他函数
	//math包还提供了其他一些有用的函数，
	//如立方根函数Cbrt、判断是否为NaN的函数IsNaN、
	//判断是否为无穷大的函数IsInf等。
	cbrtValue := math.Cbrt(27)
	fmt.Println(cbrtValue) // 输出: 3

	isNaN := math.IsNaN(math.Sqrt(-1))
	fmt.Println(isNaN) // 输出: true
	//圆周率π的值，常用于圆形和球形的几何计算。
	fmt.Println(math.Pi) // 输出: 3.141592653589793
	//自然对数的底数e，常用于指数和对数计算
	fmt.Println(math.E) // 输出: 2.718281828459045

}
