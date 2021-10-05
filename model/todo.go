package model

type Todo struct {
	ID   int    `json:"id" example:"1"`
	Name string `json:"name" example:"Todos API"`
}

type TodoRequest struct {
	Name string `json:"name" example:"Todos API"`
}
