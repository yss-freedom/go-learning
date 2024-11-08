package main

import (
	"fmt"
	"os"
)

//1.2 创建文件
//Create：创建文件，如果文件存在，则清空原文件

func main() {
	//创建文件，文件不存在就创建文件，如果文件存在，则清空原文件。
	// func Create(name string) (*File, error)

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//关闭文件，延迟执行
	defer file.Close()

	//向文件中写内容
	// func (f *File) WriteString(s string) (n int, err error)
	_, err = file.WriteString("Hello, World!")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File created successfully")
}
