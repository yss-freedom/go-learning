package main

import (
	"fmt"
	"math"
)

/*
2. 使用三角函数解决实际问题
三角函数在解决例如距离计算、角度转换等实际问题中非常有用。
例如，计算一个直角三角形的斜边长度可以使用math包中的Sin、Cos函数。
*/

// 计算斜边长度
func hypotenuse(opposite, angle float64) float64 {
	return opposite / math.Sin(angle)
}

func main() {
	angle := math.Pi / 4 // 45度角
	opposite := 10.0     // 对边长度
	fmt.Printf("Hypotenuse length: %.2f\n", hypotenuse(opposite, angle))
}
