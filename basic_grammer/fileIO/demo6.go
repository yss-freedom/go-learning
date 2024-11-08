package main

import (
	"fmt"
	"os"
)

/*
使用os.File对象的Read方法可以从文件中读取数据。Read方法会将数据读取到一个字节切片中，并返回读取的字节数和可能发生的错误。
需要借助 file.Read([]byte)读取 ，将file中的数据读取到 []byte 中， n，err n读取到的行数，err 错误，EOF错误，就代表文件读取完毕了
*/
//一直调用read，就代表光标往后移动…
//以下是一个使用os.Open函数读文件的示例：

// 读取文件数据
func main() {
	// 我们习惯于在建立连接时候通过defer来关闭连接，保证程序不会出任何问题，或者忘记关闭
	// 建立连接
	file, _ := os.Open("example.txt")
	// 关闭连接
	defer file.Close()

	// 读代码 ,Go 的错误机制，让我们专心可以写业务代码。

	// 1、创建一个容器 （二进制文本文件--0100101010 => 读取流到一个容器 => 读取容器的数据）
	//一次性读取的内容长度为切片的长度，切片长度为2，容量为1024
	bs := make([]byte, 2, 1024) // 缓冲区，可以接受我们读取的数据 这里设置的是一次性读取两个字节
	fmt.Printf("字节切片数据类型%T\n", bs)
	// 2、读取到缓冲区中。 // 汉字一个汉字占 3个字节
	// func (f *File) Read(b []byte) (n int, err error) 返回读到的字节数和错误
	// At end of file, Read returns 0, io.EOF. 读到文件末尾返回0和io.EOF
	n, err := file.Read(bs)
	fmt.Println(n)
	fmt.Println(err)
	//每次读取光标移动切片的长度，将读取的内容存到切片中。如果读取的文件是文本数据，可以通过string方法将切片中的内容转换为字符串
	//如果是MP4呀，音乐，图片等文件，就不要用string转换了
	fmt.Println(string(bs)) // 读取到的字符串  ab

	//继续往下读
	// 光标不停的向下去指向，读取出来的内容就存到我们的容器中。
	file.Read(bs)
	fmt.Println(string(bs)) // 读取到的字符串  cd

	file.Read(bs)
	fmt.Println(string(bs)) // 读取到的字符串  e

	//读到文件末尾，得到的n是0，err是EOF
	n, err = file.Read(bs)
	fmt.Println(n)
	fmt.Println(err) // EOF ，读取到了文件末尾。err就会返回EOF。
	//最后读到了e，将原来切片中的cd的c覆盖了，得到的最终结果是ed
	fmt.Println(string(bs)) // 读取到的字符串

}
