package main

import "fmt"

func main() {
	str := "Hello\tGo\tWorld\n\"Go\" is Awesome!\n" // 문자열

	fmt.Print(str)
	fmt.Printf("%s", str)
	fmt.Printf("%q", str)

}
