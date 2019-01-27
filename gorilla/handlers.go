package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

// ToDoIndex ...
func ToDoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "To Do Index!")

	todos := ToDos{
		ToDo{Name: "Write Presentation"},
		ToDo{Name: "Host Meeting"},
		ToDo{Name: "Call VP"},
	}

	json.NewEncoder(w).Encode(todos)
}

// ToDoShow ...
func ToDoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Fprintln(w, "To Do Show: ", todoID)
}
