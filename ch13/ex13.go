package main

import "fmt"

type Student struct {
	Name  string
	Class int
	no    int
	Score float64
}

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	User
	VIPLevel int
	Price    int
}

func main() {
	var students [5]Student

	students[0] = Student{
		"JwNam",
		14,
		27,
		96.56,
	}

	students[1] = Student{
		Name:  "jdk",
		Class: 14,
		no:    28,
		Score: 78.66,
	}

	//user := User{
	//	Name: "HNKim",
	//	ID:   "HN123",
	//	Age:  25,
	//}
	vipUser := VIPUser{
		User{"Test1", "Test1", 14},
		3,
		257000,
	}

	fmt.Println(vipUser.Price)
	fmt.Println(vipUser.VIPLevel)
	fmt.Println(vipUser.Price)
	fmt.Println(vipUser.Name)

}
