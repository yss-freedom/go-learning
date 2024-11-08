//接口断言
/*
接口断言用于检查接口变量是否持有特定类型的值，并获取该值。
被断言的对象必须是接口类型，否则会报错
它有两种形式：不安全断言和类型安全的断言。
*/
//不安全断言
//instance := 接口对象.(实际类型)
//如果不满足类型断言，程序将发生panic报错。
package main

import "fmt"

// 断言  t := i.(T)   t：t就是i接口是T类型的  i：接口   T:类型
// 语法：t,ok:= i.(T) ok 隐藏返回值，如果断言成功 ok就是true、否则就是false

func main() {
	//assertsString("11111111111")
	assertsString(true) // panic: interface conversion: interface {} is bool, not string

}

// 判断一个变量是不是string类型的
func assertsString(i interface{}) {
	// 如果断言失败，则会抛出 panic 恐慌，程序就会停止执行。
	s := i.(string)
	fmt.Println(s)
}
