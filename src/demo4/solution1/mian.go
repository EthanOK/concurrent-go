package main

import "time"

func main() {
	// 打印100以内的素数
	for i := 1; i <= 100000; i++ {
		if isPrime(i) {
			println(i)
		}
	}

}

func isPrime(n int) bool {
	time.Sleep(time.Microsecond)
	if n <= 1 {
		return false
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true

}
