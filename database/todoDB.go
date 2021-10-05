package database

import (
	"errors"

	"github.com/arfan21/golang-todos-crud/model"
)

type ListTodos struct {
	List []model.Todo
}

func (l *ListTodos) Add(todo model.Todo) model.Todo {
	id := 1
	if len(l.List) != 0 {
		id = l.List[len(l.List)-1].ID
	}
	todo.ID = id
	l.List = append(l.List, todo)
	return todo
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
