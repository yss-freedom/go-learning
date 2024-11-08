package main

import "fmt"

// 泛型可以用在所有有类型的地方

type MyStruct[T int | string] struct {
	Id   T
	Name string
}

type IprintData[T int | float64 | string] interface {
	Print(data T)
}

// MyChan 通道
type MyChan[T int | string | float64] chan T

func main() {

	//T 泛型的参数类型的属性可以远不止一个，所有东西都可以泛型化。不一定非得用T  其他字母也可以
	// map(int)string

	// map[KEY]VALUE 类型形参（参数是不确定） KEY  、VALUE
	// KEY int | string   ， VALUE float32 | float64 约束
	// 类型的名字 MyMap[KEY,VALUE], 通过这一个类型，来代表多个类型！--> 泛型
	//map泛型就是多个参数 KEY VALUE
	type MyMap[KEY int | string | float64, VALUE float32 | float64 | int | string] map[KEY]VALUE

	//定义了泛型后，就可以实例化出不同类型的map
	// map [string]float64
	var score MyMap[string, float64] = map[string]float64{
		"go":   9.9,
		"java": 8.0,
	}
	fmt.Println(score)

	// map [int]string
	var rank MyMap[int, string] = map[int]string{
		1: "张三",
		2: "王五",
	}
	fmt.Println(rank)
}
