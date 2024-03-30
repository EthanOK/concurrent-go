package main

import (
	"fmt"
	"time"
)

func main() {
	// 定时任务 2s 执行一次
	ticker := time.NewTicker(time.Second * 3)
	for {
		go doEverything()
		t := <-ticker.C
		println(t.Format("2006-01-02 15:04:05"))
	}

}

func doEverything() {
	fmt.Println("Doing something...")
}
