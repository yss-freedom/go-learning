package main

import (
	"fmt"
	"time"
)

//5.3 判断时间先后

func main() {
	t1 := time.Now()
	t2 := t1.Add(time.Hour)
	fmt.Println(t1)
	fmt.Println(t2)

	// 判断时间先后
	fmt.Println(t1.Before(t2)) // true
	fmt.Println(t1.After(t2))  // false
	fmt.Println(t1.Equal(t2))  // false

}
