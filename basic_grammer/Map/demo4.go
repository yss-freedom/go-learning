package main

import "fmt"

//Map可以嵌套使用，即Map的值可以是另一个Map。
//这种结构在处理复杂数据结构时非常有用。

func main() {
	// 嵌套Map示例
	nestedMap := make(map[string]map[string]int)
	nestedMap["fruits"] = make(map[string]int)
	nestedMap["fruits"]["apple"] = 5
	nestedMap["fruits"]["banana"] = 3

	fmt.Println(nestedMap["fruits"]["apple"]) // 输出: 5

}
