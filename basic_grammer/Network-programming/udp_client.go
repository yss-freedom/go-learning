package main

import (
	"fmt"
	"net"
	"os"
)

/*
一个基本的UDP客户端需要完成以下步骤：
建立连接：使用net.ResolveUDPAddr和net.DialUDP函数连接到服务器。
发送数据：使用Write方法发送数据到服务器。
接收响应：使用ReadFromUDP方法接收服务器的响应。
*/

func main() {
	// 解析UDP地址
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println("解析地址错误:", err)
		os.Exit(1)
	}

	// 建立连接
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("连接错误:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 发送数据
	// 客户端发送用conn.Write(message)
	message := []byte("Hello, UDP Server!")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("发送数据错误:", err)
		os.Exit(1)
	}
	fmt.Println("发送数据:", string(message))

	// 读取响应
	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("读取响应错误:", err)
		os.Exit(1)
	}
	fmt.Println("收到响应:", string(buffer[:n]))
}
