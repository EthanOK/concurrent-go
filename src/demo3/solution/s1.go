package solution

import (
	"sync"
	"time"
)

// 用 互斥锁 解决并发 资源竞争
var mutex sync.Mutex

var mapping = make(map[int]int, 10)

func sum(num int) {
	suum := 0
	for i := 0; i <= num; i++ {
		suum += i
	}

	mutex.Lock()
	mapping[num] = suum
	mutex.Unlock()
}

func TestMutexLock() {

	for i := 0; i < 100; i++ {
		go sum(i)
	}

	time.Sleep(time.Second * 10)
	for k, v := range mapping {
		println(k, v)
	}

}
