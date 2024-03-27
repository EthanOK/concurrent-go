package solution

import (
	"fmt"
	"math/rand"
	"time"
)

// 用 Channel 解决并发 资源竞争
func sumChannel(mapChan chan interface{}, n int) {

	for i := 1; i <= n; i++ {
		// 产生一个随机数
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(1000000) + i
		mapChan <- r
		fmt.Println("生产者:", i, "生产数据:", r)
		time.Sleep(time.Second)
	}

	close(mapChan)
}

func readChannel(mapChan chan interface{}, exitChan chan bool) {

	for v := range mapChan {

		fmt.Println("读取数据:", v)
	}
	exitChan <- true
	close(exitChan)
}

func TestChannel() {

	mapChan := make(chan interface{}, 10)

	exitChan := make(chan bool, 1)

	go sumChannel(mapChan, 10)

	go readChannel(mapChan, exitChan)

	// exitChan 管道堵塞 等待读取
	if <-exitChan {
		fmt.Println("读取数据完毕")
	}

}
