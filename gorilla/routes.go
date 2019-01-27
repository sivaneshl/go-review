package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

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
