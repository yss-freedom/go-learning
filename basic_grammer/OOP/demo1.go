package main

import "fmt"

// Person 定义一个父类
type Person struct {
	name string
	age  int
}

// Student 定义一个子类，Student拥有了父类所有的属性，还有自己的特性
type Student struct {
	Person // 父类以匿名字段形式作为子类的属性，实现了继承。
	school string
}

func main() {
	// 1、 创建父类对象
	p1 := Person{name: "yssnb", age: 18}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.age)

	// 2、创建子类
	s1 := Student{Person: Person{name: "小明", age: 3}, school: "大神学校"}
	fmt.Println(s1)
	fmt.Println(s1.Person.name, s1.Person.age, s1.school)

	// 3、创建子类
	var s2 Student
	s2.Person.name = "张三"
	s2.Person.age = 3
	s2.school = "北大"
	fmt.Println(s2)

	// 概念：提升字段, 只有匿名字段才可以做到
	// Person 在Student中是一个匿名字段， Person中的属性 name age 就是提升字段
	// 所有的提升字段就可以直接使用了，不需要再通过匿名字段点来获取了

	var s3 Student
	//字段提升，可以直接用子类对象操作父类字段
	s3.name = "李四"
	s3.age = 4
	s3.school = "清华"
	fmt.Println(s3)
	fmt.Println(s3.name, s3.age)

}
