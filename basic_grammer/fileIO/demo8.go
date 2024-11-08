package main

import (
	"fmt"
	"os"
)

//2.1 创建目录
//Mkdir：使用指定的权限和名称创建一个目录，目录存在会报错。
//MkdirAll：递归创建目录，包括任何必要的上级目录，目录存在不报错。

func main() {
	// 创建单个目录
	// func Mkdir(name string, perm FileMode) error
	// os.ModePerm 是0777的权限
	//mkidr创建目录，目录存在的情况下，会报错 Cannot create a file when that file already exists.
	err := os.Mkdir("singledir", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Single directory created successfully")
	}

	// 递归创建目录，目录存在不报错
	// func MkdirAll(path string, perm FileMode) error
	err = os.MkdirAll("nested/dirs/test", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Nested directories created successfully")
	}
}
