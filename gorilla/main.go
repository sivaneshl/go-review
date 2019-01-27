// https://thenewstack.io/make-a-restful-json-api-go/
// example from the above page

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", ToDoIndex)
	router.HandleFunc("/todos/{todoId}", ToDoShow)

	log.Fatal(http.ListenAndServe(":8080", router))

}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func ToDoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "To Do Index!")
}

func ToDoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "To Do Show: ", todoId)
}
