package main

import "fmt"

/*
在使用指针之前，应该进行空指针检查，以避免出现空指针引用的错误。可以使用nil值来表示空指针。例如：
*/
func main() {
	var p *int // 声明一个int类型的指针变量，默认值为nil
	if p != nil {
		fmt.Println(*p) // 对非空指针进行操作
	} else {
		fmt.Println("Pointer is nil")
	}
}
