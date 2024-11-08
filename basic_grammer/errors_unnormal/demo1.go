package main

import (
	"fmt"
	"os"
)

// 错误是开发中必须要思考的问题
// - 某些系统错误 ，文件被占用，网络有延迟
// - 人为错误：核心就是一些不正常的用户会怎么来给你传递参数，sql注入
func main() {
	//工作目录是指在执行程序时，操作系统认为当前正在工作的目录。
	//一般情况下，我们通过在终端或文件浏览器中打开文件时所处的目录就是当前工作目录。
	//os.Getwd()只能获取当前项目的工作目录，并不是当前文件所在路径
	path, _ := os.Getwd()
	fmt.Println("Path:", path)
	err := os.Chdir("F:\\goworks\\src\\jingtian\\yufa\\错误与异常")
	if err != nil {
		return
	}
	//再次查看路径
	path2, _ := os.Getwd()
	fmt.Println("再次查看Path:", path2)
	//打开一个文件 os 系统包，所有可以用鼠标和键盘能执行的事件，都可以用程序实现
	//os.Open()返回 一个文件对象和错误error
	// func Open(name string) (*File, error)
	file, err2 := os.Open("aaa.txt")
	// 在开发中，我们需要思考这个错误的类型  PathError
	// 1、文件不存在 err
	// 2、文件被占用 err
	// 3、文件被损耗 err
	// 调用方法后，出现错误，需要解决
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	fmt.Println(file.Name())
}

// 在实际工程项目中，
// 我们希望通过程序的错误信息快速定位问题，但是又不喜欢错误处理代码写的冗余而又啰嗦。
// Go语言没有提供像Java. c#语言中的try...catch异常处理方式，
// 而是通过函数返回值逐层往上抛, 如果没有人处理这个错误，程序就终止 panic

// 这种设计,鼓励工程师在代码中显式的检查错误，而非忽略错误。
// 好处就是避免漏掉本应处理的错误。但是带来一个弊端，让代码繁琐。

// Go中的错误也是一种类型。错误用内置的error类型表示。就像其他类型的，如int, float64。
// 错误值可以存储在变量中，从函数中返回，传递参数 等等。
