package main

import (
	"fmt"
	"reflect"
)

/*
6. 反射调用函数
reflect.ValueOf 通过函数名来进行反射 得到函数名 reflect.ValueOf(函数名)
然后通过函数名.Call() 调用函数
*/

// 反射调用函数  func
func main() {
	// 通过函数名来进行反射  reflect.ValueOf(函数名)
	// Kind func
	value1 := reflect.ValueOf(fun1)
	fmt.Println("打印拿到的数据类型", value1.Kind(), value1.Type())
	//调用无参函数
	value1.Call(nil)

	//调用有参函数
	value2 := reflect.ValueOf(fun2)
	fmt.Println("打印有参函数", value2.Kind(), value2.Type())

	//创建参数
	args1 := make([]reflect.Value, 2)
	args1[0] = reflect.ValueOf(1)
	args1[1] = reflect.ValueOf("hahahhaha")
	value2.Call(args1)

	//调用有参数，有返回值函数
	vuale3 := reflect.ValueOf(fun3)
	fmt.Println(vuale3.Kind(), vuale3.Type())
	args2 := make([]reflect.Value, 2)
	args2[0] = reflect.ValueOf(2)
	args2[1] = reflect.ValueOf("hahahhaha")
	//接收返回值 返回的是切片
	// func (v Value) Call(in []Value) []Value
	resultValue := vuale3.Call(args2)
	fmt.Println("返回值：", resultValue[0])
}

// 无参函数
func fun1() {
	fmt.Println("fun1:无参")
}

// 有参函数
func fun2(i int, s string) {
	fmt.Println("fun2:有参 i=", i, " s=", s)
}

// 有返回值函数
func fun3(i int, s string) string {
	fmt.Println("fun3:有参有返回值 i=", i, " s=", s)
	return s
}
