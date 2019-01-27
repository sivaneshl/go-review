// https://thenewstack.io/make-a-restful-json-api-go/
// example from the above page

package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}
