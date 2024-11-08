package main

import (
	"fmt"
	"os"
)

/*
在现代软件开发中，高效的输入输出（I/O）操作是提高程序性能的关键之一。Go语言提供了丰富的I/O操作接口，使得文件读写、网络通信等任务变得简单而高效。
Go语言的I/O操作主要通过标准库中的io包和os包实现。io包提供了一系列用于输入和输出操作的基本接口和原语。这些接口和原语构成了Go语言中处理I/O操作的基础，支持多种I/O设备的读写操作，包括文件、网络连接、内存缓冲区等。
而os包则提供了对操作系统功能的封装，包括文件操作、进程管理等。
os包是Go语言标准库中的一个重要包，提供了操作系统相关的功能接口。通过os包，我们可以执行各种底层操作，
如读取和写入文件、获取系统信息、执行外部命令等。
*/
//1.1 获取文件信息
//Stat：获取文件信息对象，符号链接会跳转。
//Lstat：获取文件信息对象，符号链接不跳转。
//看下fileinfo接口
// file
// fileInfo
/*
type FileInfo interface {
   Name() string       // base name of the file
   Size() int64        // length in bytes for regular files; system-dependent for others
   Mode() FileMode     // file mode bits ： 权限
   ModTime() time.Time // modification time
   IsDir() bool        // abbreviation for Mode().IsDir()

   // 获取更加详细的文件信息， *syscall.Stat_t  反射来获取
   Sys() any           // underlying data source (can return nil)
*/

func main() {
	// 获取某个文件的状态
	//func Stat(name string) (FileInfo, error)
	fileinfo, err := os.Stat("D:\\goLang\\GoWorks\\src\\projects\\basic_grammer\\errors_unnormal\\demo9.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fileinfo.Name())    // 文件名
	fmt.Println(fileinfo.IsDir())   // 是否目录 false
	fmt.Println(fileinfo.ModTime()) // 文件的修改时间
	fmt.Println(fileinfo.Size())    // 文件大小
	fmt.Println(fileinfo.Mode())    // 权限
}
