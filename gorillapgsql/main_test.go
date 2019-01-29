package main_test

import (
	"testing"
)

var a main.App

func TestMain(m *testing.M) {
	a = main.App{}

	a.Initialize("admin", "admin", "testdb")

	ensureTableExists()

}

func ensureTableExists() {

}
