// https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql

package main

func main() {

	a := App{}
	a.Initialize("postgres", "admin", "testdb")
	a.Run(":8080")

}
