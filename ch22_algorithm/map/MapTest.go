package main

import "fmt"

type testType struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]string)
	m["test1"] = "test1"

	testStruct := testType{"ABC", 5}

	a := m["test2"]

	fmt.Printf("%T\n", a)
	fmt.Printf("Key: test2, Val: %v\n", a)

	testMap := make(map[interface{}]interface{})

	testMap["a"] = 1
	testMap[testStruct] = testStruct // 오 문제 없이 실행되는구만... 쓸데가 있을까 모르겠지만..

	fmt.Printf("%v\n, %T", testMap[testStruct], testMap[testStruct])

}
