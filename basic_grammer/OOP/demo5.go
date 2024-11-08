package main

/*
1. 接口的定义
1、Go语言提供了接口数据类型。
2、接口就是把一些共性的方法集合在一起定义。
3、如果有实现类将接口定义的方法全部实现了，那么就代表实现了这个接口
4、隐式实现 Go ，假设A实现了B接口中的所有方法，不需要显示声明
5、接口是方法的定义集合，不需要实现具体的方法内容。名字约束
在Go语言中，接口（Interface）是一个重要的特性，它允许我们定义一组方法但不实现它们，任何类型只要实现了这些方法，就被认为是实现了该接口。
接口体现了程序设计的多态、高内聚、低耦合的思想，是实现面向对象编程中多态性的关键工具。
接口通过interface关键字定义，它是一组方法的集合。接口中的方法没有实现体，即它们没有具体的实现代码。一个类型只要实现了接口中的所有方法，就认为该类型实现了该接口。
如果一个结构体实现了这个接口所有的方法，那这个结构体就是这个接口类型的
*/

import (
	"fmt"
)

// 接口： USB、typec、插座
// 1、Go语言提供了接口数据类型。
// 2、接口就是把一些共性的方法集合在一起定义。
// 3、如果有实现类将接口定义的方法全部实现了，那么就代表实现了这个接口
// 4、隐式实现 Go ，假设A实现了B接口中的所有方法，不需要显示声明
// 5、接口是方法的定义集合，不需要实现具体的方法内容。名字约束

// USB 接口的定义 interface 来定义，方法太多了，要归类，方法的集合
type USB interface { // 接口，方法的集合
	input()  // 输入方法
	output() // 输出方法
}

// Mouse 结构体
type Mouse struct {
	name string
}

// 结构体实现了接口的全部方法就代表实现了这个接口，否则不算实现这个接口
func (mouse Mouse) output() {
	fmt.Println(mouse.name, "鼠标输出")
}
func (mouse Mouse) input() {
	fmt.Println(mouse.name, "鼠标输入")
}

// 接口调用测试
func test(u USB) {
	u.input()
	u.output()
}

func main() {
	// 通过传入接口实现类来进行调用
	m1 := Mouse{name: "罗技"}
	// test 参数 USB 类型，如果一个结构体实现了这个接口所有的方法，那这个结构体就是这个接口类型的
	test(m1)
	//也可以单独测试接口
	//m1.input()

	k1 := KeyBoard{name: "雷蛇"}
	test(k1)

	// 定义高级类型  k1就升级了  KeyBoard --> USB  向上转型
	var usb USB
	usb = k1
	fmt.Println(usb)
	// 接口是无法使用实现类的属性的
	//fmt.Println(usb.name)

}

// KeyBoard 结构体
type KeyBoard struct {
	name string
}

// 结构体实现了接口的全部方法就代表实现了这个接口，否则不算实现这个接口
func (key KeyBoard) output() {
	fmt.Println(key.name, "键盘输出")
}
func (key KeyBoard) input() {
	fmt.Println(key.name, "键盘输入")
}
