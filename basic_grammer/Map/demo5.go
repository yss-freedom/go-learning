package main

import (
	"fmt"
	"sync"
)

/*
在并发环境下，普通的Map不是并发安全的。
如果多个goroutine同时读写同一个Map，可能会导致竞态条件。
为了解决这个问题，Go语言标准库提供了sync.Map类型，它是专门为并发环境设计的。
*/
var g_syncMap sync.Map

func main() {
	// 添加元素
	g_syncMap.Store(1, "one")
	g_syncMap.Store(2, "two")

	// 遍历
	g_syncMap.Range(func(key, value interface{}) bool {
		fmt.Printf("key: %v, value: %v\n", key, value)
		return true
	})

	// 删除元素
	g_syncMap.Delete(1)

	// Load或LoadOrStore
	if v, ok := g_syncMap.Load(2); ok {
		fmt.Println("Loaded:", v)
	}

	if loaded, ok := g_syncMap.LoadOrStore(3, "three"); !ok {
		fmt.Println("Stored:", loaded)
	}
}
