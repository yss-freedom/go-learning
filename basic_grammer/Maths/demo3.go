package main

import (
	"fmt"
	"math"
)

// 1. 计算平均值和标准偏差
// 标准偏差是衡量数据分散程度的一种常用方法。我们可以使用math包中的函数来计算一组数据的平均值和标准偏差。
// 计算平均值
func mean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

// 计算标准偏差
func stddev(data []float64) float64 {
	m := mean(data)
	sum := 0.0
	for _, value := range data {
		sum += math.Pow(value-m, 2)
	}
	variance := sum / float64(len(data))
	return math.Sqrt(variance)
}

func main() {
	data := []float64{2.3, 3.5, 2.8, 4.1, 5.0, 3.3}
	fmt.Printf("Mean: %.2f\n", mean(data))
	fmt.Printf("Standard Deviation: %.2f\n", stddev(data))
}
