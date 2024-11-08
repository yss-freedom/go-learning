// 带有参数和返回值的接口
package main

import "fmt"

// Tongxin 定义接口

type Tongxin interface {
	//定义带有参数和返回值的方法
	dadianhua(youdian bool) string
	jieidanhua(youdian bool) string
}

// People 定义结构体
type People struct {
	name  string
	age   int
	phone string
}

// 实现接口
func (p People) dadianhua(youdian bool) string {
	if youdian {
		return fmt.Sprintf("%v 打了电话", p.name)
	} else {
		return fmt.Sprintf("打电话时手机没电了")
	}

}

func (p People) jieidanhua(youdian bool) string {
	if youdian {
		return fmt.Sprintf("%v 接了电话", p.name)
	} else {
		return fmt.Sprintf("接电话时手机没电了")
	}

}

// 接口测试，有传参，有返回值
func testdianhua(phone Tongxin) {
	str1 := phone.dadianhua(false)
	str2 := phone.jieidanhua(true)
	fmt.Println(str1, str2)
}

func main() {
	//创建对象
	p := People{"jingtian", 18, "18898985898"}
	//如果一个结构体实现了这个接口所有的方法，那这个结构体就是这个接口类型的
	testdianhua(p)

}
