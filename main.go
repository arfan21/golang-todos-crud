package main

import (
	"fmt"
	"net/http"

	"github.com/arfan21/golang-todos-crud/router"
	"github.com/gorilla/mux"
)

// @title Hacktiv8 TODOS API
// @version 1.0
// @description This is API for completing final project 1 hacktiv8
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host http://localhost:8000
// @BasePath /
func main() {
	mux := mux.NewRouter()
	router.New(mux)

	fmt.Println("http listen on http://localhost:8000")
	http.ListenAndServe(":8000", mux)
}
