package main

import "fmt"

/*
Map的底层实现
3.1 Hash表与Bucket
Go语言的Map使用Hash表作为底层实现。Hash表是一种通过哈希函数组织数据，以支持快速插入和搜索的数据结构。
在Go的Map中，一个哈希表可以有多个Bucket，每个Bucket可以保存一个或一组键值对。
3.2 Hash冲突与解决
Hash冲突是指不同的键经过哈希函数计算后得到相同的哈希值。
Go采用链地址法（也称为开放寻址法的一种变种）来解决Hash冲突。
当一个Bucket存放的键值对超过一定数量（Go中为8个）时，会创建一个新的Bucket，并将新的键值对添加到新的Bucket中，同时用指针将两个Bucket链接起来。
3.3 负载因子与扩容
负载因子是衡量Hash表冲突情况的一个指标，其计算公式为：负载因子 = 键数量 / Bucket数量。
Go的Map在负载因子达到6.5时会触发扩容，以减少冲突并提高访问效率。
扩容时，会创建一个新的Bucket数组，其长度是原来的两倍，然后将旧Bucket数组中的元素搬迁到新的Bucket数组中。
3.4 扩容策略
Go的Map采用逐步搬迁的策略来减少扩容时的延时。
每次访问Map时都会触发一次搬迁，但每次只搬迁两个键值对。
这种策略使得扩容过程更加平滑，减少了因一次性搬迁大量数据而导致的性能问题。
*/

/*
* 实际应用案例
在开发命令行工具时，经常需要根据不同的命令名称调用不同的函数。
使用Map可以很方便地实现这一功能。定义一个Map，其键为命令名称，值为函数指针。
这样，在接收到命令时，只需根据命令名称从Map中查找对应的函数指针并调用即可
*/

var FuncMap = map[string]func(int, string){
	"111": map1,
	"222": map2,
}

func map1(a int, b string) {
	fmt.Println("111", a, b)
}

func map2(a int, b string) {
	fmt.Println("222", a, b)
}

func test1(id string) {
	a := 250
	b := "25.250"
	if fn, ok := FuncMap[id]; ok {
		fn(a, b)
	} else {
		fmt.Println(id, "not found func")
	}
}

func main() {

}
