package main

import (
	"fmt"
	"net/http"

	"github.com/arfan21/golang-todos-crud/router"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	router.New(mux)

	fmt.Println("http listen on http://localhost:8000")
	http.ListenAndServe(":8000", mux)
}
