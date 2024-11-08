package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
bufio的主要作用是减少I/O操作的次数，提供读写性能。其工作原理是通过在内存中维护一个缓冲区，先将数据读入缓冲区，
再从缓冲区中读取数据，从而减少对底层I/O设备的访问次数。
bufio的主要对象是缓冲区，操作主要有两个：读和写。读操作时，
如果缓冲区为空，则从文件读取数据填满缓冲区；如果缓冲区不为空，则从缓冲区读取数据。写操作时，如果缓冲区未满，
则将数据写入缓冲区；如果缓冲区已满，则将缓冲区的数据写入文件，并继续写入剩余的数据。
*/
//bufio的读操作主要通过bufio.Reader对象实现，它提供了多种读取数据的方法，
//如Read、ReadLine、ReadString等。
/*
在这个示例中，我们首先使用os.Open函数打开文件，然后创建一个bufio.Reader对象，
通过循环调用ReadString方法读取文件内容，每次读取一行。
当遇到文件结尾时，ReadString方法会返回一个io.EOF错误，我们据此退出循环。
*/
func main() {
	// 打开文件
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	defer file.Close() // 使用defer关闭文件

	// 创建bufio.Reader对象。借助bufio来读取文件内容
	// func NewReader(rd io.Reader) *Reader
	reader := bufio.NewReader(file)

	// 循环读取文件内容
	for {

		// 通过Read方法Read()读取文件
		//buf := make([]byte, 1024)
		//n, err := bufioReader.Read(buf)
		//fmt.Println("读取到了多少个字节：", n)
		//读取到的内容
		//fmt.Println("读取到的内容:",string(buf[:n]))

		//通过ReadString方法读取文件
		// func (b *Reader) ReadString(delim byte) (string, error)
		//参数是以什么作为分隔符。读取得到的是字符串
		//注意，最后一行后面要是没有换行符，读取不到
		str, err := reader.ReadString('\n') // 读取字符串，以换行符为分隔符
		if err == io.EOF {                  // 判断文件是否结尾
			break
		} else if err != nil { // 判断错误，打印错误
			fmt.Printf("read file failed, err:%v", err)
			break
		}
		//循环中逐行打印出读取的内容
		fmt.Printf("read string Successfully: %s\n", str) // 打印读取的文件内容
	}
}
