package main

import "fmt"

func main() {
	var minimumWage int = 10
	var workingHour int = 20

	var income int = minimumWage * workingHour

	fmt.Println(minimumWage, workingHour, income)

	minimumWage = 50

	fmt.Println("MininumWage is changed !!")

	fmt.Println(minimumWage, workingHour, income)
}