package main

import (
	"fmt"
	"reflect"
)

//获取结构体信息

type User5 struct {
	Name string
	Age  int
	Sex  string
}

func (user User5) Say(msg string) {
	fmt.Println("User 说：", msg)
}

// PrintInfo  打印结构体信息
func (user User5) PrintInfo() {
	fmt.Printf("姓名：%s，年龄：%d，性别：%s\n", user.Name, user.Age, user.Sex)
}

func main() {
	user := User5{"yss", 18, "男"}

	reflectGetInfo(user)
}

// 通过反射，获取变量的信息
func reflectGetInfo(v interface{}) {
	// 1、获取参数的类型Type , 可能是用户自己定义的，但是Kind一定是内部类型struct
	getType := reflect.TypeOf(v)
	fmt.Println(getType.Name()) // 类型信息 User5
	fmt.Println(getType.Kind()) // 找到上级的种类Kind  struct

	// 2、获取值
	getValue := reflect.ValueOf(v)
	//查看类型 Type得到的是我们自定义的类型main.User5   Kind得到的是go内置的类型 struct
	fmt.Println("获取到的value--type值类型", getValue.Type()) // main.User5
	fmt.Println("获取到的value--kind值类型", getValue.Kind()) // struct
	fmt.Println("获取到value", getValue)

	// 获取字段，通过Type扒出字段
	// Type.NumField() 获取这个类型中有几个字段  3
	// field(index) 得到字段的值
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i) // 类型
		//通过valueof.Field().Interface拿到值
		value := getValue.Field(i).Interface() // value
		// 打印
		fmt.Printf("字段名：%s,字段类型：%s，字段值：%v\n", field.Name, field.Type, value)
	}

	// 获取这个结构体的方法 , NumMethod 可以获取方法的数量
	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Printf("方法的名字：%s\t,方法类型：%s", method.Name, method.Type)
		//执行方法
		//method.Name()
	}

}

/*
由上面反射，就可以推导出结构体字段以及其方法
type User struct{
   Name string
   Age int
   Sex string
}
func (main.User) PrintInfo(){}
func (main.User) Say(string)
*/
