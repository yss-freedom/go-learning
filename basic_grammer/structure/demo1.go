package main

import "fmt"

/*
结构体是Go语言中的一种复合数据类型，可以视为一个字段的集合。
这些字段可以是不同的数据类型，包括基本数据类型（如int、float64、string等）、指针、数组、切片、其他结构体等。
结构体是用户自定义的类型，它们使得数据处理更加模块化，并增强了代码的可读性和可维护性。
*/
//结构体的定义使用type关键字和struct关键字。定义结构体时，需要在{}中列出所有的字段及其类型。

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	//一旦定义了结构体，就可以创建结构体的实例（即变量），并通过点（.）操作符访问或修改其字段。
	//var p Person
	//p.Name = "Alice"
	//p.Age = 30
	//p.Email = "alice@example.com"
	//
	//fmt.Println(p.Email) // 输出: Alice

	//结构体可以使用字面量进行初始化，即直接在声明变量时指定字段的值
	p1 := Person{
		Name: "Alice",
		Age:  30,
	}
	//也可以指定字段名来初始化结构体，这样可以忽略字段的声明顺序
	p2 := Person{
		Name: "Bob",
		Age:  25,
	}
	//
	//// 或者只初始化部分字段
	//p3 := Person{
	//	Name: "Charlie",
	//	Age: 35,
	//}
	//p3.Job = "Artist"
	//p3.Salary = 4500
	//fmt.Println(p1.Name)  // 输出: Alice
	//fmt.Println(p2.Age)   // 输出: 25
	p := Person{Name: "David", Age: 40}
	printPerson(p)
	printPersonPtr(&p)
	fmt.Println("Age after function call:", p.Age) // 输出: 41

	//结构体切片是一个包含多个结构体实例的切片，用于管理结构体的集合
	var people []Person
	people = append(people, p1)
	people = append(people, p2)

	for _, person := range people {
		fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
	}
}
func printPerson(p Person) {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}

func printPersonPtr(p *Person) {
	fmt.Println("Name:", p.Name)
	p.Age++ // 修改指针指向的结构体中的字段
}
