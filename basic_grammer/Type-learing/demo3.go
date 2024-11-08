/*
结构体类型
结构体是Go语言中的一种复合数据类型，用于将多个不同类型的变量组合成一个单一的类型。
结构体类型允许我们创建具有多个属性的自定义类型。
*/
package main

import "fmt"

//结构体字段还可以包含标签（也称为元数据），
//这些标签用于为字段提供额外的信息，常用于JSON序列化/反序列化等场景。

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p) // 输出: {Alice 30}
	p.Name = "Bob"
	fmt.Println(p.Name) // 输出: Bob
	var person = struct {
		Name string
		Age  int
	}{"Alice", 30}
	fmt.Println(person) // 输出: {Alice 30}

}
