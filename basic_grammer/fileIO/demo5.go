package main

import (
	"fmt"
	"os"
)

//使用os.Open函数可以打开一个文件，返回一个*os.File类型的文件对象。如果文件打开失败，会返回一个错误

func main() {
	//func Open(name string) (*File, error)
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	fmt.Println("File opened successfully")
	//打印打开的文件，是个内存地址
	fmt.Println("Reading file:", file)
}
