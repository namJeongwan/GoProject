package main

import (
	"fmt"
	"time"
)

func printSquare(ch chan int) {
	n := <-ch
	fmt.Println(n * n)
}

func main() {
	ch := make(chan int)
	go printSquare(ch)

	ch <- 5
	time.Sleep(5 * time.Second)
}
