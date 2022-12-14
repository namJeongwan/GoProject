package main

import "fmt"

type opFunc func(a, b int) int

func getOperator2(op string) opFunc {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	fn := getOperator2("*")

	result := fn(3, 4)

	fmt.Println(result)
}
