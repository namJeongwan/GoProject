package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type flyable interface {
	Fly()
}

type runable interface {
	Run()
}

type airplane struct {
}

func (a *airplane) Fly() {
	fmt.Println("Flying..")
}

func (a *airplane) Run() {
	fmt.Println("Run..")
}

func testFunc(f flyable) {
	f.Fly()
}

func testFunc2(r runable) {
	r.Run()
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")
	if name == "" {
		name = "World"
	}
	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "Hello %s! id:%d", name, id)
}

func main() {
	ap := &airplane{}
	testFunc(ap)
	testFunc2(ap)

	http.HandleFunc("/bar", barHandler)
	http.HandleFunc("/bar2", barHandler)
	http.ListenAndServe(":3000", nil)
}
