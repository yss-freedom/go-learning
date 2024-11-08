package main

import (
	"bufio"
	"fmt"
)

//下面是一个自定义io.Writer对象进行缓冲写的示例：
// 自定义一个io.Writer对象
/*
我们自定义了一个StringWriter对象，它实现了io.Writer接口的Write方法。
然后，我们创建一个bufio.Writer对象，并传入StringWriter对象。
接着，我们调用WriteString方法写入数据，并调用Flush方法刷新缓冲区，将数据写入StringWriter对象。
*/

type StringWriter struct{}

func (s StringWriter) Write(p []byte) (n int, err error) {
	// 这里简单地将数据写入一个字符串变量（实际使用时，可以将其写入内存、文件等）
	fmt.Print(string(p))
	return len(p), nil
}

func main() {
	// 创建一个StringWriter对象
	sw := StringWriter{}

	// 创建一个bufio.Writer对象，并传入StringWriter对象
	writer := bufio.NewWriterSize(sw, 10)

	// 写入数据
	writer.WriteString("Hello, ")
	writer.WriteString("world!\n")

	// 刷新缓冲区，将数据写入StringWriter对象
	writer.Flush()
}
