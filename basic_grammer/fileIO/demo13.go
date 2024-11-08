package main

import "basic_grammer/fileIO/utils"

//调用我们的copy方法，实现文件copy

func main() {
	source := "D:\\gofiles\\fileIO\\yssnb.txt"
	dest := "D:\\gofiles\\files\\yssnb.txt"

	utils.Copy4(source, dest, 1024)
}
