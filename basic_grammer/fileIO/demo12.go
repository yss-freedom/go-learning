package main

import (
	"fmt"
	"os/exec"
)

/*
3. 处理标准输出和标准错误
CombinedOutput
func (c *Cmd) CombinedOutput() ([]byte, error)
说明：运行命令并返回组合到一起的标准输出和标准错误
*/

func main() {

	//查看系统中是否安装有该命令
	cmd_path, err := exec.LookPath("ping")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("系统中有该命令，命令路径为:", cmd_path)

	// 创建一个新的命令，例如运行dir命令
	cmd := exec.Command("ping", "www.baidu.com")

	//获取标准输出
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("标准输出为:", string(out))

}
