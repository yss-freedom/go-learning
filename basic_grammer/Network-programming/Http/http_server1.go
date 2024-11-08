package main

import (
	"fmt"
	"net/http"
)

// 定义发送接收数据函数
func login(resp http.ResponseWriter, req *http.Request) {
	// 模拟数据库中存在一个数据
	mysqlUserData := "admin"
	mysqlPwdData := "123456"

	fmt.Println("接收到了login请求")
	// 拿到请求中的查询参数
	urlData := req.URL.Query()
	username := urlData.Get("username")
	password := urlData.Get("password")

	// 登录逻辑, 将客户端发送的数据和系统数据比对实现登录业务
	if username == mysqlUserData {
		if password == mysqlPwdData {
			resp.Write([]byte("登录成功"))
		} else {
			resp.Write([]byte("密码错误"))
		}
	} else {
		resp.Write([]byte("登录失败"))
	}

}

func main() {
	// HandleFunc http请求的处理函数
	// http程序启动之后是不会停止的，一直监听请求
	// 访问这个url就会触发 helloHandler 函数 （Request） ResponseWriter
	// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	//第一个参数是请求路径，第二个参数是一个函数
	http.HandleFunc("/login", login)
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
