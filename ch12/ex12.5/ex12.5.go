package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for idx, value := range a {
		fmt.Printf("a[%d] = %d\n", idx, value)
	}

	fmt.Println()
	for idx, value := range b {
		fmt.Printf("b[%d] = %d\n", idx, value)
	}

	b = a
	fmt.Println()
	for idx, value := range b {
		fmt.Printf("copied_b[%d] = %d\n", idx, value)
	}

	var _2DArray = [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	for _, arr := range _2DArray {
		for _, val := range arr {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}
