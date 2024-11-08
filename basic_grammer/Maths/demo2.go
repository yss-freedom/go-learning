package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
1. 设置随机数种子
在生成随机数之前，需要设置随机数种子。math/rand包生成的随机数实际上是伪随机数，其序列由种子决定。如果不设置种子，默认种子为1，每次运行程序生成的随机数序列将相同。因此，通常在生成随机数前需要设置种子，以保证每次运行程序时生成的随机数不同。种子的设置通常通过当前时间的纳秒数来完成。
Go 1.20及以上版本推荐使用rand.New(rand.NewSource(...))，而以下版本则使用rand.Seed(...)。
*/

func main() {
	// 设置随机数种子以获得不同的随机序列
	rand.New(rand.NewSource(time.Now().UnixMicro())) // 版本>=1.20
	//rand.Seed(time.Now().UnixNano()) // 版本<1.20

	randomInt := rand.Intn(100) // 生成0到99之间的随机整数
	fmt.Println(randomInt)

	randomFloat := rand.Float64() // 生成0.0到1.0之间的随机浮点数
	fmt.Println(randomFloat)

	perm := rand.Perm(10) // 生成0到9的随机排列
	fmt.Println(perm)
}
