package router

import (
	"github.com/arfan21/golang-todos-crud/controller"
	"github.com/arfan21/golang-todos-crud/database"
	"github.com/gorilla/mux"
)

func New(mux *mux.Router) {
	listTodos := new(database.ListTodos)

	ctrl := controller.New(listTodos)

	mux.HandleFunc("/", ctrl.CreateTodo).Methods("POST")
}
