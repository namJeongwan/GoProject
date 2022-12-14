package main

import "GoProject/ch20/fedex"

func SendBook2(name string, sender fedex.FSender) {
	sender.Send(name)
}

func main() {
	sender := fedex.FSender{"Joe"}
	SendBook2("어린 왕자", sender)
	SendBook2("Greece Jotbab", sender)

}
