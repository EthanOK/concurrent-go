package main

/*
defer recover 处理异常：
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}

	}()
*/

import (
	"fmt"
	"time"
)

func onIn() {
	// 定义一个只写通道
	inChan := make(chan<- int, 10)
	fmt.Println(inChan)
	inChan <- 1
	// fmt.Println(<-inChan) // 报错，只读通道不能写入数据
}

func onOut() {
	// 定义一个只读通道
	outChan := make(<-chan int, 10)
	fmt.Println(outChan)
	// outChan <- 1 // 报错，只读通道不能写入数据

}

func errrr(n int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}

	}()
	return 10 / n
}

func main() {

	go onIn()

	go onOut()

	go errrr(0)

	time.Sleep(time.Second)
	fmt.Println("main")

}
