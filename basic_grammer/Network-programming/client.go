package main

import (
	"bufio"
	"fmt"
	"net"
)

/*
一个基本的TCP客户端需要完成以下步骤：
建立连接：使用net.Dial函数连接到服务器。
发送和接收数据：通过连接发送数据，并接收服务器的响应。
关闭连接：在完成数据传输后关闭连接。
*/
// client.go
func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("连接错误:", err)
		return
	}
	defer conn.Close()

	// 发送数据
	message := "玩游戏\n"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("发送数据错误:", err)
		return
	}
	fmt.Println("发送数据:", message)

	// 读取响应
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("读取响应错误:", err)
		return
	}
	fmt.Println("收到响应:", response)
}
