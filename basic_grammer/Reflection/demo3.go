package main

import (
	"fmt"
	"reflect"
)

/*
2. 获取类型信息
通过reflect.TypeOf函数，我们可以获取变量的类型信息。
reflect.Type类型提供了多种方法来获取类型的详细信息，如类型名称、类型种类、字段信息等。
*/

type Person struct {
	Name string
	Age  int
}

func main() {
	var p Person = Person{Name: "Alice", Age: 30}

	// 获取类型信息
	t := reflect.TypeOf(p)

	// 打印类型名称
	fmt.Println("Type name:", t.Name())

	// 打印类型种类
	fmt.Println("Type kind:", t.Kind())

	// 打印结构体字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field name: %s, Field type: %s\n", field.Name, field.Type)
	}
}
