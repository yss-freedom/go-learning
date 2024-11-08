package main

import "fmt"

/*
2. 定义泛型方法：
泛型方法可以应用于泛型类型上。
针对不同类型的切片做累加和，使用泛型比较简便
*/

// MySlice 定义一个泛型切片
type MySlice[T int | float32 | int64] []T

func main() {
	//针对不同类型的切片都可以计算
	var s MySlice[int] = []int{1, 2, 3, 4}
	fmt.Println(s.sum())

	var s1 MySlice[float32] = []float32{1.0, 2.0, 3.0, 4.4}
	fmt.Println(s1.sum())
}

// 调用者，类型是不确定的，用户传什么，她就实例化什么。  类型参数化了 ， 泛型
// 没有泛型之前， 反射:     reflect.ValueOf().Kind() , 也需要很多if，本质是逻辑相同的，只是类型不同！
func (s MySlice[T]) sum() T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}
