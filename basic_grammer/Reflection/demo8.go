package main

import (
	"fmt"
	"reflect"
)

/*
5. 反射调用方法
反射还可以用于在运行时动态调用对象的方法。这在需要根据字符串名称调用方法的场景下非常有用，
例如实现一个简单的命令行接口或基于插件的架构。
通过方法名，找到这个方法, 然后调用 Call() 方法来执行 包括无参方法和有参方法
*/

type User6 struct {
	Name string
	Age  int
	Sex  string
}

func (user User6) Say2(msg string) {
	fmt.Println(user.Name, "说：", msg)
}

// PrintInfo2  打印结构体信息
func (user User6) PrintInfo2() {
	fmt.Printf("姓名：%s，年龄：%d，性别：%s\n", user.Name, user.Age, user.Sex)
}

func main() {
	user := User6{"yss", 18, "男"}

	// 通过方法名，找到这个方法, 然后调用 Call() 方法来执行
	// 反射调用方法
	value := reflect.ValueOf(user)
	fmt.Printf("kind:%s,  type:%s\n", value.Kind(), value.Type())

	//根据方法名找到方法，通过call来调用方法，call里面跟参数，。无参使用nil
	// func (v Value) Call(in []Value) []Value  Call的参数是个reflect.Value类型的切片
	value.MethodByName("PrintInfo2").Call(nil) // 无参方法调用

	// 有参方法调用。先创建个reflect.Value类型的切片，长度为参数的个数
	args := make([]reflect.Value, 1)
	//给参数赋值
	// func ValueOf(i any) Value
	args[0] = reflect.ValueOf("这反射来调用的")
	//找到方法名，调用传参
	value.MethodByName("Say2").Call(args) // 有参方法调用
}
