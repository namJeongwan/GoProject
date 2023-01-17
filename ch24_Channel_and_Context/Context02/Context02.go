package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "number", 9)
	ctx = context.WithValue(ctx, "number2", 22)
	go square(ctx)
	cancel()

	wg.Wait()

}

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("Sqare(1) :%d", n*n)
	}
	if v := ctx.Value("number2"); v != nil {
		n := v.(int)
		fmt.Printf("Sqare(2) :%d", n*n)
	} else {
		fmt.Printf("Not in Sqaure")
	}
	wg.Done()
}
