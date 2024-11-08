package main

import "fmt"

//func Add(a, b int) int {
//	return a + b
//}

// 匿名函数
var Add = func(a, b int) int {
	return a + b
}

func main() {
	result := Add(3, 2)
	fmt.Println(result)

}
