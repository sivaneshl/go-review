package main

import (
	"net/http"
)

// Route ...
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

// Routes ...
type Routes []Route

var routes = Routes{
	Route{
		"Index", "GET", "/", Index,
	},
	Route{
		"ToDoIndex", "GET", "/todos", ToDoIndex,
	},
	Route{
		"ToDoShow", "GET", "/todos/{todoId}", ToDoShow,
	},
}
