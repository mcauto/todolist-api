package todo

import "todo-list-back/models"

// Repository Todo의 Repository 인터페이스; 의존 역전 원칙
type Repository interface {
	FindAll() ([]*models.TodoItem, error)
	Create(*models.TodoItem) error
	FindByID(id uint) (*models.TodoItem, error)
	Update(*models.TodoItem) error
	Delete(*models.TodoItem) error
}
