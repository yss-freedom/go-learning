package main

import "fmt"

// 函数可以返回指针类型的值，以便在函数外部访问函数内部创建的变量。
// 指针函数， 指针是可以用作函数的返回值
func main() {
	// 调用了这个函数后，可以得到一个指针类型的变量。
	ptr := f5()

	//内存地址正常打印前面带个&
	fmt.Println("ptr:", ptr)
	fmt.Printf("ptr:%p\n", ptr)
	fmt.Printf("ptr类型:%T\n", ptr)
	fmt.Println("ptr的地址:", &ptr)
	fmt.Println("ptr地址中的值:", *ptr)

	// 使用
	fmt.Println((*ptr)[0])
	ptr[0] = 10
	fmt.Println(ptr[0])

}

// 调用该函数后返回一个指针，此时返回个数组指针
func f5() *[4]int {
	arr := [4]int{1, 2, 3, 4}
	//内存地址往往被赋值给指针
	return &arr

}
