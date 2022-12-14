package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

type operator func(int, int) int

func getOperator(op string) operator {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	op := getOperator("+")
	fmt.Println(op(20, 50))
}
