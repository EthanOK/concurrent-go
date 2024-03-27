package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go goroutine()

	for i := 0; i < 10; i++ {
		fmt.Println("main")
		time.Sleep(time.Second * 1)
	}
	// cpu 核数
	fmt.Println(runtime.NumCPU())

}

func goroutine() {
	for i := 0; i < 10; i++ {
		fmt.Println("goroutine")
		time.Sleep(time.Second)
	}

}
