package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"sort"
	"strconv"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

var students map[int]Student
var lastId int

func MakeWebHandler() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/students", GetStudentListHandler).Methods("GET")
	muxRouter.HandleFunc("/students/{id:[0-9]+}", GetStudentHandler).Methods("GET")
	muxRouter.HandleFunc("/students", PostStudentHandler).Methods("POST")
	muxRouter.HandleFunc("/students/{id:[0-9]+}", DeleteStudentHandler).Methods("DELETE")

	students = make(map[int]Student)
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 17, 45}
	lastId = 2

	return muxRouter
}

func GetStudentListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Students, 0)
	for _, student := range students {
		list = append(list, student)
	}

	sort.Sort(list)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	student, ok := students[id]
	w.Header().Set("Content-Type", "application/json")
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Not Found Student Object.."})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func PostStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student Student
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Bad Request.. Check your parameter"})
		return
	}
	lastId++
	student.Id = lastId
	json.NewEncoder(w).Encode(map[string]string{"result": "OK"})
	students[lastId] = student
	w.WriteHeader(http.StatusCreated)
}

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(vars["id"])
	_, ok := students[id]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Not Found Student Object.."})
	}

	delete(students, id)
	w.WriteHeader(http.StatusOK)

}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
