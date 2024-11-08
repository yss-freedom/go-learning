package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

/*
断点续传是在下载或上传时，将下载或上传任务（一个文件或一个压缩包）人为的划分为几个部分，每一个部分采用一个线程进行上传或下载，如果碰到网络故障，
可以从已经上传或下载的部分开始继续上传或者下载未完成的部分，而没有必要从头开始上传或者下载。
go语言实现断点续传的思路：
使用临时文件记录中断位置.
1.文件上传时,先创建上传一个新的文件
2.创建记录中断位置的临时文件，需要记住上一次传递了多少数据、temp.txt
3.设置文件读写偏移量，如果被暂停或者中断了，我们就可以读取这个temp.txt的记录，恢复上传
4.上传完成后,删除临时文件
*/

// 断点续传
func main() {

	// 传输源文件地址
	srcFile := "D:\\gofiles\\fileIO\\jdk-8u171-linux-x64.tar.gz"
	// 传输的目标位置
	destFile := "D:\\gofiles\\files\\jdk-8u171-linux-x64.tar.gz"
	// 临时记录文件
	tempFile := "D:\\gofiles\\temp.txt"

	// 创建对应的file对象，连接起来
	file1, _ := os.Open(srcFile)
	file2, _ := os.OpenFile(destFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	file3, _ := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer file1.Close()
	defer file2.Close()
	fmt.Println("file1/2/3 文件连接建立完毕")

	// 1、读取temp.txt
	file3.Seek(0, io.SeekStart)
	buf := make([]byte, 1024)
	n, _ := file3.Read(buf) //这里的n是读取file3中的字符的个数，比如1024，得到的n是4.所以要借助string转
	//查看返回的n的数据类型
	fmt.Printf("查看n的数据类型%T\n", n)
	fmt.Println("n的值为", n)

	// 2、先转换成string，然后再转换成数字。
	countStr := string(buf[:n])
	count, _ := strconv.ParseInt(countStr, 10, 64)
	fmt.Println("temp.txt中记录的值为：", count) // 5120

	// 3、设置读写的偏移量,offset是int64数据类型
	file1.Seek(count, io.SeekStart)
	file2.Seek(count, io.SeekStart)
	fmt.Println("file1/2 光标已经移动到了目标位置")

	// 4、开始读写（复制、上传）
	bufData := make([]byte, 204800)
	// 5、需要记录读取了多少个字节
	total := int(count)

	for {
		fmt.Println("传输了,", total)
		// 读取数据
		readNum, err := file1.Read(bufData)
		if err == io.EOF || readNum == 0 { // file1 读取完毕了
			fmt.Println("文件传输完毕了")
			//上传完文件再关闭临时文件file3
			file3.Close()
			os.Remove(tempFile)
			break
		}

		// 向目标文件中写入数据，返回写的字节数和错误
		writeNum, err := file2.Write(bufData[:readNum])

		// 将写入数据放到 total中, 在这里total 就是传输的进度
		total = total + writeNum

		// temp.txt 存放临时记录数据
		file3.Seek(0, io.SeekStart) // 将光标重置到开头
		//将数字转换成字符串写入，这里total逐渐变大不存在覆盖不完的问题。如果存在覆盖不完问题，使用os.Truncate(fileName, 0)来清空文件内容
		// os.Truncate(fileName, 0) 截取指定长度字节的内容，其余内容会被删除
		file3.WriteString(strconv.Itoa(total))

		//模拟故障发生
		//if total > 1000000 {
		//	panic("出故障了")
		//}

	}

}
