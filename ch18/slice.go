package main

import "fmt"

func changeArray(array2 [5]int) {
	array2[2] = 200
}

func changeSlice(slice2 []int) {
	slice2[2] = 200
}

func main() {
	//slice1 := []int{1, 2, 3}
	fmt.Println("============= 배열 공간의 여유가 있어 값 복사가 아닌 주소 복사가 된 경우 =============")
	fmt.Println("")
	slice1 := make([]int, 3, 5)

	slice2 := append(slice1, 4, 5)

	fmt.Println("slice1", slice1, ", len:", len(slice1), ",cap:", cap(slice1))
	fmt.Println("slice2", slice2, ", len:", len(slice2), ",cap:", cap(slice2))

	slice1[1] = 100

	fmt.Println("After change second element")
	fmt.Println("slice1", slice1, ", len:", len(slice1), ",cap:", cap(slice1))
	fmt.Println("slice2", slice2, ", len:", len(slice2), ",cap:", cap(slice2))

	slice1 = append(slice1, 500)

	fmt.Println("After append 500")
	fmt.Println("slice1", slice1, ", len:", len(slice1), ",cap:", cap(slice1))
	fmt.Println("slice2", slice2, ", len:", len(slice2), ",cap:", cap(slice2))

	fmt.Println("주소 값 비교")
	fmt.Println(&slice1[0], "==", &slice2[0], "?")
	fmt.Println("==============================================================================")

	fmt.Println("")

	fmt.Println("============= 배열 공간의 여유가 없어 값 복사가 된 경우 =============")
	fmt.Println("")
	poorSlice1 := make([]int, 3, 3)

	poorSlice2 := append(poorSlice1, 4, 5)

	fmt.Println("poorSlice1", poorSlice1, ", len:", len(poorSlice1), ",cap:", cap(poorSlice1))
	fmt.Println("poorSlice2", poorSlice2, ", len:", len(poorSlice2), ",cap:", cap(poorSlice2))

	poorSlice1[1] = 100

	fmt.Println("After change second element")
	fmt.Println("poorSlice1", poorSlice1, ", len:", len(poorSlice1), ",cap:", cap(poorSlice1))
	fmt.Println("poorSlice2", poorSlice2, ", len:", len(poorSlice2), ",cap:", cap(poorSlice2))

	poorSlice1 = append(poorSlice1, 500)

	fmt.Println("After append 500")
	fmt.Println("poorSlice1", poorSlice1, ", len:", len(poorSlice1), ",cap:", cap(poorSlice1))
	fmt.Println("poorSlice2", poorSlice2, ", len:", len(poorSlice2), ",cap:", cap(poorSlice2))

	fmt.Println("주소 값 비교")
	fmt.Println(&poorSlice1[0], "==", &poorSlice2[0], "?")
	fmt.Println("==============================================================================")

	slice3 := []int{1, 2, 3, 5, 6}
	slice4 := slice3[1:3:4]
	fmt.Println("slice4", slice4, ", len:", len(slice4), ",cap:", cap(slice4))

}
