package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n') // 에러 발생 시 입력 스트림 비우기
	}

	return n, err
}

func main() {
	rand.Seed(time.Now().UnixNano())

	target := rand.Intn(100)

	for {
		fmt.Printf("Input Integer value: ")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("Please input only integer type..")
		} else {
			fmt.Println("Input Value => ", n)
			if target == n {
				fmt.Println("Match!")
				break
			} else if target > n {
				fmt.Println("Up..")
			} else if target < n {
				fmt.Println("Down..")
			}
		}
	}
	fmt.Println("Target Number is ", target)
}
