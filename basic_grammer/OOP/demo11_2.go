//多个预期结果判断
/*
通过switch来判断 switch i.(T)
i 必须是接口类型
i.(type)必须在switch中使用
*/
package main

import "fmt"

// 通过switch来判断  switch i.(T)

type I interface{}

// 如果断言的类型同时实现了switch 多个case匹配，默认使用第一个case
// 所以要把范围更小的匹配放前面
func testAssert(i interface{}) {

	// switch i.(type) 接口断言
	//i.(type)必须在switch中使用
	switch i.(type) {
	case string:
		fmt.Println("变量为string类型")
	case int:
		fmt.Println("变量为int类型")
	case nil:
		fmt.Println("变量为nil类型")
	case map[string]int:
		fmt.Println("map类型")
	case interface{}:
		fmt.Println("变量为interface{}类型")
	//空接口与I一样
	case I:
		fmt.Println("变量为I类型")

	// .....
	default:
		fmt.Println("未知类型")
	}
}

func main() {

	testAssert("string")
	testAssert(1)
	var i I      // 没有初始化空接口时，默认值为 nil类型 不属于I类型
	var i2 I = 1 // 只有赋值了之后，才是对应的类型

	testAssert(i)
	testAssert(i2)

	//map类型
	j := make(map[string]int)
	testAssert(j)

}
