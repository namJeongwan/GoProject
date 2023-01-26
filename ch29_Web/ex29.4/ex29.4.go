package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type model interface {
	ToJson()
}

type Student struct {
	Name  string
	Age   int
	Score int
}

func (s *Student) ToJson() ([]byte, error) {
	return json.Marshal(s)
}

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentHandler)
	return mux
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	student := &Student{"aaa", 16, 87}
	data, _ := student.ToJson()

	// Header Set
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK
	fmt.Fprint(w, string(data))
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
