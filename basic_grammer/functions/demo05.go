package main

//定义函数时给返回值命名
/*
返回多个值，需要括号，需要表明返回值类型，返回值也可以命名。命名后，最后返回只需写return即可，也可以根据命名的返回值return。
直接调用return不带结果，那么则返回 函数返回值定义的顺序进行结果返回。
*/

import "fmt"

func main() {
	a, b := getresult(5, 6)
	fmt.Println(a, b)

}

// 要求：给出长方形的两个变长，得出长方形的周长和面积
// 可以在定义函数时，指定函数的返回值变量名
func getresult(x, y int) (zhouchang, mianji int) {
	zhouchang = (x + y) * 2
	mianji = x * y
	//函数定义处返回值命名了，return时要么和命名的变量名一样，要么可以直接不写
	//return zhouchang, mianji
	return

}
