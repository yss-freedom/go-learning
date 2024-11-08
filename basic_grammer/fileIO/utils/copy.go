package utils

import (
	"fmt"
	"io"
	"os"
)

//自定义的文件复制实现
// Copy 方法需要参数为：source 源文件 ,destination 目标文件，缓冲区大小

func Copy(source, destination string, bufferSize int) {

	// 读取文件，读文件open方法就可以
	sourceFile, err := os.Open(source)
	if err != nil {
		fmt.Println("Open错误:", err)
	}

	// 输出文件 O_WRONLY , O_CREATE 如果不不存在，则会创建
	destinationFile, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("OpenFile错误:", err)
	}
	// 关闭连接
	defer sourceFile.Close()
	defer destinationFile.Close()

	// 专注业务代码，拷贝
	// 创建缓冲区大小
	buf := make([]byte, bufferSize)
	// 循环读取，写入
	for {
		//一次读取设置的缓冲区大小的数据
		n, err := sourceFile.Read(buf)
		// 循环结束标志
		if n == 0 || err == io.EOF {
			fmt.Println("读取完毕源文件,复制完毕")
			break
		} else if err != nil {
			fmt.Println("读取错误:", err)
			return // 错误之后，必须要return终止函数执行。
		}
		// 将缓冲区的东西写出到目标文件
		// 因为循环写入过程中没有关闭文件，所以可以持续写入
		//根据实际读取的长度写入，防止最后一次读取时，读取的少于缓冲区长度
		_, err = destinationFile.Write(buf[:n])
		if err != nil {
			fmt.Println("写出错误：", err)
		}
	}

}

//2. 使用系统自带的io.Copy()方法复制
// Copy2 使用系统自带的io.Copy方法

func Copy2(source, destination string) {
	// 读取文件
	sourceFile, err := os.Open(source)
	if err != nil {
		fmt.Println("Open错误:", err)
	}
	// 输出文件 O_WRONLY , O_CREATE 如果不不存在，则会创建
	destinationFile, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("OpenFile错误:", err)
	}
	// 关闭
	defer sourceFile.Close()
	defer destinationFile.Close()

	// 具体的实现
	// 不用我们自己写循环读取写入了，直接调用系统的io.Copy()方法
	// func Copy(dst Writer, src Reader) (written int64, err error) 返回写入的字节数和错误
	written, err := io.Copy(destinationFile, sourceFile)
	if err != nil {
		fmt.Println("复制时出现错误", err)
		return
	}
	fmt.Println("文件的字节大小:", written)
}

// Copy3 使用os.ReadFile和os.WriteFile来实现读取写入，从而实现复制
func Copy3(source, destination string) {
	// func ReadFile(name string) ([]byte, error)
	fileBuffer, _ := os.ReadFile(source)
	//func WriteFile(name string, data []byte, perm FileMode) error
	err := os.WriteFile(destination, fileBuffer, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("文件复制完成")

}

// Copy4 io.CopyBuffer() 适合大文件
func Copy4(source, destination string, bufferSize int) {
	// 读取文件
	sourceFile, err := os.Open(source)
	if err != nil {
		fmt.Println("Open错误:", err)
	}
	// 输出文件 O_WRONLY , O_CREATE 如果不不存在，则会创建
	destinationFile, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("OpenFile错误:", err)
	}
	// 关闭
	defer sourceFile.Close()
	defer destinationFile.Close()

	// 具体的实现
	// 创建缓冲区大小
	buf := make([]byte, bufferSize)
	// 不用我们自己写循环读取写入了，直接调用系统的io.Copy()方法
	//io.copy这个适合小文件，大文件使用io.CopyBuffer
	// func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
	written, err := io.CopyBuffer(destinationFile, sourceFile, buf)
	if err != nil {
		fmt.Println("复制时出现错误", err)
		return
	}
	fmt.Println("文件的字节大小:", written)
}
