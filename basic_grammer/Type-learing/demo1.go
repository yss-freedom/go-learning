package main

import "fmt"

// var 变量   type 类型（结构体、接口、别名...）

// type的别名用法,全局变量中
// 这是定义了一个新类型 MyInt，是int转换过来的，和int一样，但是不能通int发生操作，类型不同
// 这里他俩MyInt int 是两个类型
// 创建了一种新类型！

type MyInt int

func main() {
	var a MyInt = 20 // MyInt
	var b int = 10   // int

	//自定义的类型和原类型也不能做运算
	// invalid operation: a + b (mismatched types MyInt and int)
	//fmt.Println(a + b)

	//可以进行强制类型转换
	// 类型转换： T(v)
	fmt.Println(int(a) + b) // 30

	//查看他俩数据类型
	fmt.Printf("%T\n", a) // main.MyInt
	fmt.Printf("%T\n", b) // int

	// 给int起一个小名，但是它还得是int   type  any
	type diyint = int //用等号赋值，这里diyint和int是一样的

	var c diyint = 30

	//查看数据类型，可以看到还是原来的数据类型
	fmt.Printf("%T\n", c) // int

	//此时的c和基本int类型的数据可以直接进行运算
	fmt.Println(c + b) //40

}

/*
type关键字的理解：
1、type 定义一个类型
 - type User struct 定义结构体类型
 - type User interface 定义接口类型
 - type Diy (int、string、....) 自定义类型，全新的类型

2、type 起别名
 - type xxx = 类型 ，将某个类型赋值给 xxx，相当于这个类型的别名
 - 别名只能在写代码的时候使用，增加代码的可阅读性。
 - 真实在项目的编译过程中，它还是原来的类型。
*/
