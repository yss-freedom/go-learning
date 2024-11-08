package main

import (
	"fmt"
	"reflect"
)

/*
1. 动态数据处理
反射的一个主要场景是处理动态数据结构，例如解析JSON或处理数据库查询结果。
在这些场景中，数据的结构在编译时可能是未知的，因此需要使用反射来动态处理。
*/

// 反射
/*
Type : reflect.TypeOf(a) , 获取变量的类型
Value ：reflect.ValueOf(a) ， 获取变量的值
*/

func main() {
	// 正常编程定义变量
	var a int = 3
	// func TypeOf(i any) Type 反射得到该变量的数据类型
	fmt.Println("type", reflect.TypeOf(a))
	// func ValueOf(i any) Value  反射得到该变量的值
	fmt.Println("value", reflect.ValueOf(a))

	// 根据反射的值，来获取对象对应的类型和数值
	// 如果我们不知道这个对象的信息，我们可以通过这个对象拿到代码中的一切。
	// 获取a的类型
	v := reflect.ValueOf(a) // string int User
	// Kind : 获取这个值的种类， 在反射中，所有数据类型判断都是使用种类。
	//根据kind来判断其类型
	if v.Kind() == reflect.Float64 {
		fmt.Println(v.Float())
	}
	if v.Kind() == reflect.Int {
		fmt.Println(v.Int())
	}
	fmt.Println(v.Kind() == reflect.Float64)
	//fmt.Println(v.Type())

}
