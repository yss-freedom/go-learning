package main

/*
panic 和 recover
与错误处理不同，Go语言中的异常处理是通过panic和recover关键字来实现的。
panic用于表示程序中的严重错误，它会导致程序中断执行；
recover用于捕获panic，并恢复程序的正常执行流程。

什么时候会发生恐慌panic，我们不知道什么时候会报错
程序运行的时候会发生panic
手动抛出panic。我们在一些能够预知到危险的情况下，可以主动抛出
panic可以接受任何类型的值作为参数，通常是一个字符串或错误对象，用于描述发生的异常。
使用 panic() 程序就会终止,停在这里强制结束
*/
import "fmt"

func main() {
	panic("发生了异常")
	//panic后面的代码，程序不会执行
	fmt.Println("hello world")
}
