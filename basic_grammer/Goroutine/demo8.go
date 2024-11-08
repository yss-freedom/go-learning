package main

import (
	"fmt"
	"sync"
)

//sync.Once用于确保某个操作只执行一次，无论有多少个goroutine调用它。
//这对于单例模式或初始化只执行一次的场景非常有用

var (
	once    sync.Once
	message string
)

func initMessage() {
	message = "Hello, World!"
}

func printMessage() {
	once.Do(initMessage)
	fmt.Println(message)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			printMessage()
		}()
	}

	wg.Wait()
}

/*
避免重复初始化：sync.Once确保某个操作只执行一次，因此它通常用于初始化全局变量或执行其他只需要执行一次的操作。
性能开销：虽然sync.Once的性能开销很小，但在高性能要求的场景下，仍然需要注意其使用。
*/
