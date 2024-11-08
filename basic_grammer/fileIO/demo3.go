package main

import (
	"fmt"
	"os"
)

func main() {
	// 删除文件
	// func Remove(name string) error
	//文件不存在会报错
	err := os.Remove("example.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File removed successfully")
	}

	// 递归删除目录
	// func RemoveAll(path string) error
	//目录不存在返回nil 不会报错
	//可以删除指定目录下的目录或文件
	err = os.RemoveAll("testdir\\mydir.go")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Directory removed successfully")
	}
}
