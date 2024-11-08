package main

import (
	"fmt"
	"reflect"
)

/*
4. 反射修改变量的值
通过反射修改值,需要操作对象的指针，拿到地址，然后拿到指针对象
通过Canset来判断值是否可以被修改
*/

// 反射设置变量的值
func main() {

	var num float64 = 3.14
	update(&num) //注意，这里传入指针地址
	fmt.Println(num)

}

func update(v any) {
	// 通过反射修改值,需要操作对象的指针，拿到地址，然后拿到指针对象
	pointer := reflect.ValueOf(v)
	newValue := pointer.Elem()

	fmt.Println("类型：", newValue.Type())
	fmt.Println("判断该类型是否可以修改：", newValue.CanSet())

	// 通过反射对象给变量赋值
	//newValue.SetFloat(999.43434)
	// 也可以对类型进行判断，根据不同类型修改成不同的值

	if newValue.Kind() == reflect.Float64 {
		// 通过反射对象给变量赋值
		newValue.SetFloat(2.21)
	}
	if newValue.Kind() == reflect.Int {
		// 通过反射对象给变量赋值
		newValue.SetInt(2)
	}
}
