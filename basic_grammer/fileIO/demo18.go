package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
3. 使用Scanner读取标准输入
下面是一个使用bufio.Scanner读取标准输入的示例：
*/

func main() {
	// 读取键盘的输入
	// 键盘的输入，实际上是流 os.Stdin
	inputReader := bufio.NewReader(os.Stdin)
	// delim 到哪里结束读取 以换行符作为分隔符，之言不换行就可以一直读
	readString, _ := inputReader.ReadString('\n')
	fmt.Println("读取键盘输入的信息:", readString)
}
