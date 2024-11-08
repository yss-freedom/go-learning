package main

import "fmt"

func GetData() (int, int) {
	return 10, 20
}

func main() {
	// 终端输出hello world
	//fmt.Println("Hello world!")

	var a int = 27
	fmt.Println(a)
	var d = true
	fmt.Println(d)

	//var intVal int
	//intVal =1

	//var x, y int
	//var c, d int = 1, 2
	//g, h := 123, "hello"

	//var (
	//	a int
	//	b bool
	//)

	a, _ = GetData()
	_, b := GetData()
	fmt.Println(a, b)

}
