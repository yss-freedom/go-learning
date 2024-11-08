package main

import "fmt"

/*
当结构体中嵌入一个匿名结构体时，匿名结构体的字段会被提升，可以直接通过外部结构体访问。
结构体中的匿名字段，没有名字的字段，这个时候属性类型不能重复。
如何打印这个匿名字段，默认使用数据类型当做字段名称。
*/

type Teacher struct {
	string
	int
}

// 匿名结构体是没有名称的结构体，通常用于一次性使用的情况。
func main() {
	hobby := struct {
		HobbyId   int
		HobbyName string
	}{
		HobbyId:   1,
		HobbyName: "Basketball",
	}
	fmt.Println(hobby)
	fmt.Println(hobby.HobbyName) // 输出: Basketball

	// 匿名字段
	t1 := Teacher{"dashan", 18}
	fmt.Println(t1)
	// 如何打印这个匿名字段，默认使用数据类型当做字段名称
	fmt.Println(t1.string) // dashan
	fmt.Println(t1.int)    // 18

}
