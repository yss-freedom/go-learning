package main

import "fmt"

//下面我们将通过一个实际的图书管理系统案例，来演示Go语言中结构体的应用。
//首先，我们定义一个Book结构体来表示图书的信息。
/*
结构体是Go语言中非常重要和强大的数据组织方式，通过结构体可以方便地表示复杂的数据结构和对象。
本文详细介绍了结构体的定义、使用、方法定义以及高级特性，如匿名结构体、嵌套结构体、结构体指针，结构体导出规则等。
通过实际案例的演示，希望能够让读者对Go语言中的结构体有更深入的理解和掌握。
结构体不仅提高了代码的可读性和可维护性，还使得数据组织更加清晰和易于管理。
*/

type Book struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Published string `json:"published"`
	ISBN      string `json:"isbn"`
}

// 接下来，我们可以创建一个包含多个Book结构体的切片来表示图书列表。
var books []Book

func initBooks() {
	books = append(books, Book{
		ID:        "001",
		Title:     "Go语言编程之旅",
		Author:    "Alice",
		Published: "2023-01-01",
		ISBN:      "978-7-121-34567-8",
	})
	// ... 可以继续添加其他图书
}

// 为了演示，我们可以添加一个功能来遍历图书列表并打印每本书的详细信息。
func printBooks() {
	for _, book := range books {
		fmt.Printf("ID: %s, Title: %s, Author: %s, Published: %s, ISBN: %s\n",
			book.ID, book.Title, book.Author, book.Published, book.ISBN)
	}
}

func main() {
	initBooks()
	// 对books进行操作，如显示图书列表等
	printBooks()
}
