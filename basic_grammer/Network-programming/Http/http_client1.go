package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	// 复杂请求
	urlStr := "http://127.0.0.1:8080/login" // ?

	// 参数如何拼接到url上，参数封装为数据url.Values{}
	data := url.Values{}
	//通过url.Values.Set()方法，将参数拼接到url上
	data.Set("username", "admin")  // ?
	data.Set("password", "123456") // ?

	// 将url字符串转化为url对象，并给携带数据
	// func ParseRequestURI(rawURL string) (*URL, error)
	urlNew, _ := url.ParseRequestURI(urlStr)
	urlNew.RawQuery = data.Encode()
	// http://127.0.0.1:8080/login?password=123456&username=admin
	// ? get的传参，多个参数之间使用 & 连接
	fmt.Println(urlNew)
	//查看类型
	fmt.Printf("%T \n", urlNew)

	// 发请求,参数是一个字符串地址
	resp, _ := http.Get(urlNew.String())
	defer resp.Body.Close()
	// 读取resp响应信息，一次性读完，返回buf
	// func ReadAll(r Reader) ([]byte, error)
	buf, _ := io.ReadAll(resp.Body)
	fmt.Println(string(buf))
}
