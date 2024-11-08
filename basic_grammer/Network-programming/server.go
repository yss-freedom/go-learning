package main

import (
	"bufio"
	"fmt"
	"net"
)

/*
一个基本的TCP服务器需要完成以下步骤：
监听端口：使用net.Listen函数监听指定的TCP端口。
接受连接：使用Accept方法接受来自客户端的连接请求。
处理连接：在一个新的goroutine中处理每个连接，以实现并发处理。
*/
func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("开启服务端错误:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接收错误", err)
			continue // 改为continue，可以处理下一个连接

		}
		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	defer conn.Close()

	//读取客户端数据
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("读取错误：", err)
		return
	}
	fmt.Println("Message Received:", message)

	//回发数据
	response := "收到了你的消息了:" + message
	conn.Write([]byte(response))
}
