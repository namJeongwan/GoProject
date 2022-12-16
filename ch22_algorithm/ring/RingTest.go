package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5) // 최초에 길이 선언 해줘야 한다

	n := r.Len() // 링 길이 반환

	for i := 0; i <= 5; i++ {
		r.Value = 'A' + i - 1
		r = r.Next() // ring cursor 변경
	}

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Next()
	}

	fmt.Println()

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Prev()
	}

}
