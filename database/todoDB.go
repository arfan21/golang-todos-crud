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
		id = l.List[len(l.List)-1].ID + 1
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

func (l *ListTodos) DeleteTodoById(id int) error {
	var index *int
	for i, val := range l.List {
		if val.ID == id {
			index = &i
			break
		} else {
			index = nil
		}
	}

	if index == nil {
		return errors.New("id not found")
	}

	l.List[*index] = l.List[len(l.List)-1]
	l.List[len(l.List)-1] = model.Todo{}
	l.List = l.List[:len(l.List)-1]

	return nil
}
