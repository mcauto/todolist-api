package todo

import "todolist-api/modules/domains/todo"

// ManyItemResponse is a response for todo
type ManyItemResponse struct {
	Todos []todo.Item `json:"todos"`
	Page  int         `json:"page"  example:"1"`
	Limit int         `json:"limit" example:"10"`
}
