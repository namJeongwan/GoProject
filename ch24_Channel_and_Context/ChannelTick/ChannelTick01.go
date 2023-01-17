package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)            // 1초 간격 시그널
	terminate := time.After(10 * time.Second) // 10초 이후 시그널

	for {
		select {
		case <-tick:
			fmt.Println("Tick!!")
		case <-terminate:
			fmt.Println("Terminated!")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Sqaure: %d\n", n*n)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	fmt.Println("Start!")
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}
