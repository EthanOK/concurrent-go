package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 协程 + 管道 应用 定时任务

func main() {

	fmt.Println(time.Now())

	// 创建一个新的 Timer，它将在至少持续时间3 s 后在其通道上发送当前时间
	timer := time.NewTimer(time.Second * 3)

	if isEven() {
		fmt.Println("偶数")
		// 停止 timer
		timer.Stop()
	} else {
		fmt.Println("奇数")
		// timer.C  只读管道
		t := <-timer.C
		fmt.Println(t)
	}

	// 方式二
	// fmt.Println(time.Now())
	// // 返回一个 Time chan
	// tt := <-time.After(time.Second * 3)
	// fmt.Println(tt)

}
func isEven() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Int31n(10)%2 == 0

}
