package main

import "fmt"

type Books struct {
	title  string
	author string
}

func main() {
	var book1 Books
	book1.title = "Go 语言入门"
	book1.author = "mars.hao"
	fmt.Println()
	fmt.Println("\xe4\x00\x00\xe7\x95\x8cabc") // �界abc

}
