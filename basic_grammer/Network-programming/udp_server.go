package main

import (
	"fmt"
	"net"
	"os"
)

/*
1. 创建UDP服务器
一个基本的UDP服务器需要完成以下步骤：
监听端口：使用net.ListenUDP函数监听指定的UDP端口。
接受数据：使用ReadFromUDP方法接收来自客户端的数据。
发送响应：使用WriteToUDP方法发送响应到客户端。
*/

func main() {
	// 解析UDP地址
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println("解析地址错误:", err)
		os.Exit(1)
	}

	// 监听UDP端口
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("监听端口错误:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("UDP服务器启动，监听端口 8088")

	buffer := make([]byte, 1024)

	for {
		// 读取数据
		n, udpaddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("读取数据错误:", err)
			continue
		}

		fmt.Printf("Received %d bytes from %s: %s\n", n, udpaddr, string(buffer[:n]))

		// 发送回复
		response := []byte("收到你的消息了!")
		_, err = conn.WriteToUDP(response, udpaddr)
		if err != nil {
			fmt.Println("发送回复错误:", err)
		}
	}
}
