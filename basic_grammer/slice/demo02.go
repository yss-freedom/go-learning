package main

import "fmt"

/*
在Go语言中，切片（Slice）是一种非常强大且灵活的数据结构，它基于数组但又提供了动态调整大小的能力。切片在Go语言中非常常用，几乎成为处理序列数据的首选方式。
本文将结合实际案例，详细介绍Go语言中切片的声明、初始化、操作、扩容等用法。

切片是Go语言中对数组的抽象表示，它提供了一种动态的方式来处理序列数据。切片不是数组，但它内部包含了对数组的引用。切片拥有三个关键属性：

指向底层数组的指针：指向切片实际存储数据的数组。
切片的长度（len）：切片当前包含的元素数量。
切片的容量（cap）：从切片开始到其底层数组末尾的元素数量。
切片的声明有几种方式，以下是一些常见的声明方法：
使用var关键字声明切片，但不初始化：
var mySlice []int
此时，mySlice是一个nil切片，即它没有指向任何底层数组，长度和容量均为0。
使用:=自动推导类型声明切片：
mySlice := []int{}
此时，mySlice是一个空的整型切片，长度为0，但已经分配了底层数组（容量为0或系统默认值）。
使用make函数创建切片：
mySlice := make([]int, 5)
通过make函数可以创建一个指定长度的切片，并且可以选择性地指定容量。如果不指定容量，则容量等于长度。

make函数是Go语言内置的函数，专门用于创建切片、map和channel等复合类型。使用make函数创建切片时，可以指定切片的长度和容量：
slice := make([]int, 5, 10) // 创建一个长度为5，容量为10的整型切片

切片也可以像数组一样使用字面量进行初始化，但不需要指定数组的长度：
myStr := []string{"Jack", "Mark", "Nick"} // 创建并初始化一个字符串切片
myNum := []int{10, 20, 30, 40} // 创建并初始化一个整型切片

切片可以基于一个已存在的数组创建，切片可以只使用数组的一部分元素或者全部元素，甚至可以创建一个比数组更大的切片（只要不超过数组的容量）。
months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
q2 := months[3:6] // 基于数组创建切片，表示第二季度
*/

func main() {
	//访问切片元素与访问数组元素类似，使用索引操作符[]。
	//mySlice := []int{1, 2, 3, 4, 5}
	//fmt.Println(mySlice[0]) // 输出: 1

	//切片支持动态地增加或减少元素，这是切片与数组最大的不同。切片扩容通常通过内置的append函数实现。
	mySlice := []int{1, 2, 3}
	mySlice = append(mySlice, 4, 5) // 向切片末尾追加两个元素
	fmt.Println(mySlice)            // 输出: [1 2 3 4 5]

	/*
		当使用append函数向切片追加元素时，如果切片的容量不足以容纳新增的元素，
		Go语言会自动进行扩容。扩容时，新切片的容量通常是原容量的两倍（当原切片长度小于1024时）。
		如果扩容后的容量仍然不够，则继续按此规律扩容，直到能够容纳所有元素。
	*/
	oldSlice := []int{1, 2, 3, 4, 5}
	newSlice := append(oldSlice, 6, 7, 8, 9)
	fmt.Println("newSlice len:", len(newSlice), "cap:", cap(newSlice))
	// 输出可能会是: newSlice len: 9 cap: 10

	//切片也可以基于另一个切片创建，这被称为切片的切片。
	//通过指定起始索引和结束索引（可选指定容量），
	//可以从一个切片中创建新的切片。
	//firstHalf := months[:6] // 前半年
	//q1 := firstHalf[:3] // 第一季度
	//Go语言提供了copy函数，用于将一个切片复制到另一个切片中。
	//如果两个切片的大小不同，则按照其中较小的那个切片的大小进行复制。
	src := []int{1, 2, 3}
	dst := make([]int, 2)
	copy(dst, src)   // 将src的前两个元素复制到dst
	fmt.Println(dst) // 输出: [1 2]

	//切片的遍历方式与数组相同，支持使用索引遍历和for range遍历。
	//slice := []int{1, 2, 3, 4, 5}
	//for i := 0; i < len(slice); i++ {
	//	fmt.Println(slice[i])
	//}
	slice := []int{1, 2, 3, 4, 5}
	for index, value := range slice {
		fmt.Println(index, value)
	}

}
