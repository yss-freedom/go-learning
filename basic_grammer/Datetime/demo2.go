package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	timestamp1 := now.Unix()     // 秒级时间戳
	timestamp2 := now.UnixNano() // 纳秒级时间戳

	fmt.Printf("current timestamp1: %v\n", timestamp1)
	fmt.Printf("current timestamp2: %v\n", timestamp2)

	// 将时间戳转换为时间对象
	timeObj := time.Unix(timestamp1, 0)
	fmt.Println(timeObj)

	// 转换为指定时区的时间对象
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	beijingTimeObj := time.Unix(timestamp1, 0).In(beijing)
	fmt.Println(beijingTimeObj)
}
