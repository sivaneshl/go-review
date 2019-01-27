package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {

		var handler http.Handler

		handler = route.HandlerFunc
		handler = logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(handler)
	}

	// the same can be written as this also
	// router.HandleFunc("/", Index).Methods("GET")
	// router.HandleFunc("/todos", ToDoIndex).Methods("GET")
	// router.HandleFunc("/todos/{todoId}", ToDoShow)

	return router
}
