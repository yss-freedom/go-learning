package main

import "fmt"

//聚合继承

type Human struct {
	name string
	age  int
}

type Man struct {
	//聚合继承，不再以匿名字段形式将父类写在这，而是命名一个变量名
	Human Human
	phone string
}

func main() {

	//创建子类对象
	man := Man{
		Human: Human{"轩辕氏", 999},
		phone: "19899999999",
	}
	fmt.Println(man)
	//访问父类属性，聚合继承不能使用变量提升。提升字段, 只有匿名字段才可以做到
	//聚合方式继承，访问父类属性需要写全，对象.变量名.父类字段名
	fmt.Println(man.Human.name)

}
