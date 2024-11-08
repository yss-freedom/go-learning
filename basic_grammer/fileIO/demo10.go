package main

import (
	"fmt"
	"os"
)

/*
5. 获取系统信息
Hostname：返回内核提供的主机名。
Getuid：返回调用者的用户ID。
Getgid：返回调用者的组ID。
Getpid：返回调用者所在进程的进程ID。
*/

func main() {
	// 获取主机名
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Hostname:", hostname)

	// 获取用户ID
	uid := os.Getuid()
	fmt.Println("User ID:", uid)

	// 获取组ID
	gid := os.Getgid()
	fmt.Println("Group ID:", gid)

	// 获取进程ID
	pid := os.Getpid()
	fmt.Println("Process ID:", pid)
}
