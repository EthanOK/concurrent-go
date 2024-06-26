package main

import (
	"time"
)

var allNumber chan int = make(chan int, 10000)

var goroutineNumber int = 100000

func initData(num int) {
	for i := 1; i <= num; i++ {
		allNumber <- i
	}
	close(allNumber)

}

func handleNumber(allNumber <-chan int, primeNumber chan<- int, exitChan chan<- bool) {
	/*
		优化:
		allNumber 只读通道
		primeNumber 只写通道
		exitChan 只写通道
	*/

	for {
		n, ok := <-allNumber
		if !ok {
			// 数据取完了
			break
		}
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
	}

	exitChan <- true

}

func main() {
	primeNumber := make(chan int, 10000)

	exitChan := make(chan bool, goroutineNumber)

	go initData(100000)

	for i := 0; i < goroutineNumber; i++ {
		go handleNumber(allNumber, primeNumber, exitChan)
	}
	go func() {
		for i := 0; i < goroutineNumber; i++ {

			<-exitChan
		}
		close(primeNumber)
	}()

	for n := range primeNumber {
		println(n)
	}

}
