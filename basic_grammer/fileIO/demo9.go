package main

import (
	"fmt"
	"os"
)

//4. 程序退出
//Exit：让当前程序以给出的状态码code退出。
//一般来说，状态码0表示成功，非0表示出错。
//程序会立刻终止，并且defer的函数不会被执行。

func main() {
	// 正常退出
	fmt.Println("Exiting the program successfully")
	//os.Exit(0)

	// 异常退出（不会执行到这里的代码）
	fmt.Println("This line will not be executed")
	os.Exit(1)
}
