package main

import (
	"fmt"
	"time"
)

// 遍历 管道不 close 会报错，怎么优化；使用 select
/*
label1:
	for {
		select {
		case n := <-primeNumber:
			fmt.Println("素数: ", n)
		default:
			break label1
		}

	}
*/

var allNumber chan int = make(chan int, 10000)

var goroutineNumber int = 100000

func initData(num int) {
	for i := 1; i <= num; i++ {
		allNumber <- i
	}

}

func handleNumber(allNumber <-chan int, primeNumber chan<- int) {

	// 捕获 panic 防止中断
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()
	/*
		优化:
		allNumber 只读通道
		primeNumber 只写通道
	*/

label:
	for {
		select {
		case n := <-allNumber:
			time.Sleep(time.Microsecond)
			if n <= 1 {
				continue
			}
			for i := 2; i <= n/2; i++ {
				if n%i == 0 {
					continue
				}
			}
			primeNumber <- n

		default:
			break label
		}

	}

}

func main() {

	start := time.Now()

	primeNumber := make(chan int, 10000)

	go initData(200000)

	for i := 0; i < goroutineNumber; i++ {
		go handleNumber(allNumber, primeNumber)
	}

label1:
	for {
		select {
		case n := <-primeNumber:
			fmt.Println("素数:", n)
		default:
			break label1
		}

	}

	end := time.Since(start)
	fmt.Println("执行时间:", end)

}
