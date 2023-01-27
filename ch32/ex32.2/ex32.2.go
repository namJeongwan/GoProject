package main

import "fmt"

func main() {
	fmt.Println("##### STR ######")
	str := "hello World"
	for i, c := range str {
		fmt.Printf("[%d] %c\n", i, c)
	}
	fmt.Println("##### STR ######")

	fmt.Println("##### SLICE ######")
	slice := []int{1, 2, 3, 4, 5}
	for i, v := range slice {
		fmt.Printf("[%d] %d\n", i, v)
	}

	fmt.Println("##### SLICE ######")
	fmt.Println("##### MAP ######")
	testMap := map[string]int{"aaa": 1, "bbb:": 2, "ccc": 333}
	for k, v := range testMap {
		fmt.Printf("%s: %d\n", k, v)
	}
	fmt.Println("##### MAP ######")

}
