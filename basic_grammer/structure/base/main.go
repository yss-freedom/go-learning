package main

import (
	"basic_grammer/structure/base/pojo"
	"fmt"
)

func main() {
	user := pojo.User{
		Name: "yssnb",
		Age:  18,
		//money : 900
		//money存在，但不可以被访问和修改
	}

	fmt.Println(user)
}
