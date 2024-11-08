package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// 构建HTTP客户端
	client := &http.Client{}

	// 定义POST URL
	myurl := "http://127.0.0.1:8080/register"

	// 定义HTTP负载
	data := url.Values{}
	data.Set("username", "yssNB")
	data.Set("password", "123456")
	// 将url.Values编码为表单格式的字符串，并创建一个Reader
	payload := strings.NewReader(data.Encode())

	// 发送POST请求
	// func NewRequest(method, url string, body io.Reader) (*Request, error)
	req, err := http.NewRequest("POST", myurl, payload)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置HTTP消息头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 执行HTTP请求
	// func (c *Client) Do(req *Request) (*Response, error)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应状态码
	fmt.Println("Response Status:", resp.Status)

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 打印响应体
	fmt.Println("Response Body:", string(body))
}
