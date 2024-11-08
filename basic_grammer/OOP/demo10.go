//接口嵌套
/*
接口可以嵌套其他接口，即一个接口可以继承多个别的接口。
这时，如果要实现这个接口，必须实现它继承的所有接口的方法
*/
package main

import (
	"fmt"
)

type AA interface {
	test1()
}
type BB interface {
	test2()
}

// CC 接口嵌套  CC :  test1()/test2()/test3()
// 如果要实现接口CC，那么需要实现这个三个方法。那这个对象就有3个接口可以转型。
type CC interface {
	AA // 导入AA接口中的方法
	BB
	test3()
}

// Dog7 编写一个结构体实现接口CC
type Dog7 struct {
}

func (dog Dog7) test1() {
	fmt.Println("test1")
}
func (dog Dog7) test2() {
	fmt.Println("test2")
}
func (dog Dog7) test3() {
	fmt.Println("test3")
}

func main() {
	// dog 拥有4种形态： Dog7 、CC 、 BB 、 AA
	var dog Dog7 = Dog7{}
	dog.test1()
	dog.test2()
	dog.test3()

	// 接口对象只能调用自己接口里面的方法
	var a AA = dog
	a.test1()
	//a.test2() // 向上转型之后只能调用它自己对应的方法
	var b BB = dog
	b.test2()

	//c三个方法都可以调用
	var c CC = dog
	c.test1()
	c.test2()
	c.test3()
}
