package main

import "fmt"

func main() {
	a := 1
	b := 1

OuterFor:
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor
			}
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)

	f := [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i := 0; i < 5; i++ {
		fmt.Printf("%f\n", f[i])
	}
}
