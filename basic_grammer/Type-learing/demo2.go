package main

import "fmt"

/*
进阶用法
类型别名不仅限于基础类型，也可以用于复合类型，
如结构体、切片、映射等。然而，重要的是要理解，类型别名和原始类型在类型系统中是被视为等价的。
*/

type Point struct {
	X, Y int
}

type Coord = Point

func main() {
	var p Point = Point{1, 2}
	var d Coord = Coord{3, 4}
	fmt.Println(p)
	// Point 和 Coord 是等价的
	var q Point = d // 正确，因为 Point 和 Coord 是等价的
	fmt.Println(q)  // 输出: {3 4}
}
