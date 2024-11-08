package main

import (
	"fmt"
	"os"
)

/*
写的话，要考虑文件的权限问题，大致流程如下
建立连接 （设置权限：可读可写，扩充这个文件的append） os.OpenFile
关闭连接
写入文件常用的方法
file.Write
file.WriteString
*/

func main() {

	fileName := "example.txt"
	// 权限：如果我们要向一个文件中追加内容， O_APPEND, 如果没有，就是从头开始写
	file, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_RDONLY|os.O_APPEND, os.ModePerm)
	defer file.Close()

	// 操作 字节切片写入
	bs := []byte{65, 66, 67, 68, 69} // A B C D E  ASCII编码
	// func (f *File) Write(b []byte) (n int, err error)
	//write方法 写入的是二进制数据字节流，即字节切片
	n, err := file.Write(bs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	// string类型的写入，可以直接将字符串写入到文件中
	n, err = file.WriteString("hhahahahah哈哈哈哈哈哈哈")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}

//如果打开文件的方式不包含os.O_APPEND 则是将源文件覆盖写。
//os.O_APPEND的意思是，在打开文件时，将光标移动到文件末尾
