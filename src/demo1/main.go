package main

import (
	"fmt"
)

func runPrintTime(num int) {
	for i := 0; i < num; i++ {
		fmt.Println(num, i)

	}

}

func main() {

	for i := 0; i < 10000000; i++ {
		go runPrintTime(i)

	}

	// runtime.NumCPU()

	// panic: too many concurrent operations on a single file or socket (max 1048575))

}
