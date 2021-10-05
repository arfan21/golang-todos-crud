package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/arfan21/golang-todos-crud/database"
	"github.com/arfan21/golang-todos-crud/helper"
	"github.com/arfan21/golang-todos-crud/model"
)

var res map[string]interface{}

type Controller struct {
	listTodos *database.ListTodos
}

func New(listTodos *database.ListTodos) *Controller {
	return &Controller{listTodos}
}

// CreateTodo
// @Tags Todos
// @Summary Create a new Todo to the database
// @Description Create a new Todo
// @Accept json
// @Produce json
// @Param todo body model.TodoRequest true "create a new todo"
// @Success 200 {object} helper.BaseResponse{data=model.Todo}
// @Failure 400 {object} helper.BaseResponse
// @Router /todos [POST]
func (c *Controller) CreateTodo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helper.Response(w, helper.BaseResponse{Status: http.StatusBadRequest, Message: err.Error(), Data: nil})
		return
	}

	todo := new(model.Todo)
	err = json.Unmarshal(body, todo)
	if err != nil {
		helper.Response(w, helper.BaseResponse{Status: http.StatusBadRequest, Message: err.Error(), Data: nil})
		return
	}
	createdTodo := c.listTodos.Add(*todo)
	helper.Response(w, helper.BaseResponse{Status: http.StatusOK, Message: "success create data", Data: createdTodo})
}
func (c *Controller) GetAllTodos(w http.ResponseWriter, r *http.Request)    {}
func (c *Controller) GetTodoById(w http.ResponseWriter, r *http.Request)    {}
func (c *Controller) UpdateTodoById(w http.ResponseWriter, r *http.Request) {}
func (c *Controller) DeleteTodoById(w http.ResponseWriter, r *http.Request) {}
