package main

/*
虽然errors.New()和fmt.Errorf()很方便，但在某些情况下，你可能需要更丰富的错误信息。这时，可以通过自定义类型来实现go内置的error接口。
自定义错误类型步骤
*/

import (
	"fmt"
)

// JingTian 定义一个结构体
type JingTian struct {
	msg  string
	code int
}

// 注意，Error()方法开头大写 string
// 实现内置error接口的Error方法
// 注意，方法的调用者用指针类型，因为结构体是值类型，不同函数里面的操作是互相独立的
func (e *JingTian) Error() string {
	//  fmt.Sprintf() 返回string
	return fmt.Sprintf("错误信息：%s,错误代码：%d\n", e.msg, e.code)
}

// 自定义错误，里面还可以写一些方法
// 处理error的逻辑
func (e *JingTian) print() bool {
	fmt.Println("hello，world")
	return true
}

// 使用错误测试
func test(i int) (int, error) {
	// 需要编写的错误
	if i != 0 {
		//自带的 errors.New()和fmt.Errorf() 只能返回字符串类型的错误信息，不常用，信息太少了
		// 更多的时候我们会使用自定义类型的错误
		//注意，返回的错误用内存地址，因为结构体是值类型
		return i, &JingTian{msg: "非预期数据", code: 500}
	}
	// 正常结果
	return i, nil
}

func main() {

	i, err := test(1)

	if err != nil {
		fmt.Println(err)
		//查看错误类型
		fmt.Printf("%T\n", err)
		ks_err, ok := err.(*JingTian)
		if ok {
			if ks_err.print() {
				// 处理这个错误中的子错误的逻辑
			}
			fmt.Println(ks_err.msg)
			fmt.Println(ks_err.code)
		}
	}
	fmt.Println(i)

}
