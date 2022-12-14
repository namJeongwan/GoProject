package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func GetStringLength(str string) int {
	return len([]rune(str))
}

type TestStruct struct {
	Data string
}

func main() {

	str := "안녕 세상"

	fmt.Printf("Byte Size => {%d}\n", len(str))
	fmt.Printf("Length => {%d}\n", GetStringLength(str))

	for i, arr := 0, []rune(str); i < len(arr); i++ {
		fmt.Printf("Type: %T, Value: %d, String-value: %c\n", arr[i], arr[i], arr[i])
	}

	testStruct1 := TestStruct{"ABC"}
	testStruct2 := testStruct1

	testStruct2.Data = "DDD"

	fmt.Println(testStruct1.Data)
	fmt.Println(testStruct2.Data)

	testString1 := "ABC"
	testString2 := testString1

	//testString2 = "DDD"

	fmt.Println(testString1)
	fmt.Println(testString2)

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&testString1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&testString2))

	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)

}
