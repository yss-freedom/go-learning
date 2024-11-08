package main

import "fmt"

/*
在Go语言中，由于没有传统的构造函数，
我们常常使用工厂模式来创建结构体的实例。工厂模式是一种创建型设计模式，它提供了一种创建对象的最佳方式
*/

type Person4 struct {
	Name string
	Age  int
}

// 工厂函数
func NewPerson(name string, age int) *Person4 {
	return &Person4{
		Name: name,
		Age:  age,
	}
}

func main() {
	p := NewPerson("Eve", 28)
	fmt.Println(p.Name) // 输出: Eve
	fmt.Println(p.Age)
	//NewPerson函数就是一个工厂函数，它接受参数并返回一个指向Person结构体的指针
}
