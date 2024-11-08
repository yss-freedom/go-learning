//类型安全的断言
/*
instance, ok := 接口对象.(实际类型)
语法：t,ok:= i.(T) ok 隐藏返回值，如果断言成功 ok就是true、否则就是false
如果断言失败，ok将会是false，而instance将会是类型的零值，并且不会触发panic。
*/
//接口断言代码示例
package main

import "fmt"

// 断言  t := i.(T)   t：t就是i接口是T类型的  i：接口   T:类型
// 语法：t,ok:= i.(T) ok 隐藏返回值，如果断言成功 ok就是true、否则就是false

func main() {
	//assertsString("11111111111")

	assertsInt("中国")
}

// 断言失败的情况，我们希望程序不会停止。
func assertsInt(i any) {
	r, ok := i.(int)
	if ok {
		fmt.Println("是我们期望的结果 int")
		fmt.Println(r)
	} else {
		fmt.Println("不是我们期望的结果，无法执行预期操作")
	}

}
