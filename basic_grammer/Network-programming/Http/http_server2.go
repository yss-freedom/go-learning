package main

import (
	"fmt"
	"net/http"
)

// 定义发送接收数据函数
func register(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("接收到了注册请求")
	// 解析表单
	// 处理表单的请求, 前端提交表单-后端解析表单
	// func (r *Request) ParseForm() error
	err := req.ParseForm()
	if err != nil {
		return
	}

	// 获取表单参数 post
	username := req.PostForm.Get("username")
	password := req.PostForm.Get("password")
	fmt.Println(username, password)
	// 很多的判断
	// 短信
	// 验证码
	// 存到数据库
	resp.Write([]byte("注册成功"))
}

func main() {
	// HandleFunc http请求的处理函数
	// http程序启动之后是不会停止的，一直监听请求
	// 访问这个url就会触发 helloHandler 函数 （Request） ResponseWriter
	// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	//第一个参数是请求路径，第二个参数是一个函数
	http.HandleFunc("/register", register)
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
