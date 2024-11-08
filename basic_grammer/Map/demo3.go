package main

import "fmt"

/*
Map可以作为函数的参数传递。需要注意的是，
传递的是Map的引用，而不是Map的副本。
因此，在函数中对Map的修改会影响到原始Map。
*/
// Map作为函数参数
func ModifyMap(m map[string]int, key string, value int) {
	m[key] = value
}

func main() {
	m := make(map[string]int)
	m["apple"] = 5

	ModifyMap(m, "apple", 10)

	fmt.Println(m["apple"]) // 输出: 10
}
