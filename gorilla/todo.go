package main

import "time"

// ToDo ...
type ToDo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// ToDos ...
type ToDos []ToDo
