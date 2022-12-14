package main

import "fmt"

const (
	Red int = iota << 1
	Blue
	Green
)

func Divide(a int, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b

	return

}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)

	d, success := Divide(9, 0)
	fmt.Println(d, success)

	fmt.Println(Red)
	fmt.Println(Blue)
	fmt.Println(Green)
}
