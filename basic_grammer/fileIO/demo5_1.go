package main

import (
	"fmt"
	"os"
)

/*
打开文件常用方式二：os.OpenFile
在Go语言中，文件的读写操作是使用非常频繁的功能之一。os.OpenFile函数是Go标准库中用于打开文件的强大工具，它允许以多种模式打开文件，并可以设置文件的权限。
os.OpenFile函数简介
os.OpenFile函数用法如下：
func OpenFile(name string, flag int, perm FileMode) (File, error)
name：要打开的文件名或路径。
flag：打开文件的模式，可以是多个标志的按位或组合。常见的标志包括：
os.O_RDONLY：只读模式。
os.O_WRONLY：只写模式。
os.O_RDWR：读写模式。
os.O_APPEND：追加模式，在文件末尾写入数据而不覆盖原有数据。只要有这个参数，就会采用追加模式
os.O_CREATE：如果文件不存在则创建文件。
os.O_EXCL：与O_CREATE一起使用，可执行权限。
os.O_SYNC：同步模式，将文件更改刷新到磁盘。
os.O_TRUNC：截断模式，清空文件内容。
perm：文件权限，表示文件的读写权限，默认为0666（即所有用户都可读写）。
返回值是一个指向打开文件的File对象和一个可能出现的错误。如果文件打开成功，则可以通过该文件对象进行读取和写入操作。
*/
//使用os.OpenFile写文件
//以下是一个使用os.OpenFile函数写文件的示例：

func main() {
	// 以追加、创建、读写模式打开文件
	file, err := os.OpenFile("./example.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close() // 确保在函数执行完毕后关闭文件

	// 写入文件
	_, err = file.Write([]byte("Hello, Go!\n"))
	if err != nil {
		fmt.Println("写入文件失败:", err)
		return
	}

	fmt.Println("文件写入成功")
}
