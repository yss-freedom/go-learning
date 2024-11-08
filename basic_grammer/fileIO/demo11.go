package main

import (
	"fmt"
	"os/exec"
)

//6. 执行外部命令
//os/exec包用于执行外部命令，它提供了更高级别的接口来创建和管理进程。

func main() {

	//查看系统中是否安装有该命令
	cmd_path, err := exec.LookPath("ipconfig")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("系统中有该命令，命令路径为:", cmd_path)
	//Command每个参数都不能包含空格。多个参数，用引号包裹，逗号隔开
	cmd := exec.Command("ipconfig")
	// 如果不需要命令的输出，直接调用cmd.Run()即可

	//err2 := cmd.Run()
	//if err2 != nil {
	//	fmt.Println("Error:", err2)
	//} else {
	//	fmt.Println("Command executed successfully")
	//}
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Command output:", string(output))
	}

}
