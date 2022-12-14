package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) User {
	var u = User{name, age}
	return u
}

func main() {
	userPointer := NewUser("AAA", 123)

	fmt.Println(userPointer)
}
