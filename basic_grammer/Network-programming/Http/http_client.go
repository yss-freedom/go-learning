package main

import (
	"fmt"
	"io"
	"net/http"
)

// 手写客户端访问
func main() {

	/*
	   http.Get()
	*/

	// 一个请求包含 请求方式   请求的url   接收响应结果
	// func Get(url string) (resp *Response, err error)
	resp, _ := http.Get("http://localhost:8080")
	// 通过defer关闭连接 resp.Body 响应的主体
	//通过Body来关闭
	defer resp.Body.Close()

	fmt.Println(resp.Body)
	fmt.Println(resp.Status) // 200 OK
	fmt.Println(resp.Header) // 响应头

	// 接收具体的响应内容，从Body里面获取
	//Body是IO流
	// The response body is streamed on demand as the Body field is read.
	// Body io.ReadCloser
	//循环从Body流中读取
	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("读取出现了错误")
			return
		} else {
			fmt.Println("读取完毕")
			res := string(buf[:n])
			fmt.Println("服务器响应的数据为：", res)
			break
		}

	}

}
