package main

import "fmt"

func getMyAge() (int, bool) {
	return 33, true
}

func main() {
	if age, ok := getMyAge(); ok {
		fmt.Println("Age is => ", age)
	} else if ok && age < 30 {

	} else if ok {

	} else {
		fmt.Println("Error!")
	}

	//var day int32 = 3
	day := 3

	switch day {
	case 1:
		fmt.Println("First")
	case 2:
		fmt.Println("Second")
	case 3:
		fmt.Println("Third")
	}

}
