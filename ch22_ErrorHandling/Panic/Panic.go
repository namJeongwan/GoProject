package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		panic("0 divide Exception")
	}
	fmt.Printf("%d / %d = %d\b\n", a, b, a/b)
}

func main() {
	divide(2, 4)
	divide(5, 0)
}
