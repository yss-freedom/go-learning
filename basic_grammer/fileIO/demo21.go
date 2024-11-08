package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
在bufio包中，Reader和Writer对象都有默认的缓冲区大小，但你也可以根据需要自定义缓冲区大小。
1. 默认缓冲区大小
对于bufio.Reader和bufio.Writer，Go语言标准库为它们提供了默认的缓冲区大小。
对于bufio.Reader，默认缓冲区大小通常是4096字节（4KB）；
对于bufio.Writer，默认缓冲区大小也是4096字节（但在某些实现中可能略有不同，
具体取决于底层操作系统的I/O性能）。

2. 自定义缓冲区大小
如果你需要更大的缓冲区来提高性能，或者更小的缓冲区来减少内存使用，
你可以使用bufio.NewReaderSize和bufio.NewWriterSize函数来创建具有自定义缓冲区大小的Reader和Writer对象。
*/

/*
在这个示例中，我们创建了一个具有8KB缓冲区大小的bufio.Writer对象，
并将其用于将数据写入文件。注意，在写入数据后，
我们调用了Flush方法来确保缓冲区中的数据被写入文件。
*/
func main() {
	// 打开文件
	file, err := os.OpenFile("D:\\goLang\\GoWorks\\src\\projects\\basic_grammer\\example.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer file.Close() // 使用defer语句关闭文件

	// 自定义缓冲区大小（例如：8KB）
	bufferSize := 8 * 1024

	// 创建具有自定义缓冲区大小的bufio.Writer对象
	fileWrite := bufio.NewWriterSize(file, bufferSize)

	inputStr := "This is a line with custom buffer size.\n"
	for i := 0; i < 10; i++ { // 循环写入数据
		_, err := fileWrite.WriteString(inputStr)
		if err != nil {
			fmt.Println("write file error:", err)
			return
		}
	}

	// 刷新缓冲区，将数据写入文件
	err = fileWrite.Flush()
	if err != nil {
		fmt.Println("flush buffer error:", err)
		return
	}

	fmt.Println("Data written to file successfully with custom buffer size.")
}
