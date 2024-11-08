package main

import "fmt"

/*
Go语言中的Map是一种内置的数据结构，它提供了一种通过键（Key）来访问值（Value）的高效方式。
Map是无序的键值对集合，其中每个键在Map中都是唯一的，
且Map的键和值可以是任意类型（但键必须是可比较的类型，如整数、浮点数、字符串等）。
在Go语言中，Map的灵活性和高效性使其成为处理复杂数据结构的首选。
*/

func main() {
	//1.1 创建Map
	//在Go语言中，可以使用内置的make函数来创建一个空的Map。
	//make函数的语法为make(map[keyType]valueType)，
	//其中keyType和valueType分别代表键和值的类型。以下是一个简单的例子：
	//m := make(map[string]int) // 创建一个键为字符串，值为整数的Map
	//另外，还可以使用字面量初始化的方式来创建并初始化Map，例如：
	//test := map[string]int{
	//	"测试1": 1,
	//	"测试2": 2,
	//}
	//1.2 添加键值对
	//向Map中添加键值对非常简单，
	//只需使用map[key] = value的语法即可。如果键已存在，则更新其对应的值。
	m := make(map[string]int)
	m["apple"] = 5
	m["banana"] = 3
	//1.3 获取值
	//通过键来获取对应的值，可以使用map[key]的语法。
	//如果键不存在，则返回一个该类型的零值（例如，对于int类型，零值是0）
	fmt.Println(m["apple"]) // 输出: 5
	//1.4 判断键是否存在
	//在获取值的时候，有时候需要判断键是否存在。
	//可以使用“逗号ok”语法来实现，如果键存在，则ok为true，否则为false。
	value, ok := m["orange"]
	if ok {
		fmt.Println("orange的值为:", value)
	} else {
		fmt.Println("orange不存在")
	}

	//1.5 删除键值对
	//要删除Map中的键值对，可以使用delete函数，
	//其语法为delete(map, key)。如果键不存在，delete函数也不会报错，相当于空操作。
	delete(m, "banana")
	//1.6 遍历Map
	//使用for循环和range关键字可以遍历Map中的键值对。
	//遍历的顺序是不确定的，因为Map是无序的。
	for key, value := range m {
		fmt.Println(key, ":", value)
	}
	//1.7 获取Map的长度
	//可以使用len函数获取Map中键值对的个数。
	fmt.Println("map的长度为:", len(m))

}
