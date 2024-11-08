package main

import (
	"fmt"
	"sync"
)

/*
读写锁（sync.RWMutex）允许多个goroutine同时读取资源，
但在写入时会阻塞所有其他读和写的goroutine。读写锁可以提高读多写少的场景下的并发性能。
*/

var (
	data map[string]int
	rwMu sync.RWMutex
)

func readData(key string) int {
	rwMu.RLock()
	defer rwMu.RUnlock()
	return data[key]
}

func writeData(key string, value int) {
	rwMu.Lock()
	defer rwMu.Unlock()
	data[key] = value
}

func main() {
	data = make(map[string]int)

	var wg sync.WaitGroup

	// 写操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			writeData(fmt.Sprintf("key%d", i), i*10)
		}(i)
	}

	// 读操作
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			value := readData(fmt.Sprintf("key%d", i%10))
			fmt.Printf("Read: key%d = %d\n", i%10, value)
		}(i)
	}

	wg.Wait()
}

/*
避免写锁长时间持有：写锁会阻塞所有其他读和写的goroutine，因此应该尽量减少写锁的持有时间。
读多写少场景：读写锁适用于读多写少的场景，如果写操作非常频繁，读写锁的性能优势可能会消失。
避免嵌套锁：与互斥锁类似，读写锁也应该避免嵌套使用。
*/
