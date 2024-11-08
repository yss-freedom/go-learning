package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	// 创建一个http.Client对象
	client := &http.Client{}

	// 构建一个LoginRequest对象
	data := LoginRequest{
		Username: "yss",
		Password: "123456",
	}

	// 将LoginRequest对象转换为JSON字符串
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON编码失败:", err)
		return
	}

	// 创建一个请求对象
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/post", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 设置请求头的内容类型为application/json
	req.Header.Set("Content-Type", "application/json")

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求发送失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取并处理服务器返回的响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}
	fmt.Println("响应状态码:", resp.StatusCode)
	fmt.Println("响应内容:", string(body))
}
