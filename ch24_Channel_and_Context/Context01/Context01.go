package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // Create Context
	go PrintEverySecond(ctx)                                                // Context 인자로 넘겨준 뒤 GoRoutine 호출
	time.Sleep(5 * time.Second)                                             // 5초 대기
	cancel()                                                                // Cancel 명령

}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}
