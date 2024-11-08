/*
空接口interface{}不包含任何方法，因此任何类型都实现了空接口。空接口可以被视为能装入任意数量、任意数据类型的数据容器。
因此空接口可以存储任何的类型
空接口不好记，因此在新版本go中起了个名字，叫any
interface{}  == any
*/
package main

import "fmt"

// A 定义空接口
type A interface{}

// Dogg 所有结构体都实现了空接口A
type Dogg struct {
	name string
}

type Catt struct {
	name string
}

func testNow(a A) {
	fmt.Println(a)
}

// 可以指定定义空接口
// // any is an alias for interface{} and is equivalent to interface{} in all ways.
// type any = interface{}
// 可以传入任何东西
func testNow2(temp interface{}) {
}

func main() {
	//A类型可以是任何类型
	var a1 A = Catt{name: "喵喵"}
	var a2 A = Dogg{name: "旺财"}
	var a3 A = 1
	var a4 A = "yssnb"
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	testNow(a1)

	// map结合空接口，就可以存储任何类型数据
	map1 := make(map[string]interface{})
	map1["name"] = "ufo"
	map1["age"] = 18
	fmt.Println(map1)

	// slice，切片定义成空接口类型，也可以存放任何类型数据
	s1 := make([]any, 0, 10)
	s1 = append(s1, 1, "12312", false, a1, a2)
	fmt.Println(s1)

	//数组空接口，数组里面的值默认是nil，也可以存放任何数据类型
	var arr [4]interface{}
	fmt.Println(arr)
	arr[0] = 3
	arr[1] = "2"
	arr[2] = s1
	arr[3] = true
	fmt.Println(arr)

}
