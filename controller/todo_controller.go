package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/arfan21/golang-todos-crud/database"
	"github.com/arfan21/golang-todos-crud/helper"
	"github.com/arfan21/golang-todos-crud/model"
)

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

// GetAllTodos
// @Tags Todos
// @Summary get all todos .
// @Description get all todso from the database
// @Accept json
// @Produce json
// @Success 200 {object} helper.BaseResponse{data=[]model.Todo}
// @Failure 400 {object} helper.BaseResponse{}
// @Router /todos [GET]
func (c *Controller) GetAllTodos(w http.ResponseWriter, r *http.Request) {

	resultTodo := c.listTodos.GetAll()

	respon := helper.BaseResponse{
		Status:  http.StatusOK,
		Message: "success get all data",
		Data:    resultTodo,
	}
	helper.Response(w, respon)
}

// GetTodoById
// @Tags Todos
// @Summary get a todo for a given todo id.
// @Description get a todo with id from the database
// @Accept json
// @Produce json
// @Param ID path int true "ID of the todo"
// @Success 200 {object} helper.BaseResponse{data=model.Todo}
// @Failure 400 {object} helper.BaseResponse{}
// @Router /todos/{id} [GET]
func (c *Controller) GetTodoById(w http.ResponseWriter, r *http.Request) {

	// handling path ID
	vars := mux.Vars(r)
	todoID, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		helper.Response(w, response)
		return
	}

	resultTodo, err := c.listTodos.GetTodoById(todoID)
	if err != nil {
		response := helper.BaseResponse{
			Status:  http.StatusOK,
			Message: err.Error(),
			Data:    nil,
		}
		helper.Response(w, response)
		return
	}

	response := helper.BaseResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    resultTodo,
	}
	helper.Response(w, response)
}

// UpdateTodoById
// @Tags Todos
// @Summary Update todo with given todo id.
// @Description Update the todo corresponding to the input id
// @Accept json
// @Produce json
// @Param todo body model.TodoRequest true "update"
// @Param ID path int true "ID of the todo"
// @Success 200 {object} helper.BaseResponse{data=model.Todo}
// @Failure 400 {object} helper.BaseResponse{}
// @Router /todos/{id} [PUT]
func (c *Controller) UpdateTodoById(w http.ResponseWriter, r *http.Request) {

	// handling path ID
	vars := mux.Vars(r)
	todoID, err := strconv.Atoi(vars["id"])
	if err != nil {
		helper.Response(w, helper.BaseResponse{Status: http.StatusBadRequest, Message: err.Error(), Data: nil})
		return
	}

	// read body request
	readBody, errRead := ioutil.ReadAll(r.Body)
	if errRead != nil {
		helper.Response(w, helper.BaseResponse{Status: http.StatusBadRequest, Message: err.Error(), Data: nil})
		return
	}
	var todo model.Todo
	// unmarshal to struct
	err = json.Unmarshal(readBody, &todo)

	if err != nil {
		helper.Response(w, helper.BaseResponse{Status: http.StatusBadRequest, Message: err.Error(), Data: nil})
		return
	}
	// call database function update
	todo.ID = todoID
	err = c.listTodos.UpdateTodoById(todo)
	if err != nil {
		helper.Response(w, helper.BaseResponse{Status: http.StatusNotFound, Message: err.Error(), Data: nil})
		return
	}

	// successreturn
	helper.Response(w, helper.BaseResponse{Status: http.StatusOK, Message: "success update", Data: todo})
}

// DeleteTodoById
// @Tags Todos
// @Summary delete todo with given todo id.
// @Description delete the todo corresponding to the input id
// @Accept json
// @Produce json
// @Param ID path int true "ID of the todo"
// @Success 200 {object} helper.BaseResponse{}
// @Failure 400 {object} helper.BaseResponse{}
// @Router /todos/{id} [DELETE]
func (c *Controller) DeleteTodoById(w http.ResponseWriter, r *http.Request) {

	// handling path ID
	vars := mux.Vars(r)
	todoID, err := strconv.Atoi(vars["id"])
	if err != nil {
		helper.Response(w, helper.BaseResponse{Status: http.StatusBadRequest, Message: err.Error(), Data: nil})
		return
	}

	err = c.listTodos.DeleteTodoById(todoID)
	if err != nil {
		helper.Response(w, helper.BaseResponse{Status: http.StatusNotFound, Message: err.Error(), Data: nil})
		return
	}

	// success
	helper.Response(w, helper.BaseResponse{Status: http.StatusOK, Message: "success Delete", Data: nil})
}
