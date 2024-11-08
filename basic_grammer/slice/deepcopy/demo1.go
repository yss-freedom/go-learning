package main

import "fmt"

//Go标准库中没有直接提供深拷贝的函数，但可以通过以下几种方式实现深拷贝
//使用json.Marshal和json.Unmarshal
//这是一种简单但效率较低的方法，适用于简单的结构体。
//通过将对象序列化为JSON字符串，然后再反序列化回一个新对象，从而实现深拷贝。

//type Person struct {
//	Name string
//	Age  int
//}

func main() {
	//person1 := Person{"Alice", 25}
	//var person2 Person
	//
	//// 序列化
	//data, err := json.Marshal(person1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// 反序列化
	//err = json.Unmarshal(data, &person2)
	//if err != nil {
	//	panic(err)
	//}
	//
	//person2.Age = 30
	//fmt.Println(person1) // 输出: {Alice 25}
	//fmt.Println(person2) // 输出: {Alice 30}
	fmt.Println()
}
