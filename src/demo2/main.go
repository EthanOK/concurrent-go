package main

import "time"

// 并发安全性 资源竞争

var mapping = make(map[int]int, 10)

func sum(num int) {
	suum := 0
	for i := 0; i <= num; i++ {
		suum += i
	}
	mapping[num] = suum
}

func main() {

	for i := 0; i < 100; i++ {
		go sum(i)
	}

	time.Sleep(time.Second * 10)
	for k, v := range mapping {
		println(k, v)
	}

	// fatal error: concurrent map writes

}
