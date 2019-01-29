package main

func main() {

	a := App{}
	a.Initialize("admin", "admin", "testdb")
	a.Run(":8080")

}
