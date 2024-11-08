package main

import (
	"fmt"
	"os"
)

func main() {
	// 重命名文件
	// func Rename(oldpath, newpath string) error
	err := os.Rename("yss.txt", "yssnb.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File renamed successfully")
	}

	// 移动文件
	err = os.Rename("yssnb.txt", "D:\\gofiles\\fileIO\\yssnb.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File moved successfully")
	}
}
