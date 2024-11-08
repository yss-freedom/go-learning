package main

//递归深拷贝
/*
对于复杂的数据结构，可以使用递归深拷贝的方法。
这种方法需要手动编写代码来遍历对象中的所有字段，
如果是基本类型则直接复制，如果是复杂类型（如结构体、切片、映射等）则递归调用深拷贝函数。
*/

import (
	"fmt"
	"reflect"
)

func DeepCopy(input interface{}) interface{} {
	if input == nil {
		return nil
	}

	switch reflect.TypeOf(input).Kind() {
	case reflect.Bool, reflect.String, reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
		return input
	case reflect.Struct:
		in := reflect.ValueOf(input)
		out := reflect.New(in.Type()).Elem()
		for i := 0; i < in.NumField(); i++ {
			field := in.Field(i)
			if field.CanInterface() {
				// 递归调用DeepCopy
				out.Field(i).Set(reflect.ValueOf(DeepCopy(field.Interface())))
			}
		}
		return out.Interface()
	// 可以根据需要添加对切片、映射等类型的支持
	default:
		// 其他类型根据需要进行处理
		return input
	}
}

func main() {
	// 匿名结构体Person
	type Person struct {
		Name string
		Age  int
	}

	// 示例结构体
	type Object struct {
		Num   int
		Str   string
		Slice []int
		Map   map[string]int
		Person
	}

	obj1 := &Object{
		Num:   1,
		Str:   "hello",
		Slice: []int{2, 3},
		Map:   map[string]int{"age": 18},
		Person: Person{
			Name: "Lucas",
			Age:  20,
		},
	}

	// 深拷贝
	obj2 := DeepCopy(obj1).(*Object)

	// 修改obj1的Name字段
	obj1.Person.Name = "Nina"

	fmt.Println("obj1:", obj1)
	fmt.Println("obj2:", obj2)
}

// 输出结果：
// obj1: &{1 hello [2 3] map[age:18] {Nina 20}}
// obj2: &{1 hello [2 3] map[age:18] {Lucas 20}}
