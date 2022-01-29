package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/arfan21/golang-todos-crud/router"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

// @title Hacktiv8 TODOS API
// @version 1.0
// @description This is API for completing final project 1 hacktiv8
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /

func main() {
	mux := mux.NewRouter()
	router.New(mux)

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8000"
	}

	fmt.Println("http listen on http://localhost:" + port)
	http.ListenAndServe(":"+port, mux)
}
