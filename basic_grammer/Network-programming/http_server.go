package main

import (
	"fmt"
	"net/http"
)

/*
1. HTTP服务器
Go语言内置的net/http包提供了简单的HTTP服务器实现。以下是一个基本的HTTP服务器示例
*/

// 定义发送接收数据函数
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
	//也可以写数据响应给客户端
	// 给浏览器响应数据，这里响应的时字符串，标签不生效
	w.Write([]byte("<h1>感谢大家来修仙的世界!</h1>"))
}

func main() {
	// 访问这个url就会触发 helloHandler 函数 （Request） ResponseWriter
	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
