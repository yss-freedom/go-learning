package main

import (
	"fmt"
	"runtime"
)

// 获取系统的信息runtime
func main() {
	// 获取goRoot目录 : 找到指定目录，存放一些项目信息。
	fmt.Println("GoRoot Path:", runtime.GOROOT())
	// 获取操作系统  windows ，可以根据操作系统类型判断盘符分隔符。 “\\”  “/”
	fmt.Println("System:", runtime.GOOS)
	// 获取cpu数量 12， 可以尝试做一些系统优化，开启更大的栈空间。
	fmt.Println("Cpu num:", runtime.NumCPU())
}
