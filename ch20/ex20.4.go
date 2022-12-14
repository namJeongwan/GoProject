package main

import (
	"GoProject/ch20/koreaPost"
	"fmt"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	koreaPostSender := &koreaPost.PostSender{}
	koreaPostSender2 := &koreaPost.PostSender{}

	koreaPostSender.SetAge(3)
	koreaPostSender2.SetAge(4)

	fmt.Println(koreaPostSender2.GetAge())
	fmt.Println(koreaPostSender.GetAge())

}
