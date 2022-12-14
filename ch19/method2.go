package main

import "fmt"

type account2 struct {
	balance   int
	firstName string
	lastName  string
}

func (a1 *account2) withdrawPointer(amount int) {
	a1.balance -= amount
}

func (a2 account2) withdrawValue(amount int) {
	a2.balance -= amount
}

func (a3 account2) withdrawReturnValue(amount int) account2 {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA = account2{100, "Joe", "Park"}
	mainB := account2{100, "Joe", "Park"}

	mainA.withdrawPointer(30)
	mainB.withdrawPointer(30)
	fmt.Println(mainA.balance)
	fmt.Println(mainB.balance)

	fmt.Println("============")

	mainA.withdrawValue(30)
	mainB.withdrawValue(30)
	fmt.Println(mainA.balance)
	fmt.Println(mainB.balance)

	fmt.Println("============")

}
