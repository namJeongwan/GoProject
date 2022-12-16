package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go diningProblem("A", fork, spoon, "fork", "spoon")
	go diningProblem("B", fork, spoon, "fork", "spoon")

	wg.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("[%v]try to take mutex...\n", name)
		first.Lock()
		fmt.Printf("[%v]take mutex(%v)\n", name, firstName)
		second.Lock()
		fmt.Printf("[%v]take mutex(%v)\n", name, secondName)

		fmt.Printf("[%v]Successfully take mutex!!\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}
