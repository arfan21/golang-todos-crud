package controller

import (
	"net/http"

	"github.com/arfan21/golang-todos-crud/database"
)

var res map[string]interface{}

type Controller struct {
	listTodos *database.ListTodos
}

func New(listTodos *database.ListTodos) *Controller {
	return &Controller{listTodos}
}

func (c *Controller) CreateTodo(w http.ResponseWriter, r *http.Request)     {}
func (c *Controller) GetAllTodos(w http.ResponseWriter, r *http.Request)    {}
func (c *Controller) GetTodoById(w http.ResponseWriter, r *http.Request)    {}
func (c *Controller) UpdateTodoById(w http.ResponseWriter, r *http.Request) {}
func (c *Controller) DeleteTodoById(w http.ResponseWriter, r *http.Request) {}
