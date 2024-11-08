package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	// 目标URL
	urlStr := "http://127.0.0.1:8080/register"

	// 构建表单数据
	formData := url.Values{}
	formData.Set("username", "admin")
	formData.Set("password", "123456")

	// 发送POST请求
	// func PostForm(url string, data url.Values) (resp *Response, err error)
	resp, err := http.PostForm(urlStr, formData)
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 打印响应状态码和响应体
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}
