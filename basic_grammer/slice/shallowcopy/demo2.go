package main

import "fmt"

//结构体的浅拷贝

type Dog struct {
	Name string
	Age  int
}

/*
虽然在这个例子中dog2看起来像是dog1的深拷贝
（因为修改dog2.Name没有影响到dog1.Name），
但实际上这是因为Dog结构体中的字段都是值类型，
它们在赋值时自然进行了深拷贝。如果Dog结构体中包含引用类型的字段，
那么就会表现出浅拷贝的特性。
*/
func main() {
	dog1 := Dog{Name: "dog1", Age: 11}
	dog2 := dog1 // 浅拷贝
	dog2.Name = "dog2"
	fmt.Println(dog1) // 输出: {dog1 11}
	fmt.Println(dog2) // 输出: {dog2 11}
}
