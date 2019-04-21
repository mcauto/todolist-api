package usecase

import (
	"todo-list-back/models"
	"todo-list-back/todo"
)

// TodoUsecase 할 일의 usecase
type TodoUsecase struct {
	repository todo.Repository
}

// NewTodoUsecase 생성자
func NewTodoUsecase(repository todo.Repository) todo.Usecase {
	return &TodoUsecase{
		repository: repository,
	}
}

// FindAll 다 가져오기
func (tu *TodoUsecase) FindAll() ([]*models.TodoItem, error) {
	return tu.repository.FindAll()
}

// Create 생성하기
func (tu *TodoUsecase) Create(item *models.TodoItem) error {
	return tu.repository.Create(item)
}

// FindByID ID를 이용하여 가져오기
func (tu *TodoUsecase) FindByID(id uint) (*models.TodoItem, error) {
	return tu.repository.FindByID(id)
}

// Update 수정하기
func (tu *TodoUsecase) Update(item *models.TodoItem) error {
	return tu.repository.Update(item)
}

// Delete 삭제하기
func (tu *TodoUsecase) Delete(item *models.TodoItem) error {
	return tu.repository.Delete(item)
}
