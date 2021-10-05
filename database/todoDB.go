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

func (l *ListTodos) GetTodoById(id int) (model.Todo, error) {
	for _, val := range l.List {
		if val.ID == id {
			return val, nil
		}
	}
	return model.Todo{}, errors.New("not found")
}

func (l *ListTodos) UpdateTodoById(todo model.Todo) error {
	for idx, val := range l.List {
		if val.ID == todo.ID {
			l.List[idx].Name = todo.Name
			return nil
		}
	}
	return errors.New("id not found")
}
