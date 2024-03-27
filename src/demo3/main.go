package main

import "concurrent-go/src/demo3/solution"

func main() {
	// 1. 用 互斥锁 实现并发访问，线程安全
	// 缺点：无法预测执行多长时间结束
	// solution.TestMutexLock()

	// 2. 用 channel 实现并发访问，线程安全
	solution.TestChannel()

}
