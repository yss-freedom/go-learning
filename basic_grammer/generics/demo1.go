package main

import "fmt"

//我们先用空接口来测试多种类型的适配实现，还是需要做类型判断

func main() {
	str := []string{"1", "2", "3", "4", "5"}
	printSlice(str)
	str2 := []int{1, 2, 3, 4, 5}
	printSlice(str2)

}

func printSlice(i interface{}) {
	//  断言 x.(T), 如果x实现了T，那么就将 x转换为T类型
	//这样的话，也是需要做好多种判断
	switch i.(type) {
	case []string:
		for _, i3 := range i.([]string) {
			fmt.Println(i3)

		}
		fmt.Println()
	case []int:
		for _, i3 := range i.([]int) {
			fmt.Println(i3)
		}

	}
	// 其他很多种类型都要做判断，所以通过这种方式不现实.........

}
