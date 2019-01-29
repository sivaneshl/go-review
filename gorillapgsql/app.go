package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

// App ...
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize ...
func (a *App) Initialize(user, password, dbname string) {

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()

}

// Run ...
func (a *App) Run(addr string) {

}
