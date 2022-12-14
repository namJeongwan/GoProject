package main

import (
	"fmt"
)

func copySlice(arr []interface{}) []interface{} {
	copiedSlice := make([]interface{}, len(arr))

	for i, v := range arr {
		copiedSlice[i] = v
	}
	return copiedSlice
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))

	for i, v := range slice1 {
		slice2[i] = v
	}

	slice1[1] = 100

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := append([]int{}, slice3...)

	slice3[1] = 100
	fmt.Println(slice3)
	fmt.Println(slice4)

}
