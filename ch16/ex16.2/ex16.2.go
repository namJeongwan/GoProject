package main

import (
	"fmt"
	"goproject/ch16/ex16.2/publicpkg"
)

func main() {
	fmt.Println("PI: ", publicpkg.PI)
	publicpkg.PublicFunc()

	var myInt publicpkg.MyInt = 10
	fmt.Println("myInt: ", myInt)

	var myStruct = publicpkg.MyStruct{Age: 19}

	fmt.Println("myStruct: ", myStruct)
}
