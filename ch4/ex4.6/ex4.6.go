package main

import "fmt"

var g int = 10 // 전역 변수

func main() {
	var m int = 20 // 지역 변수
	{
		var s int = 50 // 지역 변수2
		fmt.Println(m, s, g)
	}

	//m = s + 20
}
