package main

import "fmt"

//结构体方法
/*
结构体方法是在结构体上定义的函数，用于实现对该结构体实例的操作。
在Go语言中，通过接收者（receiver）实现这一功能。接收者可以是值接收者或指针接收者。
*/

/*在上面的例子中，Describe方法是一个值接收者方法，它返回Person实例的描述信息。
而SetAge方法是一个指针接收者方法，它允许我们修改Person实例的Age字段。
*/

type Person2 struct {
	Name string
	Age  int
}

// 使用值接收者

func (p Person2) Describe() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

// 使用指针接收者

func (p *Person2) SetAge(age int) {
	p.Age = age
}

func main() {
	p := Person2{Name: "Charlie", Age: 35}
	fmt.Println(p.Describe()) // 输出: Name: Charlie, Age: 35

	p.SetAge(40)
	fmt.Println(p.Describe()) // 输出: Name: Charlie, Age: 40
}
