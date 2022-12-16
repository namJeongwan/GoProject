package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

var mutex sync.Mutex

func main() {
	var wg sync.WaitGroup

	account := &Account{0}

	fmt.Printf("Start!\n")
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 5; j++ {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()

	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	fmt.Printf("Current Account balance => %d\n", account.Balance)
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}
