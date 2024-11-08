package main

import "fmt"

// 使用Map实现类似Set的功能
// 虽然Go语言没有直接提供Set类型，但可以使用Map来实现类似Set的功能。
// 由于Map的键是唯一的，可以将键用作Set中的元素。
func main() {
	var mySet map[string]bool
	mySet = make(map[string]bool)

	// 添加元素
	mySet["apple"] = true
	mySet["banana"] = true

	// 遍历Set
	for key := range mySet {
		fmt.Println(key)
	}

	// 检查元素是否存在
	if _, exists := mySet["apple"]; exists {
		fmt.Println("apple exists")
	}

}
