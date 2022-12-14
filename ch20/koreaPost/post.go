package koreaPost

import "fmt"

type PostSender struct {
	age int
}

func (k *PostSender) Send(parcel string) {
	fmt.Printf("Korea Post sends %v parcel\n", parcel)
}

func (k *PostSender) SetAge(age int) {
	k.age = age
}

func (k *PostSender) GetAge() int {
	return k.age
}
