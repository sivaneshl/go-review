package main

import "fmt"

var currentid int
var todos ToDos

// init ...
func init() {
	RepoCreateToDo(ToDo{Name: "Write Presentation"})
	RepoCreateToDo(ToDo{Name: "Host Meeting"})
	RepoCreateToDo(ToDo{Name: "Call VP"})
}

// RepoCreateToDo ...
func RepoCreateToDo(t ToDo) ToDo {
	currentid++
	t.ID = currentid
	todos = append(todos, t)
	return t
}

// RepoFindToDo ...
func RepoFindToDo(id int) ToDo {
	for _, t := range todos {
		if t.ID == id {
			return t
		}
	}
	return ToDo{}
}

// RepoDelToDo ...
func RepoDelToDo(id int) error {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find ToDo with ID of %d", id)
}
