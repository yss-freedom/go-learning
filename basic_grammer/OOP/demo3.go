/*
方法：需要指定调用者，约定这个方法属于谁的. 对象.方法()
函数：不需要指定调用者，定义了函数就可以直接函数名()调用
方法：可以理解为函数多了一个调用者
方法定义, func 方法调用者 方法名(){} 方法的定义比函数多了个调用者
定义中，函数与方法就不一样，函数是 func 函数名(参数) 返回值 返回值类型 {}
1、方法可以重名，只需要调用者不同
2、如果调用者相同，则不能重名
*/
package main

import "fmt"

// 方法：可以理解为函数多了一个调用者
// 方法可以重名，不同的对象，调用的结果是不一样的

type Dog struct {
	name string
	age  int
}

// 方法定义,   func 方法调用者  方法名()
// 1、方法可以重名，只需要调用者不同
// 2、如果调用者相同，则不能重名
func (dog Dog) eat() {
	fmt.Println("Dog eating...")
}

func (dog Dog) sleep() {
	fmt.Println("Dog sleep...")
}

// 在Go语言中，fmt.Sprintf函数是一个非常实用的函数，它可以将一个字符串格式化为指定的格式。
// 它的用途非常广泛，可以用来构建复杂的字符串，包括打印日志、生成报错信息等。
// 带有参数和返回值的方法
func (dog Dog) xiaohua(food string) string {
	//返回值字符串格式化
	return fmt.Sprintf("%v吃进了%s, 吐出来骨头", dog.name, food)

}

type Cat struct {
	name string
	age  int
}

func (cat Cat) eat() {
	fmt.Println("Cat eating...")
}
func (cat Cat) sleep() {
	fmt.Println("Cat sleep...")
}

func main() {

	// 创建一个对象
	dog := Dog{
		name: "旺财",
		age:  2,
	}
	fmt.Println(dog)
	// 方法的调用，通过对应的结构体对象来调用
	dog.eat()

	//调用带有参数和返回值的方法
	back := dog.xiaohua("肉")
	fmt.Println(back)

	//不同对象调用同名方法
	cat := Cat{name: "喵喵", age: 1}
	cat.eat()

	//打印查看方法所属类型，类型跟函数一样都是func()
	fmt.Printf("方法所属类型: %T\n", cat.eat)
}
