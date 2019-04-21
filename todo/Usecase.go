package todo

import "todo-list-back/models"

// Usecase Todo의 Usecase interface
type Usecase interface {
	FindAll() ([]*models.TodoItem, error)
	Create(*models.TodoItem) error
	FindByID(id uint) (*models.TodoItem, error)
	Update(*models.TodoItem) error
	Delete(*models.TodoItem) error
}
