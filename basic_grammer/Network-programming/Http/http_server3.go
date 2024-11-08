package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Result struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 定义发送接收数据函数
func postHandler(resp http.ResponseWriter, req *http.Request) {
	// 确保处理的是POST请求
	if req.Method != "POST" {
		http.Error(resp, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// 读取请求体中的数据
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(resp, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	// 打印请求体中的数据
	fmt.Printf("Received POST request with body: %s\n", string(body))

	//将json字符串反序列化为原生结构体
	var result Result
	if err := json.Unmarshal(body, &result); err != nil {
		http.Error(resp, "Error parsing request body", http.StatusInternalServerError)
	}
	fmt.Println("客户端发来的数据:", result)
	fmt.Printf("客户端发来的数据:%+v\n", result)
	fmt.Printf("数据类型%T\n", result)

	// 发送响应给客户端
	// Received your POST request. This is the response!
	resp.Write([]byte("服务端接收到数据"))
}

func main() {
	// HandleFunc http请求的处理函数
	// http程序启动之后是不会停止的，一直监听请求
	// 访问这个url就会触发 helloHandler 函数 （Request） ResponseWriter
	// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	//第一个参数是请求路径，第二个参数是一个函数
	http.HandleFunc("/post", postHandler)
	fmt.Println("Starting server at :8080")
	// func ListenAndServe(addr string, handler Handler) error
	// ListenAndServe监听TCP地址addr，并且会使用handler参数调用Serve函数处理接收到的连接。handler参数一般会设为nil，此时会使用DefaultServeMux。
	//如果用户自定义实现了Handler，那么根据相应路径在map中查询到相对应的Handler，然后再调用用户自定义的ServeHTTP处理请求。
	//如果用户没有自定义Handler，只注册了对应处理函数(使用了http.HandleFunc)，那么就会根据默认DefaultServeMux去map查询到这个函数类型Handler，然后再调用ServeHTTP处理函数。
	// 开启监听程序的代码是放在main方法的最后一行的。
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
