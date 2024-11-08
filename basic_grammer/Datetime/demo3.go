package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//时间戳字符串
	timrstr := "1617922800"

	//时间戳是int64类型数据，将时间戳字符串转化为int64类型
	timestamp, err := strconv.ParseInt(timrstr, 10, 64)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	t := time.Unix(timestamp, 0)
	fmt.Println(t)
}
