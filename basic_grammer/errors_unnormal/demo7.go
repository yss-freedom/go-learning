package main

import "fmt"

// panic  recover
//
// 出现了panic之后，如果panic语句前面有defer语句，先执行所有的defer语句。
//
// defer : 延迟函数，倒序执行，处理一些问题。
func main() {
	defer fmt.Println("main--1")
	defer fmt.Println("main--2")
	fmt.Println("main--3")
	testPanic(1) // 外部函数，执行完panic之前的所有defer
	//后面的都不执行了
	defer fmt.Println("main--4")
	fmt.Println("main--5")
}
func testPanic(num int) {
	defer fmt.Println("testPanic--1")
	defer fmt.Println("testPanic--2")
	fmt.Println("testPanic--3")
	// 如果在函数中一旦触发了 panic，会终止后面要执行的代码。
	// 如果panic语句前面存在defer，正常按照defer逻辑执行。panic语句后面的defer语句不再执行
	if num == 1 {
		panic("出现预定的异常----panic")
	}
	defer fmt.Println("testPanic--4")
	fmt.Println("testPanic--5")
}
