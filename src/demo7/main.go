package main

// 生产者 消费者 模式
import (
	"fmt"
	"strconv"
	"time"
)

type Product struct {
	Name  string
	Price float64
}

func producer(pChan chan<- Product, number int) {

	for {
		if number < 1 {
			break
		}
		pChan <- Product{Name: "Product" + strconv.Itoa(number), Price: float64(number * 2)}

		fmt.Println("生产:", "Product"+strconv.Itoa(number))

		time.Sleep(time.Second)

		number--
	}
}

func transporter(pChan <-chan Product, shopChan chan<- Product) {
	for {
		produce := <-pChan
		shopChan <- produce
		fmt.Println("运输:", produce)
	}
}

func shop(shopChan <-chan Product, number int, exitChan chan<- bool) {
	shopCount := 0
	for {
		if number < 1 {
			break
		}

		produce := <-shopChan

		fmt.Println("消费 :", produce)

		shopCount++

		number--

	}
	fmt.Println("总共消费:", shopCount)
	exitChan <- true
}

func main() {
	pChan := make(chan Product, 100)
	shopChan := make(chan Product, 100)
	exitChan := make(chan bool, 1)

	for i := 0; i < 100; i++ {
		go producer(pChan, 100)
	}

	go transporter(pChan, shopChan)

	go shop(shopChan, 100, exitChan)

	if <-exitChan {
		return
	}

}
