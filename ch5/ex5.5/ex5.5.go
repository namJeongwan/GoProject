package main

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scan(&a, &b) // 입력 두 개 받기
	/*
	 * Scan 은 정상적으로 입력받은 Item 의 개수와 입력 도중 Exception 이 발생한다면,
	 * Exception 의 Reason 을 반환한다.
	 * 즉 아래의 err != nil 은 if hasError 로 해석할 수 있겠다.
	 */
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
