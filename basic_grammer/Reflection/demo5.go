package main

import (
	"fmt"
	"reflect"
)

/*
3. 获取值信息
通过reflect.ValueOf函数，我们可以获取变量的值信息。
reflect.Value类型提供了多种方法来获取和设置值，如获取布尔值、整数值、浮点数值、字符串值等。
此外，reflect.Value类型还提供了方法来操作结构体字段、数组元素、映射键值对等。
*/

type Person3 struct {
	Name string
	Age  int
}

func main() {
	var p Person3 = Person3{Name: "yss", Age: 18}

	// 获取值信息
	v := reflect.ValueOf(p)

	// 打印值信息
	fmt.Println("Value of Name:", v.FieldByName("Name").String())
	fmt.Println("Value of Age:", v.FieldByName("Age").Int())

	// 修改结构体字段值（注意：这里只是演示，实际上无法直接修改，因为v是不可变的）
	// 要修改值，需要使用reflect.ValueOf的Elem方法配合指针变量
	// 下面的代码会报错：panic: reflect: call of reflect.Value.SetInt on zero Value
	// v.FieldByName("Age").SetInt(35)

	// 正确的修改方式
	pv := reflect.ValueOf(&p).Elem() // 获取指针指向的元素的值
	pv.FieldByName("Age").SetInt(35) // 修改字段值

	// 打印修改后的值
	fmt.Println("Modified Age:", p.Age)
}
