package main

import "fmt"

//结构体可以嵌套使用，即一个结构体的字段可以是另一个结构体类型。这有助于构建更复杂的数据结构。

type Address struct {
	City    string
	Country string
}

type Person1 struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	var p Person1
	p.Name = "Bob"
	p.Age = 25
	p.Address.City = "New York"
	p.Address.Country = "USA"

	fmt.Println(p.Address.City) // 输出: New York
	fmt.Println(p.Address.Country)
}
