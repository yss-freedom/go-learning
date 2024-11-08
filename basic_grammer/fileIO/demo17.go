package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//2. 统计文件中字符、空格、数字和其他字符的数量
//下面是一个统计文件中字符、空格、数字和其他字符数量的示例：

// CharCount 定义结构体
type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	// 打开文件
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	defer file.Close() // 使用defer关闭文件

	// 创建bufio.Reader对象
	reader := bufio.NewReader(file)

	var count CharCount

	// 循环读取文件内容
	for {
		//str, err := reader.ReadString('\n') // 读取字符串，以换行符为分隔符
		str, _, err := reader.ReadLine() // 按行读取，最后一行没有换行符也能读到
		if err == io.EOF {               // 判断文件是否结尾
			break
		}
		if err != nil { // 判断错误，打印错误
			fmt.Printf("read file failed, err:%v", err)
			break
		}

		//runeArr := []rune(str) // 将字符串转为切片类型
		runeArr := str // 将字符串转为切片类型
		fmt.Println("得到的数据是", string(runeArr))
		for _, v := range runeArr { // 循环读取数组
			switch {
			case v >= 'a' && v <= 'z': // 判断是不是小写字母
				fallthrough
			case v >= 'A' && v <= 'Z': // 判断是不是大写字母
				count.ChCount++
			case v == ' ' || v == '\t': // 判断是不是空格
				count.SpaceCount++
			case v >= '0' && v <= '9': // 判断是不是数字
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}

	fmt.Printf("char count:%d\n", count.ChCount)
	fmt.Printf("num count:%d\n", count.NumCount)
	fmt.Printf("space count:%d\n", count.SpaceCount)
	fmt.Printf("other count:%d\n", count.OtherCount)
}
