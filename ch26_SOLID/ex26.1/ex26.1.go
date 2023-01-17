package main

import "fmt"

type Event interface {
	Register(EventListener)
}

type EventListener interface {
	OnFire()
}

type Mail struct {
	listener EventListener
}

func (m *Mail) Register(listener EventListener) {
	m.listener = listener
}

func (m *Mail) OnRecv() {
	m.listener.OnFire()
}

type Alarm struct {
}

func (a *Alarm) OnFire() {
	fmt.Println("Alarm!!")
}

// Message

type Phone struct {
	listener EventListener
}

func (p *Phone) Register(listener EventListener) {
	p.listener = listener
}

func (p *Phone) OnRecv() {
	p.listener.OnFire()
}

type Message struct {
}

func (m *Message) OnFire() {
	fmt.Println("Received Message..")
}

func main() {
	mail := &Mail{}
	listener := &Alarm{}

	mail.Register(listener)
	mail.OnRecv()

	phone := &Phone{}
	message := &Message{}

	phone.Register(message)
	phone.OnRecv()

}
