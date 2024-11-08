package main

import (
	"bufio"
	"fmt"
	"os"
)

//对于bufio.Reader，你也可以使用类似的方法来创建具有自定义缓冲区大小的对象：

func main() {
	// 打开文件
	file, err := os.Open("D:\\goLang\\GoWorks\\src\\projects\\basic_grammer\\example.txt")
	if err != nil {
		fmt.Println("read file error:", err)
		return
	}
	defer file.Close() // 使用defer语句关闭文件

	// 自定义缓冲区大小（例如：8KB）
	bufferSize := 8 * 1024

	// 创建具有自定义缓冲区大小的bufio.Reader对象
	fileRead := bufio.NewReaderSize(file, bufferSize)

	// 读取文件内容（这里只是简单地读取并打印每一行）
	for {
		line, err := fileRead.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" { // 判断是否到达文件末尾
				break
			}
			fmt.Println("read file error:", err)
			return
		}
		fmt.Print(line) // 打印读取到的行
	}

	fmt.Println("File read successfully with custom buffer size.")
}
