package database

import (
	"errors"

	"github.com/arfan21/golang-todos-crud/model"
)

type ListTodos struct {
	List []model.Todo
}

func (l *ListTodos) Add(todo model.Todo) {
	l.List = append(l.List, todo)
}

func (l *ListTodos) GetAll() []model.Todo {
	return l.List
}

func (l *ListTodos) GetByID(id int) (model.Todo, error) {
	for _, val := range l.List {
		if val.ID == id {
			return val, nil
		}
	}
	return model.Todo{}, errors.New("Not Found")
}
