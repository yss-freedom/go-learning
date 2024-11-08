package main

import (
	"fmt"
	"reflect"
)

//修改结构体变量：通过属性名，来实现修改

type Person4 struct {
	Name string
	Age  int
}

// 反射设置变量的值
func main() {

	person := Person4{"yss", 18}
	//注意，这里传入指针
	update2(&person)
	fmt.Println(person)

}

func update2(v any) {
	// 通过反射修改值,需要操作对象的指针，拿到地址，然后拿到指针对象
	pointer := reflect.ValueOf(v)
	newValue := pointer.Elem()

	fmt.Println("类型：", newValue.Type())
	fmt.Println("判断该类型是否可以修改：", newValue.CanSet())

	//修改结构体数据
	// 需要找到对象的结构体字段名
	newValue.FieldByName("Name").SetString("王安石")
	//如果不知道字段的类型，可以判断下
	fmt.Println("看下字段的类型", newValue.FieldByName("Name").Kind())

	newValue.FieldByName("Age").SetInt(99)

}
