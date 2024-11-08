package main

/*
在实际开发中，经常需要将结构体实例转换为JSON格式的字符串，以便于网络传输或数据存储。
Go标准库中的encoding/json包提供了这样的功能。

在上面的例子中，通过json:"name"和json:"age"这样的标签（Tag），我们指定了结构体字段在JSON字符串中的键名。
这些标签是Go语言结构体独有的特性，使得Go能够轻松地与其他语言进行JSON数据的交换。
*/

import (
	"encoding/json"
	"fmt"
)

type Person3 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person3{Name: "David", Age: 45}

	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json encode failed:", err)
		return
	}

	fmt.Println(string(data)) // 输出: {"name":"David","age":45}
}

/**
结构体标签是结构体字段后面的元信息，它由一对反引号`包裹，格式通常为键值对形式，键值之间用冒号:分隔，值被双引号"包围。
结构体标签主要用于为结构体字段提供额外的信息，这些信息在运行时可以通过反射机制读取。
除了上述的JSON序列化场景，结构体标签还可以用于数据库ORM映射、XML序列化等场景。
*/

type Product struct {
	ID    int     `db:"id"`     // 数据库字段映射
	Name  string  `json:"name"` // JSON序列化
	Price float64 `xml:"price"` // XML序列化
}
