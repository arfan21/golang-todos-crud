package router

import (
	"github.com/arfan21/golang-todos-crud/controller"
	"github.com/arfan21/golang-todos-crud/database"
	_ "github.com/arfan21/golang-todos-crud/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func New(mux *mux.Router) {
	var listTodos database.ListTodos

	ctrl := controller.New(&listTodos)

	mux.HandleFunc("/todos", ctrl.CreateTodo).Methods("POST")

	mux.HandleFunc("/todos", ctrl.GetAllTodos).Methods("GET")
	// Handle get todo by id
	mux.HandleFunc("/todos/{id}", ctrl.GetTodoById).Methods("GET")
	// Handle update todo by id
	mux.HandleFunc("/todos/{id}", ctrl.UpdateTodoById).Methods("PUT")

	mux.HandleFunc("/todos/{id}", ctrl.DeleteTodoById).Methods("DELETE")

	// routing swagger API
	mux.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
}
