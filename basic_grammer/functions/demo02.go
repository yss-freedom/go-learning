package main

import "fmt"

func Find(m map[int]int, key int) (value int, ok bool) {
	value, ok = m[key]
	return
}

func main() {
	fmt.Println()
}
