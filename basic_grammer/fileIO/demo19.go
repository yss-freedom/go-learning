package main

import (
	"bufio"
	"fmt"
	"os"
)

// bufio的写操作主要通过bufio.Writer对象实现，它提供了多种写入数据的方法，
// 如Write、WriteString、Flush等。
func main() {
	// 打开文件句柄，以写的方式打开，如果文件不存在则创建
	file, err := os.OpenFile("D:\\goLang\\GoWorks\\src\\projects\\basic_grammer\\example.txt",
		os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer file.Close() // 使用的defer语句关闭文件

	// 创建bufio.Writer对象
	fileWrite := bufio.NewWriter(file)

	inputStr := "Hello World ...\n"
	for i := 0; i < 10; i++ { // 循环写入数据
		fileWrite.WriteString(inputStr)
	}
	// 发现并没有写出到文件，是留在了缓冲区，所以我们需要调用 flush 刷新缓冲区
	// 必须手动刷新才能写入进文件
	fileWrite.Flush() // 刷新缓冲区，将数据写入文件
}
