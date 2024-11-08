// 多态案例2
package main

import "fmt"

// Animal3 定义接口
type Animal3 interface {
	eat()
	sleep()
}

type Dog3 struct {
	name string
}

func (dog Dog3) eat() {
	fmt.Println(dog.name, "--eat")
}
func (dog Dog3) sleep() {
	fmt.Println(dog.name, "--sleep")
}

// 多态
func main() {

	// Dog 两重身份：1、Dog   2、Animal ，多态
	dog1 := Dog3{name: "旺财"}
	dog1.eat()
	dog1.sleep()

	// Dog 也可以是 Animal
	test2(dog1)

	// 定义一个类型可以为接口类型的变量
	// 实际上所有实现类都可以赋值给这个对象
	var animal Animal3 // 模糊的 -- 具体化，将具体的实现类赋值给他，才有意义
	animal = dog1
	//接口是无法使用实现类的属性的
	test2(animal)
}

// Animal 接口
func test2(a Animal3) {
	a.eat()
	a.sleep()
}

//接口的实现类都拥有多态特性：除了自己本身还是他对应接口的类型。
