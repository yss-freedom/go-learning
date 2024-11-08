package main

import "fmt"

/*
3. 支持泛型衍生类型
新符号 ~，和类型一起出现的，表示支持该类型的衍生类型！
衍生类型，就是根据该类型常见的类型
*/

// int8 衍生类型
type int8A int8
type int8B = int8

// NewInt ~ 表示可以匹配该类型的衍生类型
type NewInt interface {
	~int8
}

// ~
func main() {
	var a int8A = 8
	var b int8A = 56
	fmt.Println(GetMax(a, b))
}
func GetMax[T NewInt](a, b T) T {
	if a > b {
		return a
	}
	return b
}
