package main

import (
	"container/list"
	"fmt"
)

func main() {
	v := list.New()      // 새로운 리스트 생성
	e4 := v.PushBack(4)  // 리스트 뒤에 요소 추가
	v.PushBack(5)        // 굳이 변수에 안넣어줘도 된다!
	e1 := v.PushFront(1) // 리스트 앞에 요소 추가

	fmt.Println(e4.Value)
	// 밑에와 같은 접근은 초기화되지 않은 메모리에 접근 할 수 있는 위험이 있다... err로 잡는게 아닌갑네
	fmt.Println(e4.Next().Value)

	v.InsertBefore(3, e4) // e4 요소 앞에 요소 삽입
	v.InsertAfter(2, e1)  // e1 요소 뒤에 요소 삽입

	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}

	fmt.Println()
	// Back은 가장 뒷 요소 반환.. e.Prev는 Prev에 머물러있는 Element 참조중.. 언제까지? null이 안될 때 까지
	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

}
