package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", ToDoIndex)
	router.HandleFunc("/todos/{todoId}", ToDoShow)

	return router
}
