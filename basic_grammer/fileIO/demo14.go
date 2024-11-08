package main

import (
	"fmt"
	"io"
	"os"
)

/*
以下是一个结合具体案例的示例代码，展示了如何使用Seeker接口来定位文件位置并进行读写操作。
File对象实现了Seeker接口
*/

func main() {
	// 读取文件
	file, _ := os.OpenFile("D:\\goLang\\GoWorks\\src\\projects\\basic_grammer\\fileIO\\a.txt", os.O_RDWR, os.ModePerm)
	// defer close
	defer file.Close()

	// 测试seek
	// 相对开始位置。io.SeekStart
	// 相对于文件末尾， io.SeekEnd
	// func (f *File) Seek(offset int64, whence int) (ret int64, err error)
	//相对于开始位置，光标偏移两个字节
	file.Seek(2, io.SeekStart)
	//创建一个字节的buffer
	buf := []byte{0}
	file.Read(buf)

	fmt.Println(string(buf)) // n

	//Read读了一个字节， 光标现在在3这个位置
	// 相对于当前位置
	file.Seek(3, io.SeekCurrent)
	file.Read(buf)

	fmt.Println(string(buf)) // a

	// 在结尾追加内容
	//相对于结束位置偏移0，光标就设在了结束位置
	file.Seek(0, io.SeekEnd)
	//写入内容
	file.WriteString("hahahaha")
}
