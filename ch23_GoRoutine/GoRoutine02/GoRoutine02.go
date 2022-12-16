package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	fmt.Printf("[SumAtoB] %d to %d args init..\n", a, b)
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}

	fmt.Printf("%d to %d sum => %d\n", a, b, sum)
	wg.Done()
}

func main() {
	wg.Add(10)
	startTime := time.Now()
	for i := 0; i <= 10; i++ {
		go SumAtoB(i, 100000000)
	}
	wg.Wait()
	elapsedTime := time.Since(startTime)

	fmt.Printf("Total Second => %v", elapsedTime)
}
