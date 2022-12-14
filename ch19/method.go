package main

import "fmt"

type account struct {
	balance int
}

// 일반 function
func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

// Method, (struct) 함수명(함수 인자 값) {}
// Class 로 쓰고 싶다면.. java처럼 class 파일을 따로 만들어놓는게 좋을듯?!
func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

// 별칭 타입
type myInt int

func (a myInt) add(b int) int {
	return int(a) + b
}

func main() {
	a := account{100}

	withdrawFunc(&a, 10)

	a.withdrawMethod(30)

	fmt.Printf("%d \n", a.balance)

	var b myInt = 30

	b.add(50)

	fmt.Printf("myInt value => %d\n", b)
}
