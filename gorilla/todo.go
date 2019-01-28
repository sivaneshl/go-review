package main

import "time"

// ToDo ...
type ToDo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// ToDos ...
type ToDos []ToDo
