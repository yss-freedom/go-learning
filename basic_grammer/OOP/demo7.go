package main

import (
	"fmt"
	"math"
)

//模拟多态
/*
多态是指相同的接口（方法）可以表现出不同的行为。在Go语言中，通过接口实现多态。
在Go语言中，接口定义了一组方法的集合，但不实现它们，而是由具体的类型来实现这些方法。
任何实现了接口中所有方法的类型都被视为该接口的实现。接口是Go语言中实现多态性的关键。

多态：一个事务有多种形态
父类：动物
子类：猫
子类：狗
*/

/*在上面的代码中，我们创建了一个shapes切片，该切片包含了不同类型的形状对象（矩形和圆形）。
然后，我们遍历shapes切片，并通过Shape接口调用Area()方法。
由于这两种形状都实现了Shape接口，因此多态性使我们能够以一致的方式调用它们的Area()方法。
*/

type Shape interface {
	Area() float64
}

// Rectangle 矩形
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle 圆形
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	r := Rectangle{Width: 4, Height: 5}
	c := Circle{Radius: 4}

	shapes := []Shape{r, c}

	for _, shape := range shapes {
		fmt.Printf("Area: %f\n", shape.Area())
	}
}
