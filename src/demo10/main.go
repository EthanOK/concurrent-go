package main

import (
	"fmt"
	"sync"
)

// 100 个人抢 20 个包子
func main() {

	var wg sync.WaitGroup

	// 初始化 20 个包子
	// 使用通道 解决 资源竞争的问题
	ch := make(chan int, 20)

	for i := 1; i <= 20; i++ {
		ch <- i
	}

	// 100个人 去抢
	for i := 1; i <= 100; i++ {
		wg.Add(1)

		go func(id int) {

			defer wg.Done()
			select {

			case t := <-ch:
				// 抢到了包子
				println("Consumer.", id, "抢到了包子No.", t)

			default:

			}

		}(i)

	}

	wg.Wait()
	fmt.Println("包子被抢完了")

}
