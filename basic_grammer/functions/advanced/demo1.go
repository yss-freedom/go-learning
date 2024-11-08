package main

/*
defer语句在Go语言中是一个强大的特性，它允许你延迟函数的执行直到包含它的函数即将返回。这通常用于清理资源、解锁互斥锁、记录时间、关闭文件等操作。
比如说我们打开一个文件，我们判断它有没关闭，我们希望所有代码执行完，最后调用defer函数来关闭文件
defer语句会将其后的函数调用压入一个栈中，当外层函数即将返回时，这些被defer的函数会按照后进先出（LIFO）的顺序执行。
defer可以在函数中多次使用，以注册多个需要延迟执行的函数。

defer函数或者方法：一个函数或方法的执行被延迟了
你可以在函数中添加多个defer语句，当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回，特别是当你在进行一些打开资源的操作时i/o 流，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题
如果有很多调用 defer，那么 defer 是采用 后进先出（栈） 模式。
你可以在函数中添加多个defer语句，当没有加defer的语句执行完毕后，这些defer语句会按照逆序执行
*/

//import "fmt"
//
//// defer
//func main() {
//	f("1")
//	fmt.Println("2")
//	//其他语句执行完，再执行defer包含的函数
//	defer f("3")
//	fmt.Println("4")
//}
//
//func f(s string) {
//	fmt.Println(s)
//}

//import "fmt"
//
//// defer 作用：处理一些善后的问题，比如错误，文件、网络流关闭等等操作。
//// 特点，多个defer的问题
//// 不加defer的语句还是按顺序从上而下执行
//// 你可以在函数中添加多个defer语句，当没有加defer的语句执行完毕后，这些defer语句会按照逆序执行
//func main() {
//	f("1")
//	defer fmt.Println("2")
//	defer f("3")
//	fmt.Println("4")
//	defer f("5")
//	fmt.Println("6")
//	defer f("7")
//	fmt.Println("8")
//}
//
//func f(s string) {
//	fmt.Println(s)
//}

//defer存在传递参数
//在defer的时候，函数其实已经被调用了，但是没有执行。参数是已经传递进去的了

import "fmt"

// defer传参的调用时机
func main() {
	n := 10
	fmt.Println("main n=", n)
	// 在defer的时候，函数其实已经被调用了，但是没有执行。参数是已经传递进去的了
	defer ff(n) // 问题，defer延迟执行函数，参数时什么时候传递进去？在其他地方虽然n已经加1了，ff函数最后执行，但是ff函数中的n还是10，说明参数是先加载进来，只是没执行函数
	n++
	fmt.Println("main end n=", n)

	defer ff(n) //如果ff在n加1之后defer延迟执行，参数中的n是加1后的n
}

func ff(n int) {
	fmt.Println("ff函数中n=", n)
}

//defer后续常用场景：程序会报错： 异常（程序执行的时候突然报错了）、错误（我们开发的时候知道的预期错误）
//善后工作：defer 处理异常。
